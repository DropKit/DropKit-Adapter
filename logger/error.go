package logger

import "github.com/sirupsen/logrus"

func ErrNormal(reason error) {
	logrus.Error(reason.Error())
}
