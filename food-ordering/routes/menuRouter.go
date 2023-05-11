package routes

import (
	controller "food-ordering/controllers"

	"github.com/gin-gonic/gin"
)

func MenuRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/menus", controller.GetMenus())
	incomingRoutes.GET("/menus/:menu_id", controller.GetMenuByID())
	incomingRoutes.POST("/menus", controller.CreateMenu())
	incomingRoutes.PATCH("/menus/:menu_id", controller.UpdateMenu())
}
