package main

import (
	"github.com/Api-Micro-Service/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	routes.Setup(r)
	r.Run(":8080")
}