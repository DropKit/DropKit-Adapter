package logger

import "github.com/sirupsen/logrus"

func WarnNormal(reason error) {
	logrus.Error(reason.Error())
}
