package adaptor

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type OrderApprovedPartOne struct {
	Hash                     common.Hash
	Exchange                 common.Address
	Maker                    common.Address
	Taker                    common.Address
	MakerRelayerFeeRecipient common.Address
	Side                     uint8
	SaleKind                 uint8
	NftAddress               common.Address
	TokenId                  *big.Int
	IpfsHash                 common.Hash
}

type OrderApprovedPartTwo struct {
	Hash               common.Hash
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
}

type OrderCancelled struct {
	Hash common.Hash
}

type OrdersMatched struct {
	BuyHash                  common.Hash
	SellHash                 common.Hash
	Maker                    common.Address
	Taker                    common.Address
	MakerRelayerFeeRecipient common.Address
	TakerRelayerFeeRecipient common.Address
	Price                    *big.Int
	Metadata                 common.Hash
}

type NonceIncremented struct {
	Maker    common.Address
	NewNonce *big.Int
}
