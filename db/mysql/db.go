package mysql

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connection() (*gorm.DB, error) {
	return gorm.Open(mysql.Open("db1_dsn"), &gorm.Config{})
}

func Migrate(db *gorm.DB, model interface{}) error {
	return db.AutoMigrate(model)
}

func FetchByID(db *gorm.DB, model interface{}, ID uint) error {
	return db.Model(model).First(model, ID).Error
}

func FetchByCond(db *gorm.DB, model interface{}, condition []string) error {
	wheres := ""
	for _, cond := range condition {
		wheres += cond + ","
	}
	wheres = wheres[0 : len(wheres)-2]
	return db.Model(model).Where(wheres).Find(model).Error
}

func FetchAll(db *gorm.DB, model interface{}, condition []string) error {
	wheres := ""
	for _, cond := range condition {
		wheres += cond + ","
	}
	wheres = wheres[0 : len(wheres)-2]
	return db.Model(model).Where(wheres).Find(model).Error
}

func Save(db *gorm.DB, model interface{}) error {
	return db.Model(model).Save(model).Error
}

func Update(db *gorm.DB, model interface{}) error {
	return db.Model(model).Updates(model).Error
}
