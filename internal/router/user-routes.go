package router

import "github.com/gin-gonic/gin"

func UserRoutes(r *gin.Engine) {
	userRoutes := r.Group("/users")

	userRoutes.POST("/register")
	userRoutes.POST("/login")
}
