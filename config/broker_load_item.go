package config

import (
	"github.com/go-celery/wide"
	"github.com/go-celery/wide/enum"
)

type RedisBrokerLoadItem struct {
	Load   bool
	Config RedisConfig
}

func (r *RedisBrokerLoadItem) GetType() enum.LoadItemType {
	return enum.BrokerType
}

func (r *RedisBrokerLoadItem) Verify() error {
	return r.Config.VerifyItemConfig()
}

func (r *RedisBrokerLoadItem) IsLoad() bool {
	return r.Load
}

func (r *RedisBrokerLoadItem) GetLoadItemConfig() celerywide.LoadItemConfig {
	return &r.Config
}

type RocketBrokerLoadItem struct {
	Load         bool
	RocketConfig RocketConfig
}

func (r *RocketBrokerLoadItem) GetType() enum.LoadItemType {
	return enum.BrokerType
}

func (r *RocketBrokerLoadItem) Verify() error {
	return r.RocketConfig.VerifyItemConfig()
}

func (r *RocketBrokerLoadItem) IsLoad() bool {
	return r.Load
}

func (r *RocketBrokerLoadItem) GetLoadItemConfig() celerywide.LoadItemConfig {
	return &r.RocketConfig
}
