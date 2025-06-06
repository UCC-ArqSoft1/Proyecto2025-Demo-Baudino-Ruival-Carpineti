package clients

import (
	"backend/dao"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBClient struct {
	DB *gorm.DB
}

func NewDBClient() *DBClient {
	dsnFormat := "%s:%s@tcp(%s:%d)/%s?parseTime=true&charset=utf8mb4&loc=Local"
	dsn := fmt.Sprintf(dsnFormat, "root", "Dacota12", "127.0.0.1", 3306, "gimnasio")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Errorf("error connecting to database: %w", err))
	}

	// Auto-migrate all models
	for _, table := range []interface{}{
		&dao.User{},
		&dao.Inscription{},
		&dao.Activities{},
		&dao.Schedules{},
	} {
		if err := db.AutoMigrate(&table); err != nil {
			panic(fmt.Errorf("error migrating table: %w", err))
		}
	}

	return &DBClient{
		DB: db,
	}
}
