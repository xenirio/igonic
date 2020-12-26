package models

import (
	"io/ioutil"
	"time"

	"gopkg.in/yaml.v2"
	"gorm.io/gorm"
)

// Article : Table name is `ars`
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

// SeedArticles load from article from config/articles.yml to database
func SeedArticles(db *gorm.DB) error {
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
