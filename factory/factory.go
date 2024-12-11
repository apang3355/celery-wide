package factory

import (
	"context"
	celerywide "github.com/apang3355/celery-wide"
	"github.com/apang3355/celery-wide/backend"
	"github.com/apang3355/celery-wide/broker"
	"github.com/apang3355/celery-wide/config"
	"github.com/apang3355/celery-wide/consumer"
	"github.com/apang3355/celery-wide/enum"
	"github.com/apang3355/celery-wide/funcs"
	"github.com/apang3355/celery-wide/logger"
	"github.com/apang3355/celery-wide/producer"
	"github.com/apang3355/celery-wide/retarder"
	"github.com/go-errors/errors"
	"github.com/gocelery/gocelery"
)

type Assert[T celerywide.LoadItemConfig, D any] func(T) (D, error)

var loader celerywide.Loader

func Init(loaderInstance celerywide.Loader) error {
	if loaderInstance == nil {
		return errors.New("factory init loader is nil")
	}

	loader = loaderInstance
	return nil
}

func CreateContext(ctx context.Context) celerywide.Context {
	return celerywide.NewContext(ctx, getLoader().GetTransmitFromContexts())
}

func CreateArgsProducer[T any]() (celerywide.Producer[T], error) {
	return producer.NewArgs[T](getLoader().GetProducerCoreClient(), getLoader().GetRetarder(), getLoader().GetLogger())
}

func CreateKwargsProducer[T any]() (celerywide.Producer[T], error) {
	return producer.NewKwargs[T](getLoader().GetProducerCoreClient(), getLoader().GetRetarder(), getLoader().GetLogger())
}

func CreateKwargsConsumer[MessageData, ResultData any](queue enum.QueueName, numWorkers int) (celerywide.Consumer[MessageData, ResultData], error) {
	coreClient, err := getLoader().CreateConsumerCoreClient(queue, numWorkers)
	if err != nil {
		return nil, err
	}

	return consumer.NewKwargs[MessageData, ResultData](coreClient, getLoader().GetLogger())
}

func CreateArgsConsumer[MessageData, ResultData any](queue enum.QueueName, numWorkers int) (celerywide.Consumer[MessageData, ResultData], error) {
	coreClient, err := getLoader().CreateConsumerCoreClient(queue, numWorkers)
	if err != nil {
		return nil, err
	}

	return consumer.NewArgs[MessageData, ResultData](coreClient, getLoader().GetLogger())
}

func CreateConsumerBroker(loadItemConfig celerywide.LoadItemConfig, queue enum.QueueName) (celerywide.Broker, error) {
	if d, ok, err := assertComponent(loadItemConfig, func(t *config.RedisConfig) (celerywide.Broker, error) {
		return broker.NewRedisConsumer(t, queue)
	}); err != nil || ok {
		return d, err
	}

	return nil, errors.New("create consumer broker assert component not found")
}

func CreateProducerBroker(loadItemConfig celerywide.LoadItemConfig) (celerywide.Broker, error) {
	if d, ok, err := assertComponent(loadItemConfig, func(t *config.RedisConfig) (celerywide.Broker, error) {
		return broker.NewRedisProducer(t)
	}); err != nil || ok {
		return d, err
	}

	return nil, errors.New("create producer broker assert component not found")
}

func CreateRetarder(loadItemConfig celerywide.LoadItemConfig) (celerywide.Retarder, error) {
	if d, ok, err := assertComponent(loadItemConfig, func(t *config.TimerConfig) (celerywide.Retarder, error) {
		return retarder.NewTimer(t)
	}); err != nil || ok {
		return d, err
	}

	return nil, errors.New("create retarder assert component not found")
}

func CreateBackend(loadItemConfig celerywide.LoadItemConfig) (gocelery.CeleryBackend, error) {
	if d, ok, err := assertComponent(loadItemConfig, func(t *config.RedisConfig) (gocelery.CeleryBackend, error) {
		return backend.NewRedis(t)
	}); err != nil || ok {
		return d, err
	}

	if d, ok, err := assertComponent(loadItemConfig, func(t *config.NilConfig) (gocelery.CeleryBackend, error) {
		return backend.NewNil(t)
	}); err != nil || ok {
		return d, err
	}

	return nil, errors.New("create backend assert component not found")
}

func CreateTransmit(loadItemConfig celerywide.LoadItemConfig) ([]funcs.TransmitFromContext, error) {
	if d, ok, err := assertComponent(loadItemConfig, func(t *config.TransmitConfig) ([]funcs.TransmitFromContext, error) {
		return t.TransmitFromContexts, nil
	}); err != nil || ok {
		return d, err
	}

	return nil, errors.New("create transmit assert component not found")
}

func CreateLogger(loadItemConfig celerywide.LoadItemConfig) (celerywide.Logger, error) {
	if d, ok, err := assertComponent(loadItemConfig, func(t *config.InnerLoggerConfig) (celerywide.Logger, error) {
		return logger.New(t)
	}); err != nil || ok {
		return d, err
	}

	return nil, errors.New("create logger assert component not found")
}

func assertComponent[T celerywide.LoadItemConfig, D any](loadItemConfig celerywide.LoadItemConfig, assert Assert[T, D]) (D, bool, error) {
	var nilD D
	if loadItemConfig == nil {
		return nilD, false, errors.New("config loader broker load item config is nil")
	}

	if err := loadItemConfig.VerifyItemConfig(); err != nil {
		return nilD, false, err
	}

	if t, ok := loadItemConfig.(T); ok {
		d, err := assert(t)
		if err != nil {
			return d, ok, err
		}

		return d, ok, nil
	}

	return nilD, false, nil
}

func getLoader() celerywide.Loader {
	return loader
}
