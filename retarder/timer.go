package retarder

import (
	celerywide "github.com/go-celery/celery-wite"
	"github.com/go-celery/celery-wite/config"
	"time"
)

type Timer struct {
}

func NewTimer(config *config.TimerConfig) (celerywide.Retarder, error) {
	return &Timer{}, nil
}

func (t *Timer) Execute(delay time.Duration, f func()) error {
	if delay == 0 {
		f()
	}

	time.AfterFunc(delay, f)
	return nil
}
