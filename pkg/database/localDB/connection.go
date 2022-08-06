package localDB

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Connection() (*gorm.DB, error) {
	return gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
}
