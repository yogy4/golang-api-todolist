package controllers

import (
	"golang-api-todolist/entities"
	"golang-api-todolist/services"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ActivityController interface {
	AddActivity(*gin.Context)
	GetAllActivities(*gin.Context)
	GetOneActivity(*gin.Context)
	UpdateActivity(*gin.Context)
	DeleteActivity(*gin.Context)
}

type activityController struct {
	activityService services.ActivityService
}

func NewActivityController(s services.ActivityService) ActivityController {
	return activityController{
		activityService: s,
	}
}

// AddActivity implements ActivityController
func (a activityController) AddActivity(c *gin.Context) {
	// panic("unimplemented")
	log.Print("[ActivityController]...Add Activity")
	var activity entities.Activity
	if err := c.ShouldBindJSON(&activity); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	activity, err := a.activityService.Create(activity)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error while saving activity"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "Success", "message": "Success", "data": activity})
}

// GetAllActivities implements ActivityController
func (a activityController) GetAllActivities(c *gin.Context) {
	// panic("unimplemented")
	log.Print("[ActivityController]...Get all activities")
	activities, err := a.activityService.GetAllActivities()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error while getting activities"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "Success", "message": "Success", "data": activities})
}

// DeleteActivity implements ActivityController
func (a activityController) DeleteActivity(c *gin.Context) {
	// panic("unimplemented")
	var activity entities.Activity
	ids := c.Param("id")
	id, _ := strconv.Atoi(ids)
	activity.ActivityID = int(id)

	status := " Not Found"
	activity, err := a.activityService.DeleteActivity(activity)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNotFound, gin.H{"status": status, "message": "Activity with ID " + ids + status})

}

// GetOneActivity implements ActivityController
func (a activityController) GetOneActivity(c *gin.Context) {
	// panic("unimplemented")
	log.Print("[ActivityController]...Get one activity")
	ids := c.Params.ByName("id")
	id, _ := strconv.Atoi(ids)
	activity, err := a.activityService.GetOneActivity(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error while getting activity"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "Success", "message": "Success", "data": activity})
}

// UpdateActivity implements ActivityController
func (a activityController) UpdateActivity(c *gin.Context) {
	// panic("unimplemented")
	log.Print("[ActivityController]...Update activity")
	var err error
	var activity entities.Activity
	if err := c.ShouldBindJSON(&activity); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ids := c.Param("id")
	id, _ := strconv.Atoi(ids)
	// activity.ActivityID = int(id)
	// activity.Title = c.Param("title")
	activity, err = a.activityService.UpdateActivity(id, activity)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error while update activity"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "Success", "message": "Success", "data": activity})
}
