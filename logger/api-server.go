package logger

import "github.com/sirupsen/logrus"

func InfoAPIDatabaseCreate(statement interface{}) {
	APIServerLogger.WithField("component", "api-server").WithFields(logrus.Fields{"command": "POST /db/create"}).Info(statement)
}

func InfoAPIDatabaseInsert(statement interface{}) {
	APIServerLogger.WithField("component", "api-server").WithFields(logrus.Fields{"command": "POST /db/insert"}).Info(statement)
}

func InfoAPIDatabaseSelect(statement interface{}) {
	APIServerLogger.WithField("component", "api-server").WithFields(logrus.Fields{"command": "POST /db/select"}).Info(statement)
}

func InfoAPIDatabaseUpdate(statement interface{}) {
	APIServerLogger.WithField("component", "api-server").WithFields(logrus.Fields{"command": "POST /db/update"}).Info(statement)
}

func InfoAPIDatabaseDelete(statement interface{}) {
	APIServerLogger.WithField("component", "api-server").WithFields(logrus.Fields{"command": "POST /db/delete"}).Info(statement)
}

func InfoAPIAuthorityGrant(statement interface{}) {
	APIServerLogger.WithField("component", "api-server").WithFields(logrus.Fields{"command": "POST /auth/grant"}).Info(statement)
}

func InfoAPIAuthorityRevoke(statement interface{}) {
	APIServerLogger.WithField("component", "api-server").WithFields(logrus.Fields{"command": "POST /auth/revoke"}).Info(statement)
}

func InfoAPIAuthorityVerify(statement interface{}) {
	APIServerLogger.WithField("component", "api-server").WithFields(logrus.Fields{"command": "POST /auth/verify"}).Info(statement)
}
