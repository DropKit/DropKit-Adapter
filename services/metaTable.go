package services

import (
	"context"
	"crypto/ecdsa"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/viper"

	"github.com/DropKit/DropKit-Adapter/package/crypto/contracts/metaTable"
)

func AddMetaTable(tableName string, tableAddress string, tableAddressStorageContractAddress string, privatekeyHex string) {
	quorumEndpoint := viper.GetString(`QUORUM.ENDPOINT`)

	quorumClient, err := ethclient.Dial(quorumEndpoint)
	if err != nil {
		print(err)
	}

	privateKey, err := crypto.HexToECDSA(privatekeyHex)
	if err != nil {
		print(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		print(err)
	}

	accountAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	nonce, err := quorumClient.PendingNonceAt(context.Background(), accountAddress)
	if err != nil {
		print(err)
	}

	gasPrice, err := quorumClient.SuggestGasPrice(context.Background())
	if err != nil {
		print(err)
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(300000)
	auth.GasPrice = gasPrice

	address := common.HexToAddress(tableAddressStorageContractAddress)

	contractInstance, err := metaTable.NewMetaTable(address, quorumClient)
	if err != nil {
		print(err)
	}

	_, err = contractInstance.Add(auth, tableName, tableAddress)
	if err != nil {
		print(err)
	}
}

func GetMetaTable(tableName string, tableAddressStorageContractAddress string) string {
	quorumEndpoint := viper.GetString(`QUORUM.ENDPOINT`)

	quorumClient, err := ethclient.Dial(quorumEndpoint)
	if err != nil {
		print(err)
	}

	address := common.HexToAddress(tableAddressStorageContractAddress)

	contractInstance, err := metaTable.NewMetaTable(address, quorumClient)
	if err != nil {
		print(err)
	}

	tableAddress, err := contractInstance.Get(nil, tableName)
	if err != nil {
		print(err)
	}

	return tableAddress
}
