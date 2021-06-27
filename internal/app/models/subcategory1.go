package types

import (
	"github.com/jinzhu/gorm"
)

type SubCategory struct {
	gorm.Model
	// ID   int   `gorm:"primaryKey"`
	Name string `gorm:"not_null;"`
	Cid  int    `gorm:"not_null;"`
}
