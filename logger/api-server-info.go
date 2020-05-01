package logger

import "github.com/sirupsen/logrus"

// server
func InfoAPIPing() {
	APIServerLogger.WithField("component", "api-server").WithFields(logrus.Fields{"command": "GET /health/ping"}).Info()
}

func InfoAPIDenpendencyOk() {
	APIServerLogger.WithField("component", "api-server").WithFields(logrus.Fields{"command": "GET /health/dependency"}).Info()
}

// db

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

// permission

func InfoAPIPermissionGrantAdmin(statement interface{}) {
	APIServerLogger.WithField("component", "api-server").WithFields(logrus.Fields{"command": "POST /permission/grant/admin"}).Info(statement)
}

func InfoAPIPermissionGrantMaintainer(statement interface{}) {
	APIServerLogger.WithField("component", "api-server").WithFields(logrus.Fields{"command": "POST /permission/grant/maintainer"}).Info(statement)
}

func InfoAPIPermissionGrantUser(statement interface{}) {
	APIServerLogger.WithField("component", "api-server").WithFields(logrus.Fields{"command": "POST /permission/grant/user"}).Info(statement)
}

func InfoAPIPermissionRevokeAdmin(statement interface{}) {
	APIServerLogger.WithField("component", "api-server").WithFields(logrus.Fields{"command": "POST /permission/revoke/admin"}).Info(statement)
}

func InfoAPIPermissionRevokeMaintainer(statement interface{}) {
	APIServerLogger.WithField("component", "api-server").WithFields(logrus.Fields{"command": "POST /permission/revoke/maintainer"}).Info(statement)
}

func InfoAPIPermissionRevokeUser(statement interface{}) {
	APIServerLogger.WithField("component", "api-server").WithFields(logrus.Fields{"command": "POST /permission/revoke/user"}).Info(statement)
}

func InfoAPIPermissionVerifyAdmin(statement interface{}) {
	APIServerLogger.WithField("component", "api-server").WithFields(logrus.Fields{"command": "POST /permission/verify/admin"}).Info(statement)
}

func InfoAPIPermissionVerifyMaintainer(statement interface{}) {
	APIServerLogger.WithField("component", "api-server").WithFields(logrus.Fields{"command": "POST /permission/verify/maintainer"}).Info(statement)
}

func InfoAPIPermissionVerifyUser(statement interface{}) {
	APIServerLogger.WithField("component", "api-server").WithFields(logrus.Fields{"command": "POST /permission/verify/user"}).Info(statement)
}

// user

func InfoAPIUserCreate() {
	APIServerLogger.WithField("component", "api-server").WithFields(logrus.Fields{"command": "GET /user/create"}).Info()
}

// payment

func InfoAPIPaymentMint(statement interface{}) {
	APIServerLogger.WithField("component", "api-server").WithFields(logrus.Fields{"command": "POST /payment/mint"}).Info(statement)
}

func InfoAPIPaymentBurn(statement interface{}) {
	APIServerLogger.WithField("component", "api-server").WithFields(logrus.Fields{"command": "POST /payment/burn"}).Info(statement)
}

func InfoAPIPaymentTransfer(statement interface{}) {
	APIServerLogger.WithField("component", "api-server").WithFields(logrus.Fields{"command": "POST /payment/transfer"}).Info(statement)
}

func InfoAPIPaymentBalance(account string) {
	APIServerLogger.WithField("component", "api-server").WithFields(logrus.Fields{"command": "POST /payment/balance"}).Info(account)
}
