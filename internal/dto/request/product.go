package request

// ProductListRequest 商品列表请求
type ProductListRequest struct {
	Page       int    `form:"page" binding:"min=1"`
	PageSize   int    `form:"pageSize" binding:"min=1,max=50"`
	Keyword    string `form:"keyword"`
	CategoryID uint   `form:"categoryId"`
	Status     *int   `form:"status"`
}

// ProductCreateRequest 创建商品请求
type ProductCreateRequest struct {
	Name            string  `json:"name" binding:"required,max=100"`
	MainImage       string  `json:"mainImage" binding:"required"`
	Images          []string `json:"images"`
	CategoryID      uint    `json:"categoryId" binding:"required"`
	Price           int     `json:"price" binding:"required,min=0"`
	OriginalPrice   int     `json:"originalPrice"`
	Stock           int     `json:"stock" binding:"min=0"`
	Description     string  `json:"description"`
	Status          int     `json:"status"`
	IsDiscount      bool    `json:"isDiscount"`
	DiscountPrice   int     `json:"discountPrice"`
	DiscountStartAt *string `json:"discountStartAt"`
	DiscountEndAt   *string `json:"discountEndAt"`
	IsPresale       bool    `json:"isPresale"`
	PresalePrice    int     `json:"presalePrice"`
	PresaleStartAt  *string `json:"presaleStartAt"`
	PresaleEndAt    *string `json:"presaleEndAt"`
	EnableCarousel  bool    `json:"enableCarousel"`
	CarouselImage   string  `json:"carouselImage"`
}

// ProductUpdateRequest 更新商品请求
type ProductUpdateRequest struct {
	Name            string   `json:"name"`
	MainImage       string   `json:"mainImage"`
	Images          []string `json:"images"`
	CategoryID      uint     `json:"categoryId"`
	Price           int      `json:"price"`
	OriginalPrice   int      `json:"originalPrice"`
	Stock           int      `json:"stock"`
	Description     string   `json:"description"`
	Status          int      `json:"status"`
	IsDiscount      *bool    `json:"isDiscount"`
	DiscountPrice   int      `json:"discountPrice"`
	DiscountStartAt *string  `json:"discountStartAt"`
	DiscountEndAt   *string  `json:"discountEndAt"`
	IsPresale       *bool    `json:"isPresale"`
	PresalePrice    int      `json:"presalePrice"`
	PresaleStartAt  *string  `json:"presaleStartAt"`
	PresaleEndAt    *string  `json:"presaleEndAt"`
	EnableCarousel  *bool    `json:"enableCarousel"`
	CarouselImage   string   `json:"carouselImage"`
}

// ProductDiscountRequest 设置打折请求
type ProductDiscountRequest struct {
	IsDiscount      bool   `json:"isDiscount" binding:"required"`
	DiscountPrice   int    `json:"discountPrice"`
	DiscountStartAt string `json:"discountStartAt"`
	DiscountEndAt   string `json:"discountEndAt"`
}

// ProductPresaleRequest 设置预售请求
type ProductPresaleRequest struct {
	IsPresale      bool   `json:"isPresale" binding:"required"`
	PresalePrice   int    `json:"presalePrice"`
	PresaleStartAt string `json:"presaleStartAt"`
	PresaleEndAt   string `json:"presaleEndAt"`
}

// ProductStatusRequest 修改商品状态请求
type ProductStatusRequest struct {
	Status int `json:"status" binding:"required,oneof=0 1"`
}

// BatchDiscountRequest 批量打折请求
type BatchDiscountRequest struct {
	Items []ProductDiscountItem `json:"items" binding:"required,min=1"`
}

type ProductDiscountItem struct {
	ProductID       uint `json:"productId" binding:"required"`
	IsDiscount      bool `json:"isDiscount" binding:"required"`
	DiscountPrice   int  `json:"discountPrice"`
	DiscountStartAt string `json:"discountStartAt"`
	DiscountEndAt   string `json:"discountEndAt"`
}

// BatchPresaleRequest 批量预售请求
type BatchPresaleRequest struct {
	Items []ProductPresaleItem `json:"items" binding:"required,min=1"`
}

type ProductPresaleItem struct {
	ProductID      uint `json:"productId" binding:"required"`
	IsPresale      bool `json:"isPresale" binding:"required"`
	PresalePrice   int  `json:"presalePrice"`
	PresaleStartAt string `json:"presaleStartAt"`
	PresaleEndAt   string `json:"presaleEndAt"`
}
