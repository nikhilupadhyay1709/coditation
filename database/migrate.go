package database

import (
	models "coditation/internal/app/models"
)

func Migrate() {
	// Using auto migration feature from GORM framework.
	DBConn.Debug().AutoMigrate(
		&models.Product{}, 
		&models.Category{}, 
		&models.Map_category_product{}, 
		&models.SubCategory{},
		models.SubCategory2{},
	)
}
