package broker

import (
	"fmt"
	"github.com/apang3355/celery-wide/config"
	"github.com/apang3355/celery-wide/enum"
	"github.com/apang3355/celery-wide/utils"
	"github.com/go-errors/errors"
	"github.com/gocelery/gocelery"
	"github.com/gomodule/redigo/redis"
	jsoniter "github.com/json-iterator/go"
)

type RedisConsumer struct {
	*redis.Pool
	QueueName enum.QueueName
}

func NewRedisConsumer(config *config.RedisConfig, queue enum.QueueName) (celerywide.Broker, error) {
	return &RedisConsumer{
		QueueName: queue,
		Pool:      utils.NewRedisPool(config),
	}, nil
}

func (cb *RedisConsumer) GetTaskMessage() (*gocelery.TaskMessage, error) {
	celeryMessage, err := cb.GetCeleryMessage()
	if err != nil {
		return nil, err
	}
	return celeryMessage.GetTaskMessage(), nil
}

func (cb *RedisConsumer) SendCeleryMessage(message *gocelery.CeleryMessage) error {
	return errors.New("redis consumer cannot send message")
}

func (cb *RedisConsumer) GetCeleryMessage() (*gocelery.CeleryMessage, error) {
	conn := cb.Get()
	defer func() {
		_ = conn.Close()
	}()
	messageJSON, err := conn.Do("BRPOP", cb.QueueName, "0")
	if err != nil {
		return nil, err
	}
	if messageJSON == nil {
		return nil, fmt.Errorf("null message received from redis")
	}
	messageList := messageJSON.([]interface{})
	if string(messageList[0].([]byte)) != cb.QueueName.Value() {
		return nil, fmt.Errorf("not a celery message: %v", messageList[0])
	}
	var message gocelery.CeleryMessage
	if err := jsoniter.Unmarshal(messageList[1].([]byte), &message); err != nil {
		return nil, err
	}
	return &message, nil
}

func (cb *RedisConsumer) GetQueueName() (enum.QueueName, error) {
	return cb.QueueName, nil
}
