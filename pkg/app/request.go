package app

import (
	"go-admin/pkg/logging"

	"github.com/astaxie/beego/validation"
	"go.uber.org/zap"
)

func MarkErrors(errors []*validation.Error) {
	for _, err := range errors {
		logging.Logger.Info(err.Key, zap.String("message", err.Message))
	}
}
