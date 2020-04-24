package logger

import "github.com/sirupsen/logrus"

func FatelDependencyService(service string, err error) {
	ServiceConnectorLogger.WithField("component", "service-connector").WithFields(logrus.Fields{"service": service, "result": "Fail"}).Fatal(err.Error())
}

func InfoDependencyService(service string) {
	ServiceConnectorLogger.WithField("component", "service-connector").WithFields(logrus.Fields{"service": service, "result": "Success"}).Info()
}