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

	"github.com/DropKit/DropKit-Adapter/logger"
	"github.com/DropKit/DropKit-Adapter/package/crypto/contracts/authority"
)

func GrantAuthority(authorityAddr string, privatekeyHex string, tableName string, addUserAddress string) {
	quorumEndpoint := viper.GetString(`QUORUM.ENDPOINT`)

	quorumClient, err := ethclient.Dial(quorumEndpoint)
	if err != nil {
		logger.InternalLogger.WithField("component", "internal").Error(err.Error())
	}

	privateKey, err := crypto.HexToECDSA(privatekeyHex)
	if err != nil {
		logger.InternalLogger.WithField("component", "internal").Error(err.Error())
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		logger.InternalLogger.WithField("component", "internal").Error(err.Error())
	}

	accountAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	nonce, err := quorumClient.PendingNonceAt(context.Background(), accountAddress)
	if err != nil {
		logger.InternalLogger.WithField("component", "internal").Error(err.Error())
	}

	gasPrice, err := quorumClient.SuggestGasPrice(context.Background())
	if err != nil {
		logger.InternalLogger.WithField("component", "internal").Error(err.Error())
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(300000)
	auth.GasPrice = gasPrice

	address := common.HexToAddress(authorityAddr)

	contractInstance, err := authority.NewAuthority(address, quorumClient)
	if err != nil {
		logger.InternalLogger.WithField("component", "internal").Error(err.Error())
	}

	_, err = contractInstance.Add(auth, tableName, addUserAddress)
	if err != nil {
		logger.InternalLogger.WithField("component", "internal").Error(err.Error())
	}
}

func RevokeAuthority(authorityAddr string, privatekeyHex string, tableName string, removeUserAddress string) {
	quorumEndpoint := viper.GetString(`QUORUM.ENDPOINT`)

	quorumClient, err := ethclient.Dial(quorumEndpoint)
	if err != nil {
		logger.InternalLogger.WithField("component", "internal").Error(err.Error())
	}

	privateKey, err := crypto.HexToECDSA(privatekeyHex)
	if err != nil {
		logger.InternalLogger.WithField("component", "internal").Error(err.Error())
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		logger.InternalLogger.WithField("component", "internal").Error(err.Error())
	}

	accountAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	nonce, err := quorumClient.PendingNonceAt(context.Background(), accountAddress)
	if err != nil {
		logger.InternalLogger.WithField("component", "internal").Error(err.Error())
	}

	gasPrice, err := quorumClient.SuggestGasPrice(context.Background())
	if err != nil {
		logger.InternalLogger.WithField("component", "internal").Error(err.Error())
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(300000)
	auth.GasPrice = gasPrice

	address := common.HexToAddress(authorityAddr)

	contractInstance, err := authority.NewAuthority(address, quorumClient)
	if err != nil {
		logger.InternalLogger.WithField("component", "internal").Error(err.Error())
	}

	_, err = contractInstance.Remove(auth, tableName, removeUserAddress)
	if err != nil {
		logger.InternalLogger.WithField("component", "internal").Error(err.Error())
	}
}

func VerifyAuthority(authorityAddr string, privatekeyHex string, tableName string, checkUserAddress string) bool {
	quorumEndpoint := viper.GetString(`QUORUM.ENDPOINT`)

	quorumClient, err := ethclient.Dial(quorumEndpoint)
	if err != nil {
		logger.InternalLogger.WithField("component", "internal").Error(err.Error())
	}

	address := common.HexToAddress(authorityAddr)

	contractInstance, err := authority.NewAuthority(address, quorumClient)
	if err != nil {
		logger.InternalLogger.WithField("component", "internal").Error(err.Error())
	}

	authority, err := contractInstance.Has(nil, tableName, checkUserAddress)
	if err != nil {
		logger.InternalLogger.WithField("component", "internal").Error(err.Error())
	}

	return authority
}
