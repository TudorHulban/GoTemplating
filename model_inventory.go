package main

import (
	"os"
	"sync"

	"github.com/TudorHulban/GoTemplating/pkg/validate"
	"github.com/TudorHulban/log"
	"gorm.io/gorm"
)

// concentrates product type related actions

type Product struct {
	ID                 uint `gorm:"primaryKey"`
	Visible            bool
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

type Inventory struct {
	m        sync.Mutex
	Store    string
	Products []Product // cached list of products, mirrors persistance
	DBConn   *gorm.DB
	l        *log.LogInfo
}

func NewInventory(store string, db *gorm.DB) *Inventory {
	return &Inventory{
		Store:    store,
		Products: []Product{},
		DBConn:   db,
		l:        log.New(log.DEBUG, os.Stderr, true),
	}
}

// AddProduct Method persist and on succesful caches product. Does not return ID of insert.
func (t *Inventory) AddProduct(p *Product) error {
	errPersist := t.DBConn.Create(p).Error
	if errPersist != nil {
		return errPersist
	}

	// cache new product
	t.m.Lock()
	defer t.m.Unlock()

	t.Products = append(t.Products, *p)
	return nil
}

type ProductImage struct {
	ProductID uint
	MetaALT   string // holds alt for image
	Path      string
}

type ProductVideo struct {
	ProductID uint
	MetaALT   string // holds alt for image
	URL       string
}

func validateProduct(p *Product) error {
	return validate.GetValidator().Struct(p)
}
