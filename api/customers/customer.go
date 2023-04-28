package customers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func listCustomersHandler(c *gin.Context) {
	c.JSON(http.StatusOK, "list customers")
}

func getCustomerHandler(c *gin.Context) {
	c.JSON(http.StatusOK, "get customers")
}

func createCustomerHandler(c *gin.Context) {
	c.JSON(http.StatusOK, "create customers")
}

func updateCustomerHandler(c *gin.Context) {
	c.JSON(http.StatusOK, "update customers")
}

func BindRoutes(rg *gin.RouterGroup, baseURL string) {
	router := rg.Group(baseURL)
	router.GET("/", listCustomersHandler)
	router.POST("/", createCustomerHandler)
	router.GET("/:id", getCustomerHandler)
	router.PUT("/:id", updateCustomerHandler)
}
