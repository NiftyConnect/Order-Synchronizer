package database

import (
	"time"

	"github.com/jinzhu/gorm"
)

type BlockLog struct {
	Id         int64
	Chain      string
	BlockHash  string
	ParentHash string
	Height     int64
	BlockTime  int64
	CreateTime int64
}

func (BlockLog) TableName() string {
	return "block_log"
}

func (l *BlockLog) BeforeCreate() (err error) {
	l.CreateTime = time.Now().Unix()
	return nil
}

func (BlockLog) InitTable(db *gorm.DB) {
	if !db.HasTable(&BlockLog{}) {
		db.CreateTable(&BlockLog{})
		db.Model(&BlockLog{}).AddUniqueIndex("idx_"+BlockLog{}.TableName()+"_blockchain", "blockchain")
		db.Model(&BlockLog{}).AddUniqueIndex("idx_"+BlockLog{}.TableName()+"_height", "height")
	}
}
