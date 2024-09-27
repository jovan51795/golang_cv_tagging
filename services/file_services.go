package services

import (
	"mime/multipart"
	"net/http"
	"strings"

	"77gsi_mynt.com/cv_tagging/models"
	"77gsi_mynt.com/cv_tagging/util"
	"code.sajari.com/docconv/v2"
	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
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
	var skills []models.Keyword

	if strings.HasSuffix(file.Filename, ".xlsx") {
		skills, err = readExcelFile(file)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "message": "Could not read file"})
			return
		}
	} else {
		skills, err = readFile(file)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "message": "Could not read file"})
			return
		}
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
		if strings.Contains(text, strings.ToLower(keyword.Keyword)) && !util.Contains(skills, keyword) {
			skill.Id = keyword.Id
			skill.Keyword = keyword.Keyword
			skill.User_id = keyword.User_id

			skills = append(skills, skill)
		}

	}
	return skills, nil
}

func readExcelFile(file *multipart.FileHeader) ([]models.Keyword, error) {
	f, err := excelize.OpenFile("./files/" + file.Filename)

	if err != nil {
		return nil, err
	}
	defer f.Close()

	sheetNames := f.GetSheetList()
	keywords, err := keywords()
	if err != nil {
		return nil, err
	}

	var skills []models.Keyword

	for _, name := range sheetNames {
		var skill models.Keyword
		rows, err := f.GetRows(name)

		if err != nil {
			return nil, err
		}

		for _, row := range rows {
			for _, cell := range row {
				for _, key := range keywords {
					cellVal := strings.ToLower(cell)
					if strings.Contains(cellVal, strings.ToLower(key.Keyword)) && !util.Contains(skills, key) {
						skill.Id = key.Id
						skill.Keyword = key.Keyword
						skill.User_id = key.User_id

						skills = append(skills, skill)
					}
				}
			}
		}
	}

	return skills, nil

}
