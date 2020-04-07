package transaction

import (
	"context"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/viper"
)

func beenConfirmed(transactionHash string) bool {
	quorumEndpoint := viper.GetString(`QUORUM.ENDPOINT`)

	client, err := ethclient.Dial(quorumEndpoint)
	if err != nil {
		print(err)
	}

	txHash := common.HexToHash(transactionHash)
	_, isPending, err := client.TransactionByHash(context.Background(), txHash)
	if err != nil {
		print(err)
	}

	if isPending == true {
		return false
	} else {
		return true
	}
}

func CheckTransactionConfirmed(transactionHash string) {
	isTransactionConfirmed := beenConfirmed(transactionHash)

	for isTransactionConfirmed == false {
		time.Sleep(1 * time.Second)
		isTransactionConfirmed = beenConfirmed(transactionHash)
	}
}
