package models

import (
	"time"

	"gorm.io/gorm"
)

// Shop 商家/店铺
type Shop struct {
	ID            uint           `gorm:"primaryKey" json:"id"`
	ShopName      string         `gorm:"type:varchar(100);not null" json:"shopName"`
	Logo          string         `gorm:"type:varchar(500)" json:"logo"`
	ContactPhone  string         `gorm:"type:varchar(20)" json:"contactPhone"`
	WechatOpenID  string         `gorm:"type:varchar(100);uniqueIndex" json:"-"`
	Status        int            `gorm:"type:tinyint;default:1" json:"status"` // 1 正常 0 禁用
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`
	CreatedAt     time.Time      `json:"createdAt"`
	UpdatedAt     time.Time      `json:"updatedAt"`
}

// TableName 指定表名
func (Shop) TableName() string {
	return "shops"
}

// ShopStatus 商家状态
type ShopStatus int

const (
	ShopStatusDisabled ShopStatus = 0
	ShopStatusNormal   ShopStatus = 1
)
