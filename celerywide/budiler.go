package celerywide

import (
	celerywide "github.com/go-celery/celery-wite"
	"github.com/go-celery/celery-wite/funcs"
)

type Builder struct {
	Logger               celerywide.Logger
	TransmitFromContexts []funcs.TransmitFromContext
}
