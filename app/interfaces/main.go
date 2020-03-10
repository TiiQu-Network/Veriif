package interfaces

import (
	"github.com/TiiQu-Network/Veriif/app/errors"
	"github.com/TiiQu-Network/Veriif/app/ethereum"
)

type Application interface {
	Init()
	Run()
	ErrLogger() errors.Logger
	EthCaller() ethereum.Caller
}
