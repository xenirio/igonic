package article

import (
	"errors"
	"io/ioutil"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"gorm.io/gorm"
)

// ModuleArticle is the Skel article module
type ModuleArticle struct{}

// Article : Table name is `articles`
type Article struct {
	ID               uint   `gorm:"primarykey"`
	Slug             string `gorm:"unique_index" yaml:"slug"`
	AuthorUID        string `gorm:"index" yaml:"author_uid"`
	Title            string `yaml:"title"`
	ShortDescription string `yaml:"short_description"`
	Body             string `yaml:"body"`
	Language         string `yaml:"language"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

// SetupRoutes configure module HTTP routes
func (m *ModuleArticle) SetupRoutes(db *gorm.DB, router *gin.Engine) error {
	router.GET("/article/:slug", articleGet(db))
	return nil
}

// Seed load from article from config/articles.yml to database
func (m *ModuleArticle) Seed(db *gorm.DB) error {
	raw, err := ioutil.ReadFile("config/articles.yml")
	if err != nil {
		return err
	}
	articles := []Article{}
	err = yaml.Unmarshal(raw, &articles)
	if err != nil {
		return err
	}

	tx := db.Create(&articles)
	if tx.Error != nil {
		return err
	}
	return nil
}

// Migrate create and modify database tables according to the module models
func (m *ModuleArticle) Migrate(db *gorm.DB) error {
	return db.AutoMigrate(&Article{})
}

// FindArticleBySlug find and return a strategy by name and uid
func FindArticleBySlug(db *gorm.DB, slug string) *Article {
	article := Article{}
	tx := db.Where("slug = ?", slug).First(&article)

	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return nil
		}
		log.Errorf("FindArticleBySlug failed: %s", tx.Error.Error())
		return nil
	}
	return &article
}
