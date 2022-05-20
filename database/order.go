package database

import (
	"time"

	"github.com/jinzhu/gorm"
)

type NiftyConnectOrder struct {
	Id int64 `json:"-"`

	Blockchain string `json:"-"`
	Height     int64  `json:"-"`
	Position   int64  `json:"-"`

	OrderHash                string `json:"order_hash"`
	TxHash                   string `json:"tx_hash"`
	Exchange                 string `json:"exchange"`
	Maker                    string `json:"maker"`
	Taker                    string `json:"taker"`
	MakerRelayerFeeRecipient string `json:"maker_relayer_fee_recipient"`
	TakerRelayerFeeRecipient string `json:"taker_relayer_fee_recipient"`
	Side                     uint8  `json:"side"`
	SaleKind                 uint8  `json:"sale_kind"`
	NftAddress               string `json:"nft_address"`
	TokenId                  string `json:"token_id"`
	IpfsHash                 string `json:"ipfs_hash"`
	Calldata                 string `json:"calldata"`
	ReplacementPattern       string `json:"replacement_pattern"`
	StaticTarget             string `json:"static_target"`
	StaticExtradata          string `json:"static_extradata"`
	PaymentToken             string `json:"payment_token"`
	OrderPrice               string `json:"order_price"`
	Extra                    string `json:"extra"`
	ListingTime              string `json:"listing_time"`
	ExpirationTime           string `json:"expiration_time"`
	Salt                     string `json:"salt"`

	IsCancelled bool `json:"is_cancelled"`
	IsExpired   bool `json:"is_expired"`
	IsFinalized bool `json:"is_finalized"`

	CreateTime int64 `json:"-"`
	UpdateTime int64 `json:"-"`
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
		db.Model(&NiftyConnectOrder{}).AddUniqueIndex("idx_"+NiftyConnectOrder{}.TableName()+"_blockchain_order_hash", "blockchain", "order_hash")

		db.Model(&NiftyConnectOrder{}).AddIndex("idx_"+NiftyConnectOrder{}.TableName()+"_blockchain_nft_address_token_id", "blockchain", "nft_address", "token_id")
		db.Model(&NiftyConnectOrder{}).AddIndex("idx_"+NiftyConnectOrder{}.TableName()+"_blockchain_nft_address", "blockchain", "nft_address")
		db.Model(&NiftyConnectOrder{}).AddIndex("idx_"+NiftyConnectOrder{}.TableName()+"_blockchain_token_id", "blockchain", "token_id")

		db.Model(&NiftyConnectOrder{}).AddIndex("idx_"+NiftyConnectOrder{}.TableName()+"_blockchain_maker", "blockchain", "maker")
		db.Model(&NiftyConnectOrder{}).AddIndex("idx_"+NiftyConnectOrder{}.TableName()+"_blockchain_side", "blockchain", "side")
		db.Model(&NiftyConnectOrder{}).AddIndex("idx_"+NiftyConnectOrder{}.TableName()+"_blockchain_sale_kind", "blockchain", "sale_kind")
		db.Model(&NiftyConnectOrder{}).AddIndex("idx_"+NiftyConnectOrder{}.TableName()+"_blockchain_expiration_time", "blockchain", "expiration_time")

		db.Model(&NiftyConnectOrder{}).AddIndex("idx_"+NiftyConnectOrder{}.TableName()+"_blockchain_is_cancelled", "blockchain", "is_cancelled")
		db.Model(&NiftyConnectOrder{}).AddIndex("idx_"+NiftyConnectOrder{}.TableName()+"_blockchain_is_expired", "blockchain", "is_expired")
		db.Model(&NiftyConnectOrder{}).AddIndex("idx_"+NiftyConnectOrder{}.TableName()+"_blockchain_is_finalized", "blockchain", "is_finalized")
	}
}
