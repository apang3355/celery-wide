package celerywide

import (
	"github.com/go-celery/celery-wite/enum"
	"github.com/go-celery/celery-wite/funcs"
	"github.com/gocelery/gocelery"
	"time"
)

type Producer[T any] interface {
	Send(ctx Context, queue enum.QueueName, task enum.ConsumerName, data T) (enum.TaskID, error)
	DelaySend(ctx Context, queue enum.QueueName, task enum.ConsumerName, delay time.Duration, data T) error
}

type Consumer[MessageData any, ResultData any] interface {
	Register(taskName enum.ConsumerName, task funcs.Task[MessageData, ResultData]) error
}

type LoadItemConfig interface {
	VerifyItemConfig() error
}

type LoadItem interface {
	IsLoad() bool
	Verify() error
	GetType() enum.LoadItemType
	GetLoadItemConfig() LoadItemConfig
}

type Loader interface {
	GetRetarder() Retarder
	GetLogger() Logger
	GetTransmitFromContexts() []funcs.TransmitFromContext
	GetProducerCoreClient() *CoreClient
	CreateConsumerCoreClient(queue enum.QueueName, numWorkers int) (*CoreClient, error)
	GetAllCoreClients() []*CoreClient
}

type Logger interface {
	Info(ctx Context, msg string, field ...any)
	Warn(ctx Context, msg string, field ...any)
	Error(ctx Context, msg string, field ...any)
	Debug(ctx Context, msg string, field ...any)
}

type Execute interface {
	Run() error
}

type Retarder interface {
	Execute(delay time.Duration, f func()) error
}

type Broker interface {
	gocelery.CeleryBroker
	GetQueueName() (enum.QueueName, error)
}
