package database

import (
	"github.com/jinzhu/gorm"
)

func InitTables(db *gorm.DB) {
	NiftyConnectOrder{}.InitTable(db)
	BlockLog{}.InitTable(db)
	SyncStatus{}.InitTable(db)

	OrderApprovedPartOne{}.InitTable(db)
	OrderApprovedPartTwo{}.InitTable(db)
	OrderCancelled{}.InitTable(db)
	OrdersMatched{}.InitTable(db)
	NonceIncremented{}.InitTable(db)
}
