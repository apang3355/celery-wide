package celerywide

import (
	jsoniter "github.com/json-iterator/go"
)

type Result[T any] struct {
	Data T `json:"data"`
}

func (r Result[T]) ToJson() (string, error) {
	json, err := jsoniter.MarshalToString(r)
	if err != nil {
		return "", err
	}

	return json, nil
}

func NewResult[T any](data T) Result[T] {
	return Result[T]{
		Data: data,
	}
}
