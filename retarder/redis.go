package retarder

import (
	celerywide "github.com/apang3355/celery-wide"
	"github.com/apang3355/celery-wide/config"
	"github.com/go-errors/errors"
)

type Redis struct{}

func NewRedis(config config.RedisConfig) (celerywide.Retarder, error) {
	return nil, errors.New("redis retarder not impl")
}
