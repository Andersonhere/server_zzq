package handlers

import (
	"net/http"
	"time"

	"server_zzq/internal/dto/request"
	"server_zzq/internal/dto/response"
	"server_zzq/internal/utils"

	"github.com/gin-gonic/gin"
)

// Login 商家登录
// POST /api/v1/manager/auth/login
func Login(c *gin.Context) {
	var req request.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, utils.BadRequest("参数错误："+err.Error()))
		return
	}

	// TODO: 实现微信 code 兑换逻辑
	// 1. 调用微信 API 用 code 换 openid
	// 2. 查询商家信息
	// 3. 生成 JWT token

	// 模拟商家 ID（实际应从数据库查询）
	shopID := uint(1)
	token, err := utils.GenerateToken(shopID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.InternalError("生成 token 失败："+err.Error()))
		return
	}

	resp := &response.LoginResponse{
		Token:    token,
		ExpireIn: 7200,
		Shop: &response.ShopInfo{
			ID:         shopID,
			ShopName:   req.ShopName,
			Logo:       "",
			ContactPhone: "",
			Status:     1,
			CreatedAt:  time.Now(),
		},
	}

	c.JSON(http.StatusOK, utils.Success(resp))
}

// RefreshToken 刷新 Token
// POST /api/v1/manager/auth/refresh
func RefreshToken(c *gin.Context) {
	// TODO: 实现刷新 token 逻辑
	c.JSON(http.StatusOK, utils.Success(gin.H{
		"token":    "new-mock-token",
		"expireIn": 7200,
	}))
}

// Logout 退出登录
// POST /api/v1/manager/auth/logout
func Logout(c *gin.Context) {
	// TODO: 将 token 加入黑名单
	c.JSON(http.StatusOK, utils.Success(nil))
}

// GetProfile 获取当前商家信息
// GET /api/v1/manager/auth/profile
func GetProfile(c *gin.Context) {
	shopID, exists := c.Get("shopID")
	if !exists {
		c.JSON(http.StatusUnauthorized, utils.Unauthorized("未登录"))
		return
	}

	// TODO: 根据 shopID 查询商家信息
	_ = shopID

	c.JSON(http.StatusOK, utils.Success(gin.H{
		"id":           1,
		"shopName":     "测试店铺",
		"logo":         "",
		"contactPhone": "",
		"status":       1,
		"createdAt":    "2025-01-01T00:00:00.000Z",
	}))
}
