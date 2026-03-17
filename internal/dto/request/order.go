package request

// OrderListRequest 订单列表请求
type OrderListRequest struct {
	Page      int    `form:"page" binding:"min=1"`
	PageSize  int    `form:"pageSize" binding:"min=1,max=50"`
	Status    *int   `form:"status"`
	Keyword   string `form:"keyword"`
	StartDate string `form:"startDate"`
	EndDate   string `form:"endDate"`
}

// OrderShipRequest 订单发货请求
type OrderShipRequest struct {
	ExpressCompany string `json:"expressCompany" binding:"required"`
	ExpressNo      string `json:"expressNo" binding:"required"`
}

// AfterSaleListRequest 售后列表请求
type AfterSaleListRequest struct {
	Page     int  `form:"page" binding:"min=1"`
	PageSize int  `form:"pageSize" binding:"min=1,max=50"`
	Status   *int `form:"status"`
}

// AfterSaleHandleRequest 处理售后请求
type AfterSaleHandleRequest struct {
	Action       string `json:"action" binding:"required,oneof=agree reject"`
	RejectReason string `json:"rejectReason"`
}
