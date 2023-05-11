package routes

import (
	controller "food-ordering/controllers"

	"github.com/gin-gonic/gin"
)

func InvoiceRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/invoices", controller.GetInvoices())
	incomingRoutes.GET("/invoices/:invoice_id", controller.GetInvoiceByID())
	incomingRoutes.POST("/invoices", controller.CreateInvoice())
	incomingRoutes.PATCH("/invoices/:invoice_id", controller.UpdateInvoice())
}
