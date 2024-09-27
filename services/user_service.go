package services

import (
	"net/http"

	"77gsi_mynt.com/cv_tagging/models"
	"77gsi_mynt.com/cv_tagging/util"
	"github.com/gin-gonic/gin"
)

func Signup(c *gin.Context) {
	var user models.User

	err := c.ShouldBindJSON(&user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "message": "Could not parse data"})
		return
	}

	err = user.Signup()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "message": "Could not savd data"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"code": 201, "message": "Signup successful"})
}

func Login(c *gin.Context) {
	var user models.User

	err := c.ShouldBindJSON(&user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "message": "Could not parse data"})
		return
	}

	err = user.Login()

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "message": err})
		return
	}

	token, err := util.GenerateToken(user.Email, user.Id)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "message": "Generating token failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 200, "message": "successfully logged in", "token": token})
}
