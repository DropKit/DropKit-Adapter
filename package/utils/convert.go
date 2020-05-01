package utils

func ToRoleName(tableName string) (string, string, string) {
	adminRoleName := tableName + "AdminRole"
	maintainerRoleName := tableName + "MaintainerRole"
	userRoleName := tableName + "UserRole"

	return adminRoleName, maintainerRoleName, userRoleName
}
