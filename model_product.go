package main

import (
	"github.com/TudorHulban/GoTemplating/pkg/validate"
)

// concentrates product type related actions

type Product struct {
	ID                 uint   `gorm:"primaryKey"`
	SKU                string `gorm:"uniqueIndex"`
	Category           string
	Name               string `validate:"required"`
	Slug               string `gorm:"uniqueIndex"`
	Description        string `validate:"required"`
	SEOMetaTitle       string
	SEOMetaDescription string
	Quantity           uint
	PriceCents         float32
	SalesPriceCents    float32
}

type ProductImage struct {
	ProductID uint
	MetaALT   string // holds alt for image
	Path      string
}

func validateProduct(p *Product) error {
	return validate.GetValidator().Struct(p)
}
