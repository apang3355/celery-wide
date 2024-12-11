package broker

import (
	"errors"
	"github.com/apang3355/celery-wide/config"
	"github.com/apang3355/celery-wide/enum"
	"github.com/apang3355/celery-wide/utils"
	"github.com/gocelery/gocelery"
	"github.com/gomodule/redigo/redis"
	jsoniter "github.com/json-iterator/go"
)

type RedisProducer struct {
	*redis.Pool
	message *gocelery.CeleryMessage
}

func NewRedisProducer(config *config.RedisConfig) (celerywide.Broker, error) {
	return &RedisProducer{
		Pool: utils.NewRedisPool(config),
	}, nil
}

func (r *RedisProducer) GetTaskMessage() (*gocelery.TaskMessage, error) {
	return nil, errors.New("redis producer cannot receive messages")
}

func (r *RedisProducer) SendCeleryMessage(message *gocelery.CeleryMessage) error {
	messageJson, err := jsoniter.Marshal(message)
	if err != nil {
		return err
	}

	conn := r.Get()
	defer func() {
		_ = conn.Close()
	}()

	r.message = message
	queue, err := r.GetQueueName()
	if err != nil {
		return err
	}

	_, err = conn.Do("LPUSH", queue, messageJson)
	if err != nil {
		return err
	}

	return nil
}

func (r *RedisProducer) GetQueueName() (enum.QueueName, error) {
	if r.message == nil {
		return "", nil
	}

	taskMessage := r.message.GetTaskMessage()
	if taskMessage == nil {
		return "", errors.New("redis producer get task message is nil")
	}

	kwargs := taskMessage.Kwargs
	if len(kwargs) != 0 {
		return enum.NewQueueName(kwargs["queue"].(string)), nil
	}

	args := taskMessage.Args
	if len(args) != 0 {
		return enum.NewQueueName(args[0].(map[string]any)["queue"].(string)), nil
	}

	return "", errors.New("queue name not found")
}
