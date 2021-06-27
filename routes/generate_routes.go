package routes

import (
	"coditation/internal/app/handlers"
	"github.com/gin-gonic/gin"
)

func GenerateRoutes() *gin.Engine {
	web := gin.Default()
	web.GET("/ping", handlers.HandlePing) //ping

	web.POST("v1/products", handlers.CreateProduct) // add product
	web.GET("v1/products", handlers.AllProduct) //get all product
	web.GET("v1/products/:cid", handlers.ProductByCategory) // products by category
	web.PUT("v1/products/:id", handlers.UpdateProduct) //update

	web.POST("v1/category", handlers.CreateCategory) //add category
	web.GET("v1/category", handlers.GetCategory) // category with subcategory

	web.POST("v1/subcategory", handlers.SubCategory) //add 
	web.POST("v1/childsubcategory", handlers.ChildSubCategory)   //add
	return web
}
