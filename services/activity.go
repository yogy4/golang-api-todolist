package services

import (
	"golang-api-todolist/entities"
	"golang-api-todolist/repositories"
)

type activityService struct {
	activityRepository repositories.ActivityRepository
}

type ActivityService interface {
	Create(entities.Activity) (entities.Activity, error)
	GetAllActivities() ([]entities.Activity, error)
	GetOneActivity(int) (entities.Activity, error)
	UpdateActivity(int, entities.Activity) (entities.Activity, error)
	DeleteActivity(entities.Activity) (entities.Activity, error)
}

func NewActivityService(r repositories.ActivityRepository) ActivityService {
	return activityService{
		activityRepository: r,
	}
}

// Create implements ActivityService
func (a activityService) Create(activity entities.Activity) (entities.Activity, error) {
	// panic("unimplemented")
	return a.activityRepository.Create(activity)
}

// GetAllActivity implements ActivityService
func (a activityService) GetAllActivities() ([]entities.Activity, error) {
	// panic("unimplemented")
	return a.activityRepository.GetAllActivities()
}

// DeleteActivity implements ActivityService
func (a activityService) DeleteActivity(activity entities.Activity) (entities.Activity, error) {
	// panic("unimplemented")
	return a.activityRepository.DeleteActivity(activity)
}

// GetOneActivity implements ActivityService
func (a activityService) GetOneActivity(id int) (entities.Activity, error) {
	// panic("unimplemented")
	return a.activityRepository.GetOneActivity(id)
}

// UpdateActivity implements ActivityService
func (a activityService) UpdateActivity(id int, activity entities.Activity) (entities.Activity, error) {
	// panic("unimplemented")
	return a.activityRepository.UpdateActivity(id, activity)
}
