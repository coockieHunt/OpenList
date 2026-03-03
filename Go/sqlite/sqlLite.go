package sqlite

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	_ "modernc.org/sqlite"
)

var DB *gorm.DB

func InitDB(models ...interface{}) {
	var err error

	DB, err = gorm.Open(sqlite.New(sqlite.Config{
		DriverName: "sqlite",
		DSN:        "OpenList.db",
	}))

	if err != nil {
		panic("failed to connect database")
	}

	sqlDB, err := DB.DB()
	if err == nil {
		sqlDB.Exec("PRAGMA journal_mode=WAL;")   //tread multi
		sqlDB.Exec("PRAGMA synchronous=NORMAL;") //sync async
		sqlDB.Exec("PRAGMA cache_size=-8000;")   //cache 8MB
		sqlDB.Exec("PRAGMA foreign_keys=ON;")    //foreign key
	}

	if len(models) > 0 {
		DB.AutoMigrate(models...)
	}
}
