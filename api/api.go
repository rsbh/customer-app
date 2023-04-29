package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/rsbh/customer-app/api/customers"
)

type ApiHandler struct {
	DB *sqlx.DB
}

func NewApiHandler(DB *sqlx.DB) *ApiHandler {
	return &ApiHandler{
		DB: DB,
	}
}

func (h *ApiHandler) BindRoutes(rg *gin.Engine, baseURL string) {
	router := rg.Group(baseURL)
	customerHandler := customers.NewCustomerHandler(h.DB)
	customerHandler.BindRoutes(router, "/customers")
}
