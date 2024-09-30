package config

import (
	"github.com/go-celery/wide"
	"github.com/go-celery/wide/enum"
)

type RedisBackendLoadItem struct {
	Load   bool
	Config RedisConfig
}

func (r *RedisBackendLoadItem) IsLoad() bool {
	return r.Load
}

func (r *RedisBackendLoadItem) Verify() error {
	return r.Config.VerifyItemConfig()
}

func (r *RedisBackendLoadItem) GetType() enum.LoadItemType {
	return enum.BackendType
}

func (r *RedisBackendLoadItem) GetLoadItemConfig() celerywide.LoadItemConfig {
	return &r.Config
}

type NilBackendLoadItem struct {
	Load   bool
	Config NilConfig
}

func (n *NilBackendLoadItem) IsLoad() bool {
	return n.Load
}

func (n *NilBackendLoadItem) Verify() error {
	return n.Config.VerifyItemConfig()
}

func (n *NilBackendLoadItem) GetType() enum.LoadItemType {
	return enum.BackendType
}

func (n *NilBackendLoadItem) GetLoadItemConfig() celerywide.LoadItemConfig {
	return &n.Config
}
