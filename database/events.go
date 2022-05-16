package database

import (
	"github.com/jinzhu/gorm"
)

type OrderApprovedPartOne struct {
	Id         int64
	Blockchain string
	TxHash     string
	Height     uint64

	Hash                     string
	Exchange                 string
	Maker                    string
	Taker                    string
	MakerRelayerFeeRecipient string
	Side                     uint8
	SaleKind                 uint8
	NftAddress               string
	TokenId                  string
	IpfsHash                 string
}

func (OrderApprovedPartOne) TableName() string {
	return "order_approved_part_one"
}

func (OrderApprovedPartOne) InitTable(db *gorm.DB) {
	if !db.HasTable(&OrderApprovedPartOne{}) {
		db.CreateTable(&OrderApprovedPartOne{})
		db.Model(&OrderApprovedPartOne{}).AddIndex("idx_"+OrderApprovedPartOne{}.TableName()+"_blockchain", "blockchain")
		db.Model(&OrderApprovedPartOne{}).AddIndex("idx_"+OrderApprovedPartOne{}.TableName()+"_height", "height")
	}
}

type OrderApprovedPartTwo struct {
	Id         int64
	Blockchain string
	TxHash     string
	Height     uint64

	Hash               string
	Calldata           string
	ReplacementPattern string
	StaticTarget       string
	StaticExtradata    string
	PaymentToken       string
	BasePrice          string
	Extra              string
	ListingTime        string
	ExpirationTime     string
	Salt               string
}

func (OrderApprovedPartTwo) TableName() string {
	return "order_approved_part_two"
}

func (OrderApprovedPartTwo) InitTable(db *gorm.DB) {
	if !db.HasTable(&OrderApprovedPartTwo{}) {
		db.CreateTable(&OrderApprovedPartTwo{})
		db.Model(&OrderApprovedPartTwo{}).AddIndex("idx_"+OrderApprovedPartTwo{}.TableName()+"_blockchain", "blockchain")
		db.Model(&OrderApprovedPartTwo{}).AddIndex("idx_"+OrderApprovedPartTwo{}.TableName()+"_height", "height")
	}
}

type OrderCancelled struct {
	Id         int64
	Blockchain string
	TxHash     string
	Height     uint64

	Hash string
}

func (OrderCancelled) TableName() string {
	return "order_cancelled"
}

func (OrderCancelled) InitTable(db *gorm.DB) {
	if !db.HasTable(&OrderCancelled{}) {
		db.CreateTable(&OrderCancelled{})
		db.Model(&OrderCancelled{}).AddIndex("idx_"+OrderCancelled{}.TableName()+"_blockchain", "blockchain")
		db.Model(&OrderCancelled{}).AddIndex("idx_"+OrderCancelled{}.TableName()+"_height", "height")
	}
}

type OrdersMatched struct {
	Id         int64
	Blockchain string
	TxHash     string
	Height     uint64

	BuyHash                  string
	SellHash                 string
	Maker                    string
	Taker                    string
	MakerRelayerFeeRecipient string
	TakerRelayerFeeRecipient string
	Price                    string
	Metadata                 string
}

func (OrdersMatched) TableName() string {
	return "orders_matched"
}

func (OrdersMatched) InitTable(db *gorm.DB) {
	if !db.HasTable(&OrdersMatched{}) {
		db.CreateTable(&OrdersMatched{})
		db.Model(&OrdersMatched{}).AddIndex("idx_"+OrdersMatched{}.TableName()+"_blockchain", "blockchain")
		db.Model(&OrdersMatched{}).AddIndex("idx_"+OrdersMatched{}.TableName()+"_height", "height")
	}
}

type NonceIncremented struct {
	Id         int64
	Blockchain string
	TxHash     string
	Height     uint64

	Maker    string
	NewNonce string
}

func (NonceIncremented) TableName() string {
	return "nonce_incremented"
}

func (NonceIncremented) InitTable(db *gorm.DB) {
	if !db.HasTable(&NonceIncremented{}) {
		db.CreateTable(&NonceIncremented{})
		db.Model(&NonceIncremented{}).AddIndex("idx_"+NonceIncremented{}.TableName()+"_blockchain", "blockchain")
		db.Model(&NonceIncremented{}).AddIndex("idx_"+NonceIncremented{}.TableName()+"_height", "height")
	}
}
