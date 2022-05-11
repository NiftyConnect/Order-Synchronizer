package database

import (
	"github.com/jinzhu/gorm"
)

func InitTables(db *gorm.DB) {
	NiftyConnectOrder{}.InitTable(db)
}