package clients

import (
	"backend/domain"

	"gorm.io/gorm"
)

type MySQLActivitiesClient struct {
	DB *gorm.DB
}

func NewMySQLActivitiesClient(db *gorm.DB) *MySQLActivitiesClient {
	return &MySQLActivitiesClient{
		DB: db,
	}
}

func (c *MySQLActivitiesClient) GetActivities() []domain.Activity {
	var activities []domain.Activity
	c.DB.Find(&activities)
	return activities
}

func (c *MySQLActivitiesClient) GetActivityByID(id int) (domain.Activity, error) {
	var activity domain.Activity
	result := c.DB.First(&activity, id)
	return activity, result.Error
}

func (c *MySQLActivitiesClient) SearchActivities(category, keyword string) []domain.Activity {
	var activities []domain.Activity
	query := c.DB.Model(&domain.Activity{})

	if category != "" {
		query = query.Where("category = ?", category)
	}
	if keyword != "" {
		query = query.Where("name LIKE ? OR description LIKE ?", "%"+keyword+"%", "%"+keyword+"%")
	}

	query.Find(&activities)
	return activities
}

func (c *MySQLActivitiesClient) GetActivitiesByUserID(userID int) []domain.Activity {
	var activities []domain.Activity
	c.DB.Joins("JOIN inscriptions ON activities.id = inscriptions.activity_id").
		Where("inscriptions.user_id = ?", userID).
		Find(&activities)
	return activities
}

func (c *MySQLActivitiesClient) EnrollUserInActivity(userID, scheduleID int) error {
	inscription := domain.Inscription{
		UserID:     userID,
		ScheduleID: scheduleID,
	}
	return c.DB.Create(&inscription).Error
}
