package models

import (
	"time"

	"gorm.io/gorm"
)

// Order 订单
type Order struct {
	ID              uint           `gorm:"primaryKey" json:"id"`
	OrderNo         string         `gorm:"type:varchar(50);uniqueIndex;not null" json:"orderNo"`
	ShopID          uint           `gorm:"not null;index" json:"-"`
	UserID          uint           `gorm:"not null;index" json:"userId"`
	UserName        string         `gorm:"type:varchar(50)" json:"userName"`
	UserPhone       string         `gorm:"type:varchar(20)" json:"userPhone"`
	TotalAmount     int            `gorm:"not null" json:"totalAmount"` // 订单总金额，单位：分
	PayAmount       int            `gorm:"not null" json:"payAmount"`    // 实付金额，单位：分
	FreightAmount   int            `gorm:"default:0" json:"freightAmount"` // 运费，单位：分
	DiscountAmount  int            `gorm:"default:0" json:"discountAmount"` // 优惠金额，单位：分
	Status          int            `gorm:"type:tinyint;not null;index" json:"status"` // 订单状态
	ReceiverName    string         `gorm:"type:varchar(50)" json:"-"`
	ReceiverPhone   string         `gorm:"type:varchar(20)" json:"-"`
	ReceiverAddress string         `gorm:"type:varchar(500)" json:"-"`
	PayTime         *time.Time     `json:"payTime"`
	ShipTime        *time.Time     `json:"shipTime"`
	ReceiveTime     *time.Time     `json:"receiveTime"`
	ExpressCompany  string         `gorm:"type:varchar(50)" json:"expressCompany"`
	ExpressNo       string         `gorm:"type:varchar(50)" json:"expressNo"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"-"`
	CreatedAt       time.Time      `json:"createdAt"`
	UpdatedAt       time.Time      `json:"updatedAt"`
	Items           []OrderItem    `gorm:"foreignKey:OrderID" json:"items,omitempty"`
	AfterSale       *AfterSale     `gorm:"foreignKey:OrderID" json:"afterSale,omitempty"`
}

// TableName 指定表名
func (Order) TableName() string {
	return "orders"
}

// OrderStatus 订单状态
type OrderStatus int

const (
	OrderStatusPending   OrderStatus = 0 // 待支付
	OrderStatusToShip    OrderStatus = 1 // 待发货
	OrderStatusToReceive OrderStatus = 2 // 待收货
	OrderStatusCompleted OrderStatus = 3 // 已完成
	OrderStatusCancelled OrderStatus = 4 // 已取消
	OrderStatusAfterSale OrderStatus = 5 // 售后中
)

// OrderItem 订单明细
type OrderItem struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	OrderID     uint      `gorm:"not null;index" json:"-"`
	ProductID   uint      `gorm:"not null" json:"productId"`
	ProductName string    `gorm:"type:varchar(100)" json:"productName"`
	MainImage   string    `gorm:"type:varchar(500)" json:"mainImage"`
	Price       int       `gorm:"not null" json:"price"`
	Quantity    int       `gorm:"not null" json:"quantity"`
	Amount      int       `gorm:"not null" json:"amount"`
	SpecText    string    `gorm:"type:varchar(100)" json:"specText"`
	CreatedAt   time.Time `json:"-"`
}

// TableName 指定表名
func (OrderItem) TableName() string {
	return "order_items"
}

// AfterSale 售后
type AfterSale struct {
	ID           uint           `gorm:"primaryKey" json:"id"`
	OrderID      uint           `gorm:"uniqueIndex;not null" json:"orderId"`
	ShopID       uint           `gorm:"not null;index" json:"-"`
	Type         int            `gorm:"type:tinyint;not null" json:"type"` // 1 仅退款 2 退货退款
	Status       int            `gorm:"type:tinyint;not null" json:"status"` // 0 待处理 1 已同意 2 已拒绝 3 已完成
	Reason       string         `gorm:"type:varchar(500)" json:"reason"`
	RejectReason string         `gorm:"type:varchar(500)" json:"rejectReason"`
	RefundAmount int            `gorm:"not null" json:"refundAmount"`
	HandledAt    *time.Time     `json:"handledAt"`
	CreatedAt    time.Time      `json:"createdAt"`
	UpdatedAt    time.Time      `json:"updatedAt"`
}

// TableName 指定表名
func (AfterSale) TableName() string {
	return "after_sales"
}

// AfterSaleType 售后类型
type AfterSaleType int

const (
	AfterSaleTypeRefundOnly      AfterSaleType = 1 // 仅退款
	AfterSaleTypeReturnAndRefund AfterSaleType = 2 // 退货退款
)

// AfterSaleStatus 售后状态
type AfterSaleStatus int

const (
	AfterSaleStatusPending   AfterSaleStatus = 0 // 待处理
	AfterSaleStatusApproved  AfterSaleStatus = 1 // 已同意
	AfterSaleStatusRejected  AfterSaleStatus = 2 // 已拒绝
	AfterSaleStatusCompleted AfterSaleStatus = 3 // 已完成
)
