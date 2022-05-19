// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// NiftyconnectexchangeMetaData contains all meta data concerning the Niftyconnectexchange contract.
var NiftyconnectexchangeMetaData = &bind.MetaData{
	ABI: "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"governor\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tokenTransferProxy\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"exchangeFeeRate\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"target\",\"type\":\"address\"},{\"name\":\"calldata\",\"type\":\"bytes\"},{\"name\":\"extradata\",\"type\":\"bytes\"}],\"name\":\"staticCall\",\"outputs\":[{\"name\":\"result\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"acceptGovernance\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newTakerRelayerFeeShare\",\"type\":\"uint256\"},{\"name\":\"newMakerRelayerFeeShare\",\"type\":\"uint256\"},{\"name\":\"newProtocolFeeShare\",\"type\":\"uint256\"}],\"name\":\"changeTakerRelayerFeeShare\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"royaltyRegisterHub\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"takerRelayerFeeShare\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newProtocolFeeRecipient\",\"type\":\"address\"}],\"name\":\"changeProtocolFeeRecipient\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"incrementNonce\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"protocolFeeRecipient\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"cancelledOrFinalized\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"protocolFeeShare\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newExchangeFeeRate\",\"type\":\"uint256\"}],\"name\":\"changeExchangeFeeRate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"INVERSE_BASIS_POINT\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pendingGovernor\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MAXIMUM_EXCHANGE_RATE\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"approvedOrders\",\"outputs\":[{\"name\":\"approved\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"makerRelayerFeeShare\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"merkleValidatorContract\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"pendingGovernor_\",\"type\":\"address\"}],\"name\":\"setPendingGovernor\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"tokenTransferProxyAddress\",\"type\":\"address\"},{\"name\":\"protocolFeeAddress\",\"type\":\"address\"},{\"name\":\"merkleValidatorAddress\",\"type\":\"address\"},{\"name\":\"royaltyRegisterHubAddress\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"hash\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"exchange\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"maker\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"taker\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"makerRelayerFeeRecipient\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"side\",\"type\":\"uint8\"},{\"indexed\":false,\"name\":\"saleKind\",\"type\":\"uint8\"},{\"indexed\":false,\"name\":\"nftAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"ipfsHash\",\"type\":\"bytes32\"}],\"name\":\"OrderApprovedPartOne\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"hash\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"calldata\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"replacementPattern\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"staticTarget\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"staticExtradata\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"paymentToken\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"basePrice\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"extra\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"listingTime\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"expirationTime\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"salt\",\"type\":\"uint256\"}],\"name\":\"OrderApprovedPartTwo\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"OrderCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"buyHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"sellHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"maker\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"taker\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"makerRelayerFeeRecipient\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"takerRelayerFeeRecipient\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"metadata\",\"type\":\"bytes32\"}],\"name\":\"OrdersMatched\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"maker\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"newNonce\",\"type\":\"uint256\"}],\"name\":\"NonceIncremented\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousGovernor\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newGovernor\",\"type\":\"address\"}],\"name\":\"GovernanceTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"newPendingGovernor\",\"type\":\"address\"}],\"name\":\"NewPendingGovernor\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"name\":\"selector\",\"type\":\"uint256\"},{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"nftAddress\",\"type\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"merkleRoot\",\"type\":\"bytes32\"},{\"name\":\"merkleProof\",\"type\":\"bytes32[]\"}],\"name\":\"buildCallData\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"array\",\"type\":\"bytes\"},{\"name\":\"desired\",\"type\":\"bytes\"},{\"name\":\"mask\",\"type\":\"bytes\"}],\"name\":\"guardedArrayReplace\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"side\",\"type\":\"uint8\"},{\"name\":\"saleKind\",\"type\":\"uint8\"},{\"name\":\"basePrice\",\"type\":\"uint256\"},{\"name\":\"extra\",\"type\":\"uint256\"},{\"name\":\"listingTime\",\"type\":\"uint256\"},{\"name\":\"expirationTime\",\"type\":\"uint256\"}],\"name\":\"calculateFinalPrice\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addrs\",\"type\":\"address[9]\"},{\"name\":\"uints\",\"type\":\"uint256[9]\"},{\"name\":\"side\",\"type\":\"uint8\"},{\"name\":\"saleKind\",\"type\":\"uint8\"},{\"name\":\"replacementPattern\",\"type\":\"bytes\"},{\"name\":\"staticExtradata\",\"type\":\"bytes\"},{\"name\":\"merkleRoot\",\"type\":\"bytes32\"}],\"name\":\"hashToSign_\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addrs\",\"type\":\"address[9]\"},{\"name\":\"uints\",\"type\":\"uint256[9]\"},{\"name\":\"side\",\"type\":\"uint8\"},{\"name\":\"saleKind\",\"type\":\"uint8\"},{\"name\":\"replacementPattern\",\"type\":\"bytes\"},{\"name\":\"staticExtradata\",\"type\":\"bytes\"},{\"name\":\"merkleRoot\",\"type\":\"bytes32\"}],\"name\":\"validateOrderParameters_\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addrs\",\"type\":\"address[9]\"},{\"name\":\"uints\",\"type\":\"uint256[9]\"},{\"name\":\"side\",\"type\":\"uint8\"},{\"name\":\"saleKind\",\"type\":\"uint8\"},{\"name\":\"replacementPattern\",\"type\":\"bytes\"},{\"name\":\"staticExtradata\",\"type\":\"bytes\"},{\"name\":\"merkleRoot\",\"type\":\"bytes32\"}],\"name\":\"validateOrder_\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addrs\",\"type\":\"address[9]\"},{\"name\":\"uints\",\"type\":\"uint256[9]\"},{\"name\":\"side\",\"type\":\"uint8\"},{\"name\":\"saleKind\",\"type\":\"uint8\"},{\"name\":\"replacementPattern\",\"type\":\"bytes\"},{\"name\":\"staticExtradata\",\"type\":\"bytes\"},{\"name\":\"merkleData\",\"type\":\"bytes32[2]\"}],\"name\":\"makeOrder_\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addrs\",\"type\":\"address[9]\"},{\"name\":\"uints\",\"type\":\"uint256[9]\"},{\"name\":\"side\",\"type\":\"uint8\"},{\"name\":\"saleKind\",\"type\":\"uint8\"},{\"name\":\"replacementPattern\",\"type\":\"bytes\"},{\"name\":\"staticExtradata\",\"type\":\"bytes\"},{\"name\":\"merkleRoot\",\"type\":\"bytes32\"}],\"name\":\"cancelOrder_\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addrs\",\"type\":\"address[9]\"},{\"name\":\"uints\",\"type\":\"uint256[9]\"},{\"name\":\"side\",\"type\":\"uint8\"},{\"name\":\"saleKind\",\"type\":\"uint8\"},{\"name\":\"replacementPattern\",\"type\":\"bytes\"},{\"name\":\"staticExtradata\",\"type\":\"bytes\"},{\"name\":\"merkleRoot\",\"type\":\"bytes32\"}],\"name\":\"calculateCurrentPrice_\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addrs\",\"type\":\"address[16]\"},{\"name\":\"uints\",\"type\":\"uint256[12]\"},{\"name\":\"sidesKinds\",\"type\":\"uint8[4]\"},{\"name\":\"calldataBuy\",\"type\":\"bytes\"},{\"name\":\"calldataSell\",\"type\":\"bytes\"},{\"name\":\"replacementPatternBuy\",\"type\":\"bytes\"},{\"name\":\"replacementPatternSell\",\"type\":\"bytes\"},{\"name\":\"staticExtradataBuy\",\"type\":\"bytes\"},{\"name\":\"staticExtradataSell\",\"type\":\"bytes\"}],\"name\":\"ordersCanMatch_\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"buyCalldata\",\"type\":\"bytes\"},{\"name\":\"buyReplacementPattern\",\"type\":\"bytes\"},{\"name\":\"sellCalldata\",\"type\":\"bytes\"},{\"name\":\"sellReplacementPattern\",\"type\":\"bytes\"}],\"name\":\"orderCalldataCanMatch\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addrs\",\"type\":\"address[16]\"},{\"name\":\"uints\",\"type\":\"uint256[12]\"},{\"name\":\"sidesKinds\",\"type\":\"uint8[4]\"},{\"name\":\"calldataBuy\",\"type\":\"bytes\"},{\"name\":\"calldataSell\",\"type\":\"bytes\"},{\"name\":\"replacementPatternBuy\",\"type\":\"bytes\"},{\"name\":\"replacementPatternSell\",\"type\":\"bytes\"},{\"name\":\"staticExtradataBuy\",\"type\":\"bytes\"},{\"name\":\"staticExtradataSell\",\"type\":\"bytes\"}],\"name\":\"calculateMatchPrice_\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addrs\",\"type\":\"address[16]\"},{\"name\":\"uints\",\"type\":\"uint256[12]\"},{\"name\":\"sidesKinds\",\"type\":\"uint8[4]\"},{\"name\":\"calldataBuy\",\"type\":\"bytes\"},{\"name\":\"calldataSell\",\"type\":\"bytes\"},{\"name\":\"replacementPatternBuy\",\"type\":\"bytes\"},{\"name\":\"replacementPatternSell\",\"type\":\"bytes\"},{\"name\":\"staticExtradataBuy\",\"type\":\"bytes\"},{\"name\":\"staticExtradataSell\",\"type\":\"bytes\"},{\"name\":\"rssMetadata\",\"type\":\"bytes32\"}],\"name\":\"takeOrder_\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// NiftyconnectexchangeABI is the input ABI used to generate the binding from.
// Deprecated: Use NiftyconnectexchangeMetaData.ABI instead.
var NiftyconnectexchangeABI = NiftyconnectexchangeMetaData.ABI

// Niftyconnectexchange is an auto generated Go binding around an Ethereum contract.
type Niftyconnectexchange struct {
	NiftyconnectexchangeCaller     // Read-only binding to the contract
	NiftyconnectexchangeTransactor // Write-only binding to the contract
	NiftyconnectexchangeFilterer   // Log filterer for contract events
}

// NiftyconnectexchangeCaller is an auto generated read-only Go binding around an Ethereum contract.
type NiftyconnectexchangeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NiftyconnectexchangeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NiftyconnectexchangeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NiftyconnectexchangeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NiftyconnectexchangeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NiftyconnectexchangeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NiftyconnectexchangeSession struct {
	Contract     *Niftyconnectexchange // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// NiftyconnectexchangeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NiftyconnectexchangeCallerSession struct {
	Contract *NiftyconnectexchangeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// NiftyconnectexchangeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NiftyconnectexchangeTransactorSession struct {
	Contract     *NiftyconnectexchangeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// NiftyconnectexchangeRaw is an auto generated low-level Go binding around an Ethereum contract.
type NiftyconnectexchangeRaw struct {
	Contract *Niftyconnectexchange // Generic contract binding to access the raw methods on
}

// NiftyconnectexchangeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NiftyconnectexchangeCallerRaw struct {
	Contract *NiftyconnectexchangeCaller // Generic read-only contract binding to access the raw methods on
}

// NiftyconnectexchangeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NiftyconnectexchangeTransactorRaw struct {
	Contract *NiftyconnectexchangeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNiftyconnectexchange creates a new instance of Niftyconnectexchange, bound to a specific deployed contract.
func NewNiftyconnectexchange(address common.Address, backend bind.ContractBackend) (*Niftyconnectexchange, error) {
	contract, err := bindNiftyconnectexchange(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Niftyconnectexchange{NiftyconnectexchangeCaller: NiftyconnectexchangeCaller{contract: contract}, NiftyconnectexchangeTransactor: NiftyconnectexchangeTransactor{contract: contract}, NiftyconnectexchangeFilterer: NiftyconnectexchangeFilterer{contract: contract}}, nil
}

// NewNiftyconnectexchangeCaller creates a new read-only instance of Niftyconnectexchange, bound to a specific deployed contract.
func NewNiftyconnectexchangeCaller(address common.Address, caller bind.ContractCaller) (*NiftyconnectexchangeCaller, error) {
	contract, err := bindNiftyconnectexchange(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NiftyconnectexchangeCaller{contract: contract}, nil
}

// NewNiftyconnectexchangeTransactor creates a new write-only instance of Niftyconnectexchange, bound to a specific deployed contract.
func NewNiftyconnectexchangeTransactor(address common.Address, transactor bind.ContractTransactor) (*NiftyconnectexchangeTransactor, error) {
	contract, err := bindNiftyconnectexchange(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NiftyconnectexchangeTransactor{contract: contract}, nil
}

// NewNiftyconnectexchangeFilterer creates a new log filterer instance of Niftyconnectexchange, bound to a specific deployed contract.
func NewNiftyconnectexchangeFilterer(address common.Address, filterer bind.ContractFilterer) (*NiftyconnectexchangeFilterer, error) {
	contract, err := bindNiftyconnectexchange(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NiftyconnectexchangeFilterer{contract: contract}, nil
}

// bindNiftyconnectexchange binds a generic wrapper to an already deployed contract.
func bindNiftyconnectexchange(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NiftyconnectexchangeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Niftyconnectexchange *NiftyconnectexchangeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Niftyconnectexchange.Contract.NiftyconnectexchangeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Niftyconnectexchange *NiftyconnectexchangeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Niftyconnectexchange.Contract.NiftyconnectexchangeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Niftyconnectexchange *NiftyconnectexchangeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Niftyconnectexchange.Contract.NiftyconnectexchangeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Niftyconnectexchange *NiftyconnectexchangeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Niftyconnectexchange.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Niftyconnectexchange *NiftyconnectexchangeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Niftyconnectexchange.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Niftyconnectexchange *NiftyconnectexchangeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Niftyconnectexchange.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Niftyconnectexchange *NiftyconnectexchangeCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Niftyconnectexchange.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Niftyconnectexchange *NiftyconnectexchangeSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _Niftyconnectexchange.Contract.DOMAINSEPARATOR(&_Niftyconnectexchange.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Niftyconnectexchange *NiftyconnectexchangeCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _Niftyconnectexchange.Contract.DOMAINSEPARATOR(&_Niftyconnectexchange.CallOpts)
}

// INVERSEBASISPOINT is a free data retrieval call binding the contract method 0xcae6047f.
//
// Solidity: function INVERSE_BASIS_POINT() view returns(uint256)
func (_Niftyconnectexchange *NiftyconnectexchangeCaller) INVERSEBASISPOINT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Niftyconnectexchange.contract.Call(opts, &out, "INVERSE_BASIS_POINT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// INVERSEBASISPOINT is a free data retrieval call binding the contract method 0xcae6047f.
//
// Solidity: function INVERSE_BASIS_POINT() view returns(uint256)
func (_Niftyconnectexchange *NiftyconnectexchangeSession) INVERSEBASISPOINT() (*big.Int, error) {
	return _Niftyconnectexchange.Contract.INVERSEBASISPOINT(&_Niftyconnectexchange.CallOpts)
}

// INVERSEBASISPOINT is a free data retrieval call binding the contract method 0xcae6047f.
//
// Solidity: function INVERSE_BASIS_POINT() view returns(uint256)
func (_Niftyconnectexchange *NiftyconnectexchangeCallerSession) INVERSEBASISPOINT() (*big.Int, error) {
	return _Niftyconnectexchange.Contract.INVERSEBASISPOINT(&_Niftyconnectexchange.CallOpts)
}

// MAXIMUMEXCHANGERATE is a free data retrieval call binding the contract method 0xe4e098f7.
//
// Solidity: function MAXIMUM_EXCHANGE_RATE() view returns(uint256)
func (_Niftyconnectexchange *NiftyconnectexchangeCaller) MAXIMUMEXCHANGERATE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Niftyconnectexchange.contract.Call(opts, &out, "MAXIMUM_EXCHANGE_RATE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXIMUMEXCHANGERATE is a free data retrieval call binding the contract method 0xe4e098f7.
//
// Solidity: function MAXIMUM_EXCHANGE_RATE() view returns(uint256)
func (_Niftyconnectexchange *NiftyconnectexchangeSession) MAXIMUMEXCHANGERATE() (*big.Int, error) {
	return _Niftyconnectexchange.Contract.MAXIMUMEXCHANGERATE(&_Niftyconnectexchange.CallOpts)
}

// MAXIMUMEXCHANGERATE is a free data retrieval call binding the contract method 0xe4e098f7.
//
// Solidity: function MAXIMUM_EXCHANGE_RATE() view returns(uint256)
func (_Niftyconnectexchange *NiftyconnectexchangeCallerSession) MAXIMUMEXCHANGERATE() (*big.Int, error) {
	return _Niftyconnectexchange.Contract.MAXIMUMEXCHANGERATE(&_Niftyconnectexchange.CallOpts)
}

// ApprovedOrders is a free data retrieval call binding the contract method 0xe57d4adb.
//
// Solidity: function approvedOrders(bytes32 hash) view returns(bool approved)
func (_Niftyconnectexchange *NiftyconnectexchangeCaller) ApprovedOrders(opts *bind.CallOpts, hash [32]byte) (bool, error) {
	var out []interface{}
	err := _Niftyconnectexchange.contract.Call(opts, &out, "approvedOrders", hash)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ApprovedOrders is a free data retrieval call binding the contract method 0xe57d4adb.
//
// Solidity: function approvedOrders(bytes32 hash) view returns(bool approved)
func (_Niftyconnectexchange *NiftyconnectexchangeSession) ApprovedOrders(hash [32]byte) (bool, error) {
	return _Niftyconnectexchange.Contract.ApprovedOrders(&_Niftyconnectexchange.CallOpts, hash)
}

// ApprovedOrders is a free data retrieval call binding the contract method 0xe57d4adb.
//
// Solidity: function approvedOrders(bytes32 hash) view returns(bool approved)
func (_Niftyconnectexchange *NiftyconnectexchangeCallerSession) ApprovedOrders(hash [32]byte) (bool, error) {
	return _Niftyconnectexchange.Contract.ApprovedOrders(&_Niftyconnectexchange.CallOpts, hash)
}

// BuildCallData is a free data retrieval call binding the contract method 0x37146f2e.
//
// Solidity: function buildCallData(uint256 selector, address from, address to, address nftAddress, uint256 tokenId, uint256 amount, bytes32 merkleRoot, bytes32[] merkleProof) view returns(bytes)
func (_Niftyconnectexchange *NiftyconnectexchangeCaller) BuildCallData(opts *bind.CallOpts, selector *big.Int, from common.Address, to common.Address, nftAddress common.Address, tokenId *big.Int, amount *big.Int, merkleRoot [32]byte, merkleProof [][32]byte) ([]byte, error) {
	var out []interface{}
	err := _Niftyconnectexchange.contract.Call(opts, &out, "buildCallData", selector, from, to, nftAddress, tokenId, amount, merkleRoot, merkleProof)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// BuildCallData is a free data retrieval call binding the contract method 0x37146f2e.
//
// Solidity: function buildCallData(uint256 selector, address from, address to, address nftAddress, uint256 tokenId, uint256 amount, bytes32 merkleRoot, bytes32[] merkleProof) view returns(bytes)
func (_Niftyconnectexchange *NiftyconnectexchangeSession) BuildCallData(selector *big.Int, from common.Address, to common.Address, nftAddress common.Address, tokenId *big.Int, amount *big.Int, merkleRoot [32]byte, merkleProof [][32]byte) ([]byte, error) {
	return _Niftyconnectexchange.Contract.BuildCallData(&_Niftyconnectexchange.CallOpts, selector, from, to, nftAddress, tokenId, amount, merkleRoot, merkleProof)
}

// BuildCallData is a free data retrieval call binding the contract method 0x37146f2e.
//
// Solidity: function buildCallData(uint256 selector, address from, address to, address nftAddress, uint256 tokenId, uint256 amount, bytes32 merkleRoot, bytes32[] merkleProof) view returns(bytes)
func (_Niftyconnectexchange *NiftyconnectexchangeCallerSession) BuildCallData(selector *big.Int, from common.Address, to common.Address, nftAddress common.Address, tokenId *big.Int, amount *big.Int, merkleRoot [32]byte, merkleProof [][32]byte) ([]byte, error) {
	return _Niftyconnectexchange.Contract.BuildCallData(&_Niftyconnectexchange.CallOpts, selector, from, to, nftAddress, tokenId, amount, merkleRoot, merkleProof)
}

// CalculateCurrentPrice is a free data retrieval call binding the contract method 0x1f86dbc0.
//
// Solidity: function calculateCurrentPrice_(address[9] addrs, uint256[9] uints, uint8 side, uint8 saleKind, bytes replacementPattern, bytes staticExtradata, bytes32 merkleRoot) view returns(uint256)
func (_Niftyconnectexchange *NiftyconnectexchangeCaller) CalculateCurrentPrice(opts *bind.CallOpts, addrs [9]common.Address, uints [9]*big.Int, side uint8, saleKind uint8, replacementPattern []byte, staticExtradata []byte, merkleRoot [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Niftyconnectexchange.contract.Call(opts, &out, "calculateCurrentPrice_", addrs, uints, side, saleKind, replacementPattern, staticExtradata, merkleRoot)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateCurrentPrice is a free data retrieval call binding the contract method 0x1f86dbc0.
//
// Solidity: function calculateCurrentPrice_(address[9] addrs, uint256[9] uints, uint8 side, uint8 saleKind, bytes replacementPattern, bytes staticExtradata, bytes32 merkleRoot) view returns(uint256)
func (_Niftyconnectexchange *NiftyconnectexchangeSession) CalculateCurrentPrice(addrs [9]common.Address, uints [9]*big.Int, side uint8, saleKind uint8, replacementPattern []byte, staticExtradata []byte, merkleRoot [32]byte) (*big.Int, error) {
	return _Niftyconnectexchange.Contract.CalculateCurrentPrice(&_Niftyconnectexchange.CallOpts, addrs, uints, side, saleKind, replacementPattern, staticExtradata, merkleRoot)
}

// CalculateCurrentPrice is a free data retrieval call binding the contract method 0x1f86dbc0.
//
// Solidity: function calculateCurrentPrice_(address[9] addrs, uint256[9] uints, uint8 side, uint8 saleKind, bytes replacementPattern, bytes staticExtradata, bytes32 merkleRoot) view returns(uint256)
func (_Niftyconnectexchange *NiftyconnectexchangeCallerSession) CalculateCurrentPrice(addrs [9]common.Address, uints [9]*big.Int, side uint8, saleKind uint8, replacementPattern []byte, staticExtradata []byte, merkleRoot [32]byte) (*big.Int, error) {
	return _Niftyconnectexchange.Contract.CalculateCurrentPrice(&_Niftyconnectexchange.CallOpts, addrs, uints, side, saleKind, replacementPattern, staticExtradata, merkleRoot)
}

// CalculateFinalPrice is a free data retrieval call binding the contract method 0x63d36c0b.
//
// Solidity: function calculateFinalPrice(uint8 side, uint8 saleKind, uint256 basePrice, uint256 extra, uint256 listingTime, uint256 expirationTime) view returns(uint256)
func (_Niftyconnectexchange *NiftyconnectexchangeCaller) CalculateFinalPrice(opts *bind.CallOpts, side uint8, saleKind uint8, basePrice *big.Int, extra *big.Int, listingTime *big.Int, expirationTime *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Niftyconnectexchange.contract.Call(opts, &out, "calculateFinalPrice", side, saleKind, basePrice, extra, listingTime, expirationTime)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateFinalPrice is a free data retrieval call binding the contract method 0x63d36c0b.
//
// Solidity: function calculateFinalPrice(uint8 side, uint8 saleKind, uint256 basePrice, uint256 extra, uint256 listingTime, uint256 expirationTime) view returns(uint256)
func (_Niftyconnectexchange *NiftyconnectexchangeSession) CalculateFinalPrice(side uint8, saleKind uint8, basePrice *big.Int, extra *big.Int, listingTime *big.Int, expirationTime *big.Int) (*big.Int, error) {
	return _Niftyconnectexchange.Contract.CalculateFinalPrice(&_Niftyconnectexchange.CallOpts, side, saleKind, basePrice, extra, listingTime, expirationTime)
}

// CalculateFinalPrice is a free data retrieval call binding the contract method 0x63d36c0b.
//
// Solidity: function calculateFinalPrice(uint8 side, uint8 saleKind, uint256 basePrice, uint256 extra, uint256 listingTime, uint256 expirationTime) view returns(uint256)
func (_Niftyconnectexchange *NiftyconnectexchangeCallerSession) CalculateFinalPrice(side uint8, saleKind uint8, basePrice *big.Int, extra *big.Int, listingTime *big.Int, expirationTime *big.Int) (*big.Int, error) {
	return _Niftyconnectexchange.Contract.CalculateFinalPrice(&_Niftyconnectexchange.CallOpts, side, saleKind, basePrice, extra, listingTime, expirationTime)
}

// CalculateMatchPrice is a free data retrieval call binding the contract method 0x028e01cf.
//
// Solidity: function calculateMatchPrice_(address[16] addrs, uint256[12] uints, uint8[4] sidesKinds, bytes calldataBuy, bytes calldataSell, bytes replacementPatternBuy, bytes replacementPatternSell, bytes staticExtradataBuy, bytes staticExtradataSell) view returns(uint256)
func (_Niftyconnectexchange *NiftyconnectexchangeCaller) CalculateMatchPrice(opts *bind.CallOpts, addrs [16]common.Address, uints [12]*big.Int, sidesKinds [4]uint8, calldataBuy []byte, calldataSell []byte, replacementPatternBuy []byte, replacementPatternSell []byte, staticExtradataBuy []byte, staticExtradataSell []byte) (*big.Int, error) {
	var out []interface{}
	err := _Niftyconnectexchange.contract.Call(opts, &out, "calculateMatchPrice_", addrs, uints, sidesKinds, calldataBuy, calldataSell, replacementPatternBuy, replacementPatternSell, staticExtradataBuy, staticExtradataSell)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateMatchPrice is a free data retrieval call binding the contract method 0x028e01cf.
//
// Solidity: function calculateMatchPrice_(address[16] addrs, uint256[12] uints, uint8[4] sidesKinds, bytes calldataBuy, bytes calldataSell, bytes replacementPatternBuy, bytes replacementPatternSell, bytes staticExtradataBuy, bytes staticExtradataSell) view returns(uint256)
func (_Niftyconnectexchange *NiftyconnectexchangeSession) CalculateMatchPrice(addrs [16]common.Address, uints [12]*big.Int, sidesKinds [4]uint8, calldataBuy []byte, calldataSell []byte, replacementPatternBuy []byte, replacementPatternSell []byte, staticExtradataBuy []byte, staticExtradataSell []byte) (*big.Int, error) {
	return _Niftyconnectexchange.Contract.CalculateMatchPrice(&_Niftyconnectexchange.CallOpts, addrs, uints, sidesKinds, calldataBuy, calldataSell, replacementPatternBuy, replacementPatternSell, staticExtradataBuy, staticExtradataSell)
}

// CalculateMatchPrice is a free data retrieval call binding the contract method 0x028e01cf.
//
// Solidity: function calculateMatchPrice_(address[16] addrs, uint256[12] uints, uint8[4] sidesKinds, bytes calldataBuy, bytes calldataSell, bytes replacementPatternBuy, bytes replacementPatternSell, bytes staticExtradataBuy, bytes staticExtradataSell) view returns(uint256)
func (_Niftyconnectexchange *NiftyconnectexchangeCallerSession) CalculateMatchPrice(addrs [16]common.Address, uints [12]*big.Int, sidesKinds [4]uint8, calldataBuy []byte, calldataSell []byte, replacementPatternBuy []byte, replacementPatternSell []byte, staticExtradataBuy []byte, staticExtradataSell []byte) (*big.Int, error) {
	return _Niftyconnectexchange.Contract.CalculateMatchPrice(&_Niftyconnectexchange.CallOpts, addrs, uints, sidesKinds, calldataBuy, calldataSell, replacementPatternBuy, replacementPatternSell, staticExtradataBuy, staticExtradataSell)
}

// CancelledOrFinalized is a free data retrieval call binding the contract method 0x8076f005.
//
// Solidity: function cancelledOrFinalized(bytes32 ) view returns(bool)
func (_Niftyconnectexchange *NiftyconnectexchangeCaller) CancelledOrFinalized(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _Niftyconnectexchange.contract.Call(opts, &out, "cancelledOrFinalized", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CancelledOrFinalized is a free data retrieval call binding the contract method 0x8076f005.
//
// Solidity: function cancelledOrFinalized(bytes32 ) view returns(bool)
func (_Niftyconnectexchange *NiftyconnectexchangeSession) CancelledOrFinalized(arg0 [32]byte) (bool, error) {
	return _Niftyconnectexchange.Contract.CancelledOrFinalized(&_Niftyconnectexchange.CallOpts, arg0)
}

// CancelledOrFinalized is a free data retrieval call binding the contract method 0x8076f005.
//
// Solidity: function cancelledOrFinalized(bytes32 ) view returns(bool)
func (_Niftyconnectexchange *NiftyconnectexchangeCallerSession) CancelledOrFinalized(arg0 [32]byte) (bool, error) {
	return _Niftyconnectexchange.Contract.CancelledOrFinalized(&_Niftyconnectexchange.CallOpts, arg0)
}

// ExchangeFeeRate is a free data retrieval call binding the contract method 0x0f9b4955.
//
// Solidity: function exchangeFeeRate() view returns(uint256)
func (_Niftyconnectexchange *NiftyconnectexchangeCaller) ExchangeFeeRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Niftyconnectexchange.contract.Call(opts, &out, "exchangeFeeRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExchangeFeeRate is a free data retrieval call binding the contract method 0x0f9b4955.
//
// Solidity: function exchangeFeeRate() view returns(uint256)
func (_Niftyconnectexchange *NiftyconnectexchangeSession) ExchangeFeeRate() (*big.Int, error) {
	return _Niftyconnectexchange.Contract.ExchangeFeeRate(&_Niftyconnectexchange.CallOpts)
}

// ExchangeFeeRate is a free data retrieval call binding the contract method 0x0f9b4955.
//
// Solidity: function exchangeFeeRate() view returns(uint256)
func (_Niftyconnectexchange *NiftyconnectexchangeCallerSession) ExchangeFeeRate() (*big.Int, error) {
	return _Niftyconnectexchange.Contract.ExchangeFeeRate(&_Niftyconnectexchange.CallOpts)
}

// Governor is a free data retrieval call binding the contract method 0x0c340a24.
//
// Solidity: function governor() view returns(address)
func (_Niftyconnectexchange *NiftyconnectexchangeCaller) Governor(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Niftyconnectexchange.contract.Call(opts, &out, "governor")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Governor is a free data retrieval call binding the contract method 0x0c340a24.
//
// Solidity: function governor() view returns(address)
func (_Niftyconnectexchange *NiftyconnectexchangeSession) Governor() (common.Address, error) {
	return _Niftyconnectexchange.Contract.Governor(&_Niftyconnectexchange.CallOpts)
}

// Governor is a free data retrieval call binding the contract method 0x0c340a24.
//
// Solidity: function governor() view returns(address)
func (_Niftyconnectexchange *NiftyconnectexchangeCallerSession) Governor() (common.Address, error) {
	return _Niftyconnectexchange.Contract.Governor(&_Niftyconnectexchange.CallOpts)
}

// GuardedArrayReplace is a free data retrieval call binding the contract method 0x239e83df.
//
// Solidity: function guardedArrayReplace(bytes array, bytes desired, bytes mask) pure returns(bytes)
func (_Niftyconnectexchange *NiftyconnectexchangeCaller) GuardedArrayReplace(opts *bind.CallOpts, array []byte, desired []byte, mask []byte) ([]byte, error) {
	var out []interface{}
	err := _Niftyconnectexchange.contract.Call(opts, &out, "guardedArrayReplace", array, desired, mask)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GuardedArrayReplace is a free data retrieval call binding the contract method 0x239e83df.
//
// Solidity: function guardedArrayReplace(bytes array, bytes desired, bytes mask) pure returns(bytes)
func (_Niftyconnectexchange *NiftyconnectexchangeSession) GuardedArrayReplace(array []byte, desired []byte, mask []byte) ([]byte, error) {
	return _Niftyconnectexchange.Contract.GuardedArrayReplace(&_Niftyconnectexchange.CallOpts, array, desired, mask)
}

// GuardedArrayReplace is a free data retrieval call binding the contract method 0x239e83df.
//
// Solidity: function guardedArrayReplace(bytes array, bytes desired, bytes mask) pure returns(bytes)
func (_Niftyconnectexchange *NiftyconnectexchangeCallerSession) GuardedArrayReplace(array []byte, desired []byte, mask []byte) ([]byte, error) {
	return _Niftyconnectexchange.Contract.GuardedArrayReplace(&_Niftyconnectexchange.CallOpts, array, desired, mask)
}

// HashToSign is a free data retrieval call binding the contract method 0x81da91a0.
//
// Solidity: function hashToSign_(address[9] addrs, uint256[9] uints, uint8 side, uint8 saleKind, bytes replacementPattern, bytes staticExtradata, bytes32 merkleRoot) view returns(bytes32)
func (_Niftyconnectexchange *NiftyconnectexchangeCaller) HashToSign(opts *bind.CallOpts, addrs [9]common.Address, uints [9]*big.Int, side uint8, saleKind uint8, replacementPattern []byte, staticExtradata []byte, merkleRoot [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Niftyconnectexchange.contract.Call(opts, &out, "hashToSign_", addrs, uints, side, saleKind, replacementPattern, staticExtradata, merkleRoot)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// HashToSign is a free data retrieval call binding the contract method 0x81da91a0.
//
// Solidity: function hashToSign_(address[9] addrs, uint256[9] uints, uint8 side, uint8 saleKind, bytes replacementPattern, bytes staticExtradata, bytes32 merkleRoot) view returns(bytes32)
func (_Niftyconnectexchange *NiftyconnectexchangeSession) HashToSign(addrs [9]common.Address, uints [9]*big.Int, side uint8, saleKind uint8, replacementPattern []byte, staticExtradata []byte, merkleRoot [32]byte) ([32]byte, error) {
	return _Niftyconnectexchange.Contract.HashToSign(&_Niftyconnectexchange.CallOpts, addrs, uints, side, saleKind, replacementPattern, staticExtradata, merkleRoot)
}

// HashToSign is a free data retrieval call binding the contract method 0x81da91a0.
//
// Solidity: function hashToSign_(address[9] addrs, uint256[9] uints, uint8 side, uint8 saleKind, bytes replacementPattern, bytes staticExtradata, bytes32 merkleRoot) view returns(bytes32)
func (_Niftyconnectexchange *NiftyconnectexchangeCallerSession) HashToSign(addrs [9]common.Address, uints [9]*big.Int, side uint8, saleKind uint8, replacementPattern []byte, staticExtradata []byte, merkleRoot [32]byte) ([32]byte, error) {
	return _Niftyconnectexchange.Contract.HashToSign(&_Niftyconnectexchange.CallOpts, addrs, uints, side, saleKind, replacementPattern, staticExtradata, merkleRoot)
}

// MakerRelayerFeeShare is a free data retrieval call binding the contract method 0xe8898e6d.
//
// Solidity: function makerRelayerFeeShare() view returns(uint256)
func (_Niftyconnectexchange *NiftyconnectexchangeCaller) MakerRelayerFeeShare(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Niftyconnectexchange.contract.Call(opts, &out, "makerRelayerFeeShare")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MakerRelayerFeeShare is a free data retrieval call binding the contract method 0xe8898e6d.
//
// Solidity: function makerRelayerFeeShare() view returns(uint256)
func (_Niftyconnectexchange *NiftyconnectexchangeSession) MakerRelayerFeeShare() (*big.Int, error) {
	return _Niftyconnectexchange.Contract.MakerRelayerFeeShare(&_Niftyconnectexchange.CallOpts)
}

// MakerRelayerFeeShare is a free data retrieval call binding the contract method 0xe8898e6d.
//
// Solidity: function makerRelayerFeeShare() view returns(uint256)
func (_Niftyconnectexchange *NiftyconnectexchangeCallerSession) MakerRelayerFeeShare() (*big.Int, error) {
	return _Niftyconnectexchange.Contract.MakerRelayerFeeShare(&_Niftyconnectexchange.CallOpts)
}

// MerkleValidatorContract is a free data retrieval call binding the contract method 0xf1113575.
//
// Solidity: function merkleValidatorContract() view returns(address)
func (_Niftyconnectexchange *NiftyconnectexchangeCaller) MerkleValidatorContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Niftyconnectexchange.contract.Call(opts, &out, "merkleValidatorContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MerkleValidatorContract is a free data retrieval call binding the contract method 0xf1113575.
//
// Solidity: function merkleValidatorContract() view returns(address)
func (_Niftyconnectexchange *NiftyconnectexchangeSession) MerkleValidatorContract() (common.Address, error) {
	return _Niftyconnectexchange.Contract.MerkleValidatorContract(&_Niftyconnectexchange.CallOpts)
}

// MerkleValidatorContract is a free data retrieval call binding the contract method 0xf1113575.
//
// Solidity: function merkleValidatorContract() view returns(address)
func (_Niftyconnectexchange *NiftyconnectexchangeCallerSession) MerkleValidatorContract() (common.Address, error) {
	return _Niftyconnectexchange.Contract.MerkleValidatorContract(&_Niftyconnectexchange.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Niftyconnectexchange *NiftyconnectexchangeCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Niftyconnectexchange.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Niftyconnectexchange *NiftyconnectexchangeSession) Name() (string, error) {
	return _Niftyconnectexchange.Contract.Name(&_Niftyconnectexchange.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Niftyconnectexchange *NiftyconnectexchangeCallerSession) Name() (string, error) {
	return _Niftyconnectexchange.Contract.Name(&_Niftyconnectexchange.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_Niftyconnectexchange *NiftyconnectexchangeCaller) Nonces(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Niftyconnectexchange.contract.Call(opts, &out, "nonces", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_Niftyconnectexchange *NiftyconnectexchangeSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _Niftyconnectexchange.Contract.Nonces(&_Niftyconnectexchange.CallOpts, arg0)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_Niftyconnectexchange *NiftyconnectexchangeCallerSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _Niftyconnectexchange.Contract.Nonces(&_Niftyconnectexchange.CallOpts, arg0)
}

// OrderCalldataCanMatch is a free data retrieval call binding the contract method 0x562b2ebc.
//
// Solidity: function orderCalldataCanMatch(bytes buyCalldata, bytes buyReplacementPattern, bytes sellCalldata, bytes sellReplacementPattern) pure returns(bool)
func (_Niftyconnectexchange *NiftyconnectexchangeCaller) OrderCalldataCanMatch(opts *bind.CallOpts, buyCalldata []byte, buyReplacementPattern []byte, sellCalldata []byte, sellReplacementPattern []byte) (bool, error) {
	var out []interface{}
	err := _Niftyconnectexchange.contract.Call(opts, &out, "orderCalldataCanMatch", buyCalldata, buyReplacementPattern, sellCalldata, sellReplacementPattern)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// OrderCalldataCanMatch is a free data retrieval call binding the contract method 0x562b2ebc.
//
// Solidity: function orderCalldataCanMatch(bytes buyCalldata, bytes buyReplacementPattern, bytes sellCalldata, bytes sellReplacementPattern) pure returns(bool)
func (_Niftyconnectexchange *NiftyconnectexchangeSession) OrderCalldataCanMatch(buyCalldata []byte, buyReplacementPattern []byte, sellCalldata []byte, sellReplacementPattern []byte) (bool, error) {
	return _Niftyconnectexchange.Contract.OrderCalldataCanMatch(&_Niftyconnectexchange.CallOpts, buyCalldata, buyReplacementPattern, sellCalldata, sellReplacementPattern)
}

// OrderCalldataCanMatch is a free data retrieval call binding the contract method 0x562b2ebc.
//
// Solidity: function orderCalldataCanMatch(bytes buyCalldata, bytes buyReplacementPattern, bytes sellCalldata, bytes sellReplacementPattern) pure returns(bool)
func (_Niftyconnectexchange *NiftyconnectexchangeCallerSession) OrderCalldataCanMatch(buyCalldata []byte, buyReplacementPattern []byte, sellCalldata []byte, sellReplacementPattern []byte) (bool, error) {
	return _Niftyconnectexchange.Contract.OrderCalldataCanMatch(&_Niftyconnectexchange.CallOpts, buyCalldata, buyReplacementPattern, sellCalldata, sellReplacementPattern)
}

// OrdersCanMatch is a free data retrieval call binding the contract method 0x1f2c56d7.
//
// Solidity: function ordersCanMatch_(address[16] addrs, uint256[12] uints, uint8[4] sidesKinds, bytes calldataBuy, bytes calldataSell, bytes replacementPatternBuy, bytes replacementPatternSell, bytes staticExtradataBuy, bytes staticExtradataSell) view returns(bool)
func (_Niftyconnectexchange *NiftyconnectexchangeCaller) OrdersCanMatch(opts *bind.CallOpts, addrs [16]common.Address, uints [12]*big.Int, sidesKinds [4]uint8, calldataBuy []byte, calldataSell []byte, replacementPatternBuy []byte, replacementPatternSell []byte, staticExtradataBuy []byte, staticExtradataSell []byte) (bool, error) {
	var out []interface{}
	err := _Niftyconnectexchange.contract.Call(opts, &out, "ordersCanMatch_", addrs, uints, sidesKinds, calldataBuy, calldataSell, replacementPatternBuy, replacementPatternSell, staticExtradataBuy, staticExtradataSell)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// OrdersCanMatch is a free data retrieval call binding the contract method 0x1f2c56d7.
//
// Solidity: function ordersCanMatch_(address[16] addrs, uint256[12] uints, uint8[4] sidesKinds, bytes calldataBuy, bytes calldataSell, bytes replacementPatternBuy, bytes replacementPatternSell, bytes staticExtradataBuy, bytes staticExtradataSell) view returns(bool)
func (_Niftyconnectexchange *NiftyconnectexchangeSession) OrdersCanMatch(addrs [16]common.Address, uints [12]*big.Int, sidesKinds [4]uint8, calldataBuy []byte, calldataSell []byte, replacementPatternBuy []byte, replacementPatternSell []byte, staticExtradataBuy []byte, staticExtradataSell []byte) (bool, error) {
	return _Niftyconnectexchange.Contract.OrdersCanMatch(&_Niftyconnectexchange.CallOpts, addrs, uints, sidesKinds, calldataBuy, calldataSell, replacementPatternBuy, replacementPatternSell, staticExtradataBuy, staticExtradataSell)
}

// OrdersCanMatch is a free data retrieval call binding the contract method 0x1f2c56d7.
//
// Solidity: function ordersCanMatch_(address[16] addrs, uint256[12] uints, uint8[4] sidesKinds, bytes calldataBuy, bytes calldataSell, bytes replacementPatternBuy, bytes replacementPatternSell, bytes staticExtradataBuy, bytes staticExtradataSell) view returns(bool)
func (_Niftyconnectexchange *NiftyconnectexchangeCallerSession) OrdersCanMatch(addrs [16]common.Address, uints [12]*big.Int, sidesKinds [4]uint8, calldataBuy []byte, calldataSell []byte, replacementPatternBuy []byte, replacementPatternSell []byte, staticExtradataBuy []byte, staticExtradataSell []byte) (bool, error) {
	return _Niftyconnectexchange.Contract.OrdersCanMatch(&_Niftyconnectexchange.CallOpts, addrs, uints, sidesKinds, calldataBuy, calldataSell, replacementPatternBuy, replacementPatternSell, staticExtradataBuy, staticExtradataSell)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Niftyconnectexchange *NiftyconnectexchangeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Niftyconnectexchange.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Niftyconnectexchange *NiftyconnectexchangeSession) Owner() (common.Address, error) {
	return _Niftyconnectexchange.Contract.Owner(&_Niftyconnectexchange.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Niftyconnectexchange *NiftyconnectexchangeCallerSession) Owner() (common.Address, error) {
	return _Niftyconnectexchange.Contract.Owner(&_Niftyconnectexchange.CallOpts)
}

// PendingGovernor is a free data retrieval call binding the contract method 0xe3056a34.
//
// Solidity: function pendingGovernor() view returns(address)
func (_Niftyconnectexchange *NiftyconnectexchangeCaller) PendingGovernor(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Niftyconnectexchange.contract.Call(opts, &out, "pendingGovernor")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingGovernor is a free data retrieval call binding the contract method 0xe3056a34.
//
// Solidity: function pendingGovernor() view returns(address)
func (_Niftyconnectexchange *NiftyconnectexchangeSession) PendingGovernor() (common.Address, error) {
	return _Niftyconnectexchange.Contract.PendingGovernor(&_Niftyconnectexchange.CallOpts)
}

// PendingGovernor is a free data retrieval call binding the contract method 0xe3056a34.
//
// Solidity: function pendingGovernor() view returns(address)
func (_Niftyconnectexchange *NiftyconnectexchangeCallerSession) PendingGovernor() (common.Address, error) {
	return _Niftyconnectexchange.Contract.PendingGovernor(&_Niftyconnectexchange.CallOpts)
}

// ProtocolFeeRecipient is a free data retrieval call binding the contract method 0x64df049e.
//
// Solidity: function protocolFeeRecipient() view returns(address)
func (_Niftyconnectexchange *NiftyconnectexchangeCaller) ProtocolFeeRecipient(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Niftyconnectexchange.contract.Call(opts, &out, "protocolFeeRecipient")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ProtocolFeeRecipient is a free data retrieval call binding the contract method 0x64df049e.
//
// Solidity: function protocolFeeRecipient() view returns(address)
func (_Niftyconnectexchange *NiftyconnectexchangeSession) ProtocolFeeRecipient() (common.Address, error) {
	return _Niftyconnectexchange.Contract.ProtocolFeeRecipient(&_Niftyconnectexchange.CallOpts)
}

// ProtocolFeeRecipient is a free data retrieval call binding the contract method 0x64df049e.
//
// Solidity: function protocolFeeRecipient() view returns(address)
func (_Niftyconnectexchange *NiftyconnectexchangeCallerSession) ProtocolFeeRecipient() (common.Address, error) {
	return _Niftyconnectexchange.Contract.ProtocolFeeRecipient(&_Niftyconnectexchange.CallOpts)
}

// ProtocolFeeShare is a free data retrieval call binding the contract method 0x960b26a2.
//
// Solidity: function protocolFeeShare() view returns(uint256)
func (_Niftyconnectexchange *NiftyconnectexchangeCaller) ProtocolFeeShare(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Niftyconnectexchange.contract.Call(opts, &out, "protocolFeeShare")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProtocolFeeShare is a free data retrieval call binding the contract method 0x960b26a2.
//
// Solidity: function protocolFeeShare() view returns(uint256)
func (_Niftyconnectexchange *NiftyconnectexchangeSession) ProtocolFeeShare() (*big.Int, error) {
	return _Niftyconnectexchange.Contract.ProtocolFeeShare(&_Niftyconnectexchange.CallOpts)
}

// ProtocolFeeShare is a free data retrieval call binding the contract method 0x960b26a2.
//
// Solidity: function protocolFeeShare() view returns(uint256)
func (_Niftyconnectexchange *NiftyconnectexchangeCallerSession) ProtocolFeeShare() (*big.Int, error) {
	return _Niftyconnectexchange.Contract.ProtocolFeeShare(&_Niftyconnectexchange.CallOpts)
}

// RoyaltyRegisterHub is a free data retrieval call binding the contract method 0x3eeb5bc8.
//
// Solidity: function royaltyRegisterHub() view returns(address)
func (_Niftyconnectexchange *NiftyconnectexchangeCaller) RoyaltyRegisterHub(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Niftyconnectexchange.contract.Call(opts, &out, "royaltyRegisterHub")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RoyaltyRegisterHub is a free data retrieval call binding the contract method 0x3eeb5bc8.
//
// Solidity: function royaltyRegisterHub() view returns(address)
func (_Niftyconnectexchange *NiftyconnectexchangeSession) RoyaltyRegisterHub() (common.Address, error) {
	return _Niftyconnectexchange.Contract.RoyaltyRegisterHub(&_Niftyconnectexchange.CallOpts)
}

// RoyaltyRegisterHub is a free data retrieval call binding the contract method 0x3eeb5bc8.
//
// Solidity: function royaltyRegisterHub() view returns(address)
func (_Niftyconnectexchange *NiftyconnectexchangeCallerSession) RoyaltyRegisterHub() (common.Address, error) {
	return _Niftyconnectexchange.Contract.RoyaltyRegisterHub(&_Niftyconnectexchange.CallOpts)
}

// StaticCall is a free data retrieval call binding the contract method 0x10796a47.
//
// Solidity: function staticCall(address target, bytes calldata, bytes extradata) view returns(bool result)
func (_Niftyconnectexchange *NiftyconnectexchangeCaller) StaticCall(opts *bind.CallOpts, target common.Address, calldata []byte, extradata []byte) (bool, error) {
	var out []interface{}
	err := _Niftyconnectexchange.contract.Call(opts, &out, "staticCall", target, calldata, extradata)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// StaticCall is a free data retrieval call binding the contract method 0x10796a47.
//
// Solidity: function staticCall(address target, bytes calldata, bytes extradata) view returns(bool result)
func (_Niftyconnectexchange *NiftyconnectexchangeSession) StaticCall(target common.Address, calldata []byte, extradata []byte) (bool, error) {
	return _Niftyconnectexchange.Contract.StaticCall(&_Niftyconnectexchange.CallOpts, target, calldata, extradata)
}

// StaticCall is a free data retrieval call binding the contract method 0x10796a47.
//
// Solidity: function staticCall(address target, bytes calldata, bytes extradata) view returns(bool result)
func (_Niftyconnectexchange *NiftyconnectexchangeCallerSession) StaticCall(target common.Address, calldata []byte, extradata []byte) (bool, error) {
	return _Niftyconnectexchange.Contract.StaticCall(&_Niftyconnectexchange.CallOpts, target, calldata, extradata)
}

// TakerRelayerFeeShare is a free data retrieval call binding the contract method 0x4a3b5e05.
//
// Solidity: function takerRelayerFeeShare() view returns(uint256)
func (_Niftyconnectexchange *NiftyconnectexchangeCaller) TakerRelayerFeeShare(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Niftyconnectexchange.contract.Call(opts, &out, "takerRelayerFeeShare")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TakerRelayerFeeShare is a free data retrieval call binding the contract method 0x4a3b5e05.
//
// Solidity: function takerRelayerFeeShare() view returns(uint256)
func (_Niftyconnectexchange *NiftyconnectexchangeSession) TakerRelayerFeeShare() (*big.Int, error) {
	return _Niftyconnectexchange.Contract.TakerRelayerFeeShare(&_Niftyconnectexchange.CallOpts)
}

// TakerRelayerFeeShare is a free data retrieval call binding the contract method 0x4a3b5e05.
//
// Solidity: function takerRelayerFeeShare() view returns(uint256)
func (_Niftyconnectexchange *NiftyconnectexchangeCallerSession) TakerRelayerFeeShare() (*big.Int, error) {
	return _Niftyconnectexchange.Contract.TakerRelayerFeeShare(&_Niftyconnectexchange.CallOpts)
}

// TokenTransferProxy is a free data retrieval call binding the contract method 0x0eefdbad.
//
// Solidity: function tokenTransferProxy() view returns(address)
func (_Niftyconnectexchange *NiftyconnectexchangeCaller) TokenTransferProxy(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Niftyconnectexchange.contract.Call(opts, &out, "tokenTransferProxy")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenTransferProxy is a free data retrieval call binding the contract method 0x0eefdbad.
//
// Solidity: function tokenTransferProxy() view returns(address)
func (_Niftyconnectexchange *NiftyconnectexchangeSession) TokenTransferProxy() (common.Address, error) {
	return _Niftyconnectexchange.Contract.TokenTransferProxy(&_Niftyconnectexchange.CallOpts)
}

// TokenTransferProxy is a free data retrieval call binding the contract method 0x0eefdbad.
//
// Solidity: function tokenTransferProxy() view returns(address)
func (_Niftyconnectexchange *NiftyconnectexchangeCallerSession) TokenTransferProxy() (common.Address, error) {
	return _Niftyconnectexchange.Contract.TokenTransferProxy(&_Niftyconnectexchange.CallOpts)
}

// ValidateOrderParameters is a free data retrieval call binding the contract method 0xe7b74b64.
//
// Solidity: function validateOrderParameters_(address[9] addrs, uint256[9] uints, uint8 side, uint8 saleKind, bytes replacementPattern, bytes staticExtradata, bytes32 merkleRoot) view returns(bool)
func (_Niftyconnectexchange *NiftyconnectexchangeCaller) ValidateOrderParameters(opts *bind.CallOpts, addrs [9]common.Address, uints [9]*big.Int, side uint8, saleKind uint8, replacementPattern []byte, staticExtradata []byte, merkleRoot [32]byte) (bool, error) {
	var out []interface{}
	err := _Niftyconnectexchange.contract.Call(opts, &out, "validateOrderParameters_", addrs, uints, side, saleKind, replacementPattern, staticExtradata, merkleRoot)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ValidateOrderParameters is a free data retrieval call binding the contract method 0xe7b74b64.
//
// Solidity: function validateOrderParameters_(address[9] addrs, uint256[9] uints, uint8 side, uint8 saleKind, bytes replacementPattern, bytes staticExtradata, bytes32 merkleRoot) view returns(bool)
func (_Niftyconnectexchange *NiftyconnectexchangeSession) ValidateOrderParameters(addrs [9]common.Address, uints [9]*big.Int, side uint8, saleKind uint8, replacementPattern []byte, staticExtradata []byte, merkleRoot [32]byte) (bool, error) {
	return _Niftyconnectexchange.Contract.ValidateOrderParameters(&_Niftyconnectexchange.CallOpts, addrs, uints, side, saleKind, replacementPattern, staticExtradata, merkleRoot)
}

// ValidateOrderParameters is a free data retrieval call binding the contract method 0xe7b74b64.
//
// Solidity: function validateOrderParameters_(address[9] addrs, uint256[9] uints, uint8 side, uint8 saleKind, bytes replacementPattern, bytes staticExtradata, bytes32 merkleRoot) view returns(bool)
func (_Niftyconnectexchange *NiftyconnectexchangeCallerSession) ValidateOrderParameters(addrs [9]common.Address, uints [9]*big.Int, side uint8, saleKind uint8, replacementPattern []byte, staticExtradata []byte, merkleRoot [32]byte) (bool, error) {
	return _Niftyconnectexchange.Contract.ValidateOrderParameters(&_Niftyconnectexchange.CallOpts, addrs, uints, side, saleKind, replacementPattern, staticExtradata, merkleRoot)
}

// ValidateOrder is a free data retrieval call binding the contract method 0x3df6be13.
//
// Solidity: function validateOrder_(address[9] addrs, uint256[9] uints, uint8 side, uint8 saleKind, bytes replacementPattern, bytes staticExtradata, bytes32 merkleRoot) view returns(bool)
func (_Niftyconnectexchange *NiftyconnectexchangeCaller) ValidateOrder(opts *bind.CallOpts, addrs [9]common.Address, uints [9]*big.Int, side uint8, saleKind uint8, replacementPattern []byte, staticExtradata []byte, merkleRoot [32]byte) (bool, error) {
	var out []interface{}
	err := _Niftyconnectexchange.contract.Call(opts, &out, "validateOrder_", addrs, uints, side, saleKind, replacementPattern, staticExtradata, merkleRoot)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ValidateOrder is a free data retrieval call binding the contract method 0x3df6be13.
//
// Solidity: function validateOrder_(address[9] addrs, uint256[9] uints, uint8 side, uint8 saleKind, bytes replacementPattern, bytes staticExtradata, bytes32 merkleRoot) view returns(bool)
func (_Niftyconnectexchange *NiftyconnectexchangeSession) ValidateOrder(addrs [9]common.Address, uints [9]*big.Int, side uint8, saleKind uint8, replacementPattern []byte, staticExtradata []byte, merkleRoot [32]byte) (bool, error) {
	return _Niftyconnectexchange.Contract.ValidateOrder(&_Niftyconnectexchange.CallOpts, addrs, uints, side, saleKind, replacementPattern, staticExtradata, merkleRoot)
}

// ValidateOrder is a free data retrieval call binding the contract method 0x3df6be13.
//
// Solidity: function validateOrder_(address[9] addrs, uint256[9] uints, uint8 side, uint8 saleKind, bytes replacementPattern, bytes staticExtradata, bytes32 merkleRoot) view returns(bool)
func (_Niftyconnectexchange *NiftyconnectexchangeCallerSession) ValidateOrder(addrs [9]common.Address, uints [9]*big.Int, side uint8, saleKind uint8, replacementPattern []byte, staticExtradata []byte, merkleRoot [32]byte) (bool, error) {
	return _Niftyconnectexchange.Contract.ValidateOrder(&_Niftyconnectexchange.CallOpts, addrs, uints, side, saleKind, replacementPattern, staticExtradata, merkleRoot)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_Niftyconnectexchange *NiftyconnectexchangeCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Niftyconnectexchange.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_Niftyconnectexchange *NiftyconnectexchangeSession) Version() (string, error) {
	return _Niftyconnectexchange.Contract.Version(&_Niftyconnectexchange.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_Niftyconnectexchange *NiftyconnectexchangeCallerSession) Version() (string, error) {
	return _Niftyconnectexchange.Contract.Version(&_Niftyconnectexchange.CallOpts)
}

// AcceptGovernance is a paid mutator transaction binding the contract method 0x238efcbc.
//
// Solidity: function acceptGovernance() returns()
func (_Niftyconnectexchange *NiftyconnectexchangeTransactor) AcceptGovernance(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Niftyconnectexchange.contract.Transact(opts, "acceptGovernance")
}

// AcceptGovernance is a paid mutator transaction binding the contract method 0x238efcbc.
//
// Solidity: function acceptGovernance() returns()
func (_Niftyconnectexchange *NiftyconnectexchangeSession) AcceptGovernance() (*types.Transaction, error) {
	return _Niftyconnectexchange.Contract.AcceptGovernance(&_Niftyconnectexchange.TransactOpts)
}

// AcceptGovernance is a paid mutator transaction binding the contract method 0x238efcbc.
//
// Solidity: function acceptGovernance() returns()
func (_Niftyconnectexchange *NiftyconnectexchangeTransactorSession) AcceptGovernance() (*types.Transaction, error) {
	return _Niftyconnectexchange.Contract.AcceptGovernance(&_Niftyconnectexchange.TransactOpts)
}

// CancelOrder is a paid mutator transaction binding the contract method 0x94146166.
//
// Solidity: function cancelOrder_(address[9] addrs, uint256[9] uints, uint8 side, uint8 saleKind, bytes replacementPattern, bytes staticExtradata, bytes32 merkleRoot) returns()
func (_Niftyconnectexchange *NiftyconnectexchangeTransactor) CancelOrder(opts *bind.TransactOpts, addrs [9]common.Address, uints [9]*big.Int, side uint8, saleKind uint8, replacementPattern []byte, staticExtradata []byte, merkleRoot [32]byte) (*types.Transaction, error) {
	return _Niftyconnectexchange.contract.Transact(opts, "cancelOrder_", addrs, uints, side, saleKind, replacementPattern, staticExtradata, merkleRoot)
}

// CancelOrder is a paid mutator transaction binding the contract method 0x94146166.
//
// Solidity: function cancelOrder_(address[9] addrs, uint256[9] uints, uint8 side, uint8 saleKind, bytes replacementPattern, bytes staticExtradata, bytes32 merkleRoot) returns()
func (_Niftyconnectexchange *NiftyconnectexchangeSession) CancelOrder(addrs [9]common.Address, uints [9]*big.Int, side uint8, saleKind uint8, replacementPattern []byte, staticExtradata []byte, merkleRoot [32]byte) (*types.Transaction, error) {
	return _Niftyconnectexchange.Contract.CancelOrder(&_Niftyconnectexchange.TransactOpts, addrs, uints, side, saleKind, replacementPattern, staticExtradata, merkleRoot)
}

// CancelOrder is a paid mutator transaction binding the contract method 0x94146166.
//
// Solidity: function cancelOrder_(address[9] addrs, uint256[9] uints, uint8 side, uint8 saleKind, bytes replacementPattern, bytes staticExtradata, bytes32 merkleRoot) returns()
func (_Niftyconnectexchange *NiftyconnectexchangeTransactorSession) CancelOrder(addrs [9]common.Address, uints [9]*big.Int, side uint8, saleKind uint8, replacementPattern []byte, staticExtradata []byte, merkleRoot [32]byte) (*types.Transaction, error) {
	return _Niftyconnectexchange.Contract.CancelOrder(&_Niftyconnectexchange.TransactOpts, addrs, uints, side, saleKind, replacementPattern, staticExtradata, merkleRoot)
}

// ChangeExchangeFeeRate is a paid mutator transaction binding the contract method 0xade0ccb2.
//
// Solidity: function changeExchangeFeeRate(uint256 newExchangeFeeRate) returns()
func (_Niftyconnectexchange *NiftyconnectexchangeTransactor) ChangeExchangeFeeRate(opts *bind.TransactOpts, newExchangeFeeRate *big.Int) (*types.Transaction, error) {
	return _Niftyconnectexchange.contract.Transact(opts, "changeExchangeFeeRate", newExchangeFeeRate)
}

// ChangeExchangeFeeRate is a paid mutator transaction binding the contract method 0xade0ccb2.
//
// Solidity: function changeExchangeFeeRate(uint256 newExchangeFeeRate) returns()
func (_Niftyconnectexchange *NiftyconnectexchangeSession) ChangeExchangeFeeRate(newExchangeFeeRate *big.Int) (*types.Transaction, error) {
	return _Niftyconnectexchange.Contract.ChangeExchangeFeeRate(&_Niftyconnectexchange.TransactOpts, newExchangeFeeRate)
}

// ChangeExchangeFeeRate is a paid mutator transaction binding the contract method 0xade0ccb2.
//
// Solidity: function changeExchangeFeeRate(uint256 newExchangeFeeRate) returns()
func (_Niftyconnectexchange *NiftyconnectexchangeTransactorSession) ChangeExchangeFeeRate(newExchangeFeeRate *big.Int) (*types.Transaction, error) {
	return _Niftyconnectexchange.Contract.ChangeExchangeFeeRate(&_Niftyconnectexchange.TransactOpts, newExchangeFeeRate)
}

// ChangeProtocolFeeRecipient is a paid mutator transaction binding the contract method 0x514f0330.
//
// Solidity: function changeProtocolFeeRecipient(address newProtocolFeeRecipient) returns()
func (_Niftyconnectexchange *NiftyconnectexchangeTransactor) ChangeProtocolFeeRecipient(opts *bind.TransactOpts, newProtocolFeeRecipient common.Address) (*types.Transaction, error) {
	return _Niftyconnectexchange.contract.Transact(opts, "changeProtocolFeeRecipient", newProtocolFeeRecipient)
}

// ChangeProtocolFeeRecipient is a paid mutator transaction binding the contract method 0x514f0330.
//
// Solidity: function changeProtocolFeeRecipient(address newProtocolFeeRecipient) returns()
func (_Niftyconnectexchange *NiftyconnectexchangeSession) ChangeProtocolFeeRecipient(newProtocolFeeRecipient common.Address) (*types.Transaction, error) {
	return _Niftyconnectexchange.Contract.ChangeProtocolFeeRecipient(&_Niftyconnectexchange.TransactOpts, newProtocolFeeRecipient)
}

// ChangeProtocolFeeRecipient is a paid mutator transaction binding the contract method 0x514f0330.
//
// Solidity: function changeProtocolFeeRecipient(address newProtocolFeeRecipient) returns()
func (_Niftyconnectexchange *NiftyconnectexchangeTransactorSession) ChangeProtocolFeeRecipient(newProtocolFeeRecipient common.Address) (*types.Transaction, error) {
	return _Niftyconnectexchange.Contract.ChangeProtocolFeeRecipient(&_Niftyconnectexchange.TransactOpts, newProtocolFeeRecipient)
}

// ChangeTakerRelayerFeeShare is a paid mutator transaction binding the contract method 0x3d1cf526.
//
// Solidity: function changeTakerRelayerFeeShare(uint256 newTakerRelayerFeeShare, uint256 newMakerRelayerFeeShare, uint256 newProtocolFeeShare) returns()
func (_Niftyconnectexchange *NiftyconnectexchangeTransactor) ChangeTakerRelayerFeeShare(opts *bind.TransactOpts, newTakerRelayerFeeShare *big.Int, newMakerRelayerFeeShare *big.Int, newProtocolFeeShare *big.Int) (*types.Transaction, error) {
	return _Niftyconnectexchange.contract.Transact(opts, "changeTakerRelayerFeeShare", newTakerRelayerFeeShare, newMakerRelayerFeeShare, newProtocolFeeShare)
}

// ChangeTakerRelayerFeeShare is a paid mutator transaction binding the contract method 0x3d1cf526.
//
// Solidity: function changeTakerRelayerFeeShare(uint256 newTakerRelayerFeeShare, uint256 newMakerRelayerFeeShare, uint256 newProtocolFeeShare) returns()
func (_Niftyconnectexchange *NiftyconnectexchangeSession) ChangeTakerRelayerFeeShare(newTakerRelayerFeeShare *big.Int, newMakerRelayerFeeShare *big.Int, newProtocolFeeShare *big.Int) (*types.Transaction, error) {
	return _Niftyconnectexchange.Contract.ChangeTakerRelayerFeeShare(&_Niftyconnectexchange.TransactOpts, newTakerRelayerFeeShare, newMakerRelayerFeeShare, newProtocolFeeShare)
}

// ChangeTakerRelayerFeeShare is a paid mutator transaction binding the contract method 0x3d1cf526.
//
// Solidity: function changeTakerRelayerFeeShare(uint256 newTakerRelayerFeeShare, uint256 newMakerRelayerFeeShare, uint256 newProtocolFeeShare) returns()
func (_Niftyconnectexchange *NiftyconnectexchangeTransactorSession) ChangeTakerRelayerFeeShare(newTakerRelayerFeeShare *big.Int, newMakerRelayerFeeShare *big.Int, newProtocolFeeShare *big.Int) (*types.Transaction, error) {
	return _Niftyconnectexchange.Contract.ChangeTakerRelayerFeeShare(&_Niftyconnectexchange.TransactOpts, newTakerRelayerFeeShare, newMakerRelayerFeeShare, newProtocolFeeShare)
}

// IncrementNonce is a paid mutator transaction binding the contract method 0x627cdcb9.
//
// Solidity: function incrementNonce() returns()
func (_Niftyconnectexchange *NiftyconnectexchangeTransactor) IncrementNonce(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Niftyconnectexchange.contract.Transact(opts, "incrementNonce")
}

// IncrementNonce is a paid mutator transaction binding the contract method 0x627cdcb9.
//
// Solidity: function incrementNonce() returns()
func (_Niftyconnectexchange *NiftyconnectexchangeSession) IncrementNonce() (*types.Transaction, error) {
	return _Niftyconnectexchange.Contract.IncrementNonce(&_Niftyconnectexchange.TransactOpts)
}

// IncrementNonce is a paid mutator transaction binding the contract method 0x627cdcb9.
//
// Solidity: function incrementNonce() returns()
func (_Niftyconnectexchange *NiftyconnectexchangeTransactorSession) IncrementNonce() (*types.Transaction, error) {
	return _Niftyconnectexchange.Contract.IncrementNonce(&_Niftyconnectexchange.TransactOpts)
}

// MakeOrder is a paid mutator transaction binding the contract method 0x97cea71b.
//
// Solidity: function makeOrder_(address[9] addrs, uint256[9] uints, uint8 side, uint8 saleKind, bytes replacementPattern, bytes staticExtradata, bytes32[2] merkleData) returns()
func (_Niftyconnectexchange *NiftyconnectexchangeTransactor) MakeOrder(opts *bind.TransactOpts, addrs [9]common.Address, uints [9]*big.Int, side uint8, saleKind uint8, replacementPattern []byte, staticExtradata []byte, merkleData [2][32]byte) (*types.Transaction, error) {
	return _Niftyconnectexchange.contract.Transact(opts, "makeOrder_", addrs, uints, side, saleKind, replacementPattern, staticExtradata, merkleData)
}

// MakeOrder is a paid mutator transaction binding the contract method 0x97cea71b.
//
// Solidity: function makeOrder_(address[9] addrs, uint256[9] uints, uint8 side, uint8 saleKind, bytes replacementPattern, bytes staticExtradata, bytes32[2] merkleData) returns()
func (_Niftyconnectexchange *NiftyconnectexchangeSession) MakeOrder(addrs [9]common.Address, uints [9]*big.Int, side uint8, saleKind uint8, replacementPattern []byte, staticExtradata []byte, merkleData [2][32]byte) (*types.Transaction, error) {
	return _Niftyconnectexchange.Contract.MakeOrder(&_Niftyconnectexchange.TransactOpts, addrs, uints, side, saleKind, replacementPattern, staticExtradata, merkleData)
}

// MakeOrder is a paid mutator transaction binding the contract method 0x97cea71b.
//
// Solidity: function makeOrder_(address[9] addrs, uint256[9] uints, uint8 side, uint8 saleKind, bytes replacementPattern, bytes staticExtradata, bytes32[2] merkleData) returns()
func (_Niftyconnectexchange *NiftyconnectexchangeTransactorSession) MakeOrder(addrs [9]common.Address, uints [9]*big.Int, side uint8, saleKind uint8, replacementPattern []byte, staticExtradata []byte, merkleData [2][32]byte) (*types.Transaction, error) {
	return _Niftyconnectexchange.Contract.MakeOrder(&_Niftyconnectexchange.TransactOpts, addrs, uints, side, saleKind, replacementPattern, staticExtradata, merkleData)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Niftyconnectexchange *NiftyconnectexchangeTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Niftyconnectexchange.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Niftyconnectexchange *NiftyconnectexchangeSession) RenounceOwnership() (*types.Transaction, error) {
	return _Niftyconnectexchange.Contract.RenounceOwnership(&_Niftyconnectexchange.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Niftyconnectexchange *NiftyconnectexchangeTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Niftyconnectexchange.Contract.RenounceOwnership(&_Niftyconnectexchange.TransactOpts)
}

// SetPendingGovernor is a paid mutator transaction binding the contract method 0xf235757f.
//
// Solidity: function setPendingGovernor(address pendingGovernor_) returns()
func (_Niftyconnectexchange *NiftyconnectexchangeTransactor) SetPendingGovernor(opts *bind.TransactOpts, pendingGovernor_ common.Address) (*types.Transaction, error) {
	return _Niftyconnectexchange.contract.Transact(opts, "setPendingGovernor", pendingGovernor_)
}

// SetPendingGovernor is a paid mutator transaction binding the contract method 0xf235757f.
//
// Solidity: function setPendingGovernor(address pendingGovernor_) returns()
func (_Niftyconnectexchange *NiftyconnectexchangeSession) SetPendingGovernor(pendingGovernor_ common.Address) (*types.Transaction, error) {
	return _Niftyconnectexchange.Contract.SetPendingGovernor(&_Niftyconnectexchange.TransactOpts, pendingGovernor_)
}

// SetPendingGovernor is a paid mutator transaction binding the contract method 0xf235757f.
//
// Solidity: function setPendingGovernor(address pendingGovernor_) returns()
func (_Niftyconnectexchange *NiftyconnectexchangeTransactorSession) SetPendingGovernor(pendingGovernor_ common.Address) (*types.Transaction, error) {
	return _Niftyconnectexchange.Contract.SetPendingGovernor(&_Niftyconnectexchange.TransactOpts, pendingGovernor_)
}

// TakeOrder is a paid mutator transaction binding the contract method 0x7da26f55.
//
// Solidity: function takeOrder_(address[16] addrs, uint256[12] uints, uint8[4] sidesKinds, bytes calldataBuy, bytes calldataSell, bytes replacementPatternBuy, bytes replacementPatternSell, bytes staticExtradataBuy, bytes staticExtradataSell, bytes32 rssMetadata) payable returns()
func (_Niftyconnectexchange *NiftyconnectexchangeTransactor) TakeOrder(opts *bind.TransactOpts, addrs [16]common.Address, uints [12]*big.Int, sidesKinds [4]uint8, calldataBuy []byte, calldataSell []byte, replacementPatternBuy []byte, replacementPatternSell []byte, staticExtradataBuy []byte, staticExtradataSell []byte, rssMetadata [32]byte) (*types.Transaction, error) {
	return _Niftyconnectexchange.contract.Transact(opts, "takeOrder_", addrs, uints, sidesKinds, calldataBuy, calldataSell, replacementPatternBuy, replacementPatternSell, staticExtradataBuy, staticExtradataSell, rssMetadata)
}

// TakeOrder is a paid mutator transaction binding the contract method 0x7da26f55.
//
// Solidity: function takeOrder_(address[16] addrs, uint256[12] uints, uint8[4] sidesKinds, bytes calldataBuy, bytes calldataSell, bytes replacementPatternBuy, bytes replacementPatternSell, bytes staticExtradataBuy, bytes staticExtradataSell, bytes32 rssMetadata) payable returns()
func (_Niftyconnectexchange *NiftyconnectexchangeSession) TakeOrder(addrs [16]common.Address, uints [12]*big.Int, sidesKinds [4]uint8, calldataBuy []byte, calldataSell []byte, replacementPatternBuy []byte, replacementPatternSell []byte, staticExtradataBuy []byte, staticExtradataSell []byte, rssMetadata [32]byte) (*types.Transaction, error) {
	return _Niftyconnectexchange.Contract.TakeOrder(&_Niftyconnectexchange.TransactOpts, addrs, uints, sidesKinds, calldataBuy, calldataSell, replacementPatternBuy, replacementPatternSell, staticExtradataBuy, staticExtradataSell, rssMetadata)
}

// TakeOrder is a paid mutator transaction binding the contract method 0x7da26f55.
//
// Solidity: function takeOrder_(address[16] addrs, uint256[12] uints, uint8[4] sidesKinds, bytes calldataBuy, bytes calldataSell, bytes replacementPatternBuy, bytes replacementPatternSell, bytes staticExtradataBuy, bytes staticExtradataSell, bytes32 rssMetadata) payable returns()
func (_Niftyconnectexchange *NiftyconnectexchangeTransactorSession) TakeOrder(addrs [16]common.Address, uints [12]*big.Int, sidesKinds [4]uint8, calldataBuy []byte, calldataSell []byte, replacementPatternBuy []byte, replacementPatternSell []byte, staticExtradataBuy []byte, staticExtradataSell []byte, rssMetadata [32]byte) (*types.Transaction, error) {
	return _Niftyconnectexchange.Contract.TakeOrder(&_Niftyconnectexchange.TransactOpts, addrs, uints, sidesKinds, calldataBuy, calldataSell, replacementPatternBuy, replacementPatternSell, staticExtradataBuy, staticExtradataSell, rssMetadata)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Niftyconnectexchange *NiftyconnectexchangeTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Niftyconnectexchange.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Niftyconnectexchange *NiftyconnectexchangeSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Niftyconnectexchange.Contract.TransferOwnership(&_Niftyconnectexchange.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Niftyconnectexchange *NiftyconnectexchangeTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Niftyconnectexchange.Contract.TransferOwnership(&_Niftyconnectexchange.TransactOpts, newOwner)
}

// NiftyconnectexchangeGovernanceTransferredIterator is returned from FilterGovernanceTransferred and is used to iterate over the raw logs and unpacked data for GovernanceTransferred events raised by the Niftyconnectexchange contract.
type NiftyconnectexchangeGovernanceTransferredIterator struct {
	Event *NiftyconnectexchangeGovernanceTransferred // Event containing the contract specifics and raw log

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
func (it *NiftyconnectexchangeGovernanceTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NiftyconnectexchangeGovernanceTransferred)
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
		it.Event = new(NiftyconnectexchangeGovernanceTransferred)
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
func (it *NiftyconnectexchangeGovernanceTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NiftyconnectexchangeGovernanceTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NiftyconnectexchangeGovernanceTransferred represents a GovernanceTransferred event raised by the Niftyconnectexchange contract.
type NiftyconnectexchangeGovernanceTransferred struct {
	PreviousGovernor common.Address
	NewGovernor      common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterGovernanceTransferred is a free log retrieval operation binding the contract event 0x5f56bee8cffbe9a78652a74a60705edede02af10b0bbb888ca44b79a0d42ce80.
//
// Solidity: event GovernanceTransferred(address indexed previousGovernor, address indexed newGovernor)
func (_Niftyconnectexchange *NiftyconnectexchangeFilterer) FilterGovernanceTransferred(opts *bind.FilterOpts, previousGovernor []common.Address, newGovernor []common.Address) (*NiftyconnectexchangeGovernanceTransferredIterator, error) {

	var previousGovernorRule []interface{}
	for _, previousGovernorItem := range previousGovernor {
		previousGovernorRule = append(previousGovernorRule, previousGovernorItem)
	}
	var newGovernorRule []interface{}
	for _, newGovernorItem := range newGovernor {
		newGovernorRule = append(newGovernorRule, newGovernorItem)
	}

	logs, sub, err := _Niftyconnectexchange.contract.FilterLogs(opts, "GovernanceTransferred", previousGovernorRule, newGovernorRule)
	if err != nil {
		return nil, err
	}
	return &NiftyconnectexchangeGovernanceTransferredIterator{contract: _Niftyconnectexchange.contract, event: "GovernanceTransferred", logs: logs, sub: sub}, nil
}

// WatchGovernanceTransferred is a free log subscription operation binding the contract event 0x5f56bee8cffbe9a78652a74a60705edede02af10b0bbb888ca44b79a0d42ce80.
//
// Solidity: event GovernanceTransferred(address indexed previousGovernor, address indexed newGovernor)
func (_Niftyconnectexchange *NiftyconnectexchangeFilterer) WatchGovernanceTransferred(opts *bind.WatchOpts, sink chan<- *NiftyconnectexchangeGovernanceTransferred, previousGovernor []common.Address, newGovernor []common.Address) (event.Subscription, error) {

	var previousGovernorRule []interface{}
	for _, previousGovernorItem := range previousGovernor {
		previousGovernorRule = append(previousGovernorRule, previousGovernorItem)
	}
	var newGovernorRule []interface{}
	for _, newGovernorItem := range newGovernor {
		newGovernorRule = append(newGovernorRule, newGovernorItem)
	}

	logs, sub, err := _Niftyconnectexchange.contract.WatchLogs(opts, "GovernanceTransferred", previousGovernorRule, newGovernorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NiftyconnectexchangeGovernanceTransferred)
				if err := _Niftyconnectexchange.contract.UnpackLog(event, "GovernanceTransferred", log); err != nil {
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

// ParseGovernanceTransferred is a log parse operation binding the contract event 0x5f56bee8cffbe9a78652a74a60705edede02af10b0bbb888ca44b79a0d42ce80.
//
// Solidity: event GovernanceTransferred(address indexed previousGovernor, address indexed newGovernor)
func (_Niftyconnectexchange *NiftyconnectexchangeFilterer) ParseGovernanceTransferred(log types.Log) (*NiftyconnectexchangeGovernanceTransferred, error) {
	event := new(NiftyconnectexchangeGovernanceTransferred)
	if err := _Niftyconnectexchange.contract.UnpackLog(event, "GovernanceTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NiftyconnectexchangeNewPendingGovernorIterator is returned from FilterNewPendingGovernor and is used to iterate over the raw logs and unpacked data for NewPendingGovernor events raised by the Niftyconnectexchange contract.
type NiftyconnectexchangeNewPendingGovernorIterator struct {
	Event *NiftyconnectexchangeNewPendingGovernor // Event containing the contract specifics and raw log

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
func (it *NiftyconnectexchangeNewPendingGovernorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NiftyconnectexchangeNewPendingGovernor)
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
		it.Event = new(NiftyconnectexchangeNewPendingGovernor)
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
func (it *NiftyconnectexchangeNewPendingGovernorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NiftyconnectexchangeNewPendingGovernorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NiftyconnectexchangeNewPendingGovernor represents a NewPendingGovernor event raised by the Niftyconnectexchange contract.
type NiftyconnectexchangeNewPendingGovernor struct {
	NewPendingGovernor common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterNewPendingGovernor is a free log retrieval operation binding the contract event 0xe6df4d3d01a6133dfdecd1b451c04ec286cb4b10e7235d2b27321b476216e6d7.
//
// Solidity: event NewPendingGovernor(address indexed newPendingGovernor)
func (_Niftyconnectexchange *NiftyconnectexchangeFilterer) FilterNewPendingGovernor(opts *bind.FilterOpts, newPendingGovernor []common.Address) (*NiftyconnectexchangeNewPendingGovernorIterator, error) {

	var newPendingGovernorRule []interface{}
	for _, newPendingGovernorItem := range newPendingGovernor {
		newPendingGovernorRule = append(newPendingGovernorRule, newPendingGovernorItem)
	}

	logs, sub, err := _Niftyconnectexchange.contract.FilterLogs(opts, "NewPendingGovernor", newPendingGovernorRule)
	if err != nil {
		return nil, err
	}
	return &NiftyconnectexchangeNewPendingGovernorIterator{contract: _Niftyconnectexchange.contract, event: "NewPendingGovernor", logs: logs, sub: sub}, nil
}

// WatchNewPendingGovernor is a free log subscription operation binding the contract event 0xe6df4d3d01a6133dfdecd1b451c04ec286cb4b10e7235d2b27321b476216e6d7.
//
// Solidity: event NewPendingGovernor(address indexed newPendingGovernor)
func (_Niftyconnectexchange *NiftyconnectexchangeFilterer) WatchNewPendingGovernor(opts *bind.WatchOpts, sink chan<- *NiftyconnectexchangeNewPendingGovernor, newPendingGovernor []common.Address) (event.Subscription, error) {

	var newPendingGovernorRule []interface{}
	for _, newPendingGovernorItem := range newPendingGovernor {
		newPendingGovernorRule = append(newPendingGovernorRule, newPendingGovernorItem)
	}

	logs, sub, err := _Niftyconnectexchange.contract.WatchLogs(opts, "NewPendingGovernor", newPendingGovernorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NiftyconnectexchangeNewPendingGovernor)
				if err := _Niftyconnectexchange.contract.UnpackLog(event, "NewPendingGovernor", log); err != nil {
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

// ParseNewPendingGovernor is a log parse operation binding the contract event 0xe6df4d3d01a6133dfdecd1b451c04ec286cb4b10e7235d2b27321b476216e6d7.
//
// Solidity: event NewPendingGovernor(address indexed newPendingGovernor)
func (_Niftyconnectexchange *NiftyconnectexchangeFilterer) ParseNewPendingGovernor(log types.Log) (*NiftyconnectexchangeNewPendingGovernor, error) {
	event := new(NiftyconnectexchangeNewPendingGovernor)
	if err := _Niftyconnectexchange.contract.UnpackLog(event, "NewPendingGovernor", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NiftyconnectexchangeNonceIncrementedIterator is returned from FilterNonceIncremented and is used to iterate over the raw logs and unpacked data for NonceIncremented events raised by the Niftyconnectexchange contract.
type NiftyconnectexchangeNonceIncrementedIterator struct {
	Event *NiftyconnectexchangeNonceIncremented // Event containing the contract specifics and raw log

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
func (it *NiftyconnectexchangeNonceIncrementedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NiftyconnectexchangeNonceIncremented)
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
		it.Event = new(NiftyconnectexchangeNonceIncremented)
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
func (it *NiftyconnectexchangeNonceIncrementedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NiftyconnectexchangeNonceIncrementedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NiftyconnectexchangeNonceIncremented represents a NonceIncremented event raised by the Niftyconnectexchange contract.
type NiftyconnectexchangeNonceIncremented struct {
	Maker    common.Address
	NewNonce *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNonceIncremented is a free log retrieval operation binding the contract event 0xa82a649bbd060c9099cd7b7326e2b0dc9e9af0836480e0f849dc9eaa79710b3b.
//
// Solidity: event NonceIncremented(address indexed maker, uint256 newNonce)
func (_Niftyconnectexchange *NiftyconnectexchangeFilterer) FilterNonceIncremented(opts *bind.FilterOpts, maker []common.Address) (*NiftyconnectexchangeNonceIncrementedIterator, error) {

	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}

	logs, sub, err := _Niftyconnectexchange.contract.FilterLogs(opts, "NonceIncremented", makerRule)
	if err != nil {
		return nil, err
	}
	return &NiftyconnectexchangeNonceIncrementedIterator{contract: _Niftyconnectexchange.contract, event: "NonceIncremented", logs: logs, sub: sub}, nil
}

// WatchNonceIncremented is a free log subscription operation binding the contract event 0xa82a649bbd060c9099cd7b7326e2b0dc9e9af0836480e0f849dc9eaa79710b3b.
//
// Solidity: event NonceIncremented(address indexed maker, uint256 newNonce)
func (_Niftyconnectexchange *NiftyconnectexchangeFilterer) WatchNonceIncremented(opts *bind.WatchOpts, sink chan<- *NiftyconnectexchangeNonceIncremented, maker []common.Address) (event.Subscription, error) {

	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}

	logs, sub, err := _Niftyconnectexchange.contract.WatchLogs(opts, "NonceIncremented", makerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NiftyconnectexchangeNonceIncremented)
				if err := _Niftyconnectexchange.contract.UnpackLog(event, "NonceIncremented", log); err != nil {
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

// ParseNonceIncremented is a log parse operation binding the contract event 0xa82a649bbd060c9099cd7b7326e2b0dc9e9af0836480e0f849dc9eaa79710b3b.
//
// Solidity: event NonceIncremented(address indexed maker, uint256 newNonce)
func (_Niftyconnectexchange *NiftyconnectexchangeFilterer) ParseNonceIncremented(log types.Log) (*NiftyconnectexchangeNonceIncremented, error) {
	event := new(NiftyconnectexchangeNonceIncremented)
	if err := _Niftyconnectexchange.contract.UnpackLog(event, "NonceIncremented", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NiftyconnectexchangeOrderApprovedPartOneIterator is returned from FilterOrderApprovedPartOne and is used to iterate over the raw logs and unpacked data for OrderApprovedPartOne events raised by the Niftyconnectexchange contract.
type NiftyconnectexchangeOrderApprovedPartOneIterator struct {
	Event *NiftyconnectexchangeOrderApprovedPartOne // Event containing the contract specifics and raw log

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
func (it *NiftyconnectexchangeOrderApprovedPartOneIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NiftyconnectexchangeOrderApprovedPartOne)
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
		it.Event = new(NiftyconnectexchangeOrderApprovedPartOne)
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
func (it *NiftyconnectexchangeOrderApprovedPartOneIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NiftyconnectexchangeOrderApprovedPartOneIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NiftyconnectexchangeOrderApprovedPartOne represents a OrderApprovedPartOne event raised by the Niftyconnectexchange contract.
type NiftyconnectexchangeOrderApprovedPartOne struct {
	Hash                     [32]byte
	Exchange                 common.Address
	Maker                    common.Address
	Taker                    common.Address
	MakerRelayerFeeRecipient common.Address
	Side                     uint8
	SaleKind                 uint8
	NftAddress               common.Address
	TokenId                  *big.Int
	IpfsHash                 [32]byte
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterOrderApprovedPartOne is a free log retrieval operation binding the contract event 0xbfc991b64000533072b5f27ccd5e8628fea28ae33286778a627005b3156c6fe1.
//
// Solidity: event OrderApprovedPartOne(bytes32 indexed hash, address exchange, address indexed maker, address taker, address indexed makerRelayerFeeRecipient, uint8 side, uint8 saleKind, address nftAddress, uint256 tokenId, bytes32 ipfsHash)
func (_Niftyconnectexchange *NiftyconnectexchangeFilterer) FilterOrderApprovedPartOne(opts *bind.FilterOpts, hash [][32]byte, maker []common.Address, makerRelayerFeeRecipient []common.Address) (*NiftyconnectexchangeOrderApprovedPartOneIterator, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}

	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}

	var makerRelayerFeeRecipientRule []interface{}
	for _, makerRelayerFeeRecipientItem := range makerRelayerFeeRecipient {
		makerRelayerFeeRecipientRule = append(makerRelayerFeeRecipientRule, makerRelayerFeeRecipientItem)
	}

	logs, sub, err := _Niftyconnectexchange.contract.FilterLogs(opts, "OrderApprovedPartOne", hashRule, makerRule, makerRelayerFeeRecipientRule)
	if err != nil {
		return nil, err
	}
	return &NiftyconnectexchangeOrderApprovedPartOneIterator{contract: _Niftyconnectexchange.contract, event: "OrderApprovedPartOne", logs: logs, sub: sub}, nil
}

// WatchOrderApprovedPartOne is a free log subscription operation binding the contract event 0xbfc991b64000533072b5f27ccd5e8628fea28ae33286778a627005b3156c6fe1.
//
// Solidity: event OrderApprovedPartOne(bytes32 indexed hash, address exchange, address indexed maker, address taker, address indexed makerRelayerFeeRecipient, uint8 side, uint8 saleKind, address nftAddress, uint256 tokenId, bytes32 ipfsHash)
func (_Niftyconnectexchange *NiftyconnectexchangeFilterer) WatchOrderApprovedPartOne(opts *bind.WatchOpts, sink chan<- *NiftyconnectexchangeOrderApprovedPartOne, hash [][32]byte, maker []common.Address, makerRelayerFeeRecipient []common.Address) (event.Subscription, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}

	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}

	var makerRelayerFeeRecipientRule []interface{}
	for _, makerRelayerFeeRecipientItem := range makerRelayerFeeRecipient {
		makerRelayerFeeRecipientRule = append(makerRelayerFeeRecipientRule, makerRelayerFeeRecipientItem)
	}

	logs, sub, err := _Niftyconnectexchange.contract.WatchLogs(opts, "OrderApprovedPartOne", hashRule, makerRule, makerRelayerFeeRecipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NiftyconnectexchangeOrderApprovedPartOne)
				if err := _Niftyconnectexchange.contract.UnpackLog(event, "OrderApprovedPartOne", log); err != nil {
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

// ParseOrderApprovedPartOne is a log parse operation binding the contract event 0xbfc991b64000533072b5f27ccd5e8628fea28ae33286778a627005b3156c6fe1.
//
// Solidity: event OrderApprovedPartOne(bytes32 indexed hash, address exchange, address indexed maker, address taker, address indexed makerRelayerFeeRecipient, uint8 side, uint8 saleKind, address nftAddress, uint256 tokenId, bytes32 ipfsHash)
func (_Niftyconnectexchange *NiftyconnectexchangeFilterer) ParseOrderApprovedPartOne(log types.Log) (*NiftyconnectexchangeOrderApprovedPartOne, error) {
	event := new(NiftyconnectexchangeOrderApprovedPartOne)
	if err := _Niftyconnectexchange.contract.UnpackLog(event, "OrderApprovedPartOne", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NiftyconnectexchangeOrderApprovedPartTwoIterator is returned from FilterOrderApprovedPartTwo and is used to iterate over the raw logs and unpacked data for OrderApprovedPartTwo events raised by the Niftyconnectexchange contract.
type NiftyconnectexchangeOrderApprovedPartTwoIterator struct {
	Event *NiftyconnectexchangeOrderApprovedPartTwo // Event containing the contract specifics and raw log

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
func (it *NiftyconnectexchangeOrderApprovedPartTwoIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NiftyconnectexchangeOrderApprovedPartTwo)
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
		it.Event = new(NiftyconnectexchangeOrderApprovedPartTwo)
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
func (it *NiftyconnectexchangeOrderApprovedPartTwoIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NiftyconnectexchangeOrderApprovedPartTwoIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NiftyconnectexchangeOrderApprovedPartTwo represents a OrderApprovedPartTwo event raised by the Niftyconnectexchange contract.
type NiftyconnectexchangeOrderApprovedPartTwo struct {
	Hash               [32]byte
	Calldata           []byte
	ReplacementPattern []byte
	StaticTarget       common.Address
	StaticExtradata    []byte
	PaymentToken       common.Address
	BasePrice          *big.Int
	Extra              *big.Int
	ListingTime        *big.Int
	ExpirationTime     *big.Int
	Salt               *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterOrderApprovedPartTwo is a free log retrieval operation binding the contract event 0xb7c210e6374e28618aff2db2406f01343e302b15476278b7795869bccc51f979.
//
// Solidity: event OrderApprovedPartTwo(bytes32 indexed hash, bytes calldata, bytes replacementPattern, address staticTarget, bytes staticExtradata, address paymentToken, uint256 basePrice, uint256 extra, uint256 listingTime, uint256 expirationTime, uint256 salt)
func (_Niftyconnectexchange *NiftyconnectexchangeFilterer) FilterOrderApprovedPartTwo(opts *bind.FilterOpts, hash [][32]byte) (*NiftyconnectexchangeOrderApprovedPartTwoIterator, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}

	logs, sub, err := _Niftyconnectexchange.contract.FilterLogs(opts, "OrderApprovedPartTwo", hashRule)
	if err != nil {
		return nil, err
	}
	return &NiftyconnectexchangeOrderApprovedPartTwoIterator{contract: _Niftyconnectexchange.contract, event: "OrderApprovedPartTwo", logs: logs, sub: sub}, nil
}

// WatchOrderApprovedPartTwo is a free log subscription operation binding the contract event 0xb7c210e6374e28618aff2db2406f01343e302b15476278b7795869bccc51f979.
//
// Solidity: event OrderApprovedPartTwo(bytes32 indexed hash, bytes calldata, bytes replacementPattern, address staticTarget, bytes staticExtradata, address paymentToken, uint256 basePrice, uint256 extra, uint256 listingTime, uint256 expirationTime, uint256 salt)
func (_Niftyconnectexchange *NiftyconnectexchangeFilterer) WatchOrderApprovedPartTwo(opts *bind.WatchOpts, sink chan<- *NiftyconnectexchangeOrderApprovedPartTwo, hash [][32]byte) (event.Subscription, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}

	logs, sub, err := _Niftyconnectexchange.contract.WatchLogs(opts, "OrderApprovedPartTwo", hashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NiftyconnectexchangeOrderApprovedPartTwo)
				if err := _Niftyconnectexchange.contract.UnpackLog(event, "OrderApprovedPartTwo", log); err != nil {
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

// ParseOrderApprovedPartTwo is a log parse operation binding the contract event 0xb7c210e6374e28618aff2db2406f01343e302b15476278b7795869bccc51f979.
//
// Solidity: event OrderApprovedPartTwo(bytes32 indexed hash, bytes calldata, bytes replacementPattern, address staticTarget, bytes staticExtradata, address paymentToken, uint256 basePrice, uint256 extra, uint256 listingTime, uint256 expirationTime, uint256 salt)
func (_Niftyconnectexchange *NiftyconnectexchangeFilterer) ParseOrderApprovedPartTwo(log types.Log) (*NiftyconnectexchangeOrderApprovedPartTwo, error) {
	event := new(NiftyconnectexchangeOrderApprovedPartTwo)
	if err := _Niftyconnectexchange.contract.UnpackLog(event, "OrderApprovedPartTwo", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NiftyconnectexchangeOrderCancelledIterator is returned from FilterOrderCancelled and is used to iterate over the raw logs and unpacked data for OrderCancelled events raised by the Niftyconnectexchange contract.
type NiftyconnectexchangeOrderCancelledIterator struct {
	Event *NiftyconnectexchangeOrderCancelled // Event containing the contract specifics and raw log

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
func (it *NiftyconnectexchangeOrderCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NiftyconnectexchangeOrderCancelled)
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
		it.Event = new(NiftyconnectexchangeOrderCancelled)
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
func (it *NiftyconnectexchangeOrderCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NiftyconnectexchangeOrderCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NiftyconnectexchangeOrderCancelled represents a OrderCancelled event raised by the Niftyconnectexchange contract.
type NiftyconnectexchangeOrderCancelled struct {
	Hash [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterOrderCancelled is a free log retrieval operation binding the contract event 0x5152abf959f6564662358c2e52b702259b78bac5ee7842a0f01937e670efcc7d.
//
// Solidity: event OrderCancelled(bytes32 indexed hash)
func (_Niftyconnectexchange *NiftyconnectexchangeFilterer) FilterOrderCancelled(opts *bind.FilterOpts, hash [][32]byte) (*NiftyconnectexchangeOrderCancelledIterator, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}

	logs, sub, err := _Niftyconnectexchange.contract.FilterLogs(opts, "OrderCancelled", hashRule)
	if err != nil {
		return nil, err
	}
	return &NiftyconnectexchangeOrderCancelledIterator{contract: _Niftyconnectexchange.contract, event: "OrderCancelled", logs: logs, sub: sub}, nil
}

// WatchOrderCancelled is a free log subscription operation binding the contract event 0x5152abf959f6564662358c2e52b702259b78bac5ee7842a0f01937e670efcc7d.
//
// Solidity: event OrderCancelled(bytes32 indexed hash)
func (_Niftyconnectexchange *NiftyconnectexchangeFilterer) WatchOrderCancelled(opts *bind.WatchOpts, sink chan<- *NiftyconnectexchangeOrderCancelled, hash [][32]byte) (event.Subscription, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}

	logs, sub, err := _Niftyconnectexchange.contract.WatchLogs(opts, "OrderCancelled", hashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NiftyconnectexchangeOrderCancelled)
				if err := _Niftyconnectexchange.contract.UnpackLog(event, "OrderCancelled", log); err != nil {
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

// ParseOrderCancelled is a log parse operation binding the contract event 0x5152abf959f6564662358c2e52b702259b78bac5ee7842a0f01937e670efcc7d.
//
// Solidity: event OrderCancelled(bytes32 indexed hash)
func (_Niftyconnectexchange *NiftyconnectexchangeFilterer) ParseOrderCancelled(log types.Log) (*NiftyconnectexchangeOrderCancelled, error) {
	event := new(NiftyconnectexchangeOrderCancelled)
	if err := _Niftyconnectexchange.contract.UnpackLog(event, "OrderCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NiftyconnectexchangeOrdersMatchedIterator is returned from FilterOrdersMatched and is used to iterate over the raw logs and unpacked data for OrdersMatched events raised by the Niftyconnectexchange contract.
type NiftyconnectexchangeOrdersMatchedIterator struct {
	Event *NiftyconnectexchangeOrdersMatched // Event containing the contract specifics and raw log

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
func (it *NiftyconnectexchangeOrdersMatchedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NiftyconnectexchangeOrdersMatched)
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
		it.Event = new(NiftyconnectexchangeOrdersMatched)
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
func (it *NiftyconnectexchangeOrdersMatchedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NiftyconnectexchangeOrdersMatchedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NiftyconnectexchangeOrdersMatched represents a OrdersMatched event raised by the Niftyconnectexchange contract.
type NiftyconnectexchangeOrdersMatched struct {
	BuyHash                  [32]byte
	SellHash                 [32]byte
	Maker                    common.Address
	Taker                    common.Address
	MakerRelayerFeeRecipient common.Address
	TakerRelayerFeeRecipient common.Address
	Price                    *big.Int
	Metadata                 [32]byte
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterOrdersMatched is a free log retrieval operation binding the contract event 0x5e89bc5bf129d9595ae14697a763c17e6acd67579b9f1f4fa548f57ec762a057.
//
// Solidity: event OrdersMatched(bytes32 buyHash, bytes32 sellHash, address indexed maker, address indexed taker, address makerRelayerFeeRecipient, address takerRelayerFeeRecipient, uint256 price, bytes32 indexed metadata)
func (_Niftyconnectexchange *NiftyconnectexchangeFilterer) FilterOrdersMatched(opts *bind.FilterOpts, maker []common.Address, taker []common.Address, metadata [][32]byte) (*NiftyconnectexchangeOrdersMatchedIterator, error) {

	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}
	var takerRule []interface{}
	for _, takerItem := range taker {
		takerRule = append(takerRule, takerItem)
	}

	var metadataRule []interface{}
	for _, metadataItem := range metadata {
		metadataRule = append(metadataRule, metadataItem)
	}

	logs, sub, err := _Niftyconnectexchange.contract.FilterLogs(opts, "OrdersMatched", makerRule, takerRule, metadataRule)
	if err != nil {
		return nil, err
	}
	return &NiftyconnectexchangeOrdersMatchedIterator{contract: _Niftyconnectexchange.contract, event: "OrdersMatched", logs: logs, sub: sub}, nil
}

// WatchOrdersMatched is a free log subscription operation binding the contract event 0x5e89bc5bf129d9595ae14697a763c17e6acd67579b9f1f4fa548f57ec762a057.
//
// Solidity: event OrdersMatched(bytes32 buyHash, bytes32 sellHash, address indexed maker, address indexed taker, address makerRelayerFeeRecipient, address takerRelayerFeeRecipient, uint256 price, bytes32 indexed metadata)
func (_Niftyconnectexchange *NiftyconnectexchangeFilterer) WatchOrdersMatched(opts *bind.WatchOpts, sink chan<- *NiftyconnectexchangeOrdersMatched, maker []common.Address, taker []common.Address, metadata [][32]byte) (event.Subscription, error) {

	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}
	var takerRule []interface{}
	for _, takerItem := range taker {
		takerRule = append(takerRule, takerItem)
	}

	var metadataRule []interface{}
	for _, metadataItem := range metadata {
		metadataRule = append(metadataRule, metadataItem)
	}

	logs, sub, err := _Niftyconnectexchange.contract.WatchLogs(opts, "OrdersMatched", makerRule, takerRule, metadataRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NiftyconnectexchangeOrdersMatched)
				if err := _Niftyconnectexchange.contract.UnpackLog(event, "OrdersMatched", log); err != nil {
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

// ParseOrdersMatched is a log parse operation binding the contract event 0x5e89bc5bf129d9595ae14697a763c17e6acd67579b9f1f4fa548f57ec762a057.
//
// Solidity: event OrdersMatched(bytes32 buyHash, bytes32 sellHash, address indexed maker, address indexed taker, address makerRelayerFeeRecipient, address takerRelayerFeeRecipient, uint256 price, bytes32 indexed metadata)
func (_Niftyconnectexchange *NiftyconnectexchangeFilterer) ParseOrdersMatched(log types.Log) (*NiftyconnectexchangeOrdersMatched, error) {
	event := new(NiftyconnectexchangeOrdersMatched)
	if err := _Niftyconnectexchange.contract.UnpackLog(event, "OrdersMatched", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NiftyconnectexchangeOwnershipRenouncedIterator is returned from FilterOwnershipRenounced and is used to iterate over the raw logs and unpacked data for OwnershipRenounced events raised by the Niftyconnectexchange contract.
type NiftyconnectexchangeOwnershipRenouncedIterator struct {
	Event *NiftyconnectexchangeOwnershipRenounced // Event containing the contract specifics and raw log

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
func (it *NiftyconnectexchangeOwnershipRenouncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NiftyconnectexchangeOwnershipRenounced)
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
		it.Event = new(NiftyconnectexchangeOwnershipRenounced)
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
func (it *NiftyconnectexchangeOwnershipRenouncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NiftyconnectexchangeOwnershipRenouncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NiftyconnectexchangeOwnershipRenounced represents a OwnershipRenounced event raised by the Niftyconnectexchange contract.
type NiftyconnectexchangeOwnershipRenounced struct {
	PreviousOwner common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipRenounced is a free log retrieval operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: event OwnershipRenounced(address indexed previousOwner)
func (_Niftyconnectexchange *NiftyconnectexchangeFilterer) FilterOwnershipRenounced(opts *bind.FilterOpts, previousOwner []common.Address) (*NiftyconnectexchangeOwnershipRenouncedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _Niftyconnectexchange.contract.FilterLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return &NiftyconnectexchangeOwnershipRenouncedIterator{contract: _Niftyconnectexchange.contract, event: "OwnershipRenounced", logs: logs, sub: sub}, nil
}

// WatchOwnershipRenounced is a free log subscription operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: event OwnershipRenounced(address indexed previousOwner)
func (_Niftyconnectexchange *NiftyconnectexchangeFilterer) WatchOwnershipRenounced(opts *bind.WatchOpts, sink chan<- *NiftyconnectexchangeOwnershipRenounced, previousOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _Niftyconnectexchange.contract.WatchLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NiftyconnectexchangeOwnershipRenounced)
				if err := _Niftyconnectexchange.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
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

// ParseOwnershipRenounced is a log parse operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: event OwnershipRenounced(address indexed previousOwner)
func (_Niftyconnectexchange *NiftyconnectexchangeFilterer) ParseOwnershipRenounced(log types.Log) (*NiftyconnectexchangeOwnershipRenounced, error) {
	event := new(NiftyconnectexchangeOwnershipRenounced)
	if err := _Niftyconnectexchange.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NiftyconnectexchangeOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Niftyconnectexchange contract.
type NiftyconnectexchangeOwnershipTransferredIterator struct {
	Event *NiftyconnectexchangeOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *NiftyconnectexchangeOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NiftyconnectexchangeOwnershipTransferred)
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
		it.Event = new(NiftyconnectexchangeOwnershipTransferred)
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
func (it *NiftyconnectexchangeOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NiftyconnectexchangeOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NiftyconnectexchangeOwnershipTransferred represents a OwnershipTransferred event raised by the Niftyconnectexchange contract.
type NiftyconnectexchangeOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Niftyconnectexchange *NiftyconnectexchangeFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*NiftyconnectexchangeOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Niftyconnectexchange.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &NiftyconnectexchangeOwnershipTransferredIterator{contract: _Niftyconnectexchange.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Niftyconnectexchange *NiftyconnectexchangeFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *NiftyconnectexchangeOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Niftyconnectexchange.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NiftyconnectexchangeOwnershipTransferred)
				if err := _Niftyconnectexchange.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Niftyconnectexchange *NiftyconnectexchangeFilterer) ParseOwnershipTransferred(log types.Log) (*NiftyconnectexchangeOwnershipTransferred, error) {
	event := new(NiftyconnectexchangeOwnershipTransferred)
	if err := _Niftyconnectexchange.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
