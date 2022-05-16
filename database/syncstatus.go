package database

import (
	"time"

	"github.com/jinzhu/gorm"
)

type SyncStatus struct {
	Id int64

	Blockchain string

	LatestSyncHeight int64

	CreateTime int64
	UpdateTime int64
}

func (SyncStatus) TableName() string {
	return "sync_status"
}

func (l *SyncStatus) BeforeCreate() (err error) {
	l.CreateTime = time.Now().Unix()
	return nil
}

func (SyncStatus) InitTable(db *gorm.DB) {
	if !db.HasTable(&SyncStatus{}) {
		db.CreateTable(&SyncStatus{})
		db.Model(&SyncStatus{}).AddIndex("idx_"+SyncStatus{}.TableName()+"_blockchain", "blockchain")
	}
}
