package db

import (
	"backend/dao"
	"fmt"
	"log"
	"os"

	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	dsnFormat := "%s:%s@tcp(%s:%s)/%s?parseTime=true&charset=utf8mb4&loc=Local"
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dsn := fmt.Sprintf(dsnFormat, user, password, host, port, dbName)

	db, err := waitForDB(dsn)
	if err != nil {
		panic(fmt.Errorf("error connecting to database: %w", err))
	}

	for _, table := range []interface{}{
		&dao.User{},
		&dao.Inscription{},
		&dao.Activities{},
		&dao.Schedules{},
	} {
		if err := db.AutoMigrate(table); err != nil {
			panic(fmt.Errorf("error migrating table: %w", err))
		}
	}
	return db
}

func waitForDB(dsn string) (*gorm.DB, error) {
	var db *gorm.DB
	var err error

	for i := 0; i < 10; i++ {
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err == nil {
			sqlDB, err := db.DB()
			if err == nil && sqlDB.Ping() == nil {
				return db, nil
			}
		}

		log.Printf("⏳ Esperando base de datos... intento %d/10\n", i+1)
		time.Sleep(3 * time.Second)
	}

	return nil, fmt.Errorf("❌ no se pudo conectar a la BD: %v", err)
}
