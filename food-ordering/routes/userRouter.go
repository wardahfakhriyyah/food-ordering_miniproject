package routes

import (
	controller "food-ordering/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/users", controller.GetUsers())
	incomingRoutes.GET("/users/:user_id", controller.GetUserByID())
	incomingRoutes.POST("/users/register", controller.Register())
	incomingRoutes.POST("/users/login", controller.Login())
	incomingRoutes.DELETE("/users/:user_id", controller.DeleteUser())
}
