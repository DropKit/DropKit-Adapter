package logger

import "github.com/sirupsen/logrus"

func InfoNormal(content string) {
	logrus.Info(content)
}

func InfoDBConnection() {
	logrus.Info("Database connection successful.")
}

func InfoQuorumConnection() {
	logrus.Info("Quorum connection successful.")
}
