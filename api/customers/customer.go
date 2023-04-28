package customers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func BindRoutes(rg *gin.RouterGroup, baseURL string) {
	router := rg.Group(baseURL)
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "customers")
	})
}
