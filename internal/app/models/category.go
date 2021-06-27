package types

import (
	"github.com/jinzhu/gorm"
)

type Category struct {
	gorm.Model
	// ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"not_null;index:name"`
}
type CategoryWithChildCategory struct {
	CategoryId        int
	CategoryName      string
	SubCategoriesId   int
	SubCategoriesName string
}
