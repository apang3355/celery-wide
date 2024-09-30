package retarder

import (
	"github.com/go-celery/wide"
	"github.com/go-celery/wide/config"
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
