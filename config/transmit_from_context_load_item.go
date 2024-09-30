package config

import (
	"github.com/go-celery/wide"
	"github.com/go-celery/wide/enum"
)

type GoContextLoadItem struct {
	Load   bool
	Config TransmitConfig
}

func (t *GoContextLoadItem) IsLoad() bool {
	return t.Load
}

func (t *GoContextLoadItem) Verify() error {
	return t.Config.VerifyItemConfig()
}

func (t *GoContextLoadItem) GetType() enum.LoadItemType {
	return enum.TransmitType
}

func (t *GoContextLoadItem) GetLoadItemConfig() celerywide.LoadItemConfig {
	return &t.Config
}
