package database

import (
	"time"

	"github.com/jinzhu/gorm"
)

type OrderAnalysisStatus struct {
	Id int64

	Blockchain string

	LatestAnalysisHeight int64

	CreateTime int64
	UpdateTime int64
}

func (OrderAnalysisStatus) TableName() string {
	return "order_analysis_status"
}

func (l *OrderAnalysisStatus) BeforeCreate() (err error) {
	l.CreateTime = time.Now().Unix()
	return nil
}

func (OrderAnalysisStatus) InitTable(db *gorm.DB) {
	if !db.HasTable(&OrderAnalysisStatus{}) {
		db.CreateTable(&OrderAnalysisStatus{})
		db.Model(&OrderAnalysisStatus{}).AddIndex("idx_"+OrderAnalysisStatus{}.TableName()+"_blockchain", "blockchain")
	}
}
