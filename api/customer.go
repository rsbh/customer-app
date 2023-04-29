package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rsbh/customer-app/db/schema"
)

type ListCustomerResponse struct {
	Customers []schema.Customer `json:"customers"`
}

func (h *ApiHandler) listCustomersHandler(c *gin.Context) {
	customers, err := h.customerRepo.ListCustomers(c)
	if err != nil {
		log.Printf(`error in querying db: %s`, err)
		c.JSON(http.StatusInternalServerError, "Something went wrong")
		return
	}
	resp := ListCustomerResponse{Customers: customers}
	c.JSON(http.StatusOK, resp)
}

func (h *ApiHandler) getCustomerHandler(c *gin.Context) {
	c.JSON(http.StatusOK, "get customers")
}

func (h *ApiHandler) createCustomerHandler(c *gin.Context) {
	c.JSON(http.StatusOK, "create customers")
}

func (h *ApiHandler) updateCustomerHandler(c *gin.Context) {
	c.JSON(http.StatusOK, "update customers")
}
