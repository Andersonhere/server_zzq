# 库存管理 接口规格

**文档路径**：api-docs/manager-backend/inventory.md  
**用途**：库存查询、库存调整  
**适用端**：manager-miniprogram ↔ backend

---

## 1. 库存列表（分页）

| 项目 | 说明 |
|------|------|
| **路径** | `GET /api/v1/manager/inventory` |
| **鉴权** | 是 |
| **说明** | 分页查询商品库存，支持筛选低库存 |

### 请求参数（Query）

| 参数 | 类型 | 必填 | 默认 | 说明 |
|------|------|------|------|------|
| page | number | 否 | 1 | 页码 |
| pageSize | number | 否 | 10 | 每页条数（最大 50） |
| keyword | string | 否 | - | 商品名称搜索 |
| lowStock | number | 否 | - | 1 仅查库存不足商品（可设阈值） |
| threshold | number | 否 | 10 | lowStock=1 时的库存阈值，默认 10 |

### 响应格式

**成功**：
```json
{
  "code": 0,
  "data": {
    "list": [
      {
        "productId": 1,
        "productName": "新鲜苹果",
        "mainImage": "https://xxx.com/apple.jpg",
        "stock": 5,
        "lockedStock": 2,
        "availableStock": 3,
        "status": 1,
        "updatedAt": "2025-03-17T10:00:00.000Z"
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
| stock | number | 总库存 |
| lockedStock | number | 锁定库存（待支付订单占用） |
| availableStock | number | 可用库存 = stock - lockedStock |
| status | number | 1 上架 0 下架 |

### 错误码

| code | message | 说明 |
|------|---------|------|
| 401 | 未登录或 token 已失效 | 鉴权失败 |
| 400 | 参数错误 | 分页参数不合法 |

---

## 2. 单个商品库存详情

| 项目 | 说明 |
|------|------|
| **路径** | `GET /api/v1/manager/inventory/:productId` |
| **鉴权** | 是 |
| **说明** | 获取指定商品的库存详情 |

### 请求参数

| 参数 | 位置 | 类型 | 必填 | 说明 |
|------|------|------|------|------|
| productId | path | number | 是 | 商品 ID |

### 响应格式

**成功**：
```json
{
  "code": 0,
  "data": {
    "productId": 1,
    "productName": "新鲜苹果",
    "stock": 100,
    "lockedStock": 5,
    "availableStock": 95,
    "status": 1,
    "updatedAt": "2025-03-17T10:00:00.000Z"
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

## 3. 库存调整

| 项目 | 说明 |
|------|------|
| **路径** | `PUT /api/v1/manager/inventory/:productId/adjust` |
| **鉴权** | 是 |
| **说明** | 增减库存，用于盘点、补货等 |

### 请求参数

**Path**：`productId`（商品 ID）

**Body JSON**：

| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| delta | number | 是 | 调整量，正数增加、负数减少 |
| reason | string | 否 | 调整原因，便于追溯 |

### 响应格式

**成功**：
```json
{
  "code": 0,
  "data": {
    "productId": 1,
    "beforeStock": 100,
    "afterStock": 150,
    "delta": 50
  },
  "message": "success",
  "timestamp": 1710662400000
}
```

### 错误码

| code | message | 说明 |
|------|---------|------|
| 400 | 参数错误：delta 必填且不能为 0 | 参数不合法 |
| 401 | 未登录或 token 已失效 | 鉴权失败 |
| 404 | 商品不存在 | 商品不存在 |
| 1110 | 库存不足，无法减少 | 减少后库存将为负 |
| 1111 | 减少后库存不能低于已锁定数量 | 逻辑约束 |

---

## 4. 批量库存调整（可选）

| 项目 | 说明 |
|------|------|
| **路径** | `POST /api/v1/manager/inventory/batch-adjust` |
| **鉴权** | 是 |
| **说明** | 一次性调整多个商品库存 |

### 请求参数（Body JSON）

| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| items | array | 是 | 调整项数组 |
| items[].productId | number | 是 | 商品 ID |
| items[].delta | number | 是 | 调整量，正增负减 |
| reason | string | 否 | 统一调整原因 |

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
        "message": "库存不足，无法减少"
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
