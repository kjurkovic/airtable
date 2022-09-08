package config

import (
	"time"

	"github.com/google/uuid"
	"xorm.io/xorm/log"
)

type Server struct {
	Port           int       `yaml:"port"`
	Secret         string    `yaml:"secret"`
	AllowedOrigins []string  `yaml:"allowedOrigins"`
	Debug          bool      `yaml:"debug"`
	Timeout        Timeout   `yaml:"timeout"`
	SystemUUID     uuid.UUID `yaml:"systemUuid"`
}

type Logger struct {
	Level log.LogLevel `yaml:"level"`
}

type Database struct {
	Name     string `yaml:"name"`
	Ip       string `yaml:"ip"`
	Port     int    `yaml:"port"`
	Schema   string `yaml:"schema"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Logger   Logger `yaml:"logger"`
}

type Config struct {
	Server   Server   `yaml:"server"`
	Database Database `yaml:"database"`
	Services Services `yaml:"services"`
}

type Timeout struct {
	// seconds
	Idle     time.Duration `yaml:"idle"`
	Read     time.Duration `yaml:"read"`
	Write    time.Duration `yaml:"write"`
	Shutdown time.Duration `yaml:"shutdown"`
}

type Services struct {
	Meta string `yaml:"meta"`
}
