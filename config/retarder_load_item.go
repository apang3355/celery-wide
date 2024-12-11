package config

import (
	celerywide "github.com/apang3355/celery-wide"
	"github.com/apang3355/celery-wide/enum"
)

type TimerRetarderLoadItem struct {
	Load   bool
	Config TimerConfig
}

func (r *TimerRetarderLoadItem) GetType() enum.LoadItemType {
	return enum.RetarderType
}

func (r *TimerRetarderLoadItem) Verify() error {
	return r.Config.VerifyItemConfig()
}

func (r *TimerRetarderLoadItem) IsLoad() bool {
	return r.Load
}

func (r *TimerRetarderLoadItem) GetLoadItemConfig() celerywide.LoadItemConfig {
	return &r.Config
}

type RedisRetarderLoadItem struct {
	Config RedisConfig
	Load   bool
}

func (r *RedisRetarderLoadItem) GetType() enum.LoadItemType {
	return enum.RetarderType
}

func (r *RedisRetarderLoadItem) Verify() error {
	return r.Config.VerifyItemConfig()
}

func (r *RedisRetarderLoadItem) IsLoad() bool {
	return r.Load
}

func (r *RedisRetarderLoadItem) GetLoadItemConfig() celerywide.LoadItemConfig {
	return &r.Config
}
