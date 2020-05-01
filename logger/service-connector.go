package logger

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func ErrorDependencyService(service string, err error) {
	ServiceConnectorLogger.WithField("component", "service-connector").WithFields(logrus.Fields{"service": service}).Error(err.Error())
}

func FatalDependencyService(service string, err error) {
	ServiceConnectorLogger.WithField("component", "service-connector").WithFields(logrus.Fields{"service": service}).Fatal(err.Error())
}

func InfoDependencyService() {
	ServiceConnectorLogger.WithField("component", "service-connector").WithFields(logrus.Fields{"result": "Success"}).Info("API server start at port " + viper.GetString(`DROPKIT.PORT`))
}
