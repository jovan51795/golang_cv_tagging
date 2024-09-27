package services

import (
	"net/http"

	"77gsi_mynt.com/cv_tagging/models"
	"github.com/gin-gonic/gin"
)

func SaveKeyword(c *gin.Context) {
	var keyword models.Keyword

	err := c.ShouldBindJSON(&keyword)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "message": "Unable to parse data"})
		return
	}

	err = keyword.Save()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "message": "Unable to save data"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"code": 201, "message": "Successfully saved"})
}

func GetAllKeywords(c *gin.Context) {
	data, err := keywords()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 400, "message": "AN error occured"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 200, "data": data})

}

func keywords() ([]models.Keyword, error) {
	data, err := models.GetAllKeywords()
	if err != nil {
		return nil, err
	}

	return data, nil
}
