package localDB

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Connection(filename string) (*gorm.DB, error) {
	return gorm.Open(sqlite.Open(filename), &gorm.Config{})
}
