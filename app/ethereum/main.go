package ethereum

import (
	"context"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"

	"github.com/TiiQu-Network/Veriif/app/ethereum/contracts"
)

const (
	contractAddress = "0x6De37C9538f9EB6129D22C4f9372B8be2b98b96E"
	node            = "https://kovan.infura.io/v3/87abed49b20c49bd99f9bc2645023f34"
)

type Caller struct {
	CertRegContract *contracts.CertRegistryCaller
	CallOpts        *bind.CallOpts
}

func Init() (Caller, error) {
	c := Caller{}

	ca := common.HexToAddress(contractAddress)
	client, err := ethclient.Dial(node)
	if err != nil {
		return c, errors.Wrap(err, "Failed to connect to the Ethereum client: %v")
	}

	c.CertRegContract, err = contracts.NewCertRegistryCaller(ca, client)
	if err != nil {
		return c, err
	}

	c.CallOpts = &bind.CallOpts{
		Pending: false,
		Context: context.TODO(),
	}

	return c, nil
}
