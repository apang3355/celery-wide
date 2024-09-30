package logger

import (
	"context"
	"fmt"
	"github.com/go-celery/wide"
	"github.com/go-celery/wide/funcs"
	jsoniter "github.com/json-iterator/go"
)

type Default struct{}

func NewDefault() celerywide.Logger {
	return &Default{}
}

func (d *Default) Info(ctx celerywide.Context, msg string, fields ...any) {
	fmt.Println(NewMessageJson(ctx, "info", msg, fields))
}

func (d *Default) Warn(ctx celerywide.Context, msg string, fields ...any) {
	fmt.Println(NewMessageJson(ctx, "warn", msg, fields))
}

func (d *Default) Error(ctx celerywide.Context, msg string, fields ...any) {
	fmt.Println(NewMessageJson(ctx, "error", msg, fields))
}

func (d *Default) Debug(ctx celerywide.Context, msg string, fields ...any) {
	fmt.Println(NewMessageJson(ctx, "debug", msg, fields))
}

type Message struct {
	Level    string              `json:"level"`
	Msg      string              `json:"msg" json:"msg,omitempty"`
	Transmit celerywide.Transmit `json:"transmit" json:"transmit,omitempty"`
	Fields   []any               `json:"fields" json:"fields,omitempty"`
}

func NewMessageJson(ctx celerywide.Context, level string, msg string, fields []any) string {
	var transmit = make(celerywide.Transmit)
	if ctx != nil {
		transmit = ctx.GetTransmit()
	}
	return (&Message{
		Level:    level,
		Msg:      msg,
		Transmit: transmit,
		Fields:   fields,
	}).ToJson()
}

func (m *Message) ToJson() string {
	json, err := jsoniter.MarshalToString(m)
	if err != nil {
		return NewMessageJson(celerywide.NewContext(context.Background(), []funcs.TransmitFromContext{}), "error", err.Error(), []any{})
	}

	return json
}
