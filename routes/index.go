package routes

import (
	"net/http"

	"github.com/openware/igonic/middleware"

	"github.com/gin-gonic/gin"
	"github.com/openware/pkg/jwt"
)

// SetUp : Set up routes to render view HTML
func SetUp(router *gin.Engine, keyStore *jwt.KeyStore) *gin.RouterGroup {
	router.GET("/", func(ctx *gin.Context) {
		//render with master
		ctx.HTML(http.StatusOK, "index", gin.H{
			"title": "Index title!",
			"add": func(a int, b int) int {
				return a + b
			},
		})
	})

	router.GET("/page", func(ctx *gin.Context) {
		//render only file, must full name with extension
		ctx.HTML(http.StatusOK, "page.html", gin.H{"title": "Page file title!!"})
	})

	private := router.Group("/api/v2/app/private")
	private.Use(middleware.AuthJWT(keyStore))

	private.GET("/profile", func(c *gin.Context) {
		uid, _ := c.Get("uid")
		role, _ := c.Get("role")
		level, _ := c.Get("level")
		email, _ := c.Get("email")

		c.JSON(200, gin.H{
			"uid":   uid,
			"role":  role,
			"level": level,
			"email": email,
		})
	})

	return private
}
