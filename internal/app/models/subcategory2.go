package types

import (
	"github.com/jinzhu/gorm"
)

type SubCategory2 struct {
	gorm.Model
	// ID   int   `gorm:"primaryKey"`
	Name string `gorm:"not_null;"`
	Sid  int    `gorm:"not_null;"`
}
