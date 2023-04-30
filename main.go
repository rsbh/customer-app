package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	api "github.com/rsbh/customer-app/api"
	"github.com/rsbh/customer-app/config"
	"github.com/rsbh/customer-app/db"
	"github.com/rsbh/customer-app/db/repositories"
)

func getServer(conf *config.Config, db *sqlx.DB) *http.Server {
	router := gin.Default()
	customerRepo := repositories.NewCustomerRepo(db)
	apiHandler := api.NewApiHandler(customerRepo)
	apiHandler.BindRoutes(router.Group("/api"))

	addr := fmt.Sprintf("%s:%s", conf.HOST, conf.PORT)
	return &http.Server{
		Addr:    addr,
		Handler: router,
	}
}

func main() {
	if len(os.Args) != 2 {
		log.Fatal("no argument, usage: customer-app <serve | migarate>")
		return
	}

	runCmd := os.Args[1]

	conf, err := config.Load()
	if err != nil {
		log.Fatal("Unable to load config", err)
	}

	switch runCmd {
	case "serve":
		runServe(conf)
	case "migrate":
		runMigration(conf)
	default:
		fmt.Println("invalid argument, usage: customer-app <serve | migarate>")
	}
}

func runServe(conf *config.Config) {
	DB, err := db.Connect(conf)
	if err != nil {
		log.Fatalf(`db connection failed: %s`, err)
	}
	server := getServer(conf, DB)
	server.ListenAndServe()
}

func runMigration(conf *config.Config) {
	if err := db.MigrationsUp(conf); err != nil {
		log.Fatal("Unable to run Migrations")
	}
	log.Println("Migration Done")
}
