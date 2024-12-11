package broker

import (
	"github.com/apang3355/celery-wide/config"
	"github.com/go-errors/errors"
	"github.com/gocelery/gocelery"
	"github.com/gomodule/redigo/redis"
)

type Rocket struct {
	*redis.Pool
	QueueName string
}

func NewRocket(config config.RocketConfig) (gocelery.CeleryBroker, error) {
	return nil, errors.New("rocket broker not currently supported")
}
