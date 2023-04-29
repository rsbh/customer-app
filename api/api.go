package api

import (
	"github.com/gin-gonic/gin"
	"github.com/rsbh/customer-app/db/repositories"
)

type ApiHandler struct {
	customerRepo *repositories.CustomerRepo
}

func NewApiHandler(customerRepo *repositories.CustomerRepo) *ApiHandler {
	return &ApiHandler{
		customerRepo: customerRepo,
	}
}

func (h *ApiHandler) BindRoutes(rg *gin.RouterGroup) {
	customerRoutes := rg.Group("/customers")
	{
		customerRoutes.GET("/", h.listCustomersHandler)
		customerRoutes.POST("/", h.createCustomerHandler)
		customerRoutes.GET("/:id", h.getCustomerHandler)
		customerRoutes.PUT("/:id", h.updateCustomerHandler)
	}
}
