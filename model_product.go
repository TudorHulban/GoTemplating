package main

import (
	"github.com/TudorHulban/GoTemplating/pkg/validator"
)

// concentrates product type related actions

type Product struct {
	ID                 uint   `gorm:"primaryKey"`
	SKU                string `gorm:"uniqueIndex"`
	Category           string
	Slug               string `gorm:"uniqueIndex"`
	Name               string `validate:"required"`
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
	return validator.GetValidator().Struct(p)
}
