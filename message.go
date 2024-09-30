package celerywide

import (
	"errors"
	"github.com/go-celery/wide/enum"
	jsoniter "github.com/json-iterator/go"
	"strings"
)

type Message[T any] struct {
	Context  Context           `json:"context"`
	Queue    enum.QueueName    `json:"queue"`
	Consumer enum.ConsumerName `json:"consumer"`
	Data     T                 `json:"data"`
}

func NewMessage[T any](ctx Context, queue enum.QueueName, consumer enum.ConsumerName, data T) (*Message[T], error) {
	message := &Message[T]{
		Context:  ctx,
		Queue:    queue,
		Consumer: consumer,
		Data:     data,
	}
	if err := message.verify(); err != nil {
		return nil, err
	}

	return message, nil
}

func NewMessageFromMap[T any](kwargs map[string]any) (*Message[T], error) {
	if kwargs == nil {
		return nil, errors.New("kwargs is nil")
	}

	json, err := jsoniter.Marshal(kwargs)
	if err != nil {
		return nil, err
	}

	return NewMessageFromJson[T](json)
}

func NewMessageFromJson[T any](json []byte) (*Message[T], error) {
	celeryContext, err := NewContextFromJson(jsoniter.Get(json, "context").ToString())
	if err != nil {
		return nil, err
	}

	queueName := enum.NewQueueName(jsoniter.Get(json, "queue").ToString())
	if strings.TrimSpace(queueName.Value()) == "" {
		queueName = "-"
	}

	consumerName := enum.NewConsumerName(jsoniter.Get(json, "task").ToString())
	if strings.TrimSpace(consumerName.Value()) == "" {
		consumerName = "-"
	}

	var data T
	if err = jsoniter.UnmarshalFromString(jsoniter.Get(json, "data").ToString(), &data); err != nil {
		return nil, err
	}

	return NewMessage[T](celeryContext, queueName, consumerName, data)
}

func (m *Message[T]) ToMap() map[string]any {
	return map[string]any{
		"context":  m.Context,
		"queue":    m.Queue.Value(),
		"consumer": m.Consumer.Value(),
		"data":     m.Data,
	}
}

func (m *Message[T]) verify() error {
	if m.Context == nil {
		return errors.New("context is nil")
	}

	if err := m.Queue.Verify(); err != nil {
		return err
	}

	if err := m.Consumer.Verify(); err != nil {
		return err
	}

	return nil
}
