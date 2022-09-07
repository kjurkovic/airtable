package service

import (
	"crypto/sha256"
	"encoding/hex"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/kjurkovic/airtable/service/auth/config"
	"github.com/kjurkovic/airtable/service/auth/datastore"
	"github.com/kjurkovic/airtable/service/auth/models"
	"github.com/kjurkovic/airtable/service/auth/util"
	"github.com/kjurkovic/airtable/service/auth/wrappers"
)

type AuthService struct {
	Log    *util.Logger
	Config *config.Config
}

type UserService struct {
	Log    *util.Logger
	Config *config.Config
}

func (service *AuthService) generateAuthResponse(user *models.User, rw http.ResponseWriter) *models.AuthResponse {
	expirationTime := time.Now().Add(10 * time.Minute)
	claims := &models.Claims{
		UserId:      user.Id,
		Username:    user.Email,
		ProductRole: user.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(service.Config.Server.Secret))

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return nil
	}

	refreshToken, expiresAt := service.generateRefreshToken(user)

	response := &models.AuthResponse{
		AccessToken:           tokenString,
		AccessTokenExpiresAt:  expirationTime.UnixMilli(),
		RefreshToken:          refreshToken,
		RefreshTokenExpiresAt: expiresAt,
	}

	auditObj, err := util.ToJson(user)
	if err != nil {
		auditObj = user.Id.String()
	}
	wrappers.Audit.SendEvent(service.Config.Server.SystemUUID, auditObj, wrappers.Login)

	return response
}

func (service *AuthService) generateRefreshToken(user *models.User) (string, int64) {
	expirationTime := time.Now().Add(10 * 24 * time.Hour).UnixMilli()
	var sb strings.Builder
	sb.WriteString(user.Id.String())
	sb.WriteString(time.Now().GoString())
	hash := sha256.Sum256([]byte(sb.String()))
	refreshToken := hex.EncodeToString(hash[:])
	datastore.RefreshTokenDao.Save(&models.RefreshToken{
		Id:        uuid.New(),
		UserId:    user.Id,
		Token:     refreshToken,
		ExpiresAt: expirationTime,
	})
	return refreshToken, expirationTime
}
