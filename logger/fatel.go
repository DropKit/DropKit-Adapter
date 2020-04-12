package logger

import "github.com/sirupsen/logrus"

func FatelDBConnection(reason error) {
	logrus.WithFields(logrus.Fields{"reason": reason.Error()}).Fatal("Database connection failed")
}

func FatelQuorumConnection(reason error) {
	logrus.WithFields(logrus.Fields{"reason": reason.Error()}).Fatal("Quorum connection failed")
}
