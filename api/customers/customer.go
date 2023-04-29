package customers

import (
	"log"
	"net/http"

	"github.com/doug-martin/goqu/v9"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/rsbh/customer-app/db/schema"
)

type CustomerHandler struct {
	DB *sqlx.DB
}

func NewCustomerHandler(DB *sqlx.DB) *CustomerHandler {
	return &CustomerHandler{
		DB: DB,
	}
}

func (h *CustomerHandler) listCustomersHandler(c *gin.Context) {
	customers := []schema.Customer{}
	query, _, err := goqu.Select("*").From("customers").ToSQL()
	if err != nil {
		log.Printf(`error in query generation: %s`, err)
		c.JSON(http.StatusInternalServerError, "Something went wrong")
		return
	}
	if err = h.DB.SelectContext(c, &customers, query); err != nil {
		log.Printf(`error in querying db: %s`, err)
		c.JSON(http.StatusInternalServerError, "Something went wrong")
		return
	}
	c.JSON(http.StatusOK, customers)
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
