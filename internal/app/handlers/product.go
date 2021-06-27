package handlers

import (
	//"github.com/configs/database"
	//"github.com/coditation/internal/app/types"
	"coditation/database"
	models "coditation/internal/app/models"

	"github.com/gin-gonic/gin"

	"log"
)

func CreateProduct(c *gin.Context) {

	//var product models.Product
	var productWithCategory models.ProductWithCategory
	var responseCode int
	responseMessage := ""
	if c.ShouldBind(&productWithCategory) == nil {
		log.Printf("=========1=======---------------------------------------")
		product := &models.Product{
			Name: productWithCategory.Name,
		}
		if errors := database.DBConn.Debug().Create(product).GetErrors(); len(errors) > 0 {
			responseCode = 422
			for _, err := range errors {
				responseMessage = responseMessage + ", " + err.Error()
			}
		} else {
			responseCode = 201
			responseMessage = "Created"
		}
		//map category and product
		productwithcat := &models.Map_category_product{
			Cid: uint(productWithCategory.Cid),
			Pid: product.ID,
		}
		if errors := database.DBConn.Debug().Create(productwithcat).GetErrors(); len(errors) > 0 {
			responseCode = 422
			for _, err := range errors {
				responseMessage = responseMessage + ", " + err.Error()
			}
		} else {
			responseCode = 201
			responseMessage = "Map with cetagory "
		}
		log.Printf("=========1======= %+v", models.Product{})
		log.Printf("======2========== %+v", product.ID)
	}
	c.String(responseCode, responseMessage)

}
func AllProduct(c *gin.Context) {
	db := database.DBConn

	rows, _ := db.Table("products").Select("products.id as pid,products.name,c.id as cid,c.name as cname").Joins("inner join map_category_products mcp on products.id = mcp.pid").Joins("inner join  categories c on c.id  = mcp.cid").Rows()
	var productsList []models.ProductWithCategory
	for rows.Next() {
		var product models.ProductWithCategory
		db.ScanRows(rows, &product)
		productsList = append(productsList, product)
	}

	c.JSON(200, productsList)
}
func ProductByCategory(c *gin.Context) {
	db := database.DBConn
	cid := c.Param("cid")

	rows, _ := db.Table("products").
		Select("products.id as pid,products.name,c.id as cid,c.name as cname").
		Joins("inner join map_category_products mcp on products.id = mcp.pid").
		Joins("inner join  categories c on c.id  = mcp.cid").
		Where("cid = ?", cid).
		Rows()
	var productsList []models.ProductWithCategory
	for rows.Next() {
		var product models.ProductWithCategory
		db.ScanRows(rows, &product)
		productsList = append(productsList, product)
	}

	c.JSON(200, productsList)
}

func UpdateProduct(c *gin.Context) {
	db := database.DBConn
	product := models.Product{}
	id := c.Param("id")

	if err := db.First(&product, id).Error; err != nil {
		c.String(404, "Product Not Found")
		return
	}
	
	// For now user can only update product name.
	productParams := models.Product{}
	if c.ShouldBind(&productParams) == nil {
		db.Table("products").Select("*").Where("id = ?", id).Update(productParams)
	}

	c.String(200, "Product Updated")
}
