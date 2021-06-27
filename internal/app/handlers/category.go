package handlers

import (
	"coditation/database"
	models "coditation/internal/app/models"
	_"fmt"

	//"strings"

	"github.com/gin-gonic/gin"
)

func CreateCategory(c *gin.Context) {
	var category models.Category

	var responseCode int
	responseMessage := ""
	if c.ShouldBind(&category) == nil {
		if errors := database.DBConn.Create(&models.Category{
			Name: category.Name,
			// ID: category.ID,
		}).GetErrors(); len(errors) > 0 {
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
func GetCategory(c *gin.Context) {
	db := database.DBConn

	rows, _ := db.Table("categories c").
		Select("c.id as category_id, c.name as category_name, sc.id as sub_categories_id, sc.name as sub_categories_name").
		Joins("inner join sub_categories sc on c.id  = sc.cid").
		Rows()
	categorylist := make(map[int]interface{})
	for rows.Next() {
		var category models.CategoryWithChildCategory
		db.ScanRows(rows, &category)
		if _, ok := categorylist[category.CategoryId]; ok {
			mdi, _ := categorylist[category.CategoryId]
			md, _ := mdi.(map[string]interface{})
			t :=  md["child_categories"].([]string)
			t = append(t,category.SubCategoriesName)
			temp := make(map[string]interface{})
			temp["id"] = category.CategoryId
			temp["child_categories"] = t
			categorylist[category.CategoryId] = temp
		} else {
			temp := make(map[string]interface{})
			temp["id"] = category.CategoryId
			temp["child_categories"] = []string{category.SubCategoriesName}
			categorylist[category.CategoryId] = temp
		}
		
		// categorylist = append(categorylist, category)
	}

	var finalData []interface{}
	for _,v := range categorylist{
		finalData = append(finalData,v)
	}
	

	c.JSON(200, finalData)
}
