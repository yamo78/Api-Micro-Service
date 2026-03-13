package routes

import (
	"github.com/gin-gonic/gin"
)

func Setup(r *gin.Engine) {
	r.POST("/login", func(c *gin.Context) {c.JSON(200, "ok")})
	r.POST("/register", func(c *gin.Context) {c.JSON(200, "ok")})

	todos := r.Group("/todos")
	{
		todos.GET("/", func( c *gin.Context) {c.JSON(200, []string{"todos1"})})
		todos.POST("/", func( c *gin.Context) { c.JSON(200, "creer")})
	}
}