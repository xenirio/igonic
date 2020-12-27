package main

import (
	"github.com/openware/gin-skel/config"
	"github.com/openware/gin-skel/module"
	"github.com/openware/gin-skel/pkg/utils"
	"github.com/openware/gin-skel/routes"
	"gorm.io/gorm"

	"github.com/foolin/goview/supports/ginview"
	"github.com/gin-gonic/gin"
)

var app = gin.Default()

func configApp(db *gorm.DB) {
	mod := &module.Registry{}
	mod.RegisterAll(config.ActiveModules())
	mod.Migrate(db)

	if utils.GetEnv("DATABASE_SEED", "true") == "true" {
		mod.Seed(db)
	}

	// Serve static files
	app.Static("/public", "./public")

	// Set up view engine
	app.HTMLRender = ginview.Default()

	// View routes
	routes.SetUp(app)
	mod.SetupRoutes(db, app)
}

func main() {
	// config.LoadEnv()
	var cfg config.Config

	config.Parse(&cfg)

	db := config.ConnectDatabase()
	configApp(db)

	app.Run(":" + cfg.Port)
}
