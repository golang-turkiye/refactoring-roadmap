package domain

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email         string `json:"email"`
	Password      string `json:"password"`
	RememberToken string `json:"remember_token"`
	Links         []Link `json:"-" gorm:"foreignKey:OwnerID"`
}
