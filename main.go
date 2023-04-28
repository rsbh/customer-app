package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	api "github.com/rsbh/customer-app/api"
	"github.com/rsbh/customer-app/config"
	"github.com/rsbh/customer-app/db"
)

func getServer(conf *config.Config) *http.Server {
	router := gin.Default()
	api.BindRoutes(router, "/api")

	addr := fmt.Sprintf("%s:%s", conf.HOST, conf.PORT)
	return &http.Server{
		Addr:    addr,
		Handler: router,
	}
}

func main() {
	conf := config.Load()
	db.Connect(conf)
	server := getServer(conf)
	server.ListenAndServe()
}
