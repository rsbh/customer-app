package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rsbh/customer-app/db/repositories"
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

type createCustomerRequestBody struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

type createCustomerResponse struct {
	Customer schema.Customer `json:"customer"`
}

func (h *ApiHandler) createCustomerHandler(c *gin.Context) {
	var requestBody createCustomerRequestBody
	err := c.BindJSON(&requestBody)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Bad Request")
		return
	}

	customer, err := h.customerRepo.CreateCustomers(c, repositories.CreateCustomersBody{
		FirstName: requestBody.FirstName,
		LastName:  requestBody.LastName,
		Email:     requestBody.Email,
	})

	if err != nil {
		log.Printf(`error in db insert: %s`, err)
		c.JSON(http.StatusInternalServerError, "Something went wrong")
		return
	}

	c.JSON(http.StatusOK, createCustomerResponse{
		Customer: customer,
	})
}

func (h *ApiHandler) updateCustomerHandler(c *gin.Context) {
	c.JSON(http.StatusOK, "update customers")
}
