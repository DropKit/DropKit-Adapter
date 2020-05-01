package services

import (
	"database/sql"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func DependencyServicesCheck() (string, error) {
	yugabyteHost := viper.GetString(`YUGABYTEDB.HOST`)
	yugabytePort := viper.GetInt(`YUGABYTEDB.PORT`)
	yugabyteUser := viper.GetString(`YUGABYTEDB.USER`)
	yugabytePassword := viper.GetString(`YUGABYTEDB.PASSWORD`)
	yugabyteDBName := viper.GetString(`YUGABYTEDB.DBNAME`)

	contractInstance, err := CreateDropKitInstance()
	if err != nil {
		return "Quorum", err
	}

	_, err = contractInstance.BalanceOf(nil, common.HexToAddress(viper.GetString(`DROPKIT.ACCOUNT`)))
	if err != nil {
		return "Quorum", err
	}

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", yugabyteHost, yugabytePort, yugabyteUser, yugabytePassword, yugabyteDBName)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return "YugabyteDB", err
	}
	if err = db.Ping(); err != nil {
		db.Close()
		return "YugabyteDB", err
	}

	return "", nil
}
