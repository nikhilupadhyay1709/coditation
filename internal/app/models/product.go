package types

import (
	"github.com/jinzhu/gorm"
)

type Product struct {
	gorm.Model
	// ID   int   `gorm:"primaryKey"`
	Name string `gorm:"not_null;"`
}

type ProductWithCategory struct {
	Cid   int
	Name  string
	Pid   int
	Cname string
}
type ProductInView struct {
	Sku             string
	Name            string
	CurrentQuantity int
}
