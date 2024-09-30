package celerywide

import (
	"github.com/go-celery/wide/config"
	"github.com/go-celery/wide/factory"
	"github.com/go-celery/wide/loader"
)

func Init(config config.Config) error {
	loaderInstance, err := loader.New(config.GetLoadItems())
	if err != nil {
		return err
	}

	if err = factory.Init(loaderInstance); err != nil {
		return err
	}

	return nil
}
