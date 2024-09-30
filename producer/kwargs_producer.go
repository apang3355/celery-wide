package producer

import (
	"errors"
	celerywide "github.com/go-celery/wide"
	"github.com/go-celery/wide/enum"
	"github.com/go-celery/wide/errs"
	"time"
)

type KwargsProducer[T any] struct {
	coreClient *celerywide.CoreClient
	retarder   celerywide.Retarder
	logger     celerywide.Logger
}

func NewKwargs[T any](coreClient *celerywide.CoreClient, retarder celerywide.Retarder, logger celerywide.Logger) (celerywide.Producer[T], error) {
	producer := &KwargsProducer[T]{
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

func (p *KwargsProducer[T]) Send(ctx celerywide.Context, queue enum.QueueName, consumer enum.ConsumerName, data T) (enum.TaskID, error) {
	p.logger.Info(ctx, "send message", queue, consumer, data)
	message, err := celerywide.NewMessage[T](ctx, queue, consumer, data)
	if err != nil {
		return "", err
	}

	return p.delay(message)
}

func (p *KwargsProducer[T]) DelaySend(ctx celerywide.Context, queue enum.QueueName, task enum.ConsumerName, delay time.Duration, data T) error {
	p.logger.Info(ctx, "send message", queue, task, data)
	message, err := celerywide.NewMessage[T](ctx, queue, task, data)
	if err != nil {
		return err
	}

	if err = p.retarder.Execute(delay, func() {
		if _, err = p.delay(message); err != nil {
			p.logger.Error(ctx, errs.NewErrorMessage("celery producer  error: %s", err.Error()))
			return
		}
	}); err != nil {
		return err
	}

	return nil
}

func (p *KwargsProducer[T]) delay(message *celerywide.Message[T]) (enum.TaskID, error) {
	if message == nil {
		return "", errors.New("message is nil")
	}

	asyncResult, err := p.coreClient.DelayKwargs(message.Consumer.Value(), message.ToMap())
	if err != nil {
		return "", err
	}

	return enum.TaskID(asyncResult.TaskID), nil
}
