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

func (o OrderApprovedPartOne) toModelType(blockchain string, txHash string, height int64, position int64) database.OrderApprovedPartOne {
	return database.OrderApprovedPartOne{
		Blockchain:               blockchain,
		TxHash:                   txHash,
		Height:                   height,
		Position:                 position,
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

func (o OrderApprovedPartTwo) toModelType(blockchain string, txHash string, height int64, position int64) database.OrderApprovedPartTwo {
	return database.OrderApprovedPartTwo{
		Blockchain:         blockchain,
		TxHash:             txHash,
		Height:             height,
		Position:           position,
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

func (o OrderCancelled) toModelType(blockchain string, txHash string, height int64, position int64) database.OrderCancelled {
	return database.OrderCancelled{
		Blockchain: blockchain,
		TxHash:     txHash,
		Height:     height,
		Position:   position,
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

func (o OrdersMatched) toModelType(blockchain string, txHash string, height int64, position int64) database.OrdersMatched {
	return database.OrdersMatched{
		Blockchain:               blockchain,
		TxHash:                   txHash,
		Height:                   height,
		Position:                 position,
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

func (o NonceIncremented) toModelType(blockchain string, txHash string, height int64, position int64) database.NonceIncremented {
	return database.NonceIncremented{
		Blockchain: blockchain,
		TxHash:     txHash,
		Height:     height,
		Position:   position,
		Maker:      o.Maker.String(),
		NewNonce:   o.NewNonce.String(),
	}
}
