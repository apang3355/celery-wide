package enum

import (
	"github.com/go-errors/errors"
	"strings"
)

type LoadItemType string

const (
	BrokerType   LoadItemType = "broker"
	BackendType  LoadItemType = "backend"
	RetarderType LoadItemType = "retarder"
	TransmitType LoadItemType = "transmit"
	LogType      LoadItemType = "logger"
)

type TaskID string

func (t TaskID) Value() string {
	return string(t)
}

type ConsumerName string

func (t ConsumerName) Value() string {
	return string(t)
}

func (t ConsumerName) Verify() error {
	if strings.TrimSpace(t.Value()) == "" {
		return errors.New("consumer name is empty")
	}

	return nil
}

func NewConsumerName(text string) ConsumerName {
	return ConsumerName(text)
}

type QueueName string

func (t QueueName) Value() string {
	return string(t)
}

func (t QueueName) Verify() error {
	if strings.TrimSpace(t.Value()) == "" {
		return errors.New("queue name is empty")
	}

	return nil
}

func NewQueueName(queueName string) QueueName {
	return QueueName(queueName)
}
