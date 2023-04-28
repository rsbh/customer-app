package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	api "github.com/rsbh/customer-app/api"
)

const HOST = "localhost"
const PORT = "8000"

func getServer() *http.Server {
	router := gin.Default()

	api.BindRoutes(router, "/api")

	addr := fmt.Sprintf("%s:%s", HOST, PORT)
	return &http.Server{
		Addr:    addr,
		Handler: router,
	}
}

func main() {
	server := getServer()
	server.ListenAndServe()
}
