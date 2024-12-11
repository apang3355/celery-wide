package celerywide

import (
	"github.com/apang3355/celery-wide/config"
	"github.com/apang3355/celery-wide/factory"
	"github.com/apang3355/celery-wide/loader"
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
