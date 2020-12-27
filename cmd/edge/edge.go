package main

import (
	"os"

	"github.com/openware/igonic/config"
	"github.com/openware/igonic/routes"

	"github.com/openware/pkg/jwt"

	"github.com/foolin/goview/supports/ginview"
	"github.com/gin-gonic/gin"
)

var app = gin.Default()

func configApp() {

	// Serve static files
	app.Static("/public", "./public")

	// Set up view engine
	app.HTMLRender = ginview.Default()

	jwtPublicKey := os.Getenv("JWT_PUBLIC_KEY")
	if jwtPublicKey == "" {
		panic("missing JWT_PUBLIC_KEY")
	}

	keyStore := &jwt.KeyStore{}
	err := keyStore.LoadPublicKeyFromString(jwtPublicKey)
	if err != nil {
		panic("Failed to load JWT public key: " + err.Error())
	}

	// View routes
	routes.SetUp(app, keyStore)
}

func main() {
	// config.LoadEnv()
	var cfg config.Config

	config.Parse(&cfg)
	configApp()

	db := config.ConnectDatabase()
	if !cfg.SkipMigrate {
		config.RunMigrations(db)
	}
	routes.SetPageRoutes(db, app)
	app.Run(":" + cfg.Port)
}
