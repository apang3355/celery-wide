package funcs

import "context"

type Task[MessageData, ResultData any] func(context.Context, MessageData) (ResultData, error)
type TransmitFromContext func(ctx context.Context) (key string, value any)
