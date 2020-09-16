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
	dropkit "github.com/DropKit/DropKit-Adapter/package/crypto/contract"
)

func CreateDropKitInstance() (*dropkit.DropkitContract, error) {
	quorumEndpoint := viper.GetString(`QUORUM.ENDPOINT`)
	dropKitContractAddr := viper.GetString(`DROPKIT.CONTRACT`)

	quorumClient, err := ethclient.Dial(quorumEndpoint)
	if err != nil {
		logger.InternalLogger.WithField("component", "instance-creator").Error(err.Error())
		return nil, err
	}

	contractAddress := common.HexToAddress(dropKitContractAddr)

	contractInstance, err := dropkit.NewDropkitContract(contractAddress, quorumClient)
	if err != nil {
		logger.InternalLogger.WithField("component", "instance-creator").Error(err.Error())
		return nil, err
	}

	return contractInstance, nil
}

func CreateAuthInstance(privatekeyHex string) (*bind.TransactOpts, error) {
	quorumEndpoint := viper.GetString(`QUORUM.ENDPOINT`)

	quorumClient, err := ethclient.Dial(quorumEndpoint)
	if err != nil {
		logger.InternalLogger.WithField("component", "instance-creator").Error(err.Error())
		return nil, err
	}
	privateKey, err := crypto.HexToECDSA(privatekeyHex)
	if err != nil {
		logger.InternalLogger.WithField("component", "instance-creator").Error(err.Error())
		return nil, err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		logger.InternalLogger.WithField("component", "instance-creator").Error(err.Error())
		return nil, err
	}

	accountAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	nonce, err := quorumClient.PendingNonceAt(context.Background(), accountAddress)
	if err != nil {
		logger.InternalLogger.WithField("component", "instance-creator").Error(err.Error())
		return nil, err
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(1000000)
	auth.GasPrice = big.NewInt(0)

	return auth, nil
}
