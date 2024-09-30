package celerywide

import (
	"context"
	"github.com/go-celery/wide/funcs"
	jsoniter "github.com/json-iterator/go"
	"time"
)

type Context interface {
	Deadline() (deadline time.Time, ok bool)
	Done() <-chan struct{}
	Err() error
	Value(key any) any
	Set(key string, value any)
	Parent() context.Context
	GetTransmit() Transmit
}

type Transmit map[string]any

type BaseContext struct {
	ctx      context.Context
	Transmit Transmit `json:"transmit" json:"transmit,omitempty"`
}

func NewContext(ctx context.Context, transmitFromContexts []funcs.TransmitFromContext) Context {
	if ctx == nil {
		ctx = context.Background()
	}
	transmit := make(Transmit)
	for _, transmitFromContext := range transmitFromContexts {
		key, value := transmitFromContext(ctx)
		transmit[key] = value
	}

	return &BaseContext{
		ctx:      ctx,
		Transmit: transmit,
	}
}

func NewContextFromJson(json string) (Context, error) {
	var c BaseContext
	if err := jsoniter.UnmarshalFromString(json, &c); err != nil {
		return nil, err
	}

	c.ctx = context.Background()
	return &c, nil
}

func (c *BaseContext) Deadline() (deadline time.Time, ok bool) {
	return c.ctx.Deadline()
}

func (c *BaseContext) Done() <-chan struct{} {
	return c.ctx.Done()
}

func (c *BaseContext) Err() error {
	return c.ctx.Err()
}

func (c *BaseContext) Value(key any) any {
	keyStr, ok := key.(string)
	if !ok {
		return nil
	}
	return c.Transmit[keyStr]
}

func (c *BaseContext) Set(key string, value any) {
	c.Transmit[key] = value
}

func (c *BaseContext) Parent() context.Context {
	return c.ctx
}

func (c *BaseContext) GetTransmit() Transmit {
	return c.Transmit
}
