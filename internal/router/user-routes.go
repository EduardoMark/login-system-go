package router

import (
	"github.com/EduardoMark/login-system-go/internal/handler"
	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine) {
	userRoutes := r.Group("/users")

	userRoutes.POST("/register", handler.Register)
	userRoutes.POST("/login", handler.Login)
}
