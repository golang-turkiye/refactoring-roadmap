package model

import (
	. "github.com/Golang-Turkiye/refactoring-roadmap/db/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email         string `json:"email"`
	Password      string `json:"password"`
	RememberToken string `json:"remember_token"`
	Links         []Link `json:"-" gorm:"foreignKey:OwnerID"`
}

type Users []User

func (u *User) FetchOne(db *gorm.DB, ID uint) error {
	return FetchByID(db, u, ID)
}

func (u *User) Save(db *gorm.DB) error {
	return Save(db, u)
}

func (u *User) Update(db *gorm.DB) error {
	return Update(db, u)
}

func (u *User) AutoMigrate(db *gorm.DB) error {
	return Migrate(db, u)
}

func GetUserByUsernamePassword(db *gorm.DB, username string, password string) (*User, error) {
	user := &User{}
	return user, db.Where("email = ? AND password = ?", username, password).First(user).Error
}
