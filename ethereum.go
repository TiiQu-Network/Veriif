package main

import (
	"context"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"

	"github.com/TiiQu-Network/Veriif/contracts"
)

const (
	contractAddress = "0x6De37C9538f9EB6129D22C4f9372B8be2b98b96E"
	node            = "https://kovan.infura.io/v3/87abed49b20c49bd99f9bc2645023f34"
)

var (
	certRegContract *contracts.CertRegistryCaller
	callOpts        *bind.CallOpts
)

func initContract() error {
	ca := common.HexToAddress(contractAddress)
	client, err := ethclient.Dial(node)
	if err != nil {
		return errors.Wrap(err, "Failed to connect to the Ethereum client: %v")
	}

	certRegContract, err = contracts.NewCertRegistryCaller(ca, client)
	if err != nil {
		return err
	}

	callOpts = &bind.CallOpts{
		Pending: false,
		Context: context.TODO(),
	}

	return nil
}
