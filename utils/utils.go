package utils

import (
	"github.com/apang3355/celery-wide/config"
	"github.com/gomodule/redigo/redis"
	"time"
)

func NewRedisPool(config *config.RedisConfig) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     config.MaxIdle,
		MaxActive:   config.MaxActive,
		IdleTimeout: time.Duration(config.IdleTimeout) * time.Millisecond,
		Dial: func() (redis.Conn, error) {
			redisConn, err := redis.DialURL(config.Dsn)
			if err != nil {
				return nil, err
			}

			return redisConn, nil
		},
	}
}
