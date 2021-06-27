package types

import (
	"github.com/jinzhu/gorm"
)

type Map_category_product struct {
	gorm.Model
	// ID  int `gorm:"primaryKey"`
	Cid uint `gorm:"not_null;"`
	Pid uint `gorm:"not_null;"`
}
