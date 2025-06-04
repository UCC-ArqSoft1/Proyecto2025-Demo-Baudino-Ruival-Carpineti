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
	dsn := fmt.Sprintf(dsnFormat, "root", "", "127.0.0.1", 3306, "gimnasio")
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
	// SELECT * FROM activities
	result := c.DB.Preload("Horarios").Find(&activities)
	if result.Error != nil {
		return nil, fmt.Errorf("error getting activities: %w", result.Error)
	}
	return activities, nil
}

func (c *MySQLClient) GetActivityByID(id int) (dao.Activities, error) {
	var activity dao.Activities
	// SELECT * FROM activities WHERE id = ? LIMIT 1
	result := c.DB.Preload("Horarios").First(&activity, id)
	if result.Error != nil {
		return dao.Activities{}, fmt.Errorf("error getting activity by ID: %w", result.Error)
	}
	return activity, nil
}


func (c *MySQLClient) GetUserInscriptions(userID int) ([]dao.Inscription, error) {
	var inscriptions []dao.Inscription
	// SELECT * FROM inscriptions WHERE usuario_id = ?
	result := c.DB.Where("usuario_id = ?", userID).Find(&inscriptions)
	if result.Error != nil {
		return nil, fmt.Errorf("error getting user inscriptions: %w", result.Error)
	}
	return inscriptions, nil
}

func (c *MySQLClient) GetScheduleByID(scheduleID int) (dao.Schedules, error) {
	var schedule dao.Schedules
	// SELECT * FROM schedules WHERE id = ? LIMIT 1
	result := c.DB.First(&schedule, scheduleID)
	if result.Error != nil {
		return dao.Schedules{}, fmt.Errorf("error getting schedule: %w", result.Error)
	}
	return schedule, nil
}

func (c *MySQLClient) CheckExistingEnrollment(userID, scheduleID int) (bool, error) {
	var enrollment dao.Inscription
	// SELECT * FROM inscriptions WHERE usuario_id = ? AND horario_id = ? LIMIT 1
	result := c.DB.Where("usuario_id = ? AND horario_id = ?", userID, scheduleID).First(&enrollment)
	if result.Error == nil {
		return true, nil
	}
	if result.Error == gorm.ErrRecordNotFound {
		return false, nil
	}
	return false, fmt.Errorf("error checking enrollment: %w", result.Error)
}

func (c *MySQLClient) CreateEnrollment(enrollment dao.Inscription) error {
	// INSERT INTO inscriptions (usuario_id, horario_id) VALUES (?, ?)
	result := c.DB.Create(&enrollment)
	if result.Error != nil {
		return fmt.Errorf("error creating enrollment: %w", result.Error)
	}
	return nil
}

func (c *MySQLClient) UpdateScheduleCapacity(scheduleID int) error {
	// UPDATE schedules SET cupo = cupo - 1 WHERE id = ?
	result := c.DB.Model(&dao.Schedules{}).Where("id = ?", scheduleID).
		Update("cupo", gorm.Expr("cupo - 1"))
	if result.Error != nil {
		return fmt.Errorf("error updating schedule capacity: %w", result.Error)
	}
	return nil
}
