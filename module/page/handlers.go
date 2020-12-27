package page

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gomarkdown/markdown"
)

func pageGet(p *Page) func(c *gin.Context) {
	return func(c *gin.Context) {
		body := string(markdown.ToHTML([]byte(p.Body), nil, nil))

		c.HTML(http.StatusOK, "page.html", gin.H{
			"title":       p.Title,
			"description": p.ShortDescription,
			"body":        template.HTML(body),
		})
	}
}
