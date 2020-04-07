package transaction

import (
	"context"
	"encoding/hex"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/spf13/viper"
)

func SendRawTransaction(receiverAddress string, txMessage string, txValue int64, privatekeyHex string) string {
	quorumEndpoint := viper.GetString(`QUORUM.ENDPOINT`)

	client, err := ethclient.Dial(quorumEndpoint)
	if err != nil {
		print(err)
	}

	rawTx := CreateRawTransaction(receiverAddress, txMessage, txValue, privatekeyHex)

	rawTxBytes, err := hex.DecodeString(rawTx)

	tx := new(types.Transaction)
	rlp.DecodeBytes(rawTxBytes, &tx)

	err = client.SendTransaction(context.Background(), tx)
	if err != nil {
		print(err)
	}

	return tx.Hash().Hex()
}
