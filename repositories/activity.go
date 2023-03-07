package repositories

import (
	"golang-api-todolist/entities"
	"log"

	// "github.com/jinzhu/gorm"
	"gorm.io/gorm"
)

type activityRepository struct {
	DB *gorm.DB
}

type ActivityRepository interface {
	Create(entities.Activity) (entities.Activity, error)
	GetAllActivities() ([]entities.Activity, error)
	GetOneActivity(int) (entities.Activity, error)
	UpdateActivity(int, entities.Activity) (entities.Activity, error)
	DeleteActivity(entities.Activity) (entities.Activity, error)
}

func NewActivityRepository(db *gorm.DB) activityRepository {
	return activityRepository{
		DB: db,
	}
}

func (a activityRepository) Create(activity entities.Activity) (entities.Activity, error) {
	log.Print("[ActivityRepository]...Create")
	err := a.DB.Create(&activity).Error
	return activity, err
}

func (a activityRepository) GetAllActivities() (activities []entities.Activity, err error) {
	log.Print("[ActivityRepository]...GetAllActivities")
	err = a.DB.Find(&activities).Error
	return activities, err
}

func (a activityRepository) GetOneActivity(id int) (activity entities.Activity, err error) {
	log.Print("[ActivityRepository]...GetOneActivity")
	err = a.DB.Where("activity_id = ?", id).First(&activity).Error
	return activity, err

}

func (a activityRepository) UpdateActivity(id int, activity entities.Activity) (entities.Activity, error) {
	log.Print("[ActivityRepository]...UpdateActivity")
	var activitysave entities.Activity
	err := a.DB.Where("activity_id = ?", id).First(&activitysave).Error
	// if err := a.DB.First(&activitysave, activity.ActivityID).Error; err != nil {
	// return activity, err
	// }
	if err != nil {
		// fmt.Println(err)
		return activity, err
	}
	// a.DB.Save(&activity)
	return activity, a.DB.Model(&activitysave).Updates(activity).Error

}

func (a activityRepository) DeleteActivity(activity entities.Activity) (entities.Activity, error) {
	log.Print("[ActivityRepository]...DeleteActivity")
	// var activity entities.Activity
	// del := a.DB.Where("activity_id = ?", id).Delete(&activity)
	if err := a.DB.First(&activity, activity.ActivityID).Error; err != nil {
		return activity, err
	}
	return activity, a.DB.Delete(&activity).Error
}
