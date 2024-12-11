package celerywide

import (
	celerywide "github.com/apang3355/celery-wide"
	"github.com/apang3355/celery-wide/funcs"
)

type Builder struct {
	Logger               celerywide.Logger
	TransmitFromContexts []funcs.TransmitFromContext
}
