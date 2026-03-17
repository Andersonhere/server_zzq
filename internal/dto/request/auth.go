package request

// LoginRequest 商家登录请求
type LoginRequest struct {
	Code     string `json:"code" binding:"required"`
	ShopName string `json:"shopName"`
}

// RefreshTokenRequest 刷新 Token 请求（无需 body）
type RefreshTokenRequest struct{}

// LogoutRequest 退出登录请求（无需 body）
type LogoutRequest struct{}
