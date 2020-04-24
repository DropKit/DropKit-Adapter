package services

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/DropKit/DropKit-Adapter/logger"
	"github.com/ethereum/go-ethereum/ethclient"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func DependencyServicesCheck() {
	quorumEndpoint := viper.GetString(`QUORUM.ENDPOINT`)
	yugabyteHost := viper.GetString(`YUGABYTEDB.HOST`)
	yugabytePort := viper.GetInt(`YUGABYTEDB.PORT`)
	yugabyteUser := viper.GetString(`YUGABYTEDB.USER`)
	yugabytePassword := viper.GetString(`YUGABYTEDB.PASSWORD`)
	yugabyteDBName := viper.GetString(`YUGABYTEDB.DBNAME`)

	client, _ := ethclient.Dial(quorumEndpoint)

	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		logger.FatelDependencyService("Quorum", err)
	} else {
		logger.InfoDependencyService("Quorum")
	}

	_ = header

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", yugabyteHost, yugabytePort, yugabyteUser, yugabytePassword, yugabyteDBName)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		logger.InternalLogger.WithField("component", "internal").Error(err.Error())
	}
	if err = db.Ping(); err != nil {
		logger.FatelDependencyService("YugabyteDB", err)
		db.Close()
	} else {
		logger.InfoDependencyService("YugabyteDB")
	}
}
