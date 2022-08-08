package domain

import "gorm.io/gorm"

type Link struct {
	gorm.Model
	LongUrl   string `json:"long_url"`
	ShortUrl  string `json:"short_url"`
	IsDeleted bool   `json:"is_deleted"`
	OwnerID   uint   `json:"owner_id"`
	User      User   `json:"-" gorm:"foreignKey:OwnerID"`
}
