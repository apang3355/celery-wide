package retarder

import (
	celerywide "github.com/go-celery/wide"
	"github.com/go-celery/wide/config"
	"github.com/go-errors/errors"
)

type Redis struct{}

func NewRedis(config config.RedisConfig) (celerywide.Retarder, error) {
	return nil, errors.New("redis retarder not impl")
}
