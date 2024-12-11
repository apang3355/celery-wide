package producer

import (
	"errors"
	celerywide "github.com/apang3355/celery-wide"
	"github.com/apang3355/celery-wide/enum"
	"github.com/apang3355/celery-wide/errs"
	"time"
)

type ArgsProducer[T any] struct {
	coreClient *celerywide.CoreClient
	retarder   celerywide.Retarder
	logger     celerywide.Logger
}

func NewArgs[T any](coreClient *celerywide.CoreClient, retarder celerywide.Retarder, logger celerywide.Logger) (celerywide.Producer[T], error) {
	producer := &ArgsProducer[T]{
		coreClient: coreClient,
		retarder:   retarder,
		logger:     logger,
	}

	if coreClient == nil {
		return producer, errors.New("new producer core celery client is nil")
	}

	if retarder == nil {
		return producer, errors.New("new producer retarder is nil")
	}

	if logger == nil {
		return producer, errors.New("new producer logger is nil")
	}

	return producer, nil
}

func (a *ArgsProducer[T]) Send(ctx celerywide.Context, queue enum.QueueName, consumer enum.ConsumerName, data T) (enum.TaskID, error) {
	a.logger.Info(ctx, "send message", queue, consumer, data)
	message, err := celerywide.NewMessage[T](ctx, queue, consumer, data)
	if err != nil {
		return "", err
	}

	return a.delay(message)
}

func (a *ArgsProducer[T]) DelaySend(ctx celerywide.Context, queue enum.QueueName, consumer enum.ConsumerName, delay time.Duration, data T) error {
	a.logger.Info(ctx, "send message", queue, consumer, data)
	message, err := celerywide.NewMessage[T](ctx, queue, consumer, data)
	if err != nil {
		return err
	}

	if err = a.retarder.Execute(delay, func() {
		if _, err = a.delay(message); err != nil {
			a.logger.Error(ctx, errs.NewErrorMessage("celery producer  error: %s", err.Error()))
			return
		}
	}); err != nil {
		return err
	}

	return nil
}

func (a *ArgsProducer[T]) delay(message *celerywide.Message[T]) (enum.TaskID, error) {
	if message == nil {
		return "", errors.New("message is nil")
	}

	asyncResult, err := a.coreClient.Delay(message.Consumer.Value(), message.ToMap())
	if err != nil {
		return "", err
	}

	return enum.TaskID(asyncResult.TaskID), nil
}
