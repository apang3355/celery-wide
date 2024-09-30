package celerywide

import (
	celerywide "github.com/go-celery/wide"
	"github.com/go-celery/wide/funcs"
)

type Builder struct {
	Logger               celerywide.Logger
	TransmitFromContexts []funcs.TransmitFromContext
}
