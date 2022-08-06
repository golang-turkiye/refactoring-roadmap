package domain

import "gorm.io/gorm"

type Link struct {
	gorm.Model
	LongUrl        string `json:"long_url"`
	ShortenURLPath string `json:"shorten_url_path"`
	IsDeleted      bool   `json:"is_deleted"`
	OwnerID        uint   `json:"owner_id"`
	User           User   `json:"-" gorm:"foreignKey:OwnerID"`
}
