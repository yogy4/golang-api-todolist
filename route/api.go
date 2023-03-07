package route

import (
	"golang-api-todolist/controllers"
	"golang-api-todolist/repositories"
	"golang-api-todolist/services"

	"github.com/gin-gonic/gin"
	// "github.com/jinzhu/gorm"
	"gorm.io/gorm"
)

func SetupRouteApi(db *gorm.DB) {
	routingapi := gin.Default()

	activityRepository := repositories.NewActivityRepository(db)
	activityService := services.NewActivityService(activityRepository)
	activityController := controllers.NewActivityController(activityService)

	todoRepository := repositories.NewTodoRepository(db)
	todoServices := services.NewTodoService(todoRepository)
	todoController := controllers.NewTodoController(todoServices)

	activity := routingapi.Group("activity-groups")
	activity.GET("/", activityController.GetAllActivities)
	activity.POST("/", activityController.AddActivity)
	activity.GET("/:id", activityController.GetOneActivity)
	activity.PATCH("/:id", activityController.UpdateActivity)
	activity.DELETE("/:id", activityController.DeleteActivity)

	todo := routingapi.Group("todo-items")
	todo.GET("/", todoController.GetAlltodos)
	todo.POST("/", todoController.AddTodo)
	todo.GET("/:id", todoController.GetOnetodo)
	todo.PATCH("/:id", todoController.UpdateTodo)
	todo.DELETE("/:id", todoController.DeleteTodo)

	routingapi.Run(":3030")
}
