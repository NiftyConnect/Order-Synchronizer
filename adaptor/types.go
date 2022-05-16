package adaptor

import (
	"encoding/hex"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/niftyConnect/order-synchronizer/database"
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

func (o OrderApprovedPartOne) toModelType(Blockchain string, TxHash string, Height uint64) database.OrderApprovedPartOne {
	return database.OrderApprovedPartOne{
		Blockchain:               Blockchain,
		TxHash:                   TxHash,
		Height:                   Height,
		Hash:                     o.Hash.String(),
		Exchange:                 o.Exchange.String(),
		Maker:                    o.Maker.String(),
		Taker:                    o.Taker.String(),
		MakerRelayerFeeRecipient: o.MakerRelayerFeeRecipient.String(),
		Side:                     o.Side,
		SaleKind:                 o.SaleKind,
		NftAddress:               o.NftAddress.String(),
		TokenId:                  o.TokenId.String(),
		IpfsHash:                 o.IpfsHash.String(),
	}
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

func (o OrderApprovedPartTwo) toModelType(Blockchain string, TxHash string, Height uint64) database.OrderApprovedPartTwo {
	return database.OrderApprovedPartTwo{
		Blockchain:         Blockchain,
		TxHash:             TxHash,
		Height:             Height,
		Hash:               o.Hash.String(),
		Calldata:           "0x" + hex.EncodeToString(o.Calldata),
		ReplacementPattern: "0x" + hex.EncodeToString(o.ReplacementPattern),
		StaticTarget:       o.StaticTarget.String(),
		StaticExtradata:    "0x" + hex.EncodeToString(o.StaticExtradata),
		PaymentToken:       o.PaymentToken.String(),
		BasePrice:          o.BasePrice.String(),
		Extra:              o.Extra.String(),
		ListingTime:        o.ListingTime.String(),
		ExpirationTime:     o.ExpirationTime.String(),
		Salt:               o.Salt.String(),
	}
}

type OrderCancelled struct {
	Hash common.Hash
}

func (o OrderCancelled) toModelType(Blockchain string, TxHash string, Height uint64) database.OrderCancelled {
	return database.OrderCancelled{
		Blockchain: Blockchain,
		TxHash:     TxHash,
		Height:     Height,
		Hash:       o.Hash.String(),
	}
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

func (o OrdersMatched) toModelType(Blockchain string, TxHash string, Height uint64) database.OrdersMatched {
	return database.OrdersMatched{
		Blockchain:               Blockchain,
		TxHash:                   TxHash,
		Height:                   Height,
		BuyHash:                  o.BuyHash.String(),
		SellHash:                 o.SellHash.String(),
		Maker:                    o.Maker.String(),
		Taker:                    o.Taker.String(),
		MakerRelayerFeeRecipient: o.MakerRelayerFeeRecipient.String(),
		TakerRelayerFeeRecipient: o.TakerRelayerFeeRecipient.String(),
		Price:                    o.Price.String(),
		Metadata:                 o.Metadata.String(),
	}
}

type NonceIncremented struct {
	Maker    common.Address
	NewNonce *big.Int
}

func (o NonceIncremented) toModelType(Blockchain string, TxHash string, Height uint64) database.NonceIncremented {
	return database.NonceIncremented{
		Blockchain: Blockchain,
		TxHash:     TxHash,
		Height:     Height,
		Maker:      o.Maker.String(),
		NewNonce:   o.NewNonce.String(),
	}
}
