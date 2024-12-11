package celerywide

import (
	"github.com/go-celery/celery-wite/config"
	"github.com/go-celery/celery-wite/factory"
	"github.com/go-celery/celery-wite/loader"
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
