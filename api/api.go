package api

import (
	"github.com/gin-gonic/gin"
	"github.com/rsbh/customer-app/api/customers"
)

func BindRoutes(rg *gin.Engine, baseURL string) {
	router := rg.Group(baseURL)
	customers.BindRoutes(router, "/customers")
}
