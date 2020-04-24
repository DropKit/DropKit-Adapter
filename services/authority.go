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

func GrantAuthority(authorityAddr string, privatekeyHex string, tableName string, addUserAddress string) error {
	quorumEndpoint := viper.GetString(`QUORUM.ENDPOINT`)

	quorumClient, err := ethclient.Dial(quorumEndpoint)
	if err != nil {
		logger.InternalLogger.WithField("component", "internal").Error(err.Error())
		return err
	}

	privateKey, err := crypto.HexToECDSA(privatekeyHex)
	if err != nil {
		logger.InternalLogger.WithField("component", "internal").Error(err.Error())
		return err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		logger.InternalLogger.WithField("component", "internal").Error(err.Error())
		return err
	}

	accountAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	nonce, err := quorumClient.PendingNonceAt(context.Background(), accountAddress)
	if err != nil {
		logger.InternalLogger.WithField("component", "internal").Error(err.Error())
		return err
	}

	gasPrice, err := quorumClient.SuggestGasPrice(context.Background())
	if err != nil {
		logger.InternalLogger.WithField("component", "internal").Error(err.Error())
		return err
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
		return err
	}

	_, err = contractInstance.Add(auth, tableName, addUserAddress)
	if err != nil {
		logger.InternalLogger.WithField("component", "internal").Error(err.Error())
		return err
	}

	return nil
}

func RevokeAuthority(authorityAddr string, privatekeyHex string, tableName string, removeUserAddress string) error {
	quorumEndpoint := viper.GetString(`QUORUM.ENDPOINT`)

	quorumClient, err := ethclient.Dial(quorumEndpoint)
	if err != nil {
		logger.InternalLogger.WithField("component", "internal").Error(err.Error())
		return err
	}

	privateKey, err := crypto.HexToECDSA(privatekeyHex)
	if err != nil {
		logger.InternalLogger.WithField("component", "internal").Error(err.Error())
		return err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		logger.InternalLogger.WithField("component", "internal").Error(err.Error())
		return err
	}

	accountAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	nonce, err := quorumClient.PendingNonceAt(context.Background(), accountAddress)
	if err != nil {
		logger.InternalLogger.WithField("component", "internal").Error(err.Error())
		return err
	}

	gasPrice, err := quorumClient.SuggestGasPrice(context.Background())
	if err != nil {
		logger.InternalLogger.WithField("component", "internal").Error(err.Error())
		return err
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
		return err
	}

	_, err = contractInstance.Remove(auth, tableName, removeUserAddress)
	if err != nil {
		logger.InternalLogger.WithField("component", "internal").Error(err.Error())
		return err
	}

	return nil
}

func VerifyAuthority(authorityAddr string, privatekeyHex string, tableName string, checkUserAddress string) (bool, error) {
	quorumEndpoint := viper.GetString(`QUORUM.ENDPOINT`)

	quorumClient, err := ethclient.Dial(quorumEndpoint)
	if err != nil {
		logger.InternalLogger.WithField("component", "internal").Error(err.Error())
		return false, err
	}

	address := common.HexToAddress(authorityAddr)

	contractInstance, err := authority.NewAuthority(address, quorumClient)
	if err != nil {
		logger.InternalLogger.WithField("component", "internal").Error(err.Error())
		return false, err
	}

	authority, err := contractInstance.Has(nil, tableName, checkUserAddress)
	if err != nil {
		logger.InternalLogger.WithField("component", "internal").Error(err.Error())
		return false, err
	}

	return authority, nil
}
