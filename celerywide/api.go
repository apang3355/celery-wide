package celerywide

import (
	"context"
	"github.com/go-celery/wide/enum"
	"github.com/go-celery/wide/factory"
	"github.com/go-celery/wide/funcs"
	"time"
)

// RegisterConsumer 注册消费者: Kwargs 参数模式
func RegisterConsumer[MessageData, ResultData any](queueName enum.QueueName, consumerName enum.ConsumerName, numWorkers int, task funcs.Task[MessageData, ResultData]) error {
	consumer, err := factory.CreateKwargsConsumer[MessageData, ResultData](queueName, numWorkers)
	if err != nil {
		return err
	}

	return consumer.Register(consumerName, task)
}

// Send 发送消息: Kwargs 参数模式
func Send[T any](ctx context.Context, queue enum.QueueName, consumerName enum.ConsumerName, data T) (enum.TaskID, error) {
	producer, err := factory.CreateKwargsProducer[T]()
	if err != nil {
		return "", err
	}

	return producer.Send(factory.CreateContext(ctx), queue, consumerName, data)
}

// DelaySend 发送延迟消息: Kwargs 参数模式
func DelaySend[T any](ctx context.Context, queue enum.QueueName, consumerName enum.ConsumerName, delay time.Duration, data T) error {
	producer, err := factory.CreateKwargsProducer[T]()
	if err != nil {
		return err
	}

	return producer.DelaySend(factory.CreateContext(ctx), queue, consumerName, delay, data)
}

// RegisterArgsConsumer 注册消费者: Args 参数模式
func RegisterArgsConsumer[MessageData, ResultData any](queueName enum.QueueName, consumerName enum.ConsumerName, numWorkers int, task funcs.Task[MessageData, ResultData]) error {
	consumer, err := factory.CreateArgsConsumer[MessageData, ResultData](queueName, numWorkers)
	if err != nil {
		return err
	}

	return consumer.Register(consumerName, task)
}

// ArgsSend 发送消息: Args 参数模式
func ArgsSend[T any](ctx context.Context, queue enum.QueueName, task enum.ConsumerName, data T) (enum.TaskID, error) {
	producer, err := factory.CreateArgsProducer[T]()
	if err != nil {
		return "", err
	}

	return producer.Send(factory.CreateContext(ctx), queue, task, data)
}

// ArgsDelaySend 发送延迟消息: Args 参数模式
func ArgsDelaySend[T any](ctx context.Context, queue enum.QueueName, task enum.ConsumerName, delay time.Duration, data T) error {
	producer, err := factory.CreateArgsProducer[T]()
	if err != nil {
		return nil
	}

	return producer.DelaySend(factory.CreateContext(ctx), queue, task, delay, data)
}
