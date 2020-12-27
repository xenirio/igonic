package main

import (
	"os"

	"github.com/openware/gin-skel/config"
	"github.com/openware/gin-skel/module"
	"github.com/openware/gin-skel/pkg/utils"
	"github.com/openware/gin-skel/routes"
	"github.com/openware/pkg/jwt"
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
	private := routes.SetUp(app, keyStore)
	mod.SetupRoutes(db, app, private)
}

func main() {
	// config.LoadEnv()
	var cfg config.Config

	config.Parse(&cfg)

	db := config.ConnectDatabase()
	configApp(db)

	app.Run(":" + cfg.Port)
}
