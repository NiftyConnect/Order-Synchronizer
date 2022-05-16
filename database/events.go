package database

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/jinzhu/gorm"
)

type OrderApprovedPartOne struct {
	Id         int64
	Blockchain string
	TxHash     string
	Height     uint64

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

func (OrderApprovedPartOne) TableName() string {
	return "order_approved_part_one"
}

func (OrderApprovedPartOne) InitTable(db *gorm.DB) {
	if !db.HasTable(&OrderApprovedPartOne{}) {
		db.CreateTable(&OrderApprovedPartOne{})
		db.Model(&OrderApprovedPartOne{}).AddUniqueIndex("idx_"+OrderApprovedPartOne{}.TableName()+"_blockchain", "blockchain")
		db.Model(&OrderApprovedPartOne{}).AddUniqueIndex("idx_"+OrderApprovedPartOne{}.TableName()+"_height", "height")
	}
}

type OrderApprovedPartTwo struct {
	Id         int64
	Blockchain string
	TxHash     string
	Height     uint64

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

func (OrderApprovedPartTwo) TableName() string {
	return "order_approved_part_two"
}

func (OrderApprovedPartTwo) InitTable(db *gorm.DB) {
	if !db.HasTable(&OrderApprovedPartTwo{}) {
		db.CreateTable(&OrderApprovedPartTwo{})
		db.Model(&OrderApprovedPartTwo{}).AddUniqueIndex("idx_"+OrderApprovedPartTwo{}.TableName()+"_blockchain", "blockchain")
		db.Model(&OrderApprovedPartTwo{}).AddUniqueIndex("idx_"+OrderApprovedPartTwo{}.TableName()+"_height", "height")
	}
}

type OrderCancelled struct {
	Id         int64
	Blockchain string
	TxHash     string
	Height     uint64

	Hash common.Hash
}

func (OrderCancelled) TableName() string {
	return "order_cancelled"
}

func (OrderCancelled) InitTable(db *gorm.DB) {
	if !db.HasTable(&OrderCancelled{}) {
		db.CreateTable(&OrderCancelled{})
		db.Model(&OrderCancelled{}).AddUniqueIndex("idx_"+OrderCancelled{}.TableName()+"_blockchain", "blockchain")
		db.Model(&OrderCancelled{}).AddUniqueIndex("idx_"+OrderCancelled{}.TableName()+"_height", "height")
	}
}

type OrdersMatched struct {
	Id         int64
	Blockchain string
	TxHash     string
	Height     uint64

	BuyHash                  common.Hash
	SellHash                 common.Hash
	Maker                    common.Address
	Taker                    common.Address
	MakerRelayerFeeRecipient common.Address
	TakerRelayerFeeRecipient common.Address
	Price                    *big.Int
	Metadata                 common.Hash
}

func (OrdersMatched) TableName() string {
	return "orders_matched"
}

func (OrdersMatched) InitTable(db *gorm.DB) {
	if !db.HasTable(&OrdersMatched{}) {
		db.CreateTable(&OrdersMatched{})
		db.Model(&OrdersMatched{}).AddUniqueIndex("idx_"+OrdersMatched{}.TableName()+"_blockchain", "blockchain")
		db.Model(&OrdersMatched{}).AddUniqueIndex("idx_"+OrdersMatched{}.TableName()+"_height", "height")
	}
}

type NonceIncremented struct {
	Id         int64
	Blockchain string
	TxHash     string
	Height     uint64

	Maker    common.Address
	NewNonce *big.Int
}

func (NonceIncremented) TableName() string {
	return "nonce_incremented"
}

func (NonceIncremented) InitTable(db *gorm.DB) {
	if !db.HasTable(&NonceIncremented{}) {
		db.CreateTable(&NonceIncremented{})
		db.Model(&NonceIncremented{}).AddUniqueIndex("idx_"+NonceIncremented{}.TableName()+"_blockchain", "blockchain")
		db.Model(&NonceIncremented{}).AddUniqueIndex("idx_"+NonceIncremented{}.TableName()+"_height", "height")
	}
}
