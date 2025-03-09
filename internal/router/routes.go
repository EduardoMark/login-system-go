package router

import "github.com/gin-gonic/gin"

func SetupRoutes() {
	r := gin.Default()
	UserRoutes(r)
	r.Run(":8080")
}
