package domain

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code     uint
	Name     string
	Price    decimal.Decimal `gorm:"type:decimal(20,8)"`
	Cover    string
	State    string
	Discount decimal.Decimal `gorm:"type:decimal(2,2)"`
}

func NewProduct() *Product {
	return &Product{}
}