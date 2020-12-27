package page

import (
	"errors"
	"io/ioutil"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"gorm.io/gorm"
)

// ModulePage is the Skel Page module
type ModulePage struct{}

// Page : Table name is `Pages`
type Page struct {
	ID               uint   `gorm:"primarykey"`
	Path             string `gorm:"unique_index" yaml:"path"`
	AuthorUID        string `gorm:"index" yaml:"author_uid"`
	Title            string `yaml:"title"`
	ShortDescription string `yaml:"short_description"`
	Body             string `yaml:"body"`
	Language         string `yaml:"language"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

// SetupRoutes configure module HTTP routes
func (m *ModulePage) SetupRoutes(db *gorm.DB, router *gin.Engine) error {
	for _, p := range ListPages(db) {
		router.GET(p.Path, pageGet(&p))
	}
	return nil
}

// Seed load from Page from config/Pages.yml to database
func (m *ModulePage) Seed(db *gorm.DB) error {
	raw, err := ioutil.ReadFile("config/pages.yml")
	if err != nil {
		return err
	}
	Pages := []Page{}
	err = yaml.Unmarshal(raw, &Pages)
	if err != nil {
		return err
	}

	tx := db.Create(&Pages)
	if tx.Error != nil {
		return err
	}
	return nil
}

// Migrate create and modify database tables according to the module models
func (m *ModulePage) Migrate(db *gorm.DB) error {
	return db.AutoMigrate(&Page{})
}

// FindPageByPath find and return a page by path
func FindPageByPath(db *gorm.DB, path string) *Page {
	Page := Page{}
	tx := db.Where("path = ?", path).First(&Page)

	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return nil
		}
		log.Errorf("FindPageByPath failed: %s", tx.Error.Error())
		return nil
	}
	return &Page
}

// ListPages returns all pages
func ListPages(db *gorm.DB) []Page {
	pages := []Page{}
	tx := db.Find(&pages)

	if tx.Error != nil {
		log.Errorf("FindPageByPath failed: %s", tx.Error.Error())
	}
	return pages
}
