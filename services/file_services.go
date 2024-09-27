package services

import (
	"mime/multipart"
	"net/http"
	"strings"

	"77gsi_mynt.com/cv_tagging/models"
	"code.sajari.com/docconv/v2"
	"github.com/gin-gonic/gin"
)

func Scan(c *gin.Context) {
	file, err := c.FormFile("file")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "message": "File is required"})
	}

	err = c.SaveUploadedFile(file, "./files/"+file.Filename)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "message": "Failed to upload the file"})
		return
	}

	skills, err := readFile(file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "message": "Could not read file"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 200, "data": skills})
}

func readFile(file *multipart.FileHeader) ([]models.Keyword, error) {
	fileData, err := docconv.ConvertPath("./files/" + file.Filename)
	if err != nil {
		return nil, err
	}

	data, err := keywords()

	if err != nil {
		return nil, err
	}

	var skills []models.Keyword

	text := strings.ToLower(fileData.Body)

	for _, keyword := range data {
		var skill models.Keyword
		if strings.Contains(text, strings.ToLower(keyword.Keyword)) {
			skill.Id = keyword.Id
			skill.Keyword = keyword.Keyword
			skill.User_id = keyword.User_id

			skills = append(skills, skill)
		}

	}
	return skills, nil
}
