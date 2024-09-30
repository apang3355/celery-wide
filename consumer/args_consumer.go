package consumer

import (
	"errors"
	"github.com/go-celery/wide"
	"github.com/go-celery/wide/enum"
	"github.com/go-celery/wide/funcs"
	jsoniter "github.com/json-iterator/go"
)

type ArgsConsumer[MessageData, ResultData any] struct {
	coreClient *celerywide.CoreClient
	logger     celerywide.Logger
	message    *celerywide.Message[MessageData]
	task       funcs.Task[MessageData, ResultData]
}

func NewArgs[MessageData, ResultData any](coreClient *celerywide.CoreClient, logger celerywide.Logger) (celerywide.Consumer[MessageData, ResultData], error) {
	if coreClient == nil {
		return nil, errors.New("new consumer core client is nil")
	}

	if logger == nil {
		return nil, errors.New("new consumer logger is nil")
	}

	return &ArgsConsumer[MessageData, ResultData]{
		coreClient: coreClient,
		logger:     logger,
	}, nil
}

func (a *ArgsConsumer[MessageData, ResultData]) Register(consumer enum.ConsumerName, task funcs.Task[MessageData, ResultData]) error {
	a.coreClient.Register(consumer.Value(), func(messageAny any) (string, error) {
		var err error
		var resultJson string
		var ctx celerywide.Context
		defer func() {
			if err == nil {
				a.logger.Info(ctx, "args consumer", map[string]any{"message": messageAny}, map[string]any{"result": resultJson})
				return
			}

			a.logger.Error(nil, "celery args consumer error", map[string]any{"error": err.Error()}, map[string]any{"message": messageAny})
		}()
		messageJson, err := jsoniter.Marshal(messageAny)
		if err != nil {
			return "", err
		}

		message, err := celerywide.NewMessageFromJson[MessageData](messageJson)
		if err != nil {
			return "", err
		}

		ctx = message.Context
		resultData, err := task(ctx, message.Data)
		if err != nil {
			return "", err
		}

		if resultJson, err = celerywide.NewResult[ResultData](resultData).ToJson(); err != nil {
			return "", err
		}

		return resultJson, nil
	})
	return nil
}
