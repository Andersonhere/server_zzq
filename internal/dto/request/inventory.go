package request

// InventoryListRequest 库存列表请求
type InventoryListRequest struct {
	Page      int    `form:"page" binding:"min=1"`
	PageSize  int    `form:"pageSize" binding:"min=1,max=50"`
	Keyword   string `form:"keyword"`
	LowStock  int    `form:"lowStock"`
	Threshold int    `form:"threshold"`
}

// InventoryAdjustRequest 库存调整请求
type InventoryAdjustRequest struct {
	Delta  int    `json:"delta" binding:"required,ne=0"`
	Reason string `json:"reason"`
}

// InventoryBatchAdjustRequest 批量库存调整请求
type InventoryBatchAdjustRequest struct {
	Items  []InventoryAdjustItem `json:"items" binding:"required,min=1"`
	Reason string                `json:"reason"`
}

type InventoryAdjustItem struct {
	ProductID uint `json:"productId" binding:"required"`
	Delta     int  `json:"delta" binding:"required,ne=0"`
}
