package clients

import (
	"backend/domain"
)

type ActivitiesClient interface {
	GetActivities() []domain.Activity
	GetActivityByID(id int) (domain.Activity, error)
	SearchActivities(category, keyword string) []domain.Activity
	GetActivitiesByUserID(userID int) []domain.Activity
	EnrollUserInActivity(userID, scheduleID int) error
}

type MySQLActivitiesClient struct {
	client *MySQLClient
}

func NewMySQLActivitiesClient(client *MySQLClient) *MySQLActivitiesClient {
	return &MySQLActivitiesClient{
		client: client,
	}
}

func (c *MySQLActivitiesClient) GetActivities() []domain.Activity {
	var activities []domain.Activity
	c.client.DB.Find(&activities)
	return activities
}

func (c *MySQLActivitiesClient) GetActivityByID(id int) (domain.Activity, error) {
	var activity domain.Activity
	result := c.client.DB.First(&activity, id)
	return activity, result.Error
}

func (c *MySQLActivitiesClient) SearchActivities(category, keyword string) []domain.Activity {
	var activities []domain.Activity
	query := c.client.DB.Model(&domain.Activity{})

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
	c.client.DB.Joins("JOIN inscriptions ON activities.id = inscriptions.activity_id").
		Where("inscriptions.user_id = ?", userID).
		Find(&activities)
	return activities
}

func (c *MySQLActivitiesClient) EnrollUserInActivity(userID, scheduleID int) error {
	inscription := domain.Inscription{
		UserID:     userID,
		ScheduleID: scheduleID,
	}
	return c.client.DB.Create(&inscription).Error
}
