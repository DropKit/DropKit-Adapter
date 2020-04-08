package services

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/ethereum/go-ethereum/ethclient"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func DBAlive() {
	yugabyteHost := viper.GetString(`YUGABYTEDB.HOST`)
	yugabytePort := viper.GetInt(`YUGABYTEDB.PORT`)
	yugabyteUser := viper.GetString(`YUGABYTEDB.USER`)
	yugabytePassword := viper.GetString(`YUGABYTEDB.PASSWORD`)
	yugabyteDBName := viper.GetString(`YUGABYTEDB.DBNAME`)

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", yugabyteHost, yugabytePort, yugabyteUser, yugabytePassword, yugabyteDBName)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		print(err)
	}
	if err = db.Ping(); err != nil {
		print("Database connection failed.")
		db.Close()
	}
}

func QuorumAlive() {
	quorumEndpoint := viper.GetString(`QUORUM.ENDPOINT`)

	client, err := ethclient.Dial(quorumEndpoint)
	if err != nil {
		print("Quorum connection failed.")
	}

	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		print("Quorum connection failed.")
	}

	_ = header
}
