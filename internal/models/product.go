package models

import (
	"time"

	"gorm.io/gorm"
)

// Product 商品
type Product struct {
	ID              uint           `gorm:"primaryKey" json:"id"`
	ShopID          uint           `gorm:"not null;index" json:"-"`
	CategoryID      uint           `gorm:"not null;index" json:"categoryId"`
	Name            string         `gorm:"type:varchar(100);not null" json:"name"`
	MainImage       string         `gorm:"type:varchar(500)" json:"mainImage"`
	Images          string         `gorm:"type:text" json:"images"` // JSON 数组存储
	Price           int            `gorm:"not null" json:"price"`
	OriginalPrice   int            `gorm:"not null" json:"originalPrice"`
	Stock           int            `gorm:"default:0" json:"stock"`
	LockedStock     int            `gorm:"default:0" json:"lockedStock"`
	Status          int            `gorm:"type:tinyint;default:0" json:"status"` // 0 下架 1 上架
	Description     string         `gorm:"type:text" json:"description"`
	SalesCount      int            `gorm:"default:0" json:"salesCount"`
	IsDiscount      bool           `gorm:"default:false" json:"isDiscount"`
	DiscountPrice   int            `json:"discountPrice"`
	DiscountStartAt *time.Time     `json:"discountStartAt"`
	DiscountEndAt   *time.Time     `json:"discountEndAt"`
	IsPresale       bool           `gorm:"default:false" json:"isPresale"`
	PresalePrice    int            `json:"presalePrice"`
	PresaleStartAt  *time.Time     `json:"presaleStartAt"`
	PresaleEndAt    *time.Time     `json:"presaleEndAt"`
	EnableCarousel  bool           `gorm:"default:false" json:"enableCarousel"`
	CarouselImage   string         `gorm:"type:varchar(500)" json:"carouselImage"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"-"`
	CreatedAt       time.Time      `json:"createdAt"`
	UpdatedAt       time.Time      `json:"updatedAt"`
}

// TableName 指定表名
func (Product) TableName() string {
	return "products"
}

// ProductStatus 商品状态
type ProductStatus int

const (
	ProductStatusOffShelf ProductStatus = 0
	ProductStatusOnShelf  ProductStatus = 1
)
