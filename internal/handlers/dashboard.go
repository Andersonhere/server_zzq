package handlers

import (
	"net/http"

	"server_zzq/internal/utils"

	"github.com/gin-gonic/gin"
)

// DashboardOverview 数据概览
// GET /api/v1/manager/dashboard/overview
func DashboardOverview(c *gin.Context) {
	period := c.DefaultQuery("period", "today")

	// 验证 period 参数
	validPeriods := map[string]bool{
		"today": true, "yesterday": true, "7d": true, "30d": true,
	}
	if !validPeriods[period] {
		c.JSON(http.StatusBadRequest, utils.BadRequest("参数错误：period 不合法"))
		return
	}

	// TODO: 实现数据概览查询逻辑
	_ = period

	c.JSON(http.StatusOK, utils.Success(gin.H{
		"salesAmount":    0,
		"orderCount":     0,
		"newOrderCount":  0,
		"refundAmount":   0,
		"refundCount":    0,
	}))
}

// DashboardCompare 多周期数据对比
// GET /api/v1/manager/dashboard/compare
func DashboardCompare(c *gin.Context) {
	// TODO: 实现多周期数据对比逻辑
	c.JSON(http.StatusOK, utils.Success(gin.H{
		"today": gin.H{
			"salesAmount": 0,
			"orderCount":  0,
		},
		"yesterday": gin.H{
			"salesAmount": 0,
			"orderCount":  0,
		},
		"last7Days": gin.H{
			"salesAmount": 0,
			"orderCount":  0,
		},
		"last30Days": gin.H{
			"salesAmount": 0,
			"orderCount":  0,
		},
	}))
}

// DashboardTrend 销售趋势
// GET /api/v1/manager/dashboard/trend
func DashboardTrend(c *gin.Context) {
	days := 7
	// TODO: 实现销售趋势查询逻辑

	// 生成 mock 数据
	list := make([]gin.H, 0, days)
	for i := 0; i < days; i++ {
		list = append(list, gin.H{
			"date":        "2025-03-17",
			"salesAmount": 0,
			"orderCount":  0,
		})
	}

	c.JSON(http.StatusOK, utils.Success(gin.H{
		"list": list,
	}))
}

// DashboardPending 待处理事项
// GET /api/v1/manager/dashboard/pending
func DashboardPending(c *gin.Context) {
	// TODO: 实现待处理事项查询逻辑
	c.JSON(http.StatusOK, utils.Success(gin.H{
		"pendingShipCount":      0,
		"pendingAfterSaleCount": 0,
		"lowStockCount":         0,
	}))
}
