package interfaces

import (
	"github.com/TiiQu-Network/Veriif/errors"
	"github.com/TiiQu-Network/Veriif/ethereum"
)

type Application interface {
	Init()
	Run()
	ErrLogger() errors.Logger
	EthCaller() ethereum.Caller
}
