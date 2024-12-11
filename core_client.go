package celerywide

import (
	"github.com/apang3355/celery-wide/enum"
	"github.com/gocelery/gocelery"
)

type CoreClient struct {
	*gocelery.CeleryClient
	Queue enum.QueueName
}
