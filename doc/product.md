# 商品管理 接口规格

**文档路径**：api-docs/manager-backend/product.md  
**用途**：商品增删改查、上下架  
**适用端**：manager-miniprogram ↔ backend

---

## 1. 商品列表（分页）

| 项目 | 说明 |
|------|------|
| **路径** | `GET /api/v1/manager/products` |
| **鉴权** | 是 |
| **说明** | 分页查询当前商家的商品列表 |

### 请求参数（Query）

| 参数 | 类型 | 必填 | 默认 | 说明 |
|------|------|------|------|------|
| page | number | 否 | 1 | 页码 |
| pageSize | number | 否 | 10 | 每页条数（最大 50） |
| keyword | string | 否 | - | 商品名称关键词搜索 |
| categoryId | number | 否 | - | 分类 ID 筛选 |
| status | number | 否 | - | 0 下架 1 上架，不传查全部 |

### 响应格式

**成功**：
```json
{
  "code": 0,
  "data": {
    "list": [
      {
        "id": 1,
        "name": "新鲜苹果",
        "mainImage": "https://xxx.com/apple.jpg",
        "categoryId": 1,
        "categoryName": "水果",
        "price": 1500,
        "originalPrice": 2000,
        "stock": 100,
        "status": 1,
        "salesCount": 50,
        "isDiscount": true,
        "discountPrice": 1200,
        "discountStartAt": "2025-03-17T00:00:00.000Z",
        "discountEndAt": "2025-03-20T23:59:59.000Z",
        "isPresale": false,
        "presalePrice": null,
        "presaleStartAt": null,
        "presaleEndAt": null,
        "enableCarousel": true,
        "carouselImage": "https://xxx.com/banner1.jpg",
        "createdAt": "2025-01-01T00:00:00.000Z"
      }
    ],
    "total": 20,
    "page": 1,
    "pageSize": 10
  },
  "message": "success",
  "timestamp": 1710662400000
}
```

| 字段 | 类型 | 说明 |
|------|------|------|
| price | number | 当前展示价格，单位：分（打折或预售生效时为优惠价，否则为原价） |
| originalPrice | number | 原价，单位：分 |
| stock | number | 库存 |
| status | number | 1 上架 0 下架 |
| salesCount | number | 销量 |
| isDiscount | boolean | 是否开启限时打折属性 |
| discountPrice | number | 打折价格，单位：分（isDiscount=true 时必有值） |
| discountStartAt | string | 打折开始时间（ISO 时间） |
| discountEndAt | string | 打折结束时间，到期后台自动恢复原价并关闭打折属性 |
| isPresale | boolean | 是否开启预售属性 |
| presalePrice | number/null | 预售价，单位：分（isPresale=true 时必有值） |
| presaleStartAt | string/null | 预售开始时间 |
| presaleEndAt | string/null | 预售结束时间，到期后台自动恢复原价并关闭预售属性 |
| enableCarousel | boolean | 是否投放到首页轮播图（true 表示该商品出现在首页轮播位） |
| carouselImage | string/null | 该商品在首页轮播中使用的一张图片 URL，enableCarousel=true 时必有值；全店最多支持 5 个 enableCarousel=true 的商品 |

### 错误码

| code | message | 说明 |
|------|---------|------|
| 401 | 未登录或 token 已失效 | 鉴权失败 |
| 400 | 参数错误 | page/pageSize 不合法 |

---

## 2. 商品详情

| 项目 | 说明 |
|------|------|
| **路径** | `GET /api/v1/manager/products/:id` |
| **鉴权** | 是 |
| **说明** | 根据 ID 获取商品详情 |

### 请求参数

| 参数 | 位置 | 类型 | 必填 | 说明 |
|------|------|------|------|------|
| id | path | number | 是 | 商品 ID |

### 响应格式

**成功**：
```json
{
  "code": 0,
  "data": {
    "id": 1,
    "name": "新鲜苹果",
    "mainImage": "https://xxx.com/apple.jpg",
    "images": ["https://xxx.com/1.jpg", "https://xxx.com/2.jpg"],
    "categoryId": 1,
    "categoryName": "水果",
    "price": 1500,
    "originalPrice": 2000,
    "stock": 100,
    "status": 1,
    "description": "商品描述",
    "salesCount": 50,
    "isDiscount": true,
    "discountPrice": 1200,
    "discountStartAt": "2025-03-17T00:00:00.000Z",
    "discountEndAt": "2025-03-20T23:59:59.000Z",
    "isPresale": false,
    "presalePrice": null,
    "presaleStartAt": null,
    "presaleEndAt": null,
    "enableCarousel": true,
    "carouselImage": "https://xxx.com/banner1.jpg",
    "createdAt": "2025-01-01T00:00:00.000Z",
    "updatedAt": "2025-01-01T00:00:00.000Z"
  },
  "message": "success",
  "timestamp": 1710662400000
}
```

### 错误码

| code | message | 说明 |
|------|---------|------|
| 401 | 未登录或 token 已失效 | 鉴权失败 |
| 404 | 商品不存在 | 商品不存在或无权限 |
| 1100 | 商品不存在 | 业务错误码 |

---

## 3. 新增商品

| 项目 | 说明 |
|------|------|
| **路径** | `POST /api/v1/manager/products` |
| **鉴权** | 是 |
| **说明** | 创建新商品 |

### 请求参数（Body JSON）

| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| name | string | 是 | 商品名称，1-100 字 |
| mainImage | string | 是 | 主图 URL |
| images | array | 否 | 商品详情轮播图 URL 数组（与首页轮播无关） |
| categoryId | number | 是 | 分类 ID |
| price | number | 是 | 售价，单位：分（原价） |
| originalPrice | number | 否 | 原价，单位：分，不传则与 price 相同 |
| stock | number | 是 | 初始库存，≥0 |
| description | string | 否 | 商品描述 |
| status | number | 否 | 0 下架 1 上架，默认 0 |
| isDiscount | boolean | 否 | 是否初始化为打折商品，默认 false；**与 isPresale 互斥，同一商品只能二选一或都为 false** |
| discountPrice | number | 否 | 打折价格，单位：分，isDiscount=true 时必填且小于 price |
| discountStartAt | string | 否 | 打折开始时间，ISO 字符串 |
| discountEndAt | string | 否 | 打折结束时间，必大于开始时间，到期后台自动恢复原价并关闭打折属性 |
| isPresale | boolean | 否 | 是否开启预售属性，默认 false；**与 isDiscount 互斥，同一商品只能二选一或都为 false** |
| presalePrice | number | 否 | 预售价，单位：分，isPresale=true 时必填且可小于或等于 price |
| presaleStartAt | string | 否 | 预售开始时间，ISO 字符串 |
| presaleEndAt | string | 否 | 预售结束时间，必大于开始时间，到期后台自动恢复原价并关闭预售属性 |
| enableCarousel | boolean | 否 | 是否投放到首页轮播图，默认 false，全店最多允许 5 个商品 enableCarousel=true |
| carouselImage | string | 否 | 首页轮播使用的一张图片 URL，enableCarousel=true 时必填 |

### 响应格式

**成功**：
```json
{
  "code": 0,
  "data": {
    "id": 1
  },
  "message": "success",
  "timestamp": 1710662400000
}
```

### 错误码

| code | message | 说明 |
|------|---------|------|
| 400 | 参数错误：xxx | 必填缺失或格式错误 |
| 401 | 未登录或 token 已失效 | 鉴权失败 |
| 1101 | 分类不存在 | categoryId 无效 |
| 1102 | 库存不能为负 | stock < 0 |
| 1112 | 打折配置不合法 | isDiscount=true 但价格/时间不合法，或与 isPresale 同时为 true |
| 1113 | 预售配置不合法 | isPresale=true 但价格/时间不合法，或与 isDiscount 同时为 true |
| 1114 | 首页轮播位已达上限（5 个） | 当前 enableCarousel=true 的商品数量已达 5 个 |

---

## 4. 更新商品

| 项目 | 说明 |
|------|------|
| **路径** | `PUT /api/v1/manager/products/:id` |
| **鉴权** | 是 |
| **说明** | 修改商品信息 |

### 请求参数

**Path**：`id`（商品 ID）

**Body JSON**（所有字段均为可选，仅传需修改的字段）：

| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| name | string | 否 | 商品名称 |
| mainImage | string | 否 | 主图 URL |
| images | array | 否 | 商品详情轮播图 URL 数组（与首页轮播无关） |
| categoryId | number | 否 | 分类 ID |
| price | number | 否 | 售价，单位：分（原价） |
| originalPrice | number | 否 | 原价，单位：分 |
| stock | number | 否 | 库存 |
| description | string | 否 | 商品描述 |
| status | number | 否 | 0 下架 1 上架 |
| isDiscount | boolean | 否 | 是否开启/关闭打折属性，true 开启，false 关闭；**与 isPresale 互斥** |
| discountPrice | number | 否 | 打折价格，单位：分，isDiscount=true 时必填且小于 price 或 originalPrice |
| discountStartAt | string | 否 | 打折开始时间，ISO 字符串 |
| discountEndAt | string | 否 | 打折结束时间，到期后台自动恢复原价并关闭打折属性 |
| isPresale | boolean | 否 | 是否开启/关闭预售属性；**与 isDiscount 互斥** |
| presalePrice | number | 否 | 预售价，单位：分，isPresale=true 时必填 |
| presaleStartAt | string | 否 | 预售开始时间，ISO 字符串 |
| presaleEndAt | string | 否 | 预售结束时间，到期后台自动恢复原价并关闭预售属性 |
| enableCarousel | boolean | 否 | 是否开启/关闭首页轮播属性，true 表示首页轮播展示该商品，全店最多 5 个 |
| carouselImage | string | 否 | 首页轮播使用的一张图片 URL，enableCarousel=true 时必填 |

### 响应格式

**成功**：
```json
{
  "code": 0,
  "data": null,
  "message": "success",
  "timestamp": 1710662400000
}
```

### 错误码

| code | message | 说明 |
|------|---------|------|
| 400 | 参数错误 | 参数不合法 |
| 401 | 未登录或 token 已失效 | 鉴权失败 |
| 404 | 商品不存在 | 商品不存在或无权限 |
| 1100 | 商品不存在 | 业务错误码 |
| 1112 | 打折配置不合法 | isDiscount/discountPrice/时间段非法，或与 isPresale 同时为 true |
| 1113 | 预售配置不合法 | isPresale/presalePrice/时间段非法，或与 isDiscount 同时为 true |
| 1114 | 首页轮播位已达上限（5 个） | enableCarousel=true 的商品数量超过 5 个 |

---

## 5. 删除商品（软删除）

| 项目 | 说明 |
|------|------|
| **路径** | `DELETE /api/v1/manager/products/:id` |
| **鉴权** | 是 |
| **说明** | 软删除商品（is_deleted=1） |

### 请求参数

| 参数 | 位置 | 类型 | 必填 | 说明 |
|------|------|------|------|------|
| id | path | number | 是 | 商品 ID |

### 响应格式

**成功**：
```json
{
  "code": 0,
  "data": null,
  "message": "success",
  "timestamp": 1710662400000
}
```

### 错误码

| code | message | 说明 |
|------|---------|------|
| 401 | 未登录或 token 已失效 | 鉴权失败 |
| 404 | 商品不存在 | 商品不存在或无权限 |
| 1103 | 商品存在未完成订单，无法删除 | 有进行中订单关联 |

---

## 6. 上架/下架商品

| 项目 | 说明 |
|------|------|
| **路径** | `PUT /api/v1/manager/products/:id/status` |
| **鉴权** | 是 |
| **说明** | 批量或单个修改商品上下架状态 |

### 请求参数

**Path**：`id`（商品 ID）

**Body JSON**：

| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| status | number | 是 | 0 下架 1 上架 |

### 响应格式

**成功**：
```json
{
  "code": 0,
  "data": null,
  "message": "success",
  "timestamp": 1710662400000
}
```

### 错误码

| code | message | 说明 |
|------|---------|------|
| 400 | 参数错误：status 必填且为 0 或 1 | 参数不合法 |
| 401 | 未登录或 token 已失效 | 鉴权失败 |
| 404 | 商品不存在 | 商品不存在或无权限 |
| 1104 | 库存为 0 无法上架 | 库存不足不允许上架 |

---

## 7. 设置商品打折属性（单个）

> 说明：为单个商品设置/关闭限时打折价格及有效期，到期后后台自动恢复原价并去除打折属性。每个商家最多允许同时存在 **10 个处于打折中的商品**。

| 项目 | 说明 |
|------|------|
| **路径** | `PUT /api/v1/manager/products/:id/discount` |
| **鉴权** | 是 |
| **说明** | 为单个商品设置/取消打折属性 |

### 请求参数

**Path**：`id`（商品 ID）

**Body JSON**：

| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| isDiscount | boolean | 是 | true 开启打折，false 关闭打折 |
| discountPrice | number | 否 | 打折价格，单位：分，isDiscount=true 时必填且小于当前原价 |
| discountStartAt | string | 否 | 打折开始时间，ISO 字符串，不传则默认立即生效 |
| discountEndAt | string | 否 | 打折结束时间，必大于开始时间，到期后台自动恢复原价并关闭打折属性 |

### 响应格式

**成功**：
```json
{
  "code": 0,
  "data": null,
  "message": "success",
  "timestamp": 1710662400000
}
```

### 错误码

| code | message | 说明 |
|------|---------|------|
| 400 | 参数错误 | 参数格式或时间段不合法 |
| 401 | 未登录或 token 已失效 | 鉴权失败 |
| 404 | 商品不存在 | 商品不存在或无权限 |
| 1112 | 打折配置不合法 | discountPrice/时间段不合法 |
| 1115 | 当前打折商品已达上限（10 个） | 开启打折时超出数量上限 |

---

## 8. 批量设置商品打折属性

> 说明：批量为多个商品设置打折属性，同样受「同时处于打折中的商品最多 10 个」限制。

| 项目 | 说明 |
|------|------|
| **路径** | `POST /api/v1/manager/products/discount/batch` |
| **鉴权** | 是 |
| **说明** | 批量设置多个商品打折属性 |

### 请求参数（Body JSON）

| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| items | array | 是 | 打折配置数组 |
| items[].productId | number | 是 | 商品 ID |
| items[].isDiscount | boolean | 是 | true 开启/更新打折，false 关闭打折 |
| items[].discountPrice | number | 否 | 打折价格，单位：分，isDiscount=true 时必填 |
| items[].discountStartAt | string | 否 | 打折开始时间，ISO 字符串 |
| items[].discountEndAt | string | 否 | 打折结束时间，必须大于开始时间 |

### 响应格式

**成功**：
```json
{
  "code": 0,
  "data": {
    "successCount": 3,
    "failCount": 1,
    "failItems": [
      {
        "productId": 4,
        "message": "当前打折商品已达上限（10 个）"
      }
    ]
  },
  "message": "success",
  "timestamp": 1710662400000
}
```

### 错误码

| code | message | 说明 |
|------|---------|------|
| 400 | 参数错误：items 必填且非空数组 | 参数不合法 |
| 401 | 未登录或 token 已失效 | 鉴权失败 |
| 1112 | 部分打折配置不合法 | discountPrice/时间段不合法 |
| 1115 | 当前打折商品已达上限（10 个） | 对于开启打折的商品，当最终打折中的商品超过 10 个 |

---

## 9. 设置商品预售属性（单个）

> 说明：为单个商品设置/关闭预售属性和预售价，到预售结束时间后台自动恢复原价并关闭预售属性。每个商家最多允许同时存在 **10 个处于预售中的商品**。

| 项目 | 说明 |
|------|------|
| **路径** | `PUT /api/v1/manager/products/:id/presale` |
| **鉴权** | 是 |
| **说明** | 为单个商品设置/取消预售属性 |

### 请求参数

**Path**：`id`（商品 ID）

**Body JSON**：

| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| isPresale | boolean | 是 | true 开启预售，false 关闭预售 |
| presalePrice | number | 否 | 预售价，单位：分，isPresale=true 时必填 |
| presaleStartAt | string | 否 | 预售开始时间，ISO 字符串，不传则立即生效 |
| presaleEndAt | string | 否 | 预售结束时间，必大于开始时间，到期后台自动恢复原价并关闭预售属性 |

### 响应格式

**成功**：
```json
{
  "code": 0,
  "data": null,
  "message": "success",
  "timestamp": 1710662400000
}
```

### 错误码

| code | message | 说明 |
|------|---------|------|
| 400 | 参数错误 | 参数格式或时间段不合法 |
| 401 | 未登录或 token 已失效 | 鉴权失败 |
| 404 | 商品不存在 | 商品不存在或无权限 |
| 1113 | 预售配置不合法 | presalePrice/时间段不合法 |
| 1116 | 当前预售商品已达上限（10 个） | 开启预售时超出数量上限 |

---

## 10. 批量设置商品预售属性

> 说明：批量为多个商品设置预售属性，同样受「同时处于预售中的商品最多 10 个」限制。

| 项目 | 说明 |
|------|------|
| **路径** | `POST /api/v1/manager/products/presale/batch` |
| **鉴权** | 是 |
| **说明** | 批量设置多个商品预售属性 |

### 请求参数（Body JSON）

| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| items | array | 是 | 预售配置数组 |
| items[].productId | number | 是 | 商品 ID |
| items[].isPresale | boolean | 是 | true 开启/更新预售，false 关闭预售 |
| items[].presalePrice | number | 否 | 预售价，单位：分，isPresale=true 时必填 |
| items[].presaleStartAt | string | 否 | 预售开始时间，ISO 字符串 |
| items[].presaleEndAt | string | 否 | 预售结束时间，必须大于开始时间 |

### 响应格式

**成功**：
```json
{
  "code": 0,
  "data": {
    "successCount": 3,
    "failCount": 1,
    "failItems": [
      {
        "productId": 5,
        "message": "当前预售商品已达上限（10 个）"
      }
    ]
  },
  "message": "success",
  "timestamp": 1710662400000
}
```

### 错误码

| code | message | 说明 |
|------|---------|------|
| 400 | 参数错误：items 必填且非空数组 | 参数不合法 |
| 401 | 未登录或 token 已失效 | 鉴权失败 |
| 1113 | 部分预售配置不合法 | presalePrice/时间段不合法 |
| 1116 | 当前预售商品已达上限（10 个） | 对于开启预售的商品，当最终预售中的商品超过 10 个 |
