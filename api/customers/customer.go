package customers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rsbh/customer-app/db/repositories"
	"github.com/rsbh/customer-app/db/schema"
)

type CustomerHandler struct {
	customerRepo *repositories.CustomerRepo
}

func NewCustomerHandler(customerRepo *repositories.CustomerRepo) *CustomerHandler {
	return &CustomerHandler{
		customerRepo: customerRepo,
	}
}

type ListCustomerResponse struct {
	Customers []schema.Customer `json:"customers"`
}

func (h *CustomerHandler) listCustomersHandler(c *gin.Context) {
	customers, err := h.customerRepo.ListCustomers(c)
	if err != nil {
		log.Printf(`error in querying db: %s`, err)
		c.JSON(http.StatusInternalServerError, "Something went wrong")
		return
	}
	resp := ListCustomerResponse{Customers: customers}
	c.JSON(http.StatusOK, resp)
}

func (h *CustomerHandler) getCustomerHandler(c *gin.Context) {
	c.JSON(http.StatusOK, "get customers")
}

func (h *CustomerHandler) createCustomerHandler(c *gin.Context) {
	c.JSON(http.StatusOK, "create customers")
}

func (h *CustomerHandler) updateCustomerHandler(c *gin.Context) {
	c.JSON(http.StatusOK, "update customers")
}

func (h *CustomerHandler) BindRoutes(rg *gin.RouterGroup, baseURL string) {
	router := rg.Group(baseURL)
	router.GET("/", h.listCustomersHandler)
	router.POST("/", h.createCustomerHandler)
	router.GET("/:id", h.getCustomerHandler)
	router.PUT("/:id", h.updateCustomerHandler)
}
