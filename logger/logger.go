package logger

import (
	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/sirupsen/logrus"
)

var ServiceConnectorLogger = logrus.New()
var APIServerLogger = logrus.New()
var InternalLogger = logrus.New()

func init() {
	ServiceConnectorLogger.SetFormatter(&nested.Formatter{
		HideKeys:        true,
		FieldsOrder:     []string{"component", "service", "result"},
		TimestampFormat: "2006-01-02 15:04:05",
	})

	APIServerLogger.SetFormatter(&nested.Formatter{
		HideKeys:        true,
		FieldsOrder:     []string{"component", "command"},
		TimestampFormat: "2006-01-02 15:04:05",
	})

	InternalLogger.SetFormatter(&nested.Formatter{
		HideKeys:        true,
		FieldsOrder:     []string{"component", "path"},
		TimestampFormat: "2006-01-02 15:04:05",
	})
}
