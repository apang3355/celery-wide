package backend

import (
	"github.com/go-celery/celery-wite/config"
	"github.com/go-celery/celery-wite/utils"
	"github.com/gocelery/gocelery"
)

type Redis struct {
	*gocelery.RedisCeleryBackend
}

func NewRedis(redisConfig *config.RedisConfig) (gocelery.CeleryBackend, error) {
	return &Redis{
		RedisCeleryBackend: gocelery.NewRedisBackend(utils.NewRedisPool(redisConfig)),
	}, nil
}

func (r *Redis) GetResult(s string) (*gocelery.ResultMessage, error) {
	return r.RedisCeleryBackend.GetResult(s)
}

func (r *Redis) SetResult(taskID string, result *gocelery.ResultMessage) error {
	return r.RedisCeleryBackend.SetResult(taskID, result)
}
