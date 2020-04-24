package logger

import "github.com/sirupsen/logrus"

func InfoAPIDatabaseCreate(statement interface{}) {
	APIServerLogger.WithField("component", "api-server").WithFields(logrus.Fields{"command": "POST /db/create"}).Info(statement)
}

func WarnAPIDatabaseCreate(err error) {
	APIServerLogger.WithField("component", "api-server").WithFields(logrus.Fields{"command": "POST /db/create"}).Warn(err.Error())
}

func InfoAPIDatabaseInsert(statement interface{}) {
	APIServerLogger.WithField("component", "api-server").WithFields(logrus.Fields{"command": "POST /db/insert"}).Info(statement)
}

func WarnAPIDatabaseInsert(err error) {
	APIServerLogger.WithField("component", "api-server").WithFields(logrus.Fields{"command": "POST /db/insert"}).Warn(err.Error())
}

func InfoAPIDatabaseSelect(statement interface{}) {
	APIServerLogger.WithField("component", "api-server").WithFields(logrus.Fields{"command": "POST /db/select"}).Info(statement)
}

func WarnAPIDatabaseSelect(err error) {
	APIServerLogger.WithField("component", "api-server").WithFields(logrus.Fields{"command": "POST /db/select"}).Warn(err.Error())
}

func InfoAPIDatabaseUpdate(statement interface{}) {
	APIServerLogger.WithField("component", "api-server").WithFields(logrus.Fields{"command": "POST /db/update"}).Info(statement)
}

func WarnAPIDatabaseUpdate(err error) {
	APIServerLogger.WithField("component", "api-server").WithFields(logrus.Fields{"command": "POST /db/update"}).Warn(err.Error())
}

func InfoAPIDatabaseDelete(statement interface{}) {
	APIServerLogger.WithField("component", "api-server").WithFields(logrus.Fields{"command": "POST /db/delete"}).Info(statement)
}

func WarnAPIDatabaseDelete(err error) {
	APIServerLogger.WithField("component", "api-server").WithFields(logrus.Fields{"command": "POST /db/delete"}).Warn(err.Error())
}

func InfoAPIAuthorityGrant(statement interface{}) {
	APIServerLogger.WithField("component", "api-server").WithFields(logrus.Fields{"command": "POST /auth/grant"}).Info(statement)
}

func WarnAPIAuthorityGrant(err error) {
	APIServerLogger.WithField("component", "api-server").WithFields(logrus.Fields{"command": "POST /auth/grant"}).Warn(err.Error())
}

func InfoAPIAuthorityRevoke(statement interface{}) {
	APIServerLogger.WithField("component", "api-server").WithFields(logrus.Fields{"command": "POST /auth/revoke"}).Info(statement)
}

func WarnAPIAuthorityRevoke(err error) {
	APIServerLogger.WithField("component", "api-server").WithFields(logrus.Fields{"command": "POST /auth/revoke"}).Warn(err.Error())
}

func InfoAPIAuthorityVerify(statement interface{}) {
	APIServerLogger.WithField("component", "api-server").WithFields(logrus.Fields{"command": "POST /auth/verify"}).Info(statement)
}

func WarnAPIAuthorityVerify(err error) {
	APIServerLogger.WithField("component", "api-server").WithFields(logrus.Fields{"command": "POST /auth/verify"}).Warn(err.Error())
}
