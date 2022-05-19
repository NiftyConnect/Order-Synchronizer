package database

import (
	"time"

	"github.com/jinzhu/gorm"
)

type NiftyConnectOrder struct {
	Id int64

	Blockchain string
	Height     int64
	Position   int64

	OrderHash                string
	TxHash                   string
	Exchange                 string
	Maker                    string
	Taker                    string
	MakerRelayerFeeRecipient string
	TakerRelayerFeeRecipient string
	Side                     uint8
	SaleKind                 uint8
	NftAddress               string
	TokenId                  string
	IpfsHash                 string
	Calldata                 string
	ReplacementPattern       string
	StaticTarget             string
	StaticExtradata          string
	PaymentToken             string
	OrderPrice               string
	Extra                    string
	ListingTime              string
	ExpirationTime           string
	Salt                     string

	IsCancelled bool
	IsExpired   bool
	IsFinalized bool

	CreateTime int64
	UpdateTime int64
}

func (NiftyConnectOrder) TableName() string {
	return "nifty_connect_order"
}

func (l *NiftyConnectOrder) BeforeCreate() (err error) {
	l.CreateTime = time.Now().Unix()
	return nil
}

func (NiftyConnectOrder) InitTable(db *gorm.DB) {
	if !db.HasTable(&NiftyConnectOrder{}) {
		db.CreateTable(&NiftyConnectOrder{})
		db.Model(&NiftyConnectOrder{}).AddUniqueIndex("idx_"+NiftyConnectOrder{}.TableName()+"_order_hash", "order_hash")

		db.Model(&NiftyConnectOrder{}).AddIndex("idx_"+NiftyConnectOrder{}.TableName()+"_nft_address_token_id", "nft_address", "token_id")
		db.Model(&NiftyConnectOrder{}).AddIndex("idx_"+NiftyConnectOrder{}.TableName()+"_nft_address", "nft_address")
		db.Model(&NiftyConnectOrder{}).AddIndex("idx_"+NiftyConnectOrder{}.TableName()+"_token_id", "token_id")

		db.Model(&NiftyConnectOrder{}).AddIndex("idx_"+NiftyConnectOrder{}.TableName()+"_maker", "maker")
		db.Model(&NiftyConnectOrder{}).AddIndex("idx_"+NiftyConnectOrder{}.TableName()+"_side", "side")
		db.Model(&NiftyConnectOrder{}).AddIndex("idx_"+NiftyConnectOrder{}.TableName()+"_sale_kind", "sale_kind")
		db.Model(&NiftyConnectOrder{}).AddIndex("idx_"+NiftyConnectOrder{}.TableName()+"_expiration_time", "expiration_time")

		db.Model(&NiftyConnectOrder{}).AddIndex("idx_"+NiftyConnectOrder{}.TableName()+"_is_cancelled", "is_cancelled")
		db.Model(&NiftyConnectOrder{}).AddIndex("idx_"+NiftyConnectOrder{}.TableName()+"_is_expired", "is_expired")
		db.Model(&NiftyConnectOrder{}).AddIndex("idx_"+NiftyConnectOrder{}.TableName()+"_is_finalized", "is_finalized")

		db.Model(&NiftyConnectOrder{}).AddIndex("idx_"+NiftyConnectOrder{}.TableName()+"_blockchain", "blockchain")
	}
}
