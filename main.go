package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	api "github.com/rsbh/customer-app/api"
	"github.com/rsbh/customer-app/config"
	"github.com/rsbh/customer-app/db"
)

func getServer(conf *config.Config, db *sqlx.DB) *http.Server {
	router := gin.Default()
	apiHandler := api.NewApiHandler(db)
	apiHandler.BindRoutes(router, "/api")

	addr := fmt.Sprintf("%s:%s", conf.HOST, conf.PORT)
	return &http.Server{
		Addr:    addr,
		Handler: router,
	}
}

func main() {
	conf, err := config.Load()
	if err != nil {
		log.Fatal("Unable to load config", err)
	}

	DB, err := db.Connect(conf)
	if err != nil {
		log.Fatal("db connection failed", err)
	}

	if err := db.MigrationsUp(conf); err != nil {
		log.Fatal("Unable to run Migrations")
	}
	server := getServer(conf, DB)
	server.ListenAndServe()
}
