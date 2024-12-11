package consumer

import (
	celerywide "github.com/go-celery/celery-wite"
	"github.com/go-celery/celery-wite/enum"
	"github.com/go-celery/celery-wite/funcs"
	"github.com/go-errors/errors"
)

type KwargsConsumer[MessageData any, ResultData any] struct {
	coreClient *celerywide.CoreClient
	logger     celerywide.Logger
	message    *celerywide.Message[MessageData]
	task       funcs.Task[MessageData, ResultData]
}

func NewKwargs[MessageData any, ResultData any](coreClient *celerywide.CoreClient, logger celerywide.Logger) (celerywide.Consumer[MessageData, ResultData], error) {
	if coreClient == nil {
		return nil, errors.New("new consumer core client is nil")
	}

	if logger == nil {
		return nil, errors.New("new consumer logger is nil")
	}

	return &KwargsConsumer[MessageData, ResultData]{
		coreClient: coreClient,
		logger:     logger,
	}, nil
}

func (c *KwargsConsumer[MessageData, ResultData]) Register(consumer enum.ConsumerName, task funcs.Task[MessageData, ResultData]) error {
	c.task = task
	c.coreClient.Register(consumer.Value(), c)
	return nil
}

func (c *KwargsConsumer[MessageData, ResultData]) ParseKwargs(m map[string]any) error {
	message, err := celerywide.NewMessageFromMap[MessageData](m)
	if err != nil {
		var ctx celerywide.Context
		if message != nil {
			ctx = message.Context
		}
		c.logger.Error(ctx, "kwargs consumer", map[string]any{"message": m})
		return err
	}

	c.message = message
	return nil
}

func (c *KwargsConsumer[MessageData, ResultData]) RunTask() (any, error) {
	var err error
	var result celerywide.Result[ResultData]
	defer func() {
		messageMap := map[string]any{"message": c.message}
		resultMap := map[string]any{"result": result}
		if err != nil {
			c.logger.Error(c.message.Context, "kwargs consumer", messageMap, resultMap)
			return
		}

		c.logger.Info(c.message.Context, "kwargs consumer", messageMap, resultMap)
	}()

	data, err := c.task(c.message.Context, c.message.Data)
	if err != nil {
		result = celerywide.NewResult(data)
		return result, err
	}

	result = celerywide.NewResult(data)
	return result, nil
}
