package config

import (
	"errors"
	"github.com/go-celery/wide"
	"github.com/go-celery/wide/funcs"
	"strings"
)

type RedisConfig struct {
	Dsn         string
	MaxIdle     int
	MaxActive   int
	IdleTimeout int
}

func (r *RedisConfig) VerifyItemConfig() error {
	if strings.TrimSpace(r.Dsn) == "" {
		return errors.New("redis dsn is empty")
	}

	if r.MaxIdle == 0 {
		return errors.New("redis max idle is 0")
	}

	if r.MaxActive == 0 {
		return errors.New("redis max active is 0")
	}

	if r.IdleTimeout == 0 {
		return errors.New("redis idle timeout is 0")
	}

	return nil
}

type TimerConfig struct{}

func (t *TimerConfig) VerifyItemConfig() error {
	return nil
}

type RocketConfig struct {
	Host string
}

func (r *RocketConfig) VerifyItemConfig() error {
	if strings.TrimSpace(r.Host) == "" {
		return errors.New("rocket host is empty")
	}
	return nil
}

type TransmitConfig struct {
	TransmitFromContexts []funcs.TransmitFromContext
}

func (r *TransmitConfig) VerifyItemConfig() error {
	return nil
}

type InnerLoggerConfig struct {
	Switch bool
	Logger celerywide.Logger
}

func (r *InnerLoggerConfig) VerifyItemConfig() error {
	if r.Logger == nil {
		return errors.New("inner log config logger is nil")
	}
	return nil
}

type NilConfig struct{}

func (r *NilConfig) VerifyItemConfig() error {
	return nil
}
