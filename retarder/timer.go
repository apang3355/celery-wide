package retarder

import (
	celerywide "github.com/apang3355/celery-wide"
	"github.com/apang3355/celery-wide/config"
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
