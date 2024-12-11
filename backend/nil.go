package backend

import (
	"github.com/apang3355/celery-wide/config"
	"github.com/gocelery/gocelery"
)

type Nil struct{}

func NewNil(config *config.NilConfig) (gocelery.CeleryBackend, error) {
	return new(Nil), nil
}

func (c *Nil) GetResult(s string) (*gocelery.ResultMessage, error) {
	return nil, nil
}

func (c *Nil) SetResult(taskID string, result *gocelery.ResultMessage) error {
	return nil
}
