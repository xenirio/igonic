package config

import (
	"fmt"

	"github.com/openware/gin-skel/models"
	"github.com/openware/gin-skel/pkg/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

// ConnectDatabase : connect to database MySQL using gorm
// gorm (GO ORM for SQL): http://gorm.io/docs/connecting_to_the_database.html
func ConnectDatabase() {

	dbHost := utils.DefaultGetEnv("DATABASE_HOST", "localhost")
	dbPort := utils.DefaultGetEnv("DATABASE_PORT", "3306")
	dbName := utils.DefaultGetEnv("DATABASE_NAME", "opendax_development")
	dbUser := utils.DefaultGetEnv("DATABASE_USER", "root")
	dbPass := utils.DefaultGetEnv("DATABASE_PASS", "")

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		dbUser, dbPass, dbPort, dbHost, dbName,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Connect database PostgreSQL successfully")
	}

	// Pass db connection to package controllers and models
	models.SetUpDBConnection(db)
	// controllers.SetUpDBConnection(db)

	// Store this db connection for package config
	setUpDBConnection(db)
}

func setUpDBConnection(DB *gorm.DB) {
	db = DB
}

// GetDBConnection : get db connection from package config
func GetDBConnection() *gorm.DB {
	return db
}
