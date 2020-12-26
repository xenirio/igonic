package main

import (
	"github.com/openware/gin-skel/config"
	"github.com/openware/gin-skel/pkg/utils"
	"github.com/openware/gin-skel/routes"

	"github.com/foolin/goview/supports/ginview"
	"github.com/gin-gonic/gin"
)

var app = gin.Default()

func configApp() {

	// Serve static files
	app.Static("/public", "./public")

	// Set up view engine
	app.HTMLRender = ginview.Default()

	// View routes
	routes.SetUp(app)
}

func main() {
	// config.LoadEnv()
	var cfg config.Config

	config.Parse(&cfg)
	configApp()

	db := config.ConnectDatabase()
	config.RunMigrations(db)

	dbSeed := utils.GetEnv("DATABASE_SEED", "true")
	if dbSeed == "true" {
		config.LoadSeeds(db)
	}

	app.Run(":" + cfg.Port)
}
