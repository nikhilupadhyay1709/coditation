package handlers

import (
	//"github.com/configs/database"
	//"github.com/coditation/internal/app/types"
	"coditation/database"
	models "coditation/internal/app/models"

	"github.com/gin-gonic/gin"

	"log"
)

func ChildSubCategory(c *gin.Context) {

	//var product models.Product
	var subCategory2 models.SubCategory2
	var responseCode int
	responseMessage := ""
	if c.ShouldBind(&subCategory2) == nil {
		log.Printf("=========1=======---------------------------------------")
		subCategory2 := &models.SubCategory2{
			Name: subCategory2.Name,
			Sid:  subCategory2.Sid,
		}
		if errors := database.DBConn.Debug().Create(subCategory2).GetErrors(); len(errors) > 0 {
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
