package logger

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

// server

func ErrorAPIDenpendencyError() {
	APIServerLogger.WithField("component", "api-server").WithFields(logrus.Fields{"command": "GET /health/dependency"}).Error()
}

// db

func WarnAPIDatabaseCreate(err error) {
	APIServerLogger.WithField("component", "api-server").WithFields(logrus.Fields{"command": "POST /db/create"}).Warn(err.Error())
}

func WarnAPIDatabaseInsert(err error) {
	APIServerLogger.WithField("component", "api-server").WithFields(logrus.Fields{"command": "POST /db/insert"}).Warn(err.Error())
}

func WarnAPIDatabaseSelect(err error) {
	APIServerLogger.WithField("component", "api-server").WithFields(logrus.Fields{"command": "POST /db/select"}).Warn(err.Error())
}

func WarnAPIDatabaseUpdate(err error) {
	APIServerLogger.WithField("component", "api-server").WithFields(logrus.Fields{"command": "POST /db/update"}).Warn(err.Error())
}

func WarnAPIDatabaseDelete(err error) {
	APIServerLogger.WithField("component", "api-server").WithFields(logrus.Fields{"command": "POST /db/delete"}).Warn(err.Error())
}

func WarnAPIDatabaseCreateUnAuth(account string) {
	APIServerLogger.WithField("component", "api-server").WithFields(logrus.Fields{"command": "POST /db/create"}).Warn("unauthorized user: " + account)
}

func WarnAPIDatabaseInsertUnAuth(account string) {
	APIServerLogger.WithField("component", "api-server").WithFields(logrus.Fields{"command": "POST /db/insert"}).Warn("unauthorized user: " + account)
}

func WarnAPIDatabaseSelectUnAuth(account string) {
	APIServerLogger.WithField("component", "api-server").WithFields(logrus.Fields{"command": "POST /db/select"}).Warn("unauthorized user: " + account)
}

func WarnAPIDatabaseUpdateUnAuth(account string) {
	APIServerLogger.WithField("component", "api-server").WithFields(logrus.Fields{"command": "POST /db/update"}).Warn("unauthorized user: " + account)
}

func WarnAPIDatabaseDeleteUnAuth(account string) {
	APIServerLogger.WithField("component", "api-server").WithFields(logrus.Fields{"command": "POST /db/delete"}).Warn("unauthorized user: " + account)
}

// permission

func WarnAPIPermissionGrantAdmin(err error) {
	APIServerLogger.WithField("component", "api-server").WithFields(logrus.Fields{"command": "POST /permission/grant/admin"}).Warn(err.Error())
}

func WarnAPIPermissionGrantMaintainer(err error) {
	APIServerLogger.WithField("component", "api-server").WithFields(logrus.Fields{"command": "POST /permission/grant/maintainer"}).Warn(err.Error())
}

func WarnAPIPermissionGrantUser(err error) {
	APIServerLogger.WithField("component", "api-server").WithFields(logrus.Fields{"command": "POST /permission/grant/user"}).Warn(err.Error())
}

func WarnAPIPermissionRevokeAdmin(err error) {
	APIServerLogger.WithField("component", "api-server").WithFields(logrus.Fields{"command": "POST /permission/revoke/admin"}).Warn(err.Error())
}

func WarnAPIPermissionRevokeMaintainer(err error) {
	APIServerLogger.WithField("component", "api-server").WithFields(logrus.Fields{"command": "POST /permission/revoke/maintainer"}).Warn(err.Error())
}

func WarnAPIPermissionRevokeUser(err error) {
	APIServerLogger.WithField("component", "api-server").WithFields(logrus.Fields{"command": "POST /permission/revoke/user"}).Warn(err.Error())
}

func WarnAPIPermissionVerifyAdmin(err error) {
	APIServerLogger.WithField("component", "api-server").WithFields(logrus.Fields{"command": "POST /permission/verify/admin"}).Warn(err.Error())
}

func WarnAPIPermissionVerifyMaintainer(err error) {
	APIServerLogger.WithField("component", "api-server").WithFields(logrus.Fields{"command": "POST /permission/verify/maintainer"}).Warn(err.Error())
}

func WarnAPIPermissionVerifyUser(err error) {
	APIServerLogger.WithField("component", "api-server").WithFields(logrus.Fields{"command": "POST /permission/verify/user"}).Warn(err.Error())
}

func WarnAPIPermissionGrantAdminUnAuth(account string) {
	APIServerLogger.WithField("component", "api-server").WithFields(logrus.Fields{"command": "POST /permission/grant/admin"}).Warn("unauthorized user: " + account)
}

func WarnAPIPermissionGrantMaintainerUnAuth(account string) {
	APIServerLogger.WithField("component", "api-server").WithFields(logrus.Fields{"command": "POST /permission/grant/maintainer"}).Warn("unauthorized user: " + account)
}

func WarnAPIPermissionGrantUserUnAuth(account string) {
	APIServerLogger.WithField("component", "api-server").WithFields(logrus.Fields{"command": "POST /permission/grant/user"}).Warn("unauthorized user: " + account)
}

func WarnAPIPermissionRevokeAdminUnAuth(account string) {
	APIServerLogger.WithField("component", "api-server").WithFields(logrus.Fields{"command": "POST /permission/revoke/admin"}).Warn("unauthorized user: " + account)
}

func WarnAPIPermissionRevokeMaintainerUnAuth(account string) {
	APIServerLogger.WithField("component", "api-server").WithFields(logrus.Fields{"command": "POST /permission/revoke/maintainer"}).Warn("unauthorized user: " + account)
}

func WarnAPIPermissionRevokeUserUnAuth(account string) {
	APIServerLogger.WithField("component", "api-server").WithFields(logrus.Fields{"command": "POST /permission/revoke/user"}).Warn("unauthorized user: " + account)
}

func WarnAPIPermissionVerifyAdminUnAuth(account string) {
	APIServerLogger.WithField("component", "api-server").WithFields(logrus.Fields{"command": "POST /permission/verify/admin"}).Warn("unauthorized user: " + account)
}

func WarnAPIPermissionVerifyMaintainerUnAuth(account string) {
	APIServerLogger.WithField("component", "api-server").WithFields(logrus.Fields{"command": "POST /permission/verify/maintainer"}).Warn("unauthorized user: " + account)
}

func WarnAPIPermissionVerifyUserUnAuth(account string) {
	APIServerLogger.WithField("component", "api-server").WithFields(logrus.Fields{"command": "POST /permission/verify/user"}).Warn("unauthorized user: " + account)
}

// payment

func WarnAPIPaymentMint(statement interface{}) {
	APIServerLogger.WithField("component", "api-server").WithFields(logrus.Fields{"command": "POST /payment/mint"}).Warn(statement)
}

func WarnAPIPaymentBurn(statement interface{}) {
	APIServerLogger.WithField("component", "api-server").WithFields(logrus.Fields{"command": "POST /payment/burn"}).Warn(statement)
}

func WarnAPIPaymentMintUnAuth(account string) {
	APIServerLogger.WithField("component", "api-server").WithFields(logrus.Fields{"command": "POST /payment/mint"}).Warn("unauthorized user: " + account)
}

func WarnAPIPaymentBurnUnAuth(account string) {
	APIServerLogger.WithField("component", "api-server").WithFields(logrus.Fields{"command": "POST /payment/burn"}).Warn("unauthorized user: " + account)
}

func WarnAPIPaymentTransfer(statement interface{}) {
	APIServerLogger.WithField("component", "api-server").WithFields(logrus.Fields{"command": "POST /payment/transfer"}).Warn(statement)
}

func WarnAPIPaymentTransferNotEnough(statement interface{}) {
	APIServerLogger.WithField("component", "api-server").WithFields(logrus.Fields{"command": "POST /payment/transfer"}).Warn("not enough: " + fmt.Sprintf("%v", statement))
}

func WarnAPIPaymentBalance(statement interface{}) {
	APIServerLogger.WithField("component", "api-server").WithFields(logrus.Fields{"command": "POST /payment/balance"}).Warn(statement)
}

func WarnAPIPaymentBalanceUnAuth(account string) {
	APIServerLogger.WithField("component", "api-server").WithFields(logrus.Fields{"command": "POST /payment/balance"}).Warn("unauthorized user: " + account)
}

// role

func WarnAPIRoleCreate(statement interface{}) {
	APIServerLogger.WithField("component", "api-server").WithFields(logrus.Fields{"command": "POST /role/create"}).Warn(statement)
}

func WarnAPIRoleCreateUnAuth(account string) {
	APIServerLogger.WithField("component", "api-server").WithFields(logrus.Fields{"command": "POST /role/create"}).Warn("unauthorized user: " + account)
}
