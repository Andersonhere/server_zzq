package models

import (
	"time"

	"gorm.io/gorm"
)

// Category 商品分类
type Category struct {
	ID         uint           `gorm:"primaryKey" json:"id"`
	ShopID     uint           `gorm:"not null;index" json:"-"`
	Name       string         `gorm:"type:varchar(50);not null" json:"name"`
	ParentID   uint           `gorm:"default:0;index" json:"parentId"`
	Sort       int            `gorm:"default:0" json:"sort"`
	Status     int            `gorm:"type:tinyint;default:1" json:"status"` // 0 禁用 1 启用
	ProductCount int          `gorm:"-" json:"productCount"` // 不计入数据库，查询时计算
	Children   []*Category    `gorm:"-" json:"children,omitempty"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`
	CreatedAt  time.Time      `json:"createdAt"`
	UpdatedAt  time.Time      `json:"updatedAt"`
}

// TableName 指定表名
func (Category) TableName() string {
	return "categories"
}

// CategoryStatus 分类状态
type CategoryStatus int

const (
	CategoryStatusDisabled CategoryStatus = 0
	CategoryStatusEnabled  CategoryStatus = 1
)
