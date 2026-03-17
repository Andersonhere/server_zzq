package request

// CategoryListRequest 分类列表请求
type CategoryListRequest struct {
	Tree   int `form:"tree"`
	Status *int `form:"status"`
}

// CategoryCreateRequest 创建分类请求
type CategoryCreateRequest struct {
	Name     string `json:"name" binding:"required,max=20"`
	ParentID uint   `json:"parentId"`
	Sort     int    `json:"sort"`
	Status   int    `json:"status"`
}

// CategoryUpdateRequest 更新分类请求
type CategoryUpdateRequest struct {
	Name     string `json:"name"`
	ParentID uint   `json:"parentId"`
	Sort     int    `json:"sort"`
	Status   int    `json:"status"`
}
