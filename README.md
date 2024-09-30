# celery-wide
--- ---
概述：针对gocelery在消息队列场景下的扩展，例如：生产消费、发布订阅、定时消息、延迟消息等。风格灵活、屏蔽内部消息组件复杂性。

当前版本概述: 目前初步构建了可扩展框架，支持了基本的生产消费、延迟消息能力。消息中间件方面目前仅支持redis。下个版本将扩展基于redis的延迟消息，安全关闭服务，发布订阅模式。

迭代日志:
v1.0.0: 组件支持: redis。消息队列能力支持: 生产消费、延迟消息 

## 简单使用
### 初始化
```go
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
```

### 注册消费者
```go
if err := celerywide.RegisterConsumer(queueName1, taskName1, 1, func(ctx context.Context, param Param) (Result, error) {
    fmt.Printf("消费消息: queue=%s, task=%s, num_workers=%d, param=%s\n", queueName1, taskName1, 1, param.ToJson())
    return Result{
        Code: 0,
    }, nil
}); err != nil {
    panic(err)
}
```

### 发送消息
```go
// 发送即时消息
taskID1, err := celerywide.Send[Param](ctx, queueName1, taskName1, param)

// 发送延迟消息
err := celerywide.DelaySend[Param](ctx, queueName1, taskName1, 3*time.Second, param)
```