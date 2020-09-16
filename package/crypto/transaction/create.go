package transaction

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"math/big"

	"github.com/DropKit/DropKit-Adapter/logger"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/viper"
)

func CreateRawTransaction(receiverAddress string, txMessage string, txValue int64, privatekeyHex string) string {
	quorumEndpoint := viper.GetString(`QUORUM.ENDPOINT`)

	client, err := ethclient.Dial(quorumEndpoint)
	if err != nil {
		logger.InternalLogger.WithField("component", "raw-transaction-creator").Error(err.Error())
	}

	privateKey, err := crypto.HexToECDSA(privatekeyHex)
	if err != nil {
		logger.InternalLogger.WithField("component", "raw-transaction-creator").Error(err.Error())
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		logger.InternalLogger.WithField("component", "raw-transaction-creator").Error(err.Error())
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		logger.InternalLogger.WithField("component", "raw-transaction-creator").Error(err.Error())
	}

	value := big.NewInt(txValue)
	gasLimit := uint64(5000000)
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		logger.InternalLogger.WithField("component", "raw-transaction-creator").Error(err.Error())
	}

	toAddress := common.HexToAddress(receiverAddress)
	data := []byte(txMessage)
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		logger.InternalLogger.WithField("component", "raw-transaction-creator").Error(err.Error())
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		logger.InternalLogger.WithField("component", "raw-transaction-creator").Error(err.Error())
	}

	ts := types.Transactions{signedTx}
	rawTxBytes := ts.GetRlp(0)
	rawTxHex := hex.EncodeToString(rawTxBytes)

	return rawTxHex
}
