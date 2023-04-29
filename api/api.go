package api

import (
	"github.com/gin-gonic/gin"
	"github.com/rsbh/customer-app/api/customers"
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

func (h *ApiHandler) BindRoutes(rg *gin.Engine, baseURL string) {
	router := rg.Group(baseURL)
	customerHandler := customers.NewCustomerHandler(h.customerRepo)
	customerHandler.BindRoutes(router, "/customers")
}
