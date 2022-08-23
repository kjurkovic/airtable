package service

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/kjurkovic/airtable/service/auth/datastore"
	"github.com/kjurkovic/airtable/service/auth/middleware"
	"github.com/kjurkovic/airtable/service/auth/models"
	"github.com/kjurkovic/airtable/service/auth/util"
	"github.com/kjurkovic/airtable/service/auth/wrappers"
	audit "gitlab.redox.media/theria/client-audit-service"
)

func (service *UserService) EditUserPassword(rw http.ResponseWriter, r *http.Request) {
	id := uuid.MustParse(mux.Vars(r)["id"])
	request, err := service.validateChangePasswordRequest(&rw, r)

	if err != nil {
		return
	}

	user, err := datastore.UserDao.GetById(id)

	if err != nil {
		models.UserNotFoundError.SendErrorResponse(rw, http.StatusBadRequest)
		return
	}

	err = datastore.UserDao.Update(&models.User{
		Id:       user.Id,
		Password: util.GeneratePasswordHashWithSalt(request.Password),
	})

	if err != nil {
		models.UserNotFoundError.SendErrorResponse(rw, http.StatusBadRequest)
		return
	} else {
		auditObj, err := util.ToJson(user)
		if err != nil {
			auditObj = id.String()
		}
		claims := r.Context().Value(middleware.KeyClaims{}).(*models.Claims)
		wrappers.Audit.SendEvent(claims.UserId, auditObj, audit.UpdateUserPassword)
	}

	rw.WriteHeader(http.StatusAccepted)
}

func (service *UserService) validateChangePasswordRequest(rw *http.ResponseWriter, r *http.Request) (*models.PasswordChangeRequest, error) {
	request := &models.PasswordChangeRequest{}
	err := request.Deserialize(r.Body)

	if err != nil {
		service.Log.ErrorS("Error deserializing request", r.Body)
		models.SerializationError.SendErrorResponse(*rw, http.StatusBadRequest)
		return nil, err
	}

	err = request.Validate()

	if err != nil {
		httpErr := models.PasswordValidationError
		httpErr.Message = err.Error()
		httpErr.SendErrorResponse(*rw, http.StatusBadRequest)
		return nil, err
	}

	return request, nil
}
