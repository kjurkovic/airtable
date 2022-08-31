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

type Timeout struct {
	// seconds
	Idle     time.Duration `yaml:"idle"`
	Read     time.Duration `yaml:"read"`
	Write    time.Duration `yaml:"write"`
	Shutdown time.Duration `yaml:"shutdown"`
}

type Logger struct {
	Level log.LogLevel `yaml:"level"`
}

type FromConfig struct {
	Email string `yaml:"email"`
	Name  string `yaml:"name"`
}

type Mailer struct {
	Key  string     `yaml:"key"`
	From FromConfig `yaml:"from"`
}

type Config struct {
	Server   Server   `yaml:"server"`
	Mailer   Mailer   `yaml:"mailer"`
	Services Services `yaml:"services"`
}

type Services struct {
	Audit string `yaml:"audit"`
}
