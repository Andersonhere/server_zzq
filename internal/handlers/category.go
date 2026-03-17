package handlers

import (
	"net/http"
	"strconv"

	"server_zzq/internal/dto/request"
	"server_zzq/internal/utils"

	"github.com/gin-gonic/gin"
)

// ListCategories 分类列表
// GET /api/v1/manager/categories
func ListCategories(c *gin.Context) {
	var req request.CategoryListRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, utils.BadRequest("参数错误："+err.Error()))
		return
	}

	// TODO: 实现分类列表查询逻辑
	_ = req

	c.JSON(http.StatusOK, utils.Success(gin.H{
		"list": []interface{}{},
	}))
}

// GetCategory 分类详情
// GET /api/v1/manager/categories/:id
func GetCategory(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.BadRequest("分类 ID 格式错误"))
		return
	}

	// TODO: 实现分类详情查询逻辑
	_ = id

	c.JSON(http.StatusOK, utils.Success(gin.H{}))
}

// CreateCategory 创建分类
// POST /api/v1/manager/categories
func CreateCategory(c *gin.Context) {
	var req request.CategoryCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, utils.BadRequest("参数错误："+err.Error()))
		return
	}

	// TODO: 实现创建分类逻辑

	c.JSON(http.StatusOK, utils.Success(gin.H{"id": 1}))
}

// UpdateCategory 更新分类
// PUT /api/v1/manager/categories/:id
func UpdateCategory(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.BadRequest("分类 ID 格式错误"))
		return
	}

	var req request.CategoryUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, utils.BadRequest("参数错误："+err.Error()))
		return
	}

	// TODO: 实现更新分类逻辑
	_ = id
	_ = req

	c.JSON(http.StatusOK, utils.Success(nil))
}

// DeleteCategory 删除分类
// DELETE /api/v1/manager/categories/:id
func DeleteCategory(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.BadRequest("分类 ID 格式错误"))
		return
	}

	// TODO: 实现删除分类逻辑
	_ = id

	c.JSON(http.StatusOK, utils.Success(nil))
}
