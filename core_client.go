package celerywide

import (
	"github.com/go-celery/wide/enum"
	"github.com/gocelery/gocelery"
)

type CoreClient struct {
	*gocelery.CeleryClient
	Queue enum.QueueName
}
