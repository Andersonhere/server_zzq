package response

import "time"

// LoginResponse 登录响应
type LoginResponse struct {
	Token    string     `json:"token"`
	ExpireIn int        `json:"expireIn"`
	Shop     *ShopInfo  `json:"shop"`
}

// RefreshTokenResponse 刷新 Token 响应
type RefreshTokenResponse struct {
	Token    string `json:"token"`
	ExpireIn int    `json:"expireIn"`
}

// ShopInfo 商家信息
type ShopInfo struct {
	ID         uint      `json:"id"`
	ShopName   string    `json:"shopName"`
	Logo       string    `json:"logo"`
	ContactPhone string  `json:"contactPhone"`
	Status     int       `json:"status"`
	CreatedAt  time.Time `json:"createdAt"`
}
