package handler

import (
	"net/http"

	"github.com/EduardoMark/login-system-go/internal/models"
	"github.com/EduardoMark/login-system-go/internal/utils"
	"github.com/gin-gonic/gin"
)

type BodyRequestJSON struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Register(c *gin.Context) {
	var body BodyRequestJSON

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPass, err := utils.HashPass(body.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	user := models.Create(body.Username, hashedPass)

	if err := models.Save(&user); err != nil {
		if err.Error() == "username already exists" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "username already exists"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}
