package handlers

import (
	"coditation/database"
	models "coditation/internal/app/models"

	"github.com/gin-gonic/gin"
)

func SubCategory(c *gin.Context) {

	//var product models.Product
	var subCategory models.SubCategory
	var responseCode int
	responseMessage := ""
	if c.ShouldBind(&subCategory) == nil {
		SubCategory := &models.SubCategory{
			Name: subCategory.Name,
			Cid:  subCategory.Cid,
		}
		if errors := database.DBConn.Create(SubCategory).GetErrors(); len(errors) > 0 {
			responseCode = 422
			for _, err := range errors {
				responseMessage = responseMessage + ", " + err.Error()
			}
		} else {
			responseCode = 201
			responseMessage = "Created"
		}
	}
	c.String(responseCode, responseMessage)
}
