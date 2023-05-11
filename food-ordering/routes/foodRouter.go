package routes

import (
	controller "food-ordering/controllers"

	"github.com/gin-gonic/gin"
)

func FoodRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/foods", controller.GetFoods())
	incomingRoutes.GET("/foods/:food_id", controller.GetFoodByID())
	incomingRoutes.POST("/foods", controller.CreateFood())
	incomingRoutes.PATCH("/foods/:food_id", controller.UpdateFood())
	incomingRoutes.DELETE("/foods/:food_id", controller.DeleteFood())
}
