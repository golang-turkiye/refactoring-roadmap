package model

import (
	"fmt"
	. "github.com/Golang-Turkiye/refactoring-roadmap/db/mysql"
	"gorm.io/gorm"
)

type Link struct {
	gorm.Model
	LongUrl        string `json:"long_url"`
	ShortenURLPath string `json:"shorten_url_path"`
	IsDeleted      bool   `json:"is_deleted"`
	OwnerID        uint   `json:"owner_id"`
	User           User   `json:"-" gorm:"foreignKey:OwnerID"`
}

type Links []Link

func (l *Link) FetchOne(db *gorm.DB, ID uint) error {
	return FetchByID(db, l, ID)
}

func (l *Link) FetchByCondition(db *gorm.DB, condition []string) error {
	return FetchByCond(db, l, condition)
}

func (l *Link) Save(db *gorm.DB) error {
	return Save(db, l)
}

func (l *Link) Update(db *gorm.DB) error {
	return Update(db, l)
}

func (l *Link) AutoMigrate(db *gorm.DB) error {
	return Migrate(db, &l)
}

func (ls *Links) FetchAll(db *gorm.DB, userID uint) error {
	condition := []string{fmt.Sprintf("owner_id:%d", userID), "is_deleted:false"}
	return FetchAll(db, ls, condition)
}
