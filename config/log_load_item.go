package config

import (
	"github.com/go-celery/wide"
	"github.com/go-celery/wide/enum"
)

type TextLoggerLoadItem struct {
	Load   bool
	Config InnerLoggerConfig
}

func (l *TextLoggerLoadItem) IsLoad() bool {
	return l.Load
}

func (l *TextLoggerLoadItem) Verify() error {
	return l.Config.VerifyItemConfig()
}

func (l *TextLoggerLoadItem) GetType() enum.LoadItemType {
	return enum.LogType
}

func (l *TextLoggerLoadItem) GetLoadItemConfig() celerywide.LoadItemConfig {
	return &l.Config
}
