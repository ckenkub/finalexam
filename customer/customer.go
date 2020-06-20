package customer

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(AuthMiddleware)
	r.GET("/customers", getCustomersHandler)
	r.GET("/customers/:id", getCustomerHandler)
	r.POST("/customers", createCustomerHandler)
	r.PUT("/customers/:id", updateCustomerHandler)
	r.DELETE("/customers/:id", deleteCustomerHandler)

	return r
}
