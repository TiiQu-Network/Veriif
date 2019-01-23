// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.
// Source https://github.com/TiiQu-Network/Cert-Registry

package contracts

import (
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// CertRegistryABI is the input ABI used to generate the binding from.
const CertRegistryABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"root\",\"type\":\"bytes32\"}],\"name\":\"rootExists\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addWhitelisted\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"removeWhitelisted\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isWhitelisted\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"root\",\"type\":\"bytes32\"}],\"name\":\"ownsRoot\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"isRevoked\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isPauser\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceWhitelistAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"currentOwner\",\"type\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"giveRootOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renouncePauser\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addWhitelistAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addPauser\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"removeRootOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceRootOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"getRoots\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"revoke\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isWhitelistAdmin\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceWhitelisted\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"root\",\"type\":\"bytes32\"}],\"name\":\"register\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"registerer\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"root\",\"type\":\"bytes32\"}],\"name\":\"CertRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"revoker\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"CertRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"currentOwner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"RootOwnershipGiven\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"account\",\"type\":\"address\"}],\"name\":\"WhitelistedAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"account\",\"type\":\"address\"}],\"name\":\"WhitelistedRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"account\",\"type\":\"address\"}],\"name\":\"WhitelistAdminAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"account\",\"type\":\"address\"}],\"name\":\"WhitelistAdminRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"account\",\"type\":\"address\"}],\"name\":\"PauserAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"account\",\"type\":\"address\"}],\"name\":\"PauserRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// CertRegistryBin is the compiled bytecode used for deploying new contracts.
const CertRegistryBin = `0x608060405234801561001057600080fd5b5060008054600160a060020a0319163317808255604051600160a060020a039190911691907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a361006c3364010000000061008d810204565b6002805460ff19169055610088336401000000006100dc810204565b6101b9565b6100a5600182640100000000610ce861012b82021704565b604051600160a060020a038216907f6719d08c1888103bea251a4ed56406bd0c3e69723c8a1686e017e7bbe159b6f890600090a250565b6100f4600382640100000000610ce861012b82021704565b604051600160a060020a038216907f22380c05984257a1cb900161c713dd71d39e74820f1aea43bd3f1bdd2096129990600090a250565b600160a060020a038116151561014057600080fd5b6101538282640100000000610182810204565b1561015d57600080fd5b600160a060020a0316600090815260209190915260409020805460ff19166001179055565b6000600160a060020a038216151561019957600080fd5b50600160a060020a03166000908152602091909152604090205460ff1690565b610e2d80620001c96000396000f3fe608060405234801561001057600080fd5b50600436106101ab576000357c0100000000000000000000000000000000000000000000000000000000900480637362d9c8116100fb5780639aa6b585116100b4578063bb5f747b1161008e578063bb5f747b1461045b578063d6cd947314610481578063e1fa8e8414610489578063f2fde38b146104a6576101ab565b80639aa6b585146103c0578063a82e3482146103c8578063b75c7dc61461043e576101ab565b80637362d9c81461031a57806382dc1ec4146103405780638456cb59146103665780638da5cb5b1461036e5780638f32d59b1461039257806394c2dac81461039a576101ab565b80634294857f1161016857806355457ab61161014257806355457ab6146102d45780635c975abb146103025780636ef8d66d1461030a578063715018a614610312576101ab565b80634294857f1461028957806346fbf68e146102a65780634c5a628c146102cc576101ab565b80630524b18f146101b057806310154bad146101e1578063291d9549146102095780633af32abf1461022f5780633f4ba83a1461025557806341d88e0c1461025d575b600080fd5b6101cd600480360360208110156101c657600080fd5b50356104cc565b604080519115158252519081900360200190f35b610207600480360360208110156101f757600080fd5b5035600160a060020a03166104e1565b005b6102076004803603602081101561021f57600080fd5b5035600160a060020a0316610501565b6101cd6004803603602081101561024557600080fd5b5035600160a060020a031661051e565b610207610537565b6101cd6004803603604081101561027357600080fd5b50600160a060020a03813516906020013561059b565b6101cd6004803603602081101561029f57600080fd5b50356105f3565b6101cd600480360360208110156102bc57600080fd5b5035600160a060020a0316610608565b61020761061b565b610207600480360360408110156102ea57600080fd5b50600160a060020a0381358116916020013516610626565b6101cd6106e8565b6102076106f2565b6102076106fb565b6102076004803603602081101561033057600080fd5b5035600160a060020a0316610765565b6102076004803603602081101561035657600080fd5b5035600160a060020a0316610782565b61020761079f565b610376610805565b60408051600160a060020a039092168252519081900360200190f35b6101cd610814565b610207600480360360208110156103b057600080fd5b5035600160a060020a0316610825565b610207610869565b6103ee600480360360208110156103de57600080fd5b5035600160a060020a0316610891565b60408051602080825283518183015283519192839290830191858101910280838360005b8381101561042a578181015183820152602001610412565b505050509050019250505060405180910390f35b6102076004803603602081101561045457600080fd5b50356108fd565b6101cd6004803603602081101561047157600080fd5b5035600160a060020a0316610997565b6102076109aa565b6102076004803603602081101561049f57600080fd5b50356109b3565b610207600480360360208110156104bc57600080fd5b5035600160a060020a0316610a68565b60009081526006602052604090205460ff1690565b6104ea33610997565b15156104f557600080fd5b6104fe81610a84565b50565b61050a33610997565b151561051557600080fd5b6104fe81610acc565b600061053160048363ffffffff610b1416565b92915050565b61054033610608565b151561054b57600080fd5b60025460ff16151561055c57600080fd5b6002805460ff191690556040805133815290517f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa9181900360200190a1565b600060606105a884610891565b905060005b81518110156105e8578382828151811015156105c557fe5b9060200190602002015114156105e057600192505050610531565b6001016105ad565b506000949350505050565b60009081526005602052604090205460ff1690565b600061053160018363ffffffff610b1416565b61062433610b4b565b565b6000610630610814565b15610639575060015b6106423361051e565b1561064b575060015b80151561065757600080fd5b60025460ff161561066757600080fd5b61067083610891565b600160a060020a0383166000908152600760209081526040909120825161069d9391929190910190610d82565b5060408051600160a060020a0380861682528416602082015281517f2854d8d9ea9d5530703d62c8b3e8bb2459a9136e7607511da717b2a13dcf70b8929181900390910190a1505050565b60025460ff165b90565b61062433610b93565b610703610814565b151561070e57600080fd5b60008054604051600160a060020a03909116907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a36000805473ffffffffffffffffffffffffffffffffffffffff19169055565b61076e33610997565b151561077957600080fd5b6104fe81610bdb565b61078b33610608565b151561079657600080fd5b6104fe81610c23565b6107a833610608565b15156107b357600080fd5b60025460ff16156107c357600080fd5b6002805460ff191660011790556040805133815290517f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2589181900360200190a1565b600054600160a060020a031690565b600054600160a060020a0316331490565b61082d610814565b151561083857600080fd5b60025460ff161561084857600080fd5b600160a060020a03811660009081526007602052604081206104fe91610dcd565b60025460ff161561087957600080fd5b33600090815260076020526040812061062491610dcd565b600160a060020a0381166000908152600760209081526040918290208054835181840281018401909452808452606093928301828280156108f157602002820191906000526020600020905b8154815260200190600101908083116108dd575b50505050509050919050565b6000610907610814565b15610910575060015b6109193361051e565b15610922575060015b80151561092e57600080fd5b60025460ff161561093e57600080fd5b600082815260056020908152604091829020805460ff19166001179055815133815290810184905281517f95805cdad68eccc1a7e73b9e1dd10d0561be298bdab9e5609e309df2e3c3d7b7929181900390910190a15050565b600061053160038363ffffffff610b1416565b61062433610acc565b6109bc3361051e565b15156109c757600080fd5b60025460ff16156109d757600080fd5b60008181526006602052604090205460ff16156109f357600080fd5b6000818152600660209081526040808320805460ff1916600190811790915533808552600784528285208054928301815585529383902001849055805192835290820183905280517f54553ef3b4e1d60c041c89079b49d51c62e39543028b4861c9e25529acaa70e89281900390910190a150565b610a70610814565b1515610a7b57600080fd5b6104fe81610c6b565b610a9560048263ffffffff610ce816565b604051600160a060020a038216907fee1504a83b6d4a361f4c1dc78ab59bfa30d6a3b6612c403e86bb01ef2984295f90600090a250565b610add60048263ffffffff610d3616565b604051600160a060020a038216907f270d9b30cf5b0793bbfd54c9d5b94aeb49462b8148399000265144a8722da6b690600090a250565b6000600160a060020a0382161515610b2b57600080fd5b50600160a060020a03166000908152602091909152604090205460ff1690565b610b5c60038263ffffffff610d3616565b604051600160a060020a038216907f0a8eb35e5ca14b3d6f28e4abf2f128dbab231a58b56e89beb5d636115001e16590600090a250565b610ba460018263ffffffff610d3616565b604051600160a060020a038216907fcd265ebaf09df2871cc7bd4133404a235ba12eff2041bb89d9c714a2621c7c7e90600090a250565b610bec60038263ffffffff610ce816565b604051600160a060020a038216907f22380c05984257a1cb900161c713dd71d39e74820f1aea43bd3f1bdd2096129990600090a250565b610c3460018263ffffffff610ce816565b604051600160a060020a038216907f6719d08c1888103bea251a4ed56406bd0c3e69723c8a1686e017e7bbe159b6f890600090a250565b600160a060020a0381161515610c8057600080fd5b60008054604051600160a060020a03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a36000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600160a060020a0381161515610cfd57600080fd5b610d078282610b14565b15610d1157600080fd5b600160a060020a0316600090815260209190915260409020805460ff19166001179055565b600160a060020a0381161515610d4b57600080fd5b610d558282610b14565b1515610d6057600080fd5b600160a060020a0316600090815260209190915260409020805460ff19169055565b828054828255906000526020600020908101928215610dbd579160200282015b82811115610dbd578251825591602001919060010190610da2565b50610dc9929150610de7565b5090565b50805460008255906000526020600020908101906104fe91905b6106ef91905b80821115610dc95760008155600101610ded56fea165627a7a7230582023651e3fdc8d4d6d0adefe0db4933c54b4a8d67f7e84d43249fa15eb8b733bba0029`

// DeployCertRegistry deploys a new Ethereum contract, binding an instance of CertRegistry to it.
func DeployCertRegistry(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *CertRegistry, error) {
	parsed, err := abi.JSON(strings.NewReader(CertRegistryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(CertRegistryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CertRegistry{CertRegistryCaller: CertRegistryCaller{contract: contract}, CertRegistryTransactor: CertRegistryTransactor{contract: contract}, CertRegistryFilterer: CertRegistryFilterer{contract: contract}}, nil
}

// CertRegistry is an auto generated Go binding around an Ethereum contract.
type CertRegistry struct {
	CertRegistryCaller     // Read-only binding to the contract
	CertRegistryTransactor // Write-only binding to the contract
	CertRegistryFilterer   // Log filterer for contract events
}

// CertRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type CertRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CertRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CertRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CertRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CertRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CertRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CertRegistrySession struct {
	Contract     *CertRegistry     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CertRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CertRegistryCallerSession struct {
	Contract *CertRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// CertRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CertRegistryTransactorSession struct {
	Contract     *CertRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// CertRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type CertRegistryRaw struct {
	Contract *CertRegistry // Generic contract binding to access the raw methods on
}

// CertRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CertRegistryCallerRaw struct {
	Contract *CertRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// CertRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CertRegistryTransactorRaw struct {
	Contract *CertRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCertRegistry creates a new instance of CertRegistry, bound to a specific deployed contract.
func NewCertRegistry(address common.Address, backend bind.ContractBackend) (*CertRegistry, error) {
	contract, err := bindCertRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CertRegistry{CertRegistryCaller: CertRegistryCaller{contract: contract}, CertRegistryTransactor: CertRegistryTransactor{contract: contract}, CertRegistryFilterer: CertRegistryFilterer{contract: contract}}, nil
}

// NewCertRegistryCaller creates a new read-only instance of CertRegistry, bound to a specific deployed contract.
func NewCertRegistryCaller(address common.Address, caller bind.ContractCaller) (*CertRegistryCaller, error) {
	contract, err := bindCertRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CertRegistryCaller{contract: contract}, nil
}

// NewCertRegistryTransactor creates a new write-only instance of CertRegistry, bound to a specific deployed contract.
func NewCertRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*CertRegistryTransactor, error) {
	contract, err := bindCertRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CertRegistryTransactor{contract: contract}, nil
}

// NewCertRegistryFilterer creates a new log filterer instance of CertRegistry, bound to a specific deployed contract.
func NewCertRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*CertRegistryFilterer, error) {
	contract, err := bindCertRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CertRegistryFilterer{contract: contract}, nil
}

// bindCertRegistry binds a generic wrapper to an already deployed contract.
func bindCertRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CertRegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CertRegistry *CertRegistryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _CertRegistry.Contract.CertRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CertRegistry *CertRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CertRegistry.Contract.CertRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CertRegistry *CertRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CertRegistry.Contract.CertRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CertRegistry *CertRegistryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _CertRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CertRegistry *CertRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CertRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CertRegistry *CertRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CertRegistry.Contract.contract.Transact(opts, method, params...)
}

// GetRoots is a free data retrieval call binding the contract method 0xa82e3482.
//
// Solidity: function getRoots(owner address) constant returns(bytes32[])
func (_CertRegistry *CertRegistryCaller) GetRoots(opts *bind.CallOpts, owner common.Address) ([][32]byte, error) {
	var (
		ret0 = new([][32]byte)
	)
	out := ret0
	err := _CertRegistry.contract.Call(opts, out, "getRoots", owner)
	return *ret0, err
}

// GetRoots is a free data retrieval call binding the contract method 0xa82e3482.
//
// Solidity: function getRoots(owner address) constant returns(bytes32[])
func (_CertRegistry *CertRegistrySession) GetRoots(owner common.Address) ([][32]byte, error) {
	return _CertRegistry.Contract.GetRoots(&_CertRegistry.CallOpts, owner)
}

// GetRoots is a free data retrieval call binding the contract method 0xa82e3482.
//
// Solidity: function getRoots(owner address) constant returns(bytes32[])
func (_CertRegistry *CertRegistryCallerSession) GetRoots(owner common.Address) ([][32]byte, error) {
	return _CertRegistry.Contract.GetRoots(&_CertRegistry.CallOpts, owner)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_CertRegistry *CertRegistryCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _CertRegistry.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_CertRegistry *CertRegistrySession) IsOwner() (bool, error) {
	return _CertRegistry.Contract.IsOwner(&_CertRegistry.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_CertRegistry *CertRegistryCallerSession) IsOwner() (bool, error) {
	return _CertRegistry.Contract.IsOwner(&_CertRegistry.CallOpts)
}

// IsPauser is a free data retrieval call binding the contract method 0x46fbf68e.
//
// Solidity: function isPauser(account address) constant returns(bool)
func (_CertRegistry *CertRegistryCaller) IsPauser(opts *bind.CallOpts, account common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _CertRegistry.contract.Call(opts, out, "isPauser", account)
	return *ret0, err
}

// IsPauser is a free data retrieval call binding the contract method 0x46fbf68e.
//
// Solidity: function isPauser(account address) constant returns(bool)
func (_CertRegistry *CertRegistrySession) IsPauser(account common.Address) (bool, error) {
	return _CertRegistry.Contract.IsPauser(&_CertRegistry.CallOpts, account)
}

// IsPauser is a free data retrieval call binding the contract method 0x46fbf68e.
//
// Solidity: function isPauser(account address) constant returns(bool)
func (_CertRegistry *CertRegistryCallerSession) IsPauser(account common.Address) (bool, error) {
	return _CertRegistry.Contract.IsPauser(&_CertRegistry.CallOpts, account)
}

// IsRevoked is a free data retrieval call binding the contract method 0x4294857f.
//
// Solidity: function isRevoked(hash bytes32) constant returns(bool)
func (_CertRegistry *CertRegistryCaller) IsRevoked(opts *bind.CallOpts, hash [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _CertRegistry.contract.Call(opts, out, "isRevoked", hash)
	return *ret0, err
}

// IsRevoked is a free data retrieval call binding the contract method 0x4294857f.
//
// Solidity: function isRevoked(hash bytes32) constant returns(bool)
func (_CertRegistry *CertRegistrySession) IsRevoked(hash [32]byte) (bool, error) {
	return _CertRegistry.Contract.IsRevoked(&_CertRegistry.CallOpts, hash)
}

// IsRevoked is a free data retrieval call binding the contract method 0x4294857f.
//
// Solidity: function isRevoked(hash bytes32) constant returns(bool)
func (_CertRegistry *CertRegistryCallerSession) IsRevoked(hash [32]byte) (bool, error) {
	return _CertRegistry.Contract.IsRevoked(&_CertRegistry.CallOpts, hash)
}

// IsWhitelistAdmin is a free data retrieval call binding the contract method 0xbb5f747b.
//
// Solidity: function isWhitelistAdmin(account address) constant returns(bool)
func (_CertRegistry *CertRegistryCaller) IsWhitelistAdmin(opts *bind.CallOpts, account common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _CertRegistry.contract.Call(opts, out, "isWhitelistAdmin", account)
	return *ret0, err
}

// IsWhitelistAdmin is a free data retrieval call binding the contract method 0xbb5f747b.
//
// Solidity: function isWhitelistAdmin(account address) constant returns(bool)
func (_CertRegistry *CertRegistrySession) IsWhitelistAdmin(account common.Address) (bool, error) {
	return _CertRegistry.Contract.IsWhitelistAdmin(&_CertRegistry.CallOpts, account)
}

// IsWhitelistAdmin is a free data retrieval call binding the contract method 0xbb5f747b.
//
// Solidity: function isWhitelistAdmin(account address) constant returns(bool)
func (_CertRegistry *CertRegistryCallerSession) IsWhitelistAdmin(account common.Address) (bool, error) {
	return _CertRegistry.Contract.IsWhitelistAdmin(&_CertRegistry.CallOpts, account)
}

// IsWhitelisted is a free data retrieval call binding the contract method 0x3af32abf.
//
// Solidity: function isWhitelisted(account address) constant returns(bool)
func (_CertRegistry *CertRegistryCaller) IsWhitelisted(opts *bind.CallOpts, account common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _CertRegistry.contract.Call(opts, out, "isWhitelisted", account)
	return *ret0, err
}

// IsWhitelisted is a free data retrieval call binding the contract method 0x3af32abf.
//
// Solidity: function isWhitelisted(account address) constant returns(bool)
func (_CertRegistry *CertRegistrySession) IsWhitelisted(account common.Address) (bool, error) {
	return _CertRegistry.Contract.IsWhitelisted(&_CertRegistry.CallOpts, account)
}

// IsWhitelisted is a free data retrieval call binding the contract method 0x3af32abf.
//
// Solidity: function isWhitelisted(account address) constant returns(bool)
func (_CertRegistry *CertRegistryCallerSession) IsWhitelisted(account common.Address) (bool, error) {
	return _CertRegistry.Contract.IsWhitelisted(&_CertRegistry.CallOpts, account)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_CertRegistry *CertRegistryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _CertRegistry.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_CertRegistry *CertRegistrySession) Owner() (common.Address, error) {
	return _CertRegistry.Contract.Owner(&_CertRegistry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_CertRegistry *CertRegistryCallerSession) Owner() (common.Address, error) {
	return _CertRegistry.Contract.Owner(&_CertRegistry.CallOpts)
}

// OwnsRoot is a free data retrieval call binding the contract method 0x41d88e0c.
//
// Solidity: function ownsRoot(owner address, root bytes32) constant returns(bool)
func (_CertRegistry *CertRegistryCaller) OwnsRoot(opts *bind.CallOpts, owner common.Address, root [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _CertRegistry.contract.Call(opts, out, "ownsRoot", owner, root)
	return *ret0, err
}

// OwnsRoot is a free data retrieval call binding the contract method 0x41d88e0c.
//
// Solidity: function ownsRoot(owner address, root bytes32) constant returns(bool)
func (_CertRegistry *CertRegistrySession) OwnsRoot(owner common.Address, root [32]byte) (bool, error) {
	return _CertRegistry.Contract.OwnsRoot(&_CertRegistry.CallOpts, owner, root)
}

// OwnsRoot is a free data retrieval call binding the contract method 0x41d88e0c.
//
// Solidity: function ownsRoot(owner address, root bytes32) constant returns(bool)
func (_CertRegistry *CertRegistryCallerSession) OwnsRoot(owner common.Address, root [32]byte) (bool, error) {
	return _CertRegistry.Contract.OwnsRoot(&_CertRegistry.CallOpts, owner, root)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_CertRegistry *CertRegistryCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _CertRegistry.contract.Call(opts, out, "paused")
	return *ret0, err
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_CertRegistry *CertRegistrySession) Paused() (bool, error) {
	return _CertRegistry.Contract.Paused(&_CertRegistry.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_CertRegistry *CertRegistryCallerSession) Paused() (bool, error) {
	return _CertRegistry.Contract.Paused(&_CertRegistry.CallOpts)
}

// RootExists is a free data retrieval call binding the contract method 0x0524b18f.
//
// Solidity: function rootExists(root bytes32) constant returns(bool)
func (_CertRegistry *CertRegistryCaller) RootExists(opts *bind.CallOpts, root [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _CertRegistry.contract.Call(opts, out, "rootExists", root)
	return *ret0, err
}

// RootExists is a free data retrieval call binding the contract method 0x0524b18f.
//
// Solidity: function rootExists(root bytes32) constant returns(bool)
func (_CertRegistry *CertRegistrySession) RootExists(root [32]byte) (bool, error) {
	return _CertRegistry.Contract.RootExists(&_CertRegistry.CallOpts, root)
}

// RootExists is a free data retrieval call binding the contract method 0x0524b18f.
//
// Solidity: function rootExists(root bytes32) constant returns(bool)
func (_CertRegistry *CertRegistryCallerSession) RootExists(root [32]byte) (bool, error) {
	return _CertRegistry.Contract.RootExists(&_CertRegistry.CallOpts, root)
}

// AddPauser is a paid mutator transaction binding the contract method 0x82dc1ec4.
//
// Solidity: function addPauser(account address) returns()
func (_CertRegistry *CertRegistryTransactor) AddPauser(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _CertRegistry.contract.Transact(opts, "addPauser", account)
}

// AddPauser is a paid mutator transaction binding the contract method 0x82dc1ec4.
//
// Solidity: function addPauser(account address) returns()
func (_CertRegistry *CertRegistrySession) AddPauser(account common.Address) (*types.Transaction, error) {
	return _CertRegistry.Contract.AddPauser(&_CertRegistry.TransactOpts, account)
}

// AddPauser is a paid mutator transaction binding the contract method 0x82dc1ec4.
//
// Solidity: function addPauser(account address) returns()
func (_CertRegistry *CertRegistryTransactorSession) AddPauser(account common.Address) (*types.Transaction, error) {
	return _CertRegistry.Contract.AddPauser(&_CertRegistry.TransactOpts, account)
}

// AddWhitelistAdmin is a paid mutator transaction binding the contract method 0x7362d9c8.
//
// Solidity: function addWhitelistAdmin(account address) returns()
func (_CertRegistry *CertRegistryTransactor) AddWhitelistAdmin(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _CertRegistry.contract.Transact(opts, "addWhitelistAdmin", account)
}

// AddWhitelistAdmin is a paid mutator transaction binding the contract method 0x7362d9c8.
//
// Solidity: function addWhitelistAdmin(account address) returns()
func (_CertRegistry *CertRegistrySession) AddWhitelistAdmin(account common.Address) (*types.Transaction, error) {
	return _CertRegistry.Contract.AddWhitelistAdmin(&_CertRegistry.TransactOpts, account)
}

// AddWhitelistAdmin is a paid mutator transaction binding the contract method 0x7362d9c8.
//
// Solidity: function addWhitelistAdmin(account address) returns()
func (_CertRegistry *CertRegistryTransactorSession) AddWhitelistAdmin(account common.Address) (*types.Transaction, error) {
	return _CertRegistry.Contract.AddWhitelistAdmin(&_CertRegistry.TransactOpts, account)
}

// AddWhitelisted is a paid mutator transaction binding the contract method 0x10154bad.
//
// Solidity: function addWhitelisted(account address) returns()
func (_CertRegistry *CertRegistryTransactor) AddWhitelisted(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _CertRegistry.contract.Transact(opts, "addWhitelisted", account)
}

// AddWhitelisted is a paid mutator transaction binding the contract method 0x10154bad.
//
// Solidity: function addWhitelisted(account address) returns()
func (_CertRegistry *CertRegistrySession) AddWhitelisted(account common.Address) (*types.Transaction, error) {
	return _CertRegistry.Contract.AddWhitelisted(&_CertRegistry.TransactOpts, account)
}

// AddWhitelisted is a paid mutator transaction binding the contract method 0x10154bad.
//
// Solidity: function addWhitelisted(account address) returns()
func (_CertRegistry *CertRegistryTransactorSession) AddWhitelisted(account common.Address) (*types.Transaction, error) {
	return _CertRegistry.Contract.AddWhitelisted(&_CertRegistry.TransactOpts, account)
}

// GiveRootOwnership is a paid mutator transaction binding the contract method 0x55457ab6.
//
// Solidity: function giveRootOwnership(currentOwner address, newOwner address) returns()
func (_CertRegistry *CertRegistryTransactor) GiveRootOwnership(opts *bind.TransactOpts, currentOwner common.Address, newOwner common.Address) (*types.Transaction, error) {
	return _CertRegistry.contract.Transact(opts, "giveRootOwnership", currentOwner, newOwner)
}

// GiveRootOwnership is a paid mutator transaction binding the contract method 0x55457ab6.
//
// Solidity: function giveRootOwnership(currentOwner address, newOwner address) returns()
func (_CertRegistry *CertRegistrySession) GiveRootOwnership(currentOwner common.Address, newOwner common.Address) (*types.Transaction, error) {
	return _CertRegistry.Contract.GiveRootOwnership(&_CertRegistry.TransactOpts, currentOwner, newOwner)
}

// GiveRootOwnership is a paid mutator transaction binding the contract method 0x55457ab6.
//
// Solidity: function giveRootOwnership(currentOwner address, newOwner address) returns()
func (_CertRegistry *CertRegistryTransactorSession) GiveRootOwnership(currentOwner common.Address, newOwner common.Address) (*types.Transaction, error) {
	return _CertRegistry.Contract.GiveRootOwnership(&_CertRegistry.TransactOpts, currentOwner, newOwner)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_CertRegistry *CertRegistryTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CertRegistry.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_CertRegistry *CertRegistrySession) Pause() (*types.Transaction, error) {
	return _CertRegistry.Contract.Pause(&_CertRegistry.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_CertRegistry *CertRegistryTransactorSession) Pause() (*types.Transaction, error) {
	return _CertRegistry.Contract.Pause(&_CertRegistry.TransactOpts)
}

// Register is a paid mutator transaction binding the contract method 0xe1fa8e84.
//
// Solidity: function register(root bytes32) returns()
func (_CertRegistry *CertRegistryTransactor) Register(opts *bind.TransactOpts, root [32]byte) (*types.Transaction, error) {
	return _CertRegistry.contract.Transact(opts, "register", root)
}

// Register is a paid mutator transaction binding the contract method 0xe1fa8e84.
//
// Solidity: function register(root bytes32) returns()
func (_CertRegistry *CertRegistrySession) Register(root [32]byte) (*types.Transaction, error) {
	return _CertRegistry.Contract.Register(&_CertRegistry.TransactOpts, root)
}

// Register is a paid mutator transaction binding the contract method 0xe1fa8e84.
//
// Solidity: function register(root bytes32) returns()
func (_CertRegistry *CertRegistryTransactorSession) Register(root [32]byte) (*types.Transaction, error) {
	return _CertRegistry.Contract.Register(&_CertRegistry.TransactOpts, root)
}

// RemoveRootOwnership is a paid mutator transaction binding the contract method 0x94c2dac8.
//
// Solidity: function removeRootOwnership(owner address) returns()
func (_CertRegistry *CertRegistryTransactor) RemoveRootOwnership(opts *bind.TransactOpts, owner common.Address) (*types.Transaction, error) {
	return _CertRegistry.contract.Transact(opts, "removeRootOwnership", owner)
}

// RemoveRootOwnership is a paid mutator transaction binding the contract method 0x94c2dac8.
//
// Solidity: function removeRootOwnership(owner address) returns()
func (_CertRegistry *CertRegistrySession) RemoveRootOwnership(owner common.Address) (*types.Transaction, error) {
	return _CertRegistry.Contract.RemoveRootOwnership(&_CertRegistry.TransactOpts, owner)
}

// RemoveRootOwnership is a paid mutator transaction binding the contract method 0x94c2dac8.
//
// Solidity: function removeRootOwnership(owner address) returns()
func (_CertRegistry *CertRegistryTransactorSession) RemoveRootOwnership(owner common.Address) (*types.Transaction, error) {
	return _CertRegistry.Contract.RemoveRootOwnership(&_CertRegistry.TransactOpts, owner)
}

// RemoveWhitelisted is a paid mutator transaction binding the contract method 0x291d9549.
//
// Solidity: function removeWhitelisted(account address) returns()
func (_CertRegistry *CertRegistryTransactor) RemoveWhitelisted(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _CertRegistry.contract.Transact(opts, "removeWhitelisted", account)
}

// RemoveWhitelisted is a paid mutator transaction binding the contract method 0x291d9549.
//
// Solidity: function removeWhitelisted(account address) returns()
func (_CertRegistry *CertRegistrySession) RemoveWhitelisted(account common.Address) (*types.Transaction, error) {
	return _CertRegistry.Contract.RemoveWhitelisted(&_CertRegistry.TransactOpts, account)
}

// RemoveWhitelisted is a paid mutator transaction binding the contract method 0x291d9549.
//
// Solidity: function removeWhitelisted(account address) returns()
func (_CertRegistry *CertRegistryTransactorSession) RemoveWhitelisted(account common.Address) (*types.Transaction, error) {
	return _CertRegistry.Contract.RemoveWhitelisted(&_CertRegistry.TransactOpts, account)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CertRegistry *CertRegistryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CertRegistry.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CertRegistry *CertRegistrySession) RenounceOwnership() (*types.Transaction, error) {
	return _CertRegistry.Contract.RenounceOwnership(&_CertRegistry.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CertRegistry *CertRegistryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _CertRegistry.Contract.RenounceOwnership(&_CertRegistry.TransactOpts)
}

// RenouncePauser is a paid mutator transaction binding the contract method 0x6ef8d66d.
//
// Solidity: function renouncePauser() returns()
func (_CertRegistry *CertRegistryTransactor) RenouncePauser(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CertRegistry.contract.Transact(opts, "renouncePauser")
}

// RenouncePauser is a paid mutator transaction binding the contract method 0x6ef8d66d.
//
// Solidity: function renouncePauser() returns()
func (_CertRegistry *CertRegistrySession) RenouncePauser() (*types.Transaction, error) {
	return _CertRegistry.Contract.RenouncePauser(&_CertRegistry.TransactOpts)
}

// RenouncePauser is a paid mutator transaction binding the contract method 0x6ef8d66d.
//
// Solidity: function renouncePauser() returns()
func (_CertRegistry *CertRegistryTransactorSession) RenouncePauser() (*types.Transaction, error) {
	return _CertRegistry.Contract.RenouncePauser(&_CertRegistry.TransactOpts)
}

// RenounceRootOwnership is a paid mutator transaction binding the contract method 0x9aa6b585.
//
// Solidity: function renounceRootOwnership() returns()
func (_CertRegistry *CertRegistryTransactor) RenounceRootOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CertRegistry.contract.Transact(opts, "renounceRootOwnership")
}

// RenounceRootOwnership is a paid mutator transaction binding the contract method 0x9aa6b585.
//
// Solidity: function renounceRootOwnership() returns()
func (_CertRegistry *CertRegistrySession) RenounceRootOwnership() (*types.Transaction, error) {
	return _CertRegistry.Contract.RenounceRootOwnership(&_CertRegistry.TransactOpts)
}

// RenounceRootOwnership is a paid mutator transaction binding the contract method 0x9aa6b585.
//
// Solidity: function renounceRootOwnership() returns()
func (_CertRegistry *CertRegistryTransactorSession) RenounceRootOwnership() (*types.Transaction, error) {
	return _CertRegistry.Contract.RenounceRootOwnership(&_CertRegistry.TransactOpts)
}

// RenounceWhitelistAdmin is a paid mutator transaction binding the contract method 0x4c5a628c.
//
// Solidity: function renounceWhitelistAdmin() returns()
func (_CertRegistry *CertRegistryTransactor) RenounceWhitelistAdmin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CertRegistry.contract.Transact(opts, "renounceWhitelistAdmin")
}

// RenounceWhitelistAdmin is a paid mutator transaction binding the contract method 0x4c5a628c.
//
// Solidity: function renounceWhitelistAdmin() returns()
func (_CertRegistry *CertRegistrySession) RenounceWhitelistAdmin() (*types.Transaction, error) {
	return _CertRegistry.Contract.RenounceWhitelistAdmin(&_CertRegistry.TransactOpts)
}

// RenounceWhitelistAdmin is a paid mutator transaction binding the contract method 0x4c5a628c.
//
// Solidity: function renounceWhitelistAdmin() returns()
func (_CertRegistry *CertRegistryTransactorSession) RenounceWhitelistAdmin() (*types.Transaction, error) {
	return _CertRegistry.Contract.RenounceWhitelistAdmin(&_CertRegistry.TransactOpts)
}

// RenounceWhitelisted is a paid mutator transaction binding the contract method 0xd6cd9473.
//
// Solidity: function renounceWhitelisted() returns()
func (_CertRegistry *CertRegistryTransactor) RenounceWhitelisted(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CertRegistry.contract.Transact(opts, "renounceWhitelisted")
}

// RenounceWhitelisted is a paid mutator transaction binding the contract method 0xd6cd9473.
//
// Solidity: function renounceWhitelisted() returns()
func (_CertRegistry *CertRegistrySession) RenounceWhitelisted() (*types.Transaction, error) {
	return _CertRegistry.Contract.RenounceWhitelisted(&_CertRegistry.TransactOpts)
}

// RenounceWhitelisted is a paid mutator transaction binding the contract method 0xd6cd9473.
//
// Solidity: function renounceWhitelisted() returns()
func (_CertRegistry *CertRegistryTransactorSession) RenounceWhitelisted() (*types.Transaction, error) {
	return _CertRegistry.Contract.RenounceWhitelisted(&_CertRegistry.TransactOpts)
}

// Revoke is a paid mutator transaction binding the contract method 0xb75c7dc6.
//
// Solidity: function revoke(hash bytes32) returns()
func (_CertRegistry *CertRegistryTransactor) Revoke(opts *bind.TransactOpts, hash [32]byte) (*types.Transaction, error) {
	return _CertRegistry.contract.Transact(opts, "revoke", hash)
}

// Revoke is a paid mutator transaction binding the contract method 0xb75c7dc6.
//
// Solidity: function revoke(hash bytes32) returns()
func (_CertRegistry *CertRegistrySession) Revoke(hash [32]byte) (*types.Transaction, error) {
	return _CertRegistry.Contract.Revoke(&_CertRegistry.TransactOpts, hash)
}

// Revoke is a paid mutator transaction binding the contract method 0xb75c7dc6.
//
// Solidity: function revoke(hash bytes32) returns()
func (_CertRegistry *CertRegistryTransactorSession) Revoke(hash [32]byte) (*types.Transaction, error) {
	return _CertRegistry.Contract.Revoke(&_CertRegistry.TransactOpts, hash)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_CertRegistry *CertRegistryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _CertRegistry.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_CertRegistry *CertRegistrySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _CertRegistry.Contract.TransferOwnership(&_CertRegistry.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_CertRegistry *CertRegistryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _CertRegistry.Contract.TransferOwnership(&_CertRegistry.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_CertRegistry *CertRegistryTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CertRegistry.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_CertRegistry *CertRegistrySession) Unpause() (*types.Transaction, error) {
	return _CertRegistry.Contract.Unpause(&_CertRegistry.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_CertRegistry *CertRegistryTransactorSession) Unpause() (*types.Transaction, error) {
	return _CertRegistry.Contract.Unpause(&_CertRegistry.TransactOpts)
}

// CertRegistryCertRegisteredIterator is returned from FilterCertRegistered and is used to iterate over the raw logs and unpacked data for CertRegistered events raised by the CertRegistry contract.
type CertRegistryCertRegisteredIterator struct {
	Event *CertRegistryCertRegistered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CertRegistryCertRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CertRegistryCertRegistered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CertRegistryCertRegistered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CertRegistryCertRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CertRegistryCertRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CertRegistryCertRegistered represents a CertRegistered event raised by the CertRegistry contract.
type CertRegistryCertRegistered struct {
	Registerer common.Address
	Root       [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCertRegistered is a free log retrieval operation binding the contract event 0x54553ef3b4e1d60c041c89079b49d51c62e39543028b4861c9e25529acaa70e8.
//
// Solidity: e CertRegistered(registerer address, root bytes32)
func (_CertRegistry *CertRegistryFilterer) FilterCertRegistered(opts *bind.FilterOpts) (*CertRegistryCertRegisteredIterator, error) {

	logs, sub, err := _CertRegistry.contract.FilterLogs(opts, "CertRegistered")
	if err != nil {
		return nil, err
	}
	return &CertRegistryCertRegisteredIterator{contract: _CertRegistry.contract, event: "CertRegistered", logs: logs, sub: sub}, nil
}

// WatchCertRegistered is a free log subscription operation binding the contract event 0x54553ef3b4e1d60c041c89079b49d51c62e39543028b4861c9e25529acaa70e8.
//
// Solidity: e CertRegistered(registerer address, root bytes32)
func (_CertRegistry *CertRegistryFilterer) WatchCertRegistered(opts *bind.WatchOpts, sink chan<- *CertRegistryCertRegistered) (event.Subscription, error) {

	logs, sub, err := _CertRegistry.contract.WatchLogs(opts, "CertRegistered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CertRegistryCertRegistered)
				if err := _CertRegistry.contract.UnpackLog(event, "CertRegistered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// CertRegistryCertRevokedIterator is returned from FilterCertRevoked and is used to iterate over the raw logs and unpacked data for CertRevoked events raised by the CertRegistry contract.
type CertRegistryCertRevokedIterator struct {
	Event *CertRegistryCertRevoked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CertRegistryCertRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CertRegistryCertRevoked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CertRegistryCertRevoked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CertRegistryCertRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CertRegistryCertRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CertRegistryCertRevoked represents a CertRevoked event raised by the CertRegistry contract.
type CertRegistryCertRevoked struct {
	Revoker common.Address
	Hash    [32]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterCertRevoked is a free log retrieval operation binding the contract event 0x95805cdad68eccc1a7e73b9e1dd10d0561be298bdab9e5609e309df2e3c3d7b7.
//
// Solidity: e CertRevoked(revoker address, hash bytes32)
func (_CertRegistry *CertRegistryFilterer) FilterCertRevoked(opts *bind.FilterOpts) (*CertRegistryCertRevokedIterator, error) {

	logs, sub, err := _CertRegistry.contract.FilterLogs(opts, "CertRevoked")
	if err != nil {
		return nil, err
	}
	return &CertRegistryCertRevokedIterator{contract: _CertRegistry.contract, event: "CertRevoked", logs: logs, sub: sub}, nil
}

// WatchCertRevoked is a free log subscription operation binding the contract event 0x95805cdad68eccc1a7e73b9e1dd10d0561be298bdab9e5609e309df2e3c3d7b7.
//
// Solidity: e CertRevoked(revoker address, hash bytes32)
func (_CertRegistry *CertRegistryFilterer) WatchCertRevoked(opts *bind.WatchOpts, sink chan<- *CertRegistryCertRevoked) (event.Subscription, error) {

	logs, sub, err := _CertRegistry.contract.WatchLogs(opts, "CertRevoked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CertRegistryCertRevoked)
				if err := _CertRegistry.contract.UnpackLog(event, "CertRevoked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// CertRegistryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the CertRegistry contract.
type CertRegistryOwnershipTransferredIterator struct {
	Event *CertRegistryOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CertRegistryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CertRegistryOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CertRegistryOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CertRegistryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CertRegistryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CertRegistryOwnershipTransferred represents a OwnershipTransferred event raised by the CertRegistry contract.
type CertRegistryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_CertRegistry *CertRegistryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*CertRegistryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CertRegistry.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &CertRegistryOwnershipTransferredIterator{contract: _CertRegistry.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_CertRegistry *CertRegistryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CertRegistryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CertRegistry.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CertRegistryOwnershipTransferred)
				if err := _CertRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// CertRegistryPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the CertRegistry contract.
type CertRegistryPausedIterator struct {
	Event *CertRegistryPaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CertRegistryPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CertRegistryPaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CertRegistryPaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CertRegistryPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CertRegistryPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CertRegistryPaused represents a Paused event raised by the CertRegistry contract.
type CertRegistryPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: e Paused(account address)
func (_CertRegistry *CertRegistryFilterer) FilterPaused(opts *bind.FilterOpts) (*CertRegistryPausedIterator, error) {

	logs, sub, err := _CertRegistry.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &CertRegistryPausedIterator{contract: _CertRegistry.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: e Paused(account address)
func (_CertRegistry *CertRegistryFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *CertRegistryPaused) (event.Subscription, error) {

	logs, sub, err := _CertRegistry.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CertRegistryPaused)
				if err := _CertRegistry.contract.UnpackLog(event, "Paused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// CertRegistryPauserAddedIterator is returned from FilterPauserAdded and is used to iterate over the raw logs and unpacked data for PauserAdded events raised by the CertRegistry contract.
type CertRegistryPauserAddedIterator struct {
	Event *CertRegistryPauserAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CertRegistryPauserAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CertRegistryPauserAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CertRegistryPauserAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CertRegistryPauserAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CertRegistryPauserAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CertRegistryPauserAdded represents a PauserAdded event raised by the CertRegistry contract.
type CertRegistryPauserAdded struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPauserAdded is a free log retrieval operation binding the contract event 0x6719d08c1888103bea251a4ed56406bd0c3e69723c8a1686e017e7bbe159b6f8.
//
// Solidity: e PauserAdded(account indexed address)
func (_CertRegistry *CertRegistryFilterer) FilterPauserAdded(opts *bind.FilterOpts, account []common.Address) (*CertRegistryPauserAddedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _CertRegistry.contract.FilterLogs(opts, "PauserAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return &CertRegistryPauserAddedIterator{contract: _CertRegistry.contract, event: "PauserAdded", logs: logs, sub: sub}, nil
}

// WatchPauserAdded is a free log subscription operation binding the contract event 0x6719d08c1888103bea251a4ed56406bd0c3e69723c8a1686e017e7bbe159b6f8.
//
// Solidity: e PauserAdded(account indexed address)
func (_CertRegistry *CertRegistryFilterer) WatchPauserAdded(opts *bind.WatchOpts, sink chan<- *CertRegistryPauserAdded, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _CertRegistry.contract.WatchLogs(opts, "PauserAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CertRegistryPauserAdded)
				if err := _CertRegistry.contract.UnpackLog(event, "PauserAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// CertRegistryPauserRemovedIterator is returned from FilterPauserRemoved and is used to iterate over the raw logs and unpacked data for PauserRemoved events raised by the CertRegistry contract.
type CertRegistryPauserRemovedIterator struct {
	Event *CertRegistryPauserRemoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CertRegistryPauserRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CertRegistryPauserRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CertRegistryPauserRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CertRegistryPauserRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CertRegistryPauserRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CertRegistryPauserRemoved represents a PauserRemoved event raised by the CertRegistry contract.
type CertRegistryPauserRemoved struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPauserRemoved is a free log retrieval operation binding the contract event 0xcd265ebaf09df2871cc7bd4133404a235ba12eff2041bb89d9c714a2621c7c7e.
//
// Solidity: e PauserRemoved(account indexed address)
func (_CertRegistry *CertRegistryFilterer) FilterPauserRemoved(opts *bind.FilterOpts, account []common.Address) (*CertRegistryPauserRemovedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _CertRegistry.contract.FilterLogs(opts, "PauserRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return &CertRegistryPauserRemovedIterator{contract: _CertRegistry.contract, event: "PauserRemoved", logs: logs, sub: sub}, nil
}

// WatchPauserRemoved is a free log subscription operation binding the contract event 0xcd265ebaf09df2871cc7bd4133404a235ba12eff2041bb89d9c714a2621c7c7e.
//
// Solidity: e PauserRemoved(account indexed address)
func (_CertRegistry *CertRegistryFilterer) WatchPauserRemoved(opts *bind.WatchOpts, sink chan<- *CertRegistryPauserRemoved, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _CertRegistry.contract.WatchLogs(opts, "PauserRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CertRegistryPauserRemoved)
				if err := _CertRegistry.contract.UnpackLog(event, "PauserRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// CertRegistryRootOwnershipGivenIterator is returned from FilterRootOwnershipGiven and is used to iterate over the raw logs and unpacked data for RootOwnershipGiven events raised by the CertRegistry contract.
type CertRegistryRootOwnershipGivenIterator struct {
	Event *CertRegistryRootOwnershipGiven // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CertRegistryRootOwnershipGivenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CertRegistryRootOwnershipGiven)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CertRegistryRootOwnershipGiven)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CertRegistryRootOwnershipGivenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CertRegistryRootOwnershipGivenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CertRegistryRootOwnershipGiven represents a RootOwnershipGiven event raised by the CertRegistry contract.
type CertRegistryRootOwnershipGiven struct {
	CurrentOwner common.Address
	NewOwner     common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRootOwnershipGiven is a free log retrieval operation binding the contract event 0x2854d8d9ea9d5530703d62c8b3e8bb2459a9136e7607511da717b2a13dcf70b8.
//
// Solidity: e RootOwnershipGiven(currentOwner address, newOwner address)
func (_CertRegistry *CertRegistryFilterer) FilterRootOwnershipGiven(opts *bind.FilterOpts) (*CertRegistryRootOwnershipGivenIterator, error) {

	logs, sub, err := _CertRegistry.contract.FilterLogs(opts, "RootOwnershipGiven")
	if err != nil {
		return nil, err
	}
	return &CertRegistryRootOwnershipGivenIterator{contract: _CertRegistry.contract, event: "RootOwnershipGiven", logs: logs, sub: sub}, nil
}

// WatchRootOwnershipGiven is a free log subscription operation binding the contract event 0x2854d8d9ea9d5530703d62c8b3e8bb2459a9136e7607511da717b2a13dcf70b8.
//
// Solidity: e RootOwnershipGiven(currentOwner address, newOwner address)
func (_CertRegistry *CertRegistryFilterer) WatchRootOwnershipGiven(opts *bind.WatchOpts, sink chan<- *CertRegistryRootOwnershipGiven) (event.Subscription, error) {

	logs, sub, err := _CertRegistry.contract.WatchLogs(opts, "RootOwnershipGiven")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CertRegistryRootOwnershipGiven)
				if err := _CertRegistry.contract.UnpackLog(event, "RootOwnershipGiven", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// CertRegistryUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the CertRegistry contract.
type CertRegistryUnpausedIterator struct {
	Event *CertRegistryUnpaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CertRegistryUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CertRegistryUnpaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CertRegistryUnpaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CertRegistryUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CertRegistryUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CertRegistryUnpaused represents a Unpaused event raised by the CertRegistry contract.
type CertRegistryUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: e Unpaused(account address)
func (_CertRegistry *CertRegistryFilterer) FilterUnpaused(opts *bind.FilterOpts) (*CertRegistryUnpausedIterator, error) {

	logs, sub, err := _CertRegistry.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &CertRegistryUnpausedIterator{contract: _CertRegistry.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: e Unpaused(account address)
func (_CertRegistry *CertRegistryFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *CertRegistryUnpaused) (event.Subscription, error) {

	logs, sub, err := _CertRegistry.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CertRegistryUnpaused)
				if err := _CertRegistry.contract.UnpackLog(event, "Unpaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// CertRegistryWhitelistAdminAddedIterator is returned from FilterWhitelistAdminAdded and is used to iterate over the raw logs and unpacked data for WhitelistAdminAdded events raised by the CertRegistry contract.
type CertRegistryWhitelistAdminAddedIterator struct {
	Event *CertRegistryWhitelistAdminAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CertRegistryWhitelistAdminAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CertRegistryWhitelistAdminAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CertRegistryWhitelistAdminAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CertRegistryWhitelistAdminAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CertRegistryWhitelistAdminAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CertRegistryWhitelistAdminAdded represents a WhitelistAdminAdded event raised by the CertRegistry contract.
type CertRegistryWhitelistAdminAdded struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWhitelistAdminAdded is a free log retrieval operation binding the contract event 0x22380c05984257a1cb900161c713dd71d39e74820f1aea43bd3f1bdd20961299.
//
// Solidity: e WhitelistAdminAdded(account indexed address)
func (_CertRegistry *CertRegistryFilterer) FilterWhitelistAdminAdded(opts *bind.FilterOpts, account []common.Address) (*CertRegistryWhitelistAdminAddedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _CertRegistry.contract.FilterLogs(opts, "WhitelistAdminAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return &CertRegistryWhitelistAdminAddedIterator{contract: _CertRegistry.contract, event: "WhitelistAdminAdded", logs: logs, sub: sub}, nil
}

// WatchWhitelistAdminAdded is a free log subscription operation binding the contract event 0x22380c05984257a1cb900161c713dd71d39e74820f1aea43bd3f1bdd20961299.
//
// Solidity: e WhitelistAdminAdded(account indexed address)
func (_CertRegistry *CertRegistryFilterer) WatchWhitelistAdminAdded(opts *bind.WatchOpts, sink chan<- *CertRegistryWhitelistAdminAdded, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _CertRegistry.contract.WatchLogs(opts, "WhitelistAdminAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CertRegistryWhitelistAdminAdded)
				if err := _CertRegistry.contract.UnpackLog(event, "WhitelistAdminAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// CertRegistryWhitelistAdminRemovedIterator is returned from FilterWhitelistAdminRemoved and is used to iterate over the raw logs and unpacked data for WhitelistAdminRemoved events raised by the CertRegistry contract.
type CertRegistryWhitelistAdminRemovedIterator struct {
	Event *CertRegistryWhitelistAdminRemoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CertRegistryWhitelistAdminRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CertRegistryWhitelistAdminRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CertRegistryWhitelistAdminRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CertRegistryWhitelistAdminRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CertRegistryWhitelistAdminRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CertRegistryWhitelistAdminRemoved represents a WhitelistAdminRemoved event raised by the CertRegistry contract.
type CertRegistryWhitelistAdminRemoved struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWhitelistAdminRemoved is a free log retrieval operation binding the contract event 0x0a8eb35e5ca14b3d6f28e4abf2f128dbab231a58b56e89beb5d636115001e165.
//
// Solidity: e WhitelistAdminRemoved(account indexed address)
func (_CertRegistry *CertRegistryFilterer) FilterWhitelistAdminRemoved(opts *bind.FilterOpts, account []common.Address) (*CertRegistryWhitelistAdminRemovedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _CertRegistry.contract.FilterLogs(opts, "WhitelistAdminRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return &CertRegistryWhitelistAdminRemovedIterator{contract: _CertRegistry.contract, event: "WhitelistAdminRemoved", logs: logs, sub: sub}, nil
}

// WatchWhitelistAdminRemoved is a free log subscription operation binding the contract event 0x0a8eb35e5ca14b3d6f28e4abf2f128dbab231a58b56e89beb5d636115001e165.
//
// Solidity: e WhitelistAdminRemoved(account indexed address)
func (_CertRegistry *CertRegistryFilterer) WatchWhitelistAdminRemoved(opts *bind.WatchOpts, sink chan<- *CertRegistryWhitelistAdminRemoved, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _CertRegistry.contract.WatchLogs(opts, "WhitelistAdminRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CertRegistryWhitelistAdminRemoved)
				if err := _CertRegistry.contract.UnpackLog(event, "WhitelistAdminRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// CertRegistryWhitelistedAddedIterator is returned from FilterWhitelistedAdded and is used to iterate over the raw logs and unpacked data for WhitelistedAdded events raised by the CertRegistry contract.
type CertRegistryWhitelistedAddedIterator struct {
	Event *CertRegistryWhitelistedAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CertRegistryWhitelistedAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CertRegistryWhitelistedAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CertRegistryWhitelistedAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CertRegistryWhitelistedAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CertRegistryWhitelistedAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CertRegistryWhitelistedAdded represents a WhitelistedAdded event raised by the CertRegistry contract.
type CertRegistryWhitelistedAdded struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWhitelistedAdded is a free log retrieval operation binding the contract event 0xee1504a83b6d4a361f4c1dc78ab59bfa30d6a3b6612c403e86bb01ef2984295f.
//
// Solidity: e WhitelistedAdded(account indexed address)
func (_CertRegistry *CertRegistryFilterer) FilterWhitelistedAdded(opts *bind.FilterOpts, account []common.Address) (*CertRegistryWhitelistedAddedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _CertRegistry.contract.FilterLogs(opts, "WhitelistedAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return &CertRegistryWhitelistedAddedIterator{contract: _CertRegistry.contract, event: "WhitelistedAdded", logs: logs, sub: sub}, nil
}

// WatchWhitelistedAdded is a free log subscription operation binding the contract event 0xee1504a83b6d4a361f4c1dc78ab59bfa30d6a3b6612c403e86bb01ef2984295f.
//
// Solidity: e WhitelistedAdded(account indexed address)
func (_CertRegistry *CertRegistryFilterer) WatchWhitelistedAdded(opts *bind.WatchOpts, sink chan<- *CertRegistryWhitelistedAdded, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _CertRegistry.contract.WatchLogs(opts, "WhitelistedAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CertRegistryWhitelistedAdded)
				if err := _CertRegistry.contract.UnpackLog(event, "WhitelistedAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// CertRegistryWhitelistedRemovedIterator is returned from FilterWhitelistedRemoved and is used to iterate over the raw logs and unpacked data for WhitelistedRemoved events raised by the CertRegistry contract.
type CertRegistryWhitelistedRemovedIterator struct {
	Event *CertRegistryWhitelistedRemoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CertRegistryWhitelistedRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CertRegistryWhitelistedRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CertRegistryWhitelistedRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CertRegistryWhitelistedRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CertRegistryWhitelistedRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CertRegistryWhitelistedRemoved represents a WhitelistedRemoved event raised by the CertRegistry contract.
type CertRegistryWhitelistedRemoved struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWhitelistedRemoved is a free log retrieval operation binding the contract event 0x270d9b30cf5b0793bbfd54c9d5b94aeb49462b8148399000265144a8722da6b6.
//
// Solidity: e WhitelistedRemoved(account indexed address)
func (_CertRegistry *CertRegistryFilterer) FilterWhitelistedRemoved(opts *bind.FilterOpts, account []common.Address) (*CertRegistryWhitelistedRemovedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _CertRegistry.contract.FilterLogs(opts, "WhitelistedRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return &CertRegistryWhitelistedRemovedIterator{contract: _CertRegistry.contract, event: "WhitelistedRemoved", logs: logs, sub: sub}, nil
}

// WatchWhitelistedRemoved is a free log subscription operation binding the contract event 0x270d9b30cf5b0793bbfd54c9d5b94aeb49462b8148399000265144a8722da6b6.
//
// Solidity: e WhitelistedRemoved(account indexed address)
func (_CertRegistry *CertRegistryFilterer) WatchWhitelistedRemoved(opts *bind.WatchOpts, sink chan<- *CertRegistryWhitelistedRemoved, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _CertRegistry.contract.WatchLogs(opts, "WhitelistedRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CertRegistryWhitelistedRemoved)
				if err := _CertRegistry.contract.UnpackLog(event, "WhitelistedRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// OwnableABI is the input ABI used to generate the binding from.
const OwnableABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// OwnableBin is the compiled bytecode used for deploying new contracts.
const OwnableBin = `0x`

// DeployOwnable deploys a new Ethereum contract, binding an instance of Ownable to it.
func DeployOwnable(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Ownable, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnableABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OwnableBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Ownable{OwnableCaller: OwnableCaller{contract: contract}, OwnableTransactor: OwnableTransactor{contract: contract}, OwnableFilterer: OwnableFilterer{contract: contract}}, nil
}

// Ownable is an auto generated Go binding around an Ethereum contract.
type Ownable struct {
	OwnableCaller     // Read-only binding to the contract
	OwnableTransactor // Write-only binding to the contract
	OwnableFilterer   // Log filterer for contract events
}

// OwnableCaller is an auto generated read-only Go binding around an Ethereum contract.
type OwnableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OwnableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OwnableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OwnableSession struct {
	Contract     *Ownable          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OwnableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OwnableCallerSession struct {
	Contract *OwnableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// OwnableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OwnableTransactorSession struct {
	Contract     *OwnableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// OwnableRaw is an auto generated low-level Go binding around an Ethereum contract.
type OwnableRaw struct {
	Contract *Ownable // Generic contract binding to access the raw methods on
}

// OwnableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OwnableCallerRaw struct {
	Contract *OwnableCaller // Generic read-only contract binding to access the raw methods on
}

// OwnableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OwnableTransactorRaw struct {
	Contract *OwnableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOwnable creates a new instance of Ownable, bound to a specific deployed contract.
func NewOwnable(address common.Address, backend bind.ContractBackend) (*Ownable, error) {
	contract, err := bindOwnable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ownable{OwnableCaller: OwnableCaller{contract: contract}, OwnableTransactor: OwnableTransactor{contract: contract}, OwnableFilterer: OwnableFilterer{contract: contract}}, nil
}

// NewOwnableCaller creates a new read-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableCaller(address common.Address, caller bind.ContractCaller) (*OwnableCaller, error) {
	contract, err := bindOwnable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableCaller{contract: contract}, nil
}

// NewOwnableTransactor creates a new write-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableTransactor(address common.Address, transactor bind.ContractTransactor) (*OwnableTransactor, error) {
	contract, err := bindOwnable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableTransactor{contract: contract}, nil
}

// NewOwnableFilterer creates a new log filterer instance of Ownable, bound to a specific deployed contract.
func NewOwnableFilterer(address common.Address, filterer bind.ContractFilterer) (*OwnableFilterer, error) {
	contract, err := bindOwnable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OwnableFilterer{contract: contract}, nil
}

// bindOwnable binds a generic wrapper to an already deployed contract.
func bindOwnable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.OwnableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transact(opts, method, params...)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Ownable *OwnableCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Ownable.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Ownable *OwnableSession) IsOwner() (bool, error) {
	return _Ownable.Contract.IsOwner(&_Ownable.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Ownable *OwnableCallerSession) IsOwner() (bool, error) {
	return _Ownable.Contract.IsOwner(&_Ownable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Ownable *OwnableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Ownable.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Ownable *OwnableSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Ownable *OwnableCallerSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableSession) RenounceOwnership() (*types.Transaction, error) {
	return _Ownable.Contract.RenounceOwnership(&_Ownable.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Ownable.Contract.RenounceOwnership(&_Ownable.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Ownable *OwnableTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Ownable *OwnableSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Ownable *OwnableTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, newOwner)
}

// OwnableOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Ownable contract.
type OwnableOwnershipTransferredIterator struct {
	Event *OwnableOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *OwnableOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwnableOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(OwnableOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *OwnableOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwnableOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwnableOwnershipTransferred represents a OwnershipTransferred event raised by the Ownable contract.
type OwnableOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_Ownable *OwnableFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OwnableOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OwnableOwnershipTransferredIterator{contract: _Ownable.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_Ownable *OwnableFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OwnableOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwnableOwnershipTransferred)
				if err := _Ownable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// PausableABI is the input ABI used to generate the binding from.
const PausableABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isPauser\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renouncePauser\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addPauser\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"account\",\"type\":\"address\"}],\"name\":\"PauserAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"account\",\"type\":\"address\"}],\"name\":\"PauserRemoved\",\"type\":\"event\"}]"

// PausableBin is the compiled bytecode used for deploying new contracts.
const PausableBin = `0x`

// DeployPausable deploys a new Ethereum contract, binding an instance of Pausable to it.
func DeployPausable(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Pausable, error) {
	parsed, err := abi.JSON(strings.NewReader(PausableABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PausableBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Pausable{PausableCaller: PausableCaller{contract: contract}, PausableTransactor: PausableTransactor{contract: contract}, PausableFilterer: PausableFilterer{contract: contract}}, nil
}

// Pausable is an auto generated Go binding around an Ethereum contract.
type Pausable struct {
	PausableCaller     // Read-only binding to the contract
	PausableTransactor // Write-only binding to the contract
	PausableFilterer   // Log filterer for contract events
}

// PausableCaller is an auto generated read-only Go binding around an Ethereum contract.
type PausableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PausableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PausableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PausableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PausableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PausableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PausableSession struct {
	Contract     *Pausable         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PausableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PausableCallerSession struct {
	Contract *PausableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// PausableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PausableTransactorSession struct {
	Contract     *PausableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// PausableRaw is an auto generated low-level Go binding around an Ethereum contract.
type PausableRaw struct {
	Contract *Pausable // Generic contract binding to access the raw methods on
}

// PausableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PausableCallerRaw struct {
	Contract *PausableCaller // Generic read-only contract binding to access the raw methods on
}

// PausableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PausableTransactorRaw struct {
	Contract *PausableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPausable creates a new instance of Pausable, bound to a specific deployed contract.
func NewPausable(address common.Address, backend bind.ContractBackend) (*Pausable, error) {
	contract, err := bindPausable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Pausable{PausableCaller: PausableCaller{contract: contract}, PausableTransactor: PausableTransactor{contract: contract}, PausableFilterer: PausableFilterer{contract: contract}}, nil
}

// NewPausableCaller creates a new read-only instance of Pausable, bound to a specific deployed contract.
func NewPausableCaller(address common.Address, caller bind.ContractCaller) (*PausableCaller, error) {
	contract, err := bindPausable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PausableCaller{contract: contract}, nil
}

// NewPausableTransactor creates a new write-only instance of Pausable, bound to a specific deployed contract.
func NewPausableTransactor(address common.Address, transactor bind.ContractTransactor) (*PausableTransactor, error) {
	contract, err := bindPausable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PausableTransactor{contract: contract}, nil
}

// NewPausableFilterer creates a new log filterer instance of Pausable, bound to a specific deployed contract.
func NewPausableFilterer(address common.Address, filterer bind.ContractFilterer) (*PausableFilterer, error) {
	contract, err := bindPausable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PausableFilterer{contract: contract}, nil
}

// bindPausable binds a generic wrapper to an already deployed contract.
func bindPausable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PausableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pausable *PausableRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Pausable.Contract.PausableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pausable *PausableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pausable.Contract.PausableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pausable *PausableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pausable.Contract.PausableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pausable *PausableCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Pausable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pausable *PausableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pausable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pausable *PausableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pausable.Contract.contract.Transact(opts, method, params...)
}

// IsPauser is a free data retrieval call binding the contract method 0x46fbf68e.
//
// Solidity: function isPauser(account address) constant returns(bool)
func (_Pausable *PausableCaller) IsPauser(opts *bind.CallOpts, account common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Pausable.contract.Call(opts, out, "isPauser", account)
	return *ret0, err
}

// IsPauser is a free data retrieval call binding the contract method 0x46fbf68e.
//
// Solidity: function isPauser(account address) constant returns(bool)
func (_Pausable *PausableSession) IsPauser(account common.Address) (bool, error) {
	return _Pausable.Contract.IsPauser(&_Pausable.CallOpts, account)
}

// IsPauser is a free data retrieval call binding the contract method 0x46fbf68e.
//
// Solidity: function isPauser(account address) constant returns(bool)
func (_Pausable *PausableCallerSession) IsPauser(account common.Address) (bool, error) {
	return _Pausable.Contract.IsPauser(&_Pausable.CallOpts, account)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_Pausable *PausableCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Pausable.contract.Call(opts, out, "paused")
	return *ret0, err
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_Pausable *PausableSession) Paused() (bool, error) {
	return _Pausable.Contract.Paused(&_Pausable.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_Pausable *PausableCallerSession) Paused() (bool, error) {
	return _Pausable.Contract.Paused(&_Pausable.CallOpts)
}

// AddPauser is a paid mutator transaction binding the contract method 0x82dc1ec4.
//
// Solidity: function addPauser(account address) returns()
func (_Pausable *PausableTransactor) AddPauser(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Pausable.contract.Transact(opts, "addPauser", account)
}

// AddPauser is a paid mutator transaction binding the contract method 0x82dc1ec4.
//
// Solidity: function addPauser(account address) returns()
func (_Pausable *PausableSession) AddPauser(account common.Address) (*types.Transaction, error) {
	return _Pausable.Contract.AddPauser(&_Pausable.TransactOpts, account)
}

// AddPauser is a paid mutator transaction binding the contract method 0x82dc1ec4.
//
// Solidity: function addPauser(account address) returns()
func (_Pausable *PausableTransactorSession) AddPauser(account common.Address) (*types.Transaction, error) {
	return _Pausable.Contract.AddPauser(&_Pausable.TransactOpts, account)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Pausable *PausableTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pausable.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Pausable *PausableSession) Pause() (*types.Transaction, error) {
	return _Pausable.Contract.Pause(&_Pausable.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Pausable *PausableTransactorSession) Pause() (*types.Transaction, error) {
	return _Pausable.Contract.Pause(&_Pausable.TransactOpts)
}

// RenouncePauser is a paid mutator transaction binding the contract method 0x6ef8d66d.
//
// Solidity: function renouncePauser() returns()
func (_Pausable *PausableTransactor) RenouncePauser(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pausable.contract.Transact(opts, "renouncePauser")
}

// RenouncePauser is a paid mutator transaction binding the contract method 0x6ef8d66d.
//
// Solidity: function renouncePauser() returns()
func (_Pausable *PausableSession) RenouncePauser() (*types.Transaction, error) {
	return _Pausable.Contract.RenouncePauser(&_Pausable.TransactOpts)
}

// RenouncePauser is a paid mutator transaction binding the contract method 0x6ef8d66d.
//
// Solidity: function renouncePauser() returns()
func (_Pausable *PausableTransactorSession) RenouncePauser() (*types.Transaction, error) {
	return _Pausable.Contract.RenouncePauser(&_Pausable.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Pausable *PausableTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pausable.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Pausable *PausableSession) Unpause() (*types.Transaction, error) {
	return _Pausable.Contract.Unpause(&_Pausable.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Pausable *PausableTransactorSession) Unpause() (*types.Transaction, error) {
	return _Pausable.Contract.Unpause(&_Pausable.TransactOpts)
}

// PausablePausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Pausable contract.
type PausablePausedIterator struct {
	Event *PausablePaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PausablePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PausablePaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PausablePaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PausablePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PausablePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PausablePaused represents a Paused event raised by the Pausable contract.
type PausablePaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: e Paused(account address)
func (_Pausable *PausableFilterer) FilterPaused(opts *bind.FilterOpts) (*PausablePausedIterator, error) {

	logs, sub, err := _Pausable.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &PausablePausedIterator{contract: _Pausable.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: e Paused(account address)
func (_Pausable *PausableFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *PausablePaused) (event.Subscription, error) {

	logs, sub, err := _Pausable.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PausablePaused)
				if err := _Pausable.contract.UnpackLog(event, "Paused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// PausablePauserAddedIterator is returned from FilterPauserAdded and is used to iterate over the raw logs and unpacked data for PauserAdded events raised by the Pausable contract.
type PausablePauserAddedIterator struct {
	Event *PausablePauserAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PausablePauserAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PausablePauserAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PausablePauserAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PausablePauserAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PausablePauserAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PausablePauserAdded represents a PauserAdded event raised by the Pausable contract.
type PausablePauserAdded struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPauserAdded is a free log retrieval operation binding the contract event 0x6719d08c1888103bea251a4ed56406bd0c3e69723c8a1686e017e7bbe159b6f8.
//
// Solidity: e PauserAdded(account indexed address)
func (_Pausable *PausableFilterer) FilterPauserAdded(opts *bind.FilterOpts, account []common.Address) (*PausablePauserAddedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Pausable.contract.FilterLogs(opts, "PauserAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return &PausablePauserAddedIterator{contract: _Pausable.contract, event: "PauserAdded", logs: logs, sub: sub}, nil
}

// WatchPauserAdded is a free log subscription operation binding the contract event 0x6719d08c1888103bea251a4ed56406bd0c3e69723c8a1686e017e7bbe159b6f8.
//
// Solidity: e PauserAdded(account indexed address)
func (_Pausable *PausableFilterer) WatchPauserAdded(opts *bind.WatchOpts, sink chan<- *PausablePauserAdded, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Pausable.contract.WatchLogs(opts, "PauserAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PausablePauserAdded)
				if err := _Pausable.contract.UnpackLog(event, "PauserAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// PausablePauserRemovedIterator is returned from FilterPauserRemoved and is used to iterate over the raw logs and unpacked data for PauserRemoved events raised by the Pausable contract.
type PausablePauserRemovedIterator struct {
	Event *PausablePauserRemoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PausablePauserRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PausablePauserRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PausablePauserRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PausablePauserRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PausablePauserRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PausablePauserRemoved represents a PauserRemoved event raised by the Pausable contract.
type PausablePauserRemoved struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPauserRemoved is a free log retrieval operation binding the contract event 0xcd265ebaf09df2871cc7bd4133404a235ba12eff2041bb89d9c714a2621c7c7e.
//
// Solidity: e PauserRemoved(account indexed address)
func (_Pausable *PausableFilterer) FilterPauserRemoved(opts *bind.FilterOpts, account []common.Address) (*PausablePauserRemovedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Pausable.contract.FilterLogs(opts, "PauserRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return &PausablePauserRemovedIterator{contract: _Pausable.contract, event: "PauserRemoved", logs: logs, sub: sub}, nil
}

// WatchPauserRemoved is a free log subscription operation binding the contract event 0xcd265ebaf09df2871cc7bd4133404a235ba12eff2041bb89d9c714a2621c7c7e.
//
// Solidity: e PauserRemoved(account indexed address)
func (_Pausable *PausableFilterer) WatchPauserRemoved(opts *bind.WatchOpts, sink chan<- *PausablePauserRemoved, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Pausable.contract.WatchLogs(opts, "PauserRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PausablePauserRemoved)
				if err := _Pausable.contract.UnpackLog(event, "PauserRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// PausableUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Pausable contract.
type PausableUnpausedIterator struct {
	Event *PausableUnpaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PausableUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PausableUnpaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PausableUnpaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PausableUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PausableUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PausableUnpaused represents a Unpaused event raised by the Pausable contract.
type PausableUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: e Unpaused(account address)
func (_Pausable *PausableFilterer) FilterUnpaused(opts *bind.FilterOpts) (*PausableUnpausedIterator, error) {

	logs, sub, err := _Pausable.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &PausableUnpausedIterator{contract: _Pausable.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: e Unpaused(account address)
func (_Pausable *PausableFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *PausableUnpaused) (event.Subscription, error) {

	logs, sub, err := _Pausable.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PausableUnpaused)
				if err := _Pausable.contract.UnpackLog(event, "Unpaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// PauserRoleABI is the input ABI used to generate the binding from.
const PauserRoleABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isPauser\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renouncePauser\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addPauser\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"account\",\"type\":\"address\"}],\"name\":\"PauserAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"account\",\"type\":\"address\"}],\"name\":\"PauserRemoved\",\"type\":\"event\"}]"

// PauserRoleBin is the compiled bytecode used for deploying new contracts.
const PauserRoleBin = `0x`

// DeployPauserRole deploys a new Ethereum contract, binding an instance of PauserRole to it.
func DeployPauserRole(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *PauserRole, error) {
	parsed, err := abi.JSON(strings.NewReader(PauserRoleABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PauserRoleBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PauserRole{PauserRoleCaller: PauserRoleCaller{contract: contract}, PauserRoleTransactor: PauserRoleTransactor{contract: contract}, PauserRoleFilterer: PauserRoleFilterer{contract: contract}}, nil
}

// PauserRole is an auto generated Go binding around an Ethereum contract.
type PauserRole struct {
	PauserRoleCaller     // Read-only binding to the contract
	PauserRoleTransactor // Write-only binding to the contract
	PauserRoleFilterer   // Log filterer for contract events
}

// PauserRoleCaller is an auto generated read-only Go binding around an Ethereum contract.
type PauserRoleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PauserRoleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PauserRoleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PauserRoleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PauserRoleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PauserRoleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PauserRoleSession struct {
	Contract     *PauserRole       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PauserRoleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PauserRoleCallerSession struct {
	Contract *PauserRoleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// PauserRoleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PauserRoleTransactorSession struct {
	Contract     *PauserRoleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// PauserRoleRaw is an auto generated low-level Go binding around an Ethereum contract.
type PauserRoleRaw struct {
	Contract *PauserRole // Generic contract binding to access the raw methods on
}

// PauserRoleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PauserRoleCallerRaw struct {
	Contract *PauserRoleCaller // Generic read-only contract binding to access the raw methods on
}

// PauserRoleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PauserRoleTransactorRaw struct {
	Contract *PauserRoleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPauserRole creates a new instance of PauserRole, bound to a specific deployed contract.
func NewPauserRole(address common.Address, backend bind.ContractBackend) (*PauserRole, error) {
	contract, err := bindPauserRole(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PauserRole{PauserRoleCaller: PauserRoleCaller{contract: contract}, PauserRoleTransactor: PauserRoleTransactor{contract: contract}, PauserRoleFilterer: PauserRoleFilterer{contract: contract}}, nil
}

// NewPauserRoleCaller creates a new read-only instance of PauserRole, bound to a specific deployed contract.
func NewPauserRoleCaller(address common.Address, caller bind.ContractCaller) (*PauserRoleCaller, error) {
	contract, err := bindPauserRole(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PauserRoleCaller{contract: contract}, nil
}

// NewPauserRoleTransactor creates a new write-only instance of PauserRole, bound to a specific deployed contract.
func NewPauserRoleTransactor(address common.Address, transactor bind.ContractTransactor) (*PauserRoleTransactor, error) {
	contract, err := bindPauserRole(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PauserRoleTransactor{contract: contract}, nil
}

// NewPauserRoleFilterer creates a new log filterer instance of PauserRole, bound to a specific deployed contract.
func NewPauserRoleFilterer(address common.Address, filterer bind.ContractFilterer) (*PauserRoleFilterer, error) {
	contract, err := bindPauserRole(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PauserRoleFilterer{contract: contract}, nil
}

// bindPauserRole binds a generic wrapper to an already deployed contract.
func bindPauserRole(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PauserRoleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PauserRole *PauserRoleRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PauserRole.Contract.PauserRoleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PauserRole *PauserRoleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PauserRole.Contract.PauserRoleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PauserRole *PauserRoleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PauserRole.Contract.PauserRoleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PauserRole *PauserRoleCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PauserRole.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PauserRole *PauserRoleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PauserRole.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PauserRole *PauserRoleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PauserRole.Contract.contract.Transact(opts, method, params...)
}

// IsPauser is a free data retrieval call binding the contract method 0x46fbf68e.
//
// Solidity: function isPauser(account address) constant returns(bool)
func (_PauserRole *PauserRoleCaller) IsPauser(opts *bind.CallOpts, account common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _PauserRole.contract.Call(opts, out, "isPauser", account)
	return *ret0, err
}

// IsPauser is a free data retrieval call binding the contract method 0x46fbf68e.
//
// Solidity: function isPauser(account address) constant returns(bool)
func (_PauserRole *PauserRoleSession) IsPauser(account common.Address) (bool, error) {
	return _PauserRole.Contract.IsPauser(&_PauserRole.CallOpts, account)
}

// IsPauser is a free data retrieval call binding the contract method 0x46fbf68e.
//
// Solidity: function isPauser(account address) constant returns(bool)
func (_PauserRole *PauserRoleCallerSession) IsPauser(account common.Address) (bool, error) {
	return _PauserRole.Contract.IsPauser(&_PauserRole.CallOpts, account)
}

// AddPauser is a paid mutator transaction binding the contract method 0x82dc1ec4.
//
// Solidity: function addPauser(account address) returns()
func (_PauserRole *PauserRoleTransactor) AddPauser(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _PauserRole.contract.Transact(opts, "addPauser", account)
}

// AddPauser is a paid mutator transaction binding the contract method 0x82dc1ec4.
//
// Solidity: function addPauser(account address) returns()
func (_PauserRole *PauserRoleSession) AddPauser(account common.Address) (*types.Transaction, error) {
	return _PauserRole.Contract.AddPauser(&_PauserRole.TransactOpts, account)
}

// AddPauser is a paid mutator transaction binding the contract method 0x82dc1ec4.
//
// Solidity: function addPauser(account address) returns()
func (_PauserRole *PauserRoleTransactorSession) AddPauser(account common.Address) (*types.Transaction, error) {
	return _PauserRole.Contract.AddPauser(&_PauserRole.TransactOpts, account)
}

// RenouncePauser is a paid mutator transaction binding the contract method 0x6ef8d66d.
//
// Solidity: function renouncePauser() returns()
func (_PauserRole *PauserRoleTransactor) RenouncePauser(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PauserRole.contract.Transact(opts, "renouncePauser")
}

// RenouncePauser is a paid mutator transaction binding the contract method 0x6ef8d66d.
//
// Solidity: function renouncePauser() returns()
func (_PauserRole *PauserRoleSession) RenouncePauser() (*types.Transaction, error) {
	return _PauserRole.Contract.RenouncePauser(&_PauserRole.TransactOpts)
}

// RenouncePauser is a paid mutator transaction binding the contract method 0x6ef8d66d.
//
// Solidity: function renouncePauser() returns()
func (_PauserRole *PauserRoleTransactorSession) RenouncePauser() (*types.Transaction, error) {
	return _PauserRole.Contract.RenouncePauser(&_PauserRole.TransactOpts)
}

// PauserRolePauserAddedIterator is returned from FilterPauserAdded and is used to iterate over the raw logs and unpacked data for PauserAdded events raised by the PauserRole contract.
type PauserRolePauserAddedIterator struct {
	Event *PauserRolePauserAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PauserRolePauserAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PauserRolePauserAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PauserRolePauserAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PauserRolePauserAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PauserRolePauserAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PauserRolePauserAdded represents a PauserAdded event raised by the PauserRole contract.
type PauserRolePauserAdded struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPauserAdded is a free log retrieval operation binding the contract event 0x6719d08c1888103bea251a4ed56406bd0c3e69723c8a1686e017e7bbe159b6f8.
//
// Solidity: e PauserAdded(account indexed address)
func (_PauserRole *PauserRoleFilterer) FilterPauserAdded(opts *bind.FilterOpts, account []common.Address) (*PauserRolePauserAddedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _PauserRole.contract.FilterLogs(opts, "PauserAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return &PauserRolePauserAddedIterator{contract: _PauserRole.contract, event: "PauserAdded", logs: logs, sub: sub}, nil
}

// WatchPauserAdded is a free log subscription operation binding the contract event 0x6719d08c1888103bea251a4ed56406bd0c3e69723c8a1686e017e7bbe159b6f8.
//
// Solidity: e PauserAdded(account indexed address)
func (_PauserRole *PauserRoleFilterer) WatchPauserAdded(opts *bind.WatchOpts, sink chan<- *PauserRolePauserAdded, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _PauserRole.contract.WatchLogs(opts, "PauserAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PauserRolePauserAdded)
				if err := _PauserRole.contract.UnpackLog(event, "PauserAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// PauserRolePauserRemovedIterator is returned from FilterPauserRemoved and is used to iterate over the raw logs and unpacked data for PauserRemoved events raised by the PauserRole contract.
type PauserRolePauserRemovedIterator struct {
	Event *PauserRolePauserRemoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PauserRolePauserRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PauserRolePauserRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PauserRolePauserRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PauserRolePauserRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PauserRolePauserRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PauserRolePauserRemoved represents a PauserRemoved event raised by the PauserRole contract.
type PauserRolePauserRemoved struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPauserRemoved is a free log retrieval operation binding the contract event 0xcd265ebaf09df2871cc7bd4133404a235ba12eff2041bb89d9c714a2621c7c7e.
//
// Solidity: e PauserRemoved(account indexed address)
func (_PauserRole *PauserRoleFilterer) FilterPauserRemoved(opts *bind.FilterOpts, account []common.Address) (*PauserRolePauserRemovedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _PauserRole.contract.FilterLogs(opts, "PauserRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return &PauserRolePauserRemovedIterator{contract: _PauserRole.contract, event: "PauserRemoved", logs: logs, sub: sub}, nil
}

// WatchPauserRemoved is a free log subscription operation binding the contract event 0xcd265ebaf09df2871cc7bd4133404a235ba12eff2041bb89d9c714a2621c7c7e.
//
// Solidity: e PauserRemoved(account indexed address)
func (_PauserRole *PauserRoleFilterer) WatchPauserRemoved(opts *bind.WatchOpts, sink chan<- *PauserRolePauserRemoved, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _PauserRole.contract.WatchLogs(opts, "PauserRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PauserRolePauserRemoved)
				if err := _PauserRole.contract.UnpackLog(event, "PauserRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// RolesABI is the input ABI used to generate the binding from.
const RolesABI = "[]"

// RolesBin is the compiled bytecode used for deploying new contracts.
const RolesBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea165627a7a7230582023201cbc65e103f65d457cd23ba27f6cb89e8bcff714dfe485ced426ffc29be20029`

// DeployRoles deploys a new Ethereum contract, binding an instance of Roles to it.
func DeployRoles(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Roles, error) {
	parsed, err := abi.JSON(strings.NewReader(RolesABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RolesBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Roles{RolesCaller: RolesCaller{contract: contract}, RolesTransactor: RolesTransactor{contract: contract}, RolesFilterer: RolesFilterer{contract: contract}}, nil
}

// Roles is an auto generated Go binding around an Ethereum contract.
type Roles struct {
	RolesCaller     // Read-only binding to the contract
	RolesTransactor // Write-only binding to the contract
	RolesFilterer   // Log filterer for contract events
}

// RolesCaller is an auto generated read-only Go binding around an Ethereum contract.
type RolesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RolesTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RolesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RolesFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RolesFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RolesSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RolesSession struct {
	Contract     *Roles            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RolesCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RolesCallerSession struct {
	Contract *RolesCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// RolesTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RolesTransactorSession struct {
	Contract     *RolesTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RolesRaw is an auto generated low-level Go binding around an Ethereum contract.
type RolesRaw struct {
	Contract *Roles // Generic contract binding to access the raw methods on
}

// RolesCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RolesCallerRaw struct {
	Contract *RolesCaller // Generic read-only contract binding to access the raw methods on
}

// RolesTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RolesTransactorRaw struct {
	Contract *RolesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRoles creates a new instance of Roles, bound to a specific deployed contract.
func NewRoles(address common.Address, backend bind.ContractBackend) (*Roles, error) {
	contract, err := bindRoles(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Roles{RolesCaller: RolesCaller{contract: contract}, RolesTransactor: RolesTransactor{contract: contract}, RolesFilterer: RolesFilterer{contract: contract}}, nil
}

// NewRolesCaller creates a new read-only instance of Roles, bound to a specific deployed contract.
func NewRolesCaller(address common.Address, caller bind.ContractCaller) (*RolesCaller, error) {
	contract, err := bindRoles(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RolesCaller{contract: contract}, nil
}

// NewRolesTransactor creates a new write-only instance of Roles, bound to a specific deployed contract.
func NewRolesTransactor(address common.Address, transactor bind.ContractTransactor) (*RolesTransactor, error) {
	contract, err := bindRoles(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RolesTransactor{contract: contract}, nil
}

// NewRolesFilterer creates a new log filterer instance of Roles, bound to a specific deployed contract.
func NewRolesFilterer(address common.Address, filterer bind.ContractFilterer) (*RolesFilterer, error) {
	contract, err := bindRoles(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RolesFilterer{contract: contract}, nil
}

// bindRoles binds a generic wrapper to an already deployed contract.
func bindRoles(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RolesABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Roles *RolesRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Roles.Contract.RolesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Roles *RolesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Roles.Contract.RolesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Roles *RolesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Roles.Contract.RolesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Roles *RolesCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Roles.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Roles *RolesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Roles.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Roles *RolesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Roles.Contract.contract.Transact(opts, method, params...)
}

// WhitelistAdminRoleABI is the input ABI used to generate the binding from.
const WhitelistAdminRoleABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"renounceWhitelistAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addWhitelistAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isWhitelistAdmin\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"account\",\"type\":\"address\"}],\"name\":\"WhitelistAdminAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"account\",\"type\":\"address\"}],\"name\":\"WhitelistAdminRemoved\",\"type\":\"event\"}]"

// WhitelistAdminRoleBin is the compiled bytecode used for deploying new contracts.
const WhitelistAdminRoleBin = `0x`

// DeployWhitelistAdminRole deploys a new Ethereum contract, binding an instance of WhitelistAdminRole to it.
func DeployWhitelistAdminRole(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *WhitelistAdminRole, error) {
	parsed, err := abi.JSON(strings.NewReader(WhitelistAdminRoleABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(WhitelistAdminRoleBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &WhitelistAdminRole{WhitelistAdminRoleCaller: WhitelistAdminRoleCaller{contract: contract}, WhitelistAdminRoleTransactor: WhitelistAdminRoleTransactor{contract: contract}, WhitelistAdminRoleFilterer: WhitelistAdminRoleFilterer{contract: contract}}, nil
}

// WhitelistAdminRole is an auto generated Go binding around an Ethereum contract.
type WhitelistAdminRole struct {
	WhitelistAdminRoleCaller     // Read-only binding to the contract
	WhitelistAdminRoleTransactor // Write-only binding to the contract
	WhitelistAdminRoleFilterer   // Log filterer for contract events
}

// WhitelistAdminRoleCaller is an auto generated read-only Go binding around an Ethereum contract.
type WhitelistAdminRoleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WhitelistAdminRoleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WhitelistAdminRoleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WhitelistAdminRoleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WhitelistAdminRoleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WhitelistAdminRoleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WhitelistAdminRoleSession struct {
	Contract     *WhitelistAdminRole // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// WhitelistAdminRoleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WhitelistAdminRoleCallerSession struct {
	Contract *WhitelistAdminRoleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// WhitelistAdminRoleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WhitelistAdminRoleTransactorSession struct {
	Contract     *WhitelistAdminRoleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// WhitelistAdminRoleRaw is an auto generated low-level Go binding around an Ethereum contract.
type WhitelistAdminRoleRaw struct {
	Contract *WhitelistAdminRole // Generic contract binding to access the raw methods on
}

// WhitelistAdminRoleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WhitelistAdminRoleCallerRaw struct {
	Contract *WhitelistAdminRoleCaller // Generic read-only contract binding to access the raw methods on
}

// WhitelistAdminRoleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WhitelistAdminRoleTransactorRaw struct {
	Contract *WhitelistAdminRoleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWhitelistAdminRole creates a new instance of WhitelistAdminRole, bound to a specific deployed contract.
func NewWhitelistAdminRole(address common.Address, backend bind.ContractBackend) (*WhitelistAdminRole, error) {
	contract, err := bindWhitelistAdminRole(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WhitelistAdminRole{WhitelistAdminRoleCaller: WhitelistAdminRoleCaller{contract: contract}, WhitelistAdminRoleTransactor: WhitelistAdminRoleTransactor{contract: contract}, WhitelistAdminRoleFilterer: WhitelistAdminRoleFilterer{contract: contract}}, nil
}

// NewWhitelistAdminRoleCaller creates a new read-only instance of WhitelistAdminRole, bound to a specific deployed contract.
func NewWhitelistAdminRoleCaller(address common.Address, caller bind.ContractCaller) (*WhitelistAdminRoleCaller, error) {
	contract, err := bindWhitelistAdminRole(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WhitelistAdminRoleCaller{contract: contract}, nil
}

// NewWhitelistAdminRoleTransactor creates a new write-only instance of WhitelistAdminRole, bound to a specific deployed contract.
func NewWhitelistAdminRoleTransactor(address common.Address, transactor bind.ContractTransactor) (*WhitelistAdminRoleTransactor, error) {
	contract, err := bindWhitelistAdminRole(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WhitelistAdminRoleTransactor{contract: contract}, nil
}

// NewWhitelistAdminRoleFilterer creates a new log filterer instance of WhitelistAdminRole, bound to a specific deployed contract.
func NewWhitelistAdminRoleFilterer(address common.Address, filterer bind.ContractFilterer) (*WhitelistAdminRoleFilterer, error) {
	contract, err := bindWhitelistAdminRole(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WhitelistAdminRoleFilterer{contract: contract}, nil
}

// bindWhitelistAdminRole binds a generic wrapper to an already deployed contract.
func bindWhitelistAdminRole(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(WhitelistAdminRoleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WhitelistAdminRole *WhitelistAdminRoleRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _WhitelistAdminRole.Contract.WhitelistAdminRoleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WhitelistAdminRole *WhitelistAdminRoleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WhitelistAdminRole.Contract.WhitelistAdminRoleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WhitelistAdminRole *WhitelistAdminRoleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WhitelistAdminRole.Contract.WhitelistAdminRoleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WhitelistAdminRole *WhitelistAdminRoleCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _WhitelistAdminRole.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WhitelistAdminRole *WhitelistAdminRoleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WhitelistAdminRole.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WhitelistAdminRole *WhitelistAdminRoleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WhitelistAdminRole.Contract.contract.Transact(opts, method, params...)
}

// IsWhitelistAdmin is a free data retrieval call binding the contract method 0xbb5f747b.
//
// Solidity: function isWhitelistAdmin(account address) constant returns(bool)
func (_WhitelistAdminRole *WhitelistAdminRoleCaller) IsWhitelistAdmin(opts *bind.CallOpts, account common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _WhitelistAdminRole.contract.Call(opts, out, "isWhitelistAdmin", account)
	return *ret0, err
}

// IsWhitelistAdmin is a free data retrieval call binding the contract method 0xbb5f747b.
//
// Solidity: function isWhitelistAdmin(account address) constant returns(bool)
func (_WhitelistAdminRole *WhitelistAdminRoleSession) IsWhitelistAdmin(account common.Address) (bool, error) {
	return _WhitelistAdminRole.Contract.IsWhitelistAdmin(&_WhitelistAdminRole.CallOpts, account)
}

// IsWhitelistAdmin is a free data retrieval call binding the contract method 0xbb5f747b.
//
// Solidity: function isWhitelistAdmin(account address) constant returns(bool)
func (_WhitelistAdminRole *WhitelistAdminRoleCallerSession) IsWhitelistAdmin(account common.Address) (bool, error) {
	return _WhitelistAdminRole.Contract.IsWhitelistAdmin(&_WhitelistAdminRole.CallOpts, account)
}

// AddWhitelistAdmin is a paid mutator transaction binding the contract method 0x7362d9c8.
//
// Solidity: function addWhitelistAdmin(account address) returns()
func (_WhitelistAdminRole *WhitelistAdminRoleTransactor) AddWhitelistAdmin(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _WhitelistAdminRole.contract.Transact(opts, "addWhitelistAdmin", account)
}

// AddWhitelistAdmin is a paid mutator transaction binding the contract method 0x7362d9c8.
//
// Solidity: function addWhitelistAdmin(account address) returns()
func (_WhitelistAdminRole *WhitelistAdminRoleSession) AddWhitelistAdmin(account common.Address) (*types.Transaction, error) {
	return _WhitelistAdminRole.Contract.AddWhitelistAdmin(&_WhitelistAdminRole.TransactOpts, account)
}

// AddWhitelistAdmin is a paid mutator transaction binding the contract method 0x7362d9c8.
//
// Solidity: function addWhitelistAdmin(account address) returns()
func (_WhitelistAdminRole *WhitelistAdminRoleTransactorSession) AddWhitelistAdmin(account common.Address) (*types.Transaction, error) {
	return _WhitelistAdminRole.Contract.AddWhitelistAdmin(&_WhitelistAdminRole.TransactOpts, account)
}

// RenounceWhitelistAdmin is a paid mutator transaction binding the contract method 0x4c5a628c.
//
// Solidity: function renounceWhitelistAdmin() returns()
func (_WhitelistAdminRole *WhitelistAdminRoleTransactor) RenounceWhitelistAdmin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WhitelistAdminRole.contract.Transact(opts, "renounceWhitelistAdmin")
}

// RenounceWhitelistAdmin is a paid mutator transaction binding the contract method 0x4c5a628c.
//
// Solidity: function renounceWhitelistAdmin() returns()
func (_WhitelistAdminRole *WhitelistAdminRoleSession) RenounceWhitelistAdmin() (*types.Transaction, error) {
	return _WhitelistAdminRole.Contract.RenounceWhitelistAdmin(&_WhitelistAdminRole.TransactOpts)
}

// RenounceWhitelistAdmin is a paid mutator transaction binding the contract method 0x4c5a628c.
//
// Solidity: function renounceWhitelistAdmin() returns()
func (_WhitelistAdminRole *WhitelistAdminRoleTransactorSession) RenounceWhitelistAdmin() (*types.Transaction, error) {
	return _WhitelistAdminRole.Contract.RenounceWhitelistAdmin(&_WhitelistAdminRole.TransactOpts)
}

// WhitelistAdminRoleWhitelistAdminAddedIterator is returned from FilterWhitelistAdminAdded and is used to iterate over the raw logs and unpacked data for WhitelistAdminAdded events raised by the WhitelistAdminRole contract.
type WhitelistAdminRoleWhitelistAdminAddedIterator struct {
	Event *WhitelistAdminRoleWhitelistAdminAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WhitelistAdminRoleWhitelistAdminAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WhitelistAdminRoleWhitelistAdminAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WhitelistAdminRoleWhitelistAdminAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WhitelistAdminRoleWhitelistAdminAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WhitelistAdminRoleWhitelistAdminAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WhitelistAdminRoleWhitelistAdminAdded represents a WhitelistAdminAdded event raised by the WhitelistAdminRole contract.
type WhitelistAdminRoleWhitelistAdminAdded struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWhitelistAdminAdded is a free log retrieval operation binding the contract event 0x22380c05984257a1cb900161c713dd71d39e74820f1aea43bd3f1bdd20961299.
//
// Solidity: e WhitelistAdminAdded(account indexed address)
func (_WhitelistAdminRole *WhitelistAdminRoleFilterer) FilterWhitelistAdminAdded(opts *bind.FilterOpts, account []common.Address) (*WhitelistAdminRoleWhitelistAdminAddedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _WhitelistAdminRole.contract.FilterLogs(opts, "WhitelistAdminAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return &WhitelistAdminRoleWhitelistAdminAddedIterator{contract: _WhitelistAdminRole.contract, event: "WhitelistAdminAdded", logs: logs, sub: sub}, nil
}

// WatchWhitelistAdminAdded is a free log subscription operation binding the contract event 0x22380c05984257a1cb900161c713dd71d39e74820f1aea43bd3f1bdd20961299.
//
// Solidity: e WhitelistAdminAdded(account indexed address)
func (_WhitelistAdminRole *WhitelistAdminRoleFilterer) WatchWhitelistAdminAdded(opts *bind.WatchOpts, sink chan<- *WhitelistAdminRoleWhitelistAdminAdded, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _WhitelistAdminRole.contract.WatchLogs(opts, "WhitelistAdminAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WhitelistAdminRoleWhitelistAdminAdded)
				if err := _WhitelistAdminRole.contract.UnpackLog(event, "WhitelistAdminAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// WhitelistAdminRoleWhitelistAdminRemovedIterator is returned from FilterWhitelistAdminRemoved and is used to iterate over the raw logs and unpacked data for WhitelistAdminRemoved events raised by the WhitelistAdminRole contract.
type WhitelistAdminRoleWhitelistAdminRemovedIterator struct {
	Event *WhitelistAdminRoleWhitelistAdminRemoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WhitelistAdminRoleWhitelistAdminRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WhitelistAdminRoleWhitelistAdminRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WhitelistAdminRoleWhitelistAdminRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WhitelistAdminRoleWhitelistAdminRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WhitelistAdminRoleWhitelistAdminRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WhitelistAdminRoleWhitelistAdminRemoved represents a WhitelistAdminRemoved event raised by the WhitelistAdminRole contract.
type WhitelistAdminRoleWhitelistAdminRemoved struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWhitelistAdminRemoved is a free log retrieval operation binding the contract event 0x0a8eb35e5ca14b3d6f28e4abf2f128dbab231a58b56e89beb5d636115001e165.
//
// Solidity: e WhitelistAdminRemoved(account indexed address)
func (_WhitelistAdminRole *WhitelistAdminRoleFilterer) FilterWhitelistAdminRemoved(opts *bind.FilterOpts, account []common.Address) (*WhitelistAdminRoleWhitelistAdminRemovedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _WhitelistAdminRole.contract.FilterLogs(opts, "WhitelistAdminRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return &WhitelistAdminRoleWhitelistAdminRemovedIterator{contract: _WhitelistAdminRole.contract, event: "WhitelistAdminRemoved", logs: logs, sub: sub}, nil
}

// WatchWhitelistAdminRemoved is a free log subscription operation binding the contract event 0x0a8eb35e5ca14b3d6f28e4abf2f128dbab231a58b56e89beb5d636115001e165.
//
// Solidity: e WhitelistAdminRemoved(account indexed address)
func (_WhitelistAdminRole *WhitelistAdminRoleFilterer) WatchWhitelistAdminRemoved(opts *bind.WatchOpts, sink chan<- *WhitelistAdminRoleWhitelistAdminRemoved, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _WhitelistAdminRole.contract.WatchLogs(opts, "WhitelistAdminRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WhitelistAdminRoleWhitelistAdminRemoved)
				if err := _WhitelistAdminRole.contract.UnpackLog(event, "WhitelistAdminRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// WhitelistedRoleABI is the input ABI used to generate the binding from.
const WhitelistedRoleABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addWhitelisted\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"removeWhitelisted\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isWhitelisted\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceWhitelistAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addWhitelistAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isWhitelistAdmin\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceWhitelisted\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"account\",\"type\":\"address\"}],\"name\":\"WhitelistedAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"account\",\"type\":\"address\"}],\"name\":\"WhitelistedRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"account\",\"type\":\"address\"}],\"name\":\"WhitelistAdminAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"account\",\"type\":\"address\"}],\"name\":\"WhitelistAdminRemoved\",\"type\":\"event\"}]"

// WhitelistedRoleBin is the compiled bytecode used for deploying new contracts.
const WhitelistedRoleBin = `0x60806040526100163364010000000061001b810204565b6100f8565b61003360008264010000000061037361006a82021704565b604051600160a060020a038216907f22380c05984257a1cb900161c713dd71d39e74820f1aea43bd3f1bdd2096129990600090a250565b600160a060020a038116151561007f57600080fd5b61009282826401000000006100c1810204565b1561009c57600080fd5b600160a060020a0316600090815260209190915260409020805460ff19166001179055565b6000600160a060020a03821615156100d857600080fd5b50600160a060020a03166000908152602091909152604090205460ff1690565b610439806101076000396000f3fe608060405234801561001057600080fd5b506004361061009a576000357c0100000000000000000000000000000000000000000000000000000000900480634c5a628c116100785780634c5a628c146101275780637362d9c81461012f578063bb5f747b14610155578063d6cd94731461017b5761009a565b806310154bad1461009f578063291d9549146100c75780633af32abf146100ed575b600080fd5b6100c5600480360360208110156100b557600080fd5b5035600160a060020a0316610183565b005b6100c5600480360360208110156100dd57600080fd5b5035600160a060020a03166101a3565b6101136004803603602081101561010357600080fd5b5035600160a060020a03166101c0565b604080519115158252519081900360200190f35b6100c56101d9565b6100c56004803603602081101561014557600080fd5b5035600160a060020a03166101e4565b6101136004803603602081101561016b57600080fd5b5035600160a060020a0316610201565b6100c5610213565b61018c33610201565b151561019757600080fd5b6101a08161021c565b50565b6101ac33610201565b15156101b757600080fd5b6101a081610264565b60006101d360018363ffffffff6102ac16565b92915050565b6101e2336102e3565b565b6101ed33610201565b15156101f857600080fd5b6101a08161032b565b60006101d3818363ffffffff6102ac16565b6101e233610264565b61022d60018263ffffffff61037316565b604051600160a060020a038216907fee1504a83b6d4a361f4c1dc78ab59bfa30d6a3b6612c403e86bb01ef2984295f90600090a250565b61027560018263ffffffff6103c116565b604051600160a060020a038216907f270d9b30cf5b0793bbfd54c9d5b94aeb49462b8148399000265144a8722da6b690600090a250565b6000600160a060020a03821615156102c357600080fd5b50600160a060020a03166000908152602091909152604090205460ff1690565b6102f460008263ffffffff6103c116565b604051600160a060020a038216907f0a8eb35e5ca14b3d6f28e4abf2f128dbab231a58b56e89beb5d636115001e16590600090a250565b61033c60008263ffffffff61037316565b604051600160a060020a038216907f22380c05984257a1cb900161c713dd71d39e74820f1aea43bd3f1bdd2096129990600090a250565b600160a060020a038116151561038857600080fd5b61039282826102ac565b1561039c57600080fd5b600160a060020a0316600090815260209190915260409020805460ff19166001179055565b600160a060020a03811615156103d657600080fd5b6103e082826102ac565b15156103eb57600080fd5b600160a060020a0316600090815260209190915260409020805460ff1916905556fea165627a7a723058207ddf7547708a1db663b1f5b8fc05e3bca4910d621fe17b7bc11f69703a114a6c0029`

// DeployWhitelistedRole deploys a new Ethereum contract, binding an instance of WhitelistedRole to it.
func DeployWhitelistedRole(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *WhitelistedRole, error) {
	parsed, err := abi.JSON(strings.NewReader(WhitelistedRoleABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(WhitelistedRoleBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &WhitelistedRole{WhitelistedRoleCaller: WhitelistedRoleCaller{contract: contract}, WhitelistedRoleTransactor: WhitelistedRoleTransactor{contract: contract}, WhitelistedRoleFilterer: WhitelistedRoleFilterer{contract: contract}}, nil
}

// WhitelistedRole is an auto generated Go binding around an Ethereum contract.
type WhitelistedRole struct {
	WhitelistedRoleCaller     // Read-only binding to the contract
	WhitelistedRoleTransactor // Write-only binding to the contract
	WhitelistedRoleFilterer   // Log filterer for contract events
}

// WhitelistedRoleCaller is an auto generated read-only Go binding around an Ethereum contract.
type WhitelistedRoleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WhitelistedRoleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WhitelistedRoleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WhitelistedRoleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WhitelistedRoleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WhitelistedRoleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WhitelistedRoleSession struct {
	Contract     *WhitelistedRole  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WhitelistedRoleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WhitelistedRoleCallerSession struct {
	Contract *WhitelistedRoleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// WhitelistedRoleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WhitelistedRoleTransactorSession struct {
	Contract     *WhitelistedRoleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// WhitelistedRoleRaw is an auto generated low-level Go binding around an Ethereum contract.
type WhitelistedRoleRaw struct {
	Contract *WhitelistedRole // Generic contract binding to access the raw methods on
}

// WhitelistedRoleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WhitelistedRoleCallerRaw struct {
	Contract *WhitelistedRoleCaller // Generic read-only contract binding to access the raw methods on
}

// WhitelistedRoleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WhitelistedRoleTransactorRaw struct {
	Contract *WhitelistedRoleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWhitelistedRole creates a new instance of WhitelistedRole, bound to a specific deployed contract.
func NewWhitelistedRole(address common.Address, backend bind.ContractBackend) (*WhitelistedRole, error) {
	contract, err := bindWhitelistedRole(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WhitelistedRole{WhitelistedRoleCaller: WhitelistedRoleCaller{contract: contract}, WhitelistedRoleTransactor: WhitelistedRoleTransactor{contract: contract}, WhitelistedRoleFilterer: WhitelistedRoleFilterer{contract: contract}}, nil
}

// NewWhitelistedRoleCaller creates a new read-only instance of WhitelistedRole, bound to a specific deployed contract.
func NewWhitelistedRoleCaller(address common.Address, caller bind.ContractCaller) (*WhitelistedRoleCaller, error) {
	contract, err := bindWhitelistedRole(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WhitelistedRoleCaller{contract: contract}, nil
}

// NewWhitelistedRoleTransactor creates a new write-only instance of WhitelistedRole, bound to a specific deployed contract.
func NewWhitelistedRoleTransactor(address common.Address, transactor bind.ContractTransactor) (*WhitelistedRoleTransactor, error) {
	contract, err := bindWhitelistedRole(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WhitelistedRoleTransactor{contract: contract}, nil
}

// NewWhitelistedRoleFilterer creates a new log filterer instance of WhitelistedRole, bound to a specific deployed contract.
func NewWhitelistedRoleFilterer(address common.Address, filterer bind.ContractFilterer) (*WhitelistedRoleFilterer, error) {
	contract, err := bindWhitelistedRole(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WhitelistedRoleFilterer{contract: contract}, nil
}

// bindWhitelistedRole binds a generic wrapper to an already deployed contract.
func bindWhitelistedRole(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(WhitelistedRoleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WhitelistedRole *WhitelistedRoleRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _WhitelistedRole.Contract.WhitelistedRoleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WhitelistedRole *WhitelistedRoleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WhitelistedRole.Contract.WhitelistedRoleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WhitelistedRole *WhitelistedRoleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WhitelistedRole.Contract.WhitelistedRoleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WhitelistedRole *WhitelistedRoleCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _WhitelistedRole.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WhitelistedRole *WhitelistedRoleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WhitelistedRole.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WhitelistedRole *WhitelistedRoleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WhitelistedRole.Contract.contract.Transact(opts, method, params...)
}

// IsWhitelistAdmin is a free data retrieval call binding the contract method 0xbb5f747b.
//
// Solidity: function isWhitelistAdmin(account address) constant returns(bool)
func (_WhitelistedRole *WhitelistedRoleCaller) IsWhitelistAdmin(opts *bind.CallOpts, account common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _WhitelistedRole.contract.Call(opts, out, "isWhitelistAdmin", account)
	return *ret0, err
}

// IsWhitelistAdmin is a free data retrieval call binding the contract method 0xbb5f747b.
//
// Solidity: function isWhitelistAdmin(account address) constant returns(bool)
func (_WhitelistedRole *WhitelistedRoleSession) IsWhitelistAdmin(account common.Address) (bool, error) {
	return _WhitelistedRole.Contract.IsWhitelistAdmin(&_WhitelistedRole.CallOpts, account)
}

// IsWhitelistAdmin is a free data retrieval call binding the contract method 0xbb5f747b.
//
// Solidity: function isWhitelistAdmin(account address) constant returns(bool)
func (_WhitelistedRole *WhitelistedRoleCallerSession) IsWhitelistAdmin(account common.Address) (bool, error) {
	return _WhitelistedRole.Contract.IsWhitelistAdmin(&_WhitelistedRole.CallOpts, account)
}

// IsWhitelisted is a free data retrieval call binding the contract method 0x3af32abf.
//
// Solidity: function isWhitelisted(account address) constant returns(bool)
func (_WhitelistedRole *WhitelistedRoleCaller) IsWhitelisted(opts *bind.CallOpts, account common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _WhitelistedRole.contract.Call(opts, out, "isWhitelisted", account)
	return *ret0, err
}

// IsWhitelisted is a free data retrieval call binding the contract method 0x3af32abf.
//
// Solidity: function isWhitelisted(account address) constant returns(bool)
func (_WhitelistedRole *WhitelistedRoleSession) IsWhitelisted(account common.Address) (bool, error) {
	return _WhitelistedRole.Contract.IsWhitelisted(&_WhitelistedRole.CallOpts, account)
}

// IsWhitelisted is a free data retrieval call binding the contract method 0x3af32abf.
//
// Solidity: function isWhitelisted(account address) constant returns(bool)
func (_WhitelistedRole *WhitelistedRoleCallerSession) IsWhitelisted(account common.Address) (bool, error) {
	return _WhitelistedRole.Contract.IsWhitelisted(&_WhitelistedRole.CallOpts, account)
}

// AddWhitelistAdmin is a paid mutator transaction binding the contract method 0x7362d9c8.
//
// Solidity: function addWhitelistAdmin(account address) returns()
func (_WhitelistedRole *WhitelistedRoleTransactor) AddWhitelistAdmin(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _WhitelistedRole.contract.Transact(opts, "addWhitelistAdmin", account)
}

// AddWhitelistAdmin is a paid mutator transaction binding the contract method 0x7362d9c8.
//
// Solidity: function addWhitelistAdmin(account address) returns()
func (_WhitelistedRole *WhitelistedRoleSession) AddWhitelistAdmin(account common.Address) (*types.Transaction, error) {
	return _WhitelistedRole.Contract.AddWhitelistAdmin(&_WhitelistedRole.TransactOpts, account)
}

// AddWhitelistAdmin is a paid mutator transaction binding the contract method 0x7362d9c8.
//
// Solidity: function addWhitelistAdmin(account address) returns()
func (_WhitelistedRole *WhitelistedRoleTransactorSession) AddWhitelistAdmin(account common.Address) (*types.Transaction, error) {
	return _WhitelistedRole.Contract.AddWhitelistAdmin(&_WhitelistedRole.TransactOpts, account)
}

// AddWhitelisted is a paid mutator transaction binding the contract method 0x10154bad.
//
// Solidity: function addWhitelisted(account address) returns()
func (_WhitelistedRole *WhitelistedRoleTransactor) AddWhitelisted(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _WhitelistedRole.contract.Transact(opts, "addWhitelisted", account)
}

// AddWhitelisted is a paid mutator transaction binding the contract method 0x10154bad.
//
// Solidity: function addWhitelisted(account address) returns()
func (_WhitelistedRole *WhitelistedRoleSession) AddWhitelisted(account common.Address) (*types.Transaction, error) {
	return _WhitelistedRole.Contract.AddWhitelisted(&_WhitelistedRole.TransactOpts, account)
}

// AddWhitelisted is a paid mutator transaction binding the contract method 0x10154bad.
//
// Solidity: function addWhitelisted(account address) returns()
func (_WhitelistedRole *WhitelistedRoleTransactorSession) AddWhitelisted(account common.Address) (*types.Transaction, error) {
	return _WhitelistedRole.Contract.AddWhitelisted(&_WhitelistedRole.TransactOpts, account)
}

// RemoveWhitelisted is a paid mutator transaction binding the contract method 0x291d9549.
//
// Solidity: function removeWhitelisted(account address) returns()
func (_WhitelistedRole *WhitelistedRoleTransactor) RemoveWhitelisted(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _WhitelistedRole.contract.Transact(opts, "removeWhitelisted", account)
}

// RemoveWhitelisted is a paid mutator transaction binding the contract method 0x291d9549.
//
// Solidity: function removeWhitelisted(account address) returns()
func (_WhitelistedRole *WhitelistedRoleSession) RemoveWhitelisted(account common.Address) (*types.Transaction, error) {
	return _WhitelistedRole.Contract.RemoveWhitelisted(&_WhitelistedRole.TransactOpts, account)
}

// RemoveWhitelisted is a paid mutator transaction binding the contract method 0x291d9549.
//
// Solidity: function removeWhitelisted(account address) returns()
func (_WhitelistedRole *WhitelistedRoleTransactorSession) RemoveWhitelisted(account common.Address) (*types.Transaction, error) {
	return _WhitelistedRole.Contract.RemoveWhitelisted(&_WhitelistedRole.TransactOpts, account)
}

// RenounceWhitelistAdmin is a paid mutator transaction binding the contract method 0x4c5a628c.
//
// Solidity: function renounceWhitelistAdmin() returns()
func (_WhitelistedRole *WhitelistedRoleTransactor) RenounceWhitelistAdmin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WhitelistedRole.contract.Transact(opts, "renounceWhitelistAdmin")
}

// RenounceWhitelistAdmin is a paid mutator transaction binding the contract method 0x4c5a628c.
//
// Solidity: function renounceWhitelistAdmin() returns()
func (_WhitelistedRole *WhitelistedRoleSession) RenounceWhitelistAdmin() (*types.Transaction, error) {
	return _WhitelistedRole.Contract.RenounceWhitelistAdmin(&_WhitelistedRole.TransactOpts)
}

// RenounceWhitelistAdmin is a paid mutator transaction binding the contract method 0x4c5a628c.
//
// Solidity: function renounceWhitelistAdmin() returns()
func (_WhitelistedRole *WhitelistedRoleTransactorSession) RenounceWhitelistAdmin() (*types.Transaction, error) {
	return _WhitelistedRole.Contract.RenounceWhitelistAdmin(&_WhitelistedRole.TransactOpts)
}

// RenounceWhitelisted is a paid mutator transaction binding the contract method 0xd6cd9473.
//
// Solidity: function renounceWhitelisted() returns()
func (_WhitelistedRole *WhitelistedRoleTransactor) RenounceWhitelisted(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WhitelistedRole.contract.Transact(opts, "renounceWhitelisted")
}

// RenounceWhitelisted is a paid mutator transaction binding the contract method 0xd6cd9473.
//
// Solidity: function renounceWhitelisted() returns()
func (_WhitelistedRole *WhitelistedRoleSession) RenounceWhitelisted() (*types.Transaction, error) {
	return _WhitelistedRole.Contract.RenounceWhitelisted(&_WhitelistedRole.TransactOpts)
}

// RenounceWhitelisted is a paid mutator transaction binding the contract method 0xd6cd9473.
//
// Solidity: function renounceWhitelisted() returns()
func (_WhitelistedRole *WhitelistedRoleTransactorSession) RenounceWhitelisted() (*types.Transaction, error) {
	return _WhitelistedRole.Contract.RenounceWhitelisted(&_WhitelistedRole.TransactOpts)
}

// WhitelistedRoleWhitelistAdminAddedIterator is returned from FilterWhitelistAdminAdded and is used to iterate over the raw logs and unpacked data for WhitelistAdminAdded events raised by the WhitelistedRole contract.
type WhitelistedRoleWhitelistAdminAddedIterator struct {
	Event *WhitelistedRoleWhitelistAdminAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WhitelistedRoleWhitelistAdminAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WhitelistedRoleWhitelistAdminAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WhitelistedRoleWhitelistAdminAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WhitelistedRoleWhitelistAdminAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WhitelistedRoleWhitelistAdminAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WhitelistedRoleWhitelistAdminAdded represents a WhitelistAdminAdded event raised by the WhitelistedRole contract.
type WhitelistedRoleWhitelistAdminAdded struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWhitelistAdminAdded is a free log retrieval operation binding the contract event 0x22380c05984257a1cb900161c713dd71d39e74820f1aea43bd3f1bdd20961299.
//
// Solidity: e WhitelistAdminAdded(account indexed address)
func (_WhitelistedRole *WhitelistedRoleFilterer) FilterWhitelistAdminAdded(opts *bind.FilterOpts, account []common.Address) (*WhitelistedRoleWhitelistAdminAddedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _WhitelistedRole.contract.FilterLogs(opts, "WhitelistAdminAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return &WhitelistedRoleWhitelistAdminAddedIterator{contract: _WhitelistedRole.contract, event: "WhitelistAdminAdded", logs: logs, sub: sub}, nil
}

// WatchWhitelistAdminAdded is a free log subscription operation binding the contract event 0x22380c05984257a1cb900161c713dd71d39e74820f1aea43bd3f1bdd20961299.
//
// Solidity: e WhitelistAdminAdded(account indexed address)
func (_WhitelistedRole *WhitelistedRoleFilterer) WatchWhitelistAdminAdded(opts *bind.WatchOpts, sink chan<- *WhitelistedRoleWhitelistAdminAdded, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _WhitelistedRole.contract.WatchLogs(opts, "WhitelistAdminAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WhitelistedRoleWhitelistAdminAdded)
				if err := _WhitelistedRole.contract.UnpackLog(event, "WhitelistAdminAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// WhitelistedRoleWhitelistAdminRemovedIterator is returned from FilterWhitelistAdminRemoved and is used to iterate over the raw logs and unpacked data for WhitelistAdminRemoved events raised by the WhitelistedRole contract.
type WhitelistedRoleWhitelistAdminRemovedIterator struct {
	Event *WhitelistedRoleWhitelistAdminRemoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WhitelistedRoleWhitelistAdminRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WhitelistedRoleWhitelistAdminRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WhitelistedRoleWhitelistAdminRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WhitelistedRoleWhitelistAdminRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WhitelistedRoleWhitelistAdminRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WhitelistedRoleWhitelistAdminRemoved represents a WhitelistAdminRemoved event raised by the WhitelistedRole contract.
type WhitelistedRoleWhitelistAdminRemoved struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWhitelistAdminRemoved is a free log retrieval operation binding the contract event 0x0a8eb35e5ca14b3d6f28e4abf2f128dbab231a58b56e89beb5d636115001e165.
//
// Solidity: e WhitelistAdminRemoved(account indexed address)
func (_WhitelistedRole *WhitelistedRoleFilterer) FilterWhitelistAdminRemoved(opts *bind.FilterOpts, account []common.Address) (*WhitelistedRoleWhitelistAdminRemovedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _WhitelistedRole.contract.FilterLogs(opts, "WhitelistAdminRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return &WhitelistedRoleWhitelistAdminRemovedIterator{contract: _WhitelistedRole.contract, event: "WhitelistAdminRemoved", logs: logs, sub: sub}, nil
}

// WatchWhitelistAdminRemoved is a free log subscription operation binding the contract event 0x0a8eb35e5ca14b3d6f28e4abf2f128dbab231a58b56e89beb5d636115001e165.
//
// Solidity: e WhitelistAdminRemoved(account indexed address)
func (_WhitelistedRole *WhitelistedRoleFilterer) WatchWhitelistAdminRemoved(opts *bind.WatchOpts, sink chan<- *WhitelistedRoleWhitelistAdminRemoved, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _WhitelistedRole.contract.WatchLogs(opts, "WhitelistAdminRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WhitelistedRoleWhitelistAdminRemoved)
				if err := _WhitelistedRole.contract.UnpackLog(event, "WhitelistAdminRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// WhitelistedRoleWhitelistedAddedIterator is returned from FilterWhitelistedAdded and is used to iterate over the raw logs and unpacked data for WhitelistedAdded events raised by the WhitelistedRole contract.
type WhitelistedRoleWhitelistedAddedIterator struct {
	Event *WhitelistedRoleWhitelistedAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WhitelistedRoleWhitelistedAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WhitelistedRoleWhitelistedAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WhitelistedRoleWhitelistedAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WhitelistedRoleWhitelistedAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WhitelistedRoleWhitelistedAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WhitelistedRoleWhitelistedAdded represents a WhitelistedAdded event raised by the WhitelistedRole contract.
type WhitelistedRoleWhitelistedAdded struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWhitelistedAdded is a free log retrieval operation binding the contract event 0xee1504a83b6d4a361f4c1dc78ab59bfa30d6a3b6612c403e86bb01ef2984295f.
//
// Solidity: e WhitelistedAdded(account indexed address)
func (_WhitelistedRole *WhitelistedRoleFilterer) FilterWhitelistedAdded(opts *bind.FilterOpts, account []common.Address) (*WhitelistedRoleWhitelistedAddedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _WhitelistedRole.contract.FilterLogs(opts, "WhitelistedAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return &WhitelistedRoleWhitelistedAddedIterator{contract: _WhitelistedRole.contract, event: "WhitelistedAdded", logs: logs, sub: sub}, nil
}

// WatchWhitelistedAdded is a free log subscription operation binding the contract event 0xee1504a83b6d4a361f4c1dc78ab59bfa30d6a3b6612c403e86bb01ef2984295f.
//
// Solidity: e WhitelistedAdded(account indexed address)
func (_WhitelistedRole *WhitelistedRoleFilterer) WatchWhitelistedAdded(opts *bind.WatchOpts, sink chan<- *WhitelistedRoleWhitelistedAdded, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _WhitelistedRole.contract.WatchLogs(opts, "WhitelistedAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WhitelistedRoleWhitelistedAdded)
				if err := _WhitelistedRole.contract.UnpackLog(event, "WhitelistedAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// WhitelistedRoleWhitelistedRemovedIterator is returned from FilterWhitelistedRemoved and is used to iterate over the raw logs and unpacked data for WhitelistedRemoved events raised by the WhitelistedRole contract.
type WhitelistedRoleWhitelistedRemovedIterator struct {
	Event *WhitelistedRoleWhitelistedRemoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WhitelistedRoleWhitelistedRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WhitelistedRoleWhitelistedRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WhitelistedRoleWhitelistedRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WhitelistedRoleWhitelistedRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WhitelistedRoleWhitelistedRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WhitelistedRoleWhitelistedRemoved represents a WhitelistedRemoved event raised by the WhitelistedRole contract.
type WhitelistedRoleWhitelistedRemoved struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWhitelistedRemoved is a free log retrieval operation binding the contract event 0x270d9b30cf5b0793bbfd54c9d5b94aeb49462b8148399000265144a8722da6b6.
//
// Solidity: e WhitelistedRemoved(account indexed address)
func (_WhitelistedRole *WhitelistedRoleFilterer) FilterWhitelistedRemoved(opts *bind.FilterOpts, account []common.Address) (*WhitelistedRoleWhitelistedRemovedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _WhitelistedRole.contract.FilterLogs(opts, "WhitelistedRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return &WhitelistedRoleWhitelistedRemovedIterator{contract: _WhitelistedRole.contract, event: "WhitelistedRemoved", logs: logs, sub: sub}, nil
}

// WatchWhitelistedRemoved is a free log subscription operation binding the contract event 0x270d9b30cf5b0793bbfd54c9d5b94aeb49462b8148399000265144a8722da6b6.
//
// Solidity: e WhitelistedRemoved(account indexed address)
func (_WhitelistedRole *WhitelistedRoleFilterer) WatchWhitelistedRemoved(opts *bind.WatchOpts, sink chan<- *WhitelistedRoleWhitelistedRemoved, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _WhitelistedRole.contract.WatchLogs(opts, "WhitelistedRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WhitelistedRoleWhitelistedRemoved)
				if err := _WhitelistedRole.contract.UnpackLog(event, "WhitelistedRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

