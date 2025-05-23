package clients

import (
	"backend/dao"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MySQLClient struct {
	DB *gorm.DB
}

func NewMySQLClient() *MySQLClient {
	dsnFormat := "%s:%s@tcp(%s:%d)/%s?parseTime=true&charset=utf8mb4&loc=Local"
	dsn := fmt.Sprintf(dsnFormat, "root", "root", "127.0.0.1", 3306, "gimnasio")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Errorf("error connecting to database: %w", err))
	}

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
	return &MySQLClient{
		DB: db,
	}
}

func (c *MySQLClient) GetUserByUsername(username string) (dao.User, error) {
	var userDAO dao.User
	//SELECT * FROM users WHERE username = "admin" LIMIT 1
	txn := c.DB.First(&userDAO, "username = ?", username)
	if txn.Error != nil {
		return dao.User{}, fmt.Errorf("error getting user: %w", txn.Error)
	}
	return userDAO, nil
}

func (c *MySQLClient) GetAllActivities() ([]dao.Activities, error) {
	var activities []dao.Activities
	result := c.DB.Preload("Horarios").Find(&activities)
	if result.Error != nil {
		return nil, fmt.Errorf("error getting activities: %w", result.Error)
	}
	return activities, nil
}

func (c *MySQLClient) GetActivityByID(id int) (dao.Activities, error) {
	var activity dao.Activities
	result := c.DB.Preload("Horarios").First(&activity, id)
	if result.Error != nil {
		return dao.Activities{}, fmt.Errorf("error getting activity by ID: %w", result.Error)
	}
	return activity, nil
}

/*
func (c *MySQLClient) CreateActivity(activity dao.Activity) (int,error){
	txn := c.DB.Create(&activity)
	if txn.Error != nil {
		return 0, fmt.Errorf("error creating activity: %w", txn.Error)
	}
	return activity.ID, nil
}
*/
