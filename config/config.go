package config

import (
	"github.com/go-celery/wide"
)

type Config struct {
	Broker   BrokerLoadConfig
	Backend  BackendConfig
	Retarder RetarderConfig
	Context  ContextConfig
	Logger   LoggerConfig
}

func (c *Config) GetLoadItems() []celerywide.LoadItem {
	return []celerywide.LoadItem{
		&c.Broker.Redis,
		&c.Broker.Rocket,
		&c.Retarder.Redis,
		&c.Retarder.Timer,
		&c.Context.GoContext,
		&c.Logger.Text,
		&c.Backend.Redis,
		&c.Backend.Nil,
	}
}

type BrokerLoadConfig struct {
	Redis  RedisBrokerLoadItem
	Rocket RocketBrokerLoadItem
}

type RetarderConfig struct {
	Timer TimerRetarderLoadItem
	Redis RedisRetarderLoadItem
}

type ContextConfig struct {
	GoContext GoContextLoadItem
}

type LoggerConfig struct {
	Text TextLoggerLoadItem
}

type BackendConfig struct {
	Redis RedisBackendLoadItem
	Nil   NilBackendLoadItem
}
