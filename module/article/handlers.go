package article

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gomarkdown/markdown"
	"gorm.io/gorm"
)

func articleGet(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		slug := c.Param("slug")
		article := FindArticleBySlug(db, slug)
		if article == nil {
			c.String(http.StatusNotFound, "Article Not Found\n")
			return
		}

		body := string(markdown.ToHTML([]byte(article.Body), nil, nil))

		c.HTML(http.StatusOK, "article.html", gin.H{
			"title":       article.Title,
			"description": article.ShortDescription,
			"body":        template.HTML(body),
		})
	}
}
