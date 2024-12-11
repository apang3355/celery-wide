package test

import (
	"context"
	"fmt"
	"git.kxsz.net/gopkg/utils"
	"github.com/apang3355/celery-wide/celerywide"
	"github.com/apang3355/celery-wide/config"
	"github.com/apang3355/celery-wide/enum"
	"github.com/apang3355/celery-wide/funcs"
	"github.com/apang3355/celery-wide/logger"
	jsoniter "github.com/json-iterator/go"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func init() {
	redisConfig := config.RedisConfig{
		Dsn:         "redis://:password@127.0.0.1:6379/0",
		MaxIdle:     20,
		MaxActive:   200,
		IdleTimeout: 100,
	}
	celeryConfig := celerywide.BuildRedisConfig(redisConfig, []funcs.TransmitFromContext{}, logger.NewDefault())
	if err := celerywide.Init(celeryConfig); err != nil {
		panic(err)
	}

	if err := celerywide.RegisterConsumer(queueName1, taskName1, 1, func(ctx context.Context, param Param) (Result, error) {
		fmt.Printf("消费消息: queue=%s, task=%s, num_workers=%d, param=%s\n", queueName1, taskName1, 1, param.ToJson())
		return Result{
			Code: 0,
		}, nil
	}); err != nil {
		panic(err)
	}

	if err := celerywide.RegisterConsumer(queueName2, taskName2, 1, func(ctx context.Context, param Param) (Result, error) {
		fmt.Printf("消费消息: queue=%s, task=%s, num_workers=%d, param=%s\n", queueName2, taskName2, 1, param.ToJson())
		return Result{
			Code: 0,
		}, nil
	}); err != nil {
		panic(err)
	}
}

const queueName1 enum.QueueName = "test_send_queue1"
const queueName2 enum.QueueName = "test_send_queue2"
const taskName1 enum.ConsumerName = "test_send_task1"
const taskName2 enum.ConsumerName = "test_send_task2"

type Param struct {
	Name string `json:"name"`
}

type Result struct {
	Code int `json:"code"`
}

func (p *Param) ToJson() string {
	json, _ := jsoniter.MarshalToString(p)
	return json
}

func Test_Send(t *testing.T) {
	param := Param{
		Name: "tom",
	}
	ctx := context.Background()
	ctx = context.WithValue(ctx, "trace_id", utils.GenUUID())
	taskID1, err := celerywide.Send[Param](ctx, queueName1, taskName1, param)
	fmt.Printf("task_id1:%s\n", taskID1)
	assert.NoError(t, err)

	taskID2, err := celerywide.Send[Param](ctx, queueName2, taskName2, param)
	fmt.Printf("task_id2:%s\n", taskID2)
	assert.NoError(t, err)

	time.Sleep(10 * time.Second)
}

func Test_DelaySend(t *testing.T) {
	param := Param{
		Name: "tom",
	}
	ctx := context.Background()
	err := celerywide.DelaySend[Param](ctx, queueName1, taskName1, 3*time.Second, param)
	assert.NoError(t, err)

	//taskID2, err := celeryapi.Send[Param](ctx, queueName, taskName2, param)
	//fmt.Printf("task_id2:%s\n", taskID2)
	//assert.NoError(t, err)

	time.Sleep(10 * time.Second)
}
