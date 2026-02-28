package sqlite

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(models ...interface{}) {
	var err error

	DB, err = gorm.Open(sqlite.Open("OpenList.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	if len(models) > 0 {
		DB.AutoMigrate(models...)
	}
}
