package celerywide

import (
	celerywide "github.com/apang3355/celery-wide"
	"github.com/apang3355/celery-wide/config"
	"github.com/apang3355/celery-wide/funcs"
)

type DefaultRedisConfig struct {
	Builder
	RedisConfig config.RedisConfig
}

func BuildRedisConfig(redis config.RedisConfig, transmitFromContexts []funcs.TransmitFromContext, logger celerywide.Logger) config.Config {
	return config.Config{
		Context: config.ContextConfig{
			GoContext: config.GoContextLoadItem{
				Load: true,
				Config: config.TransmitConfig{
					TransmitFromContexts: transmitFromContexts,
				},
			},
		},
		Broker: config.BrokerLoadConfig{
			Redis: config.RedisBrokerLoadItem{
				Load:   true,
				Config: redis,
			},
		},
		Backend: config.BackendConfig{
			Redis: config.RedisBackendLoadItem{
				Load:   true,
				Config: redis,
			},
		},
		Retarder: config.RetarderConfig{
			Timer: config.TimerRetarderLoadItem{
				Load:   true,
				Config: config.TimerConfig{},
			},
		},
		Logger: config.LoggerConfig{
			Text: config.TextLoggerLoadItem{
				Load: true,
				Config: config.InnerLoggerConfig{
					Switch: true,
					Logger: logger,
				},
			},
		},
	}
}
