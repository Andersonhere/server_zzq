# 数据概览 接口规格

**文档路径**：api-docs/manager-backend/dashboard.md  
**用途**：商家管理端首页数据统计（销售额、订单数等）  
**适用端**：manager-miniprogram ↔ backend

---

## 1. 首页数据概览

| 项目 | 说明 |
|------|------|
| **路径** | `GET /api/v1/manager/dashboard/overview` |
| **鉴权** | 是 |
| **说明** | 获取今日/昨日/近7天/近30天的核心经营数据 |

### 请求参数（Query）

| 参数 | 类型 | 必填 | 默认 | 说明 |
|------|------|------|------|------|
| period | string | 否 | today | today / yesterday / 7d / 30d |

### 响应格式

**成功**：
```json
{
  "code": 0,
  "data": {
    "salesAmount": 125800,
    "orderCount": 28,
    "newOrderCount": 5,
    "refundAmount": 0,
    "refundCount": 0
  },
  "message": "success",
  "timestamp": 1710662400000
}
```

| 字段 | 类型 | 说明 |
|------|------|------|
| salesAmount | number | 销售额，单位：分 |
| orderCount | number | 订单总数 |
| newOrderCount | number | 新增订单数（如今日新下单） |
| refundAmount | number | 退款金额，单位：分 |
| refundCount | number | 退款笔数 |

### 错误码

| code | message | 说明 |
|------|---------|------|
| 401 | 未登录或 token 已失效 | 鉴权失败 |
| 400 | 参数错误：period 不合法 | period 值非法 |

---

## 2. 多周期数据对比（可选）

| 项目 | 说明 |
|------|------|
| **路径** | `GET /api/v1/manager/dashboard/compare` |
| **鉴权** | 是 |
| **说明** | 今日/昨日/近7天/近30天数据一次性返回，用于对比展示 |

### 请求参数

无，固定返回各周期数据。

### 响应格式

**成功**：
```json
{
  "code": 0,
  "data": {
    "today": {
      "salesAmount": 125800,
      "orderCount": 28
    },
    "yesterday": {
      "salesAmount": 98000,
      "orderCount": 22
    },
    "last7Days": {
      "salesAmount": 650000,
      "orderCount": 150
    },
    "last30Days": {
      "salesAmount": 2300000,
      "orderCount": 520
    }
  },
  "message": "success",
  "timestamp": 1710662400000
}
```

### 错误码

| code | message | 说明 |
|------|---------|------|
| 401 | 未登录或 token 已失效 | 鉴权失败 |

---

## 3. 销售额/订单数趋势（按日）

| 项目 | 说明 |
|------|------|
| **路径** | `GET /api/v1/manager/dashboard/trend` |
| **鉴权** | 是 |
| **说明** | 近 N 天每日销售额、订单数，用于折线图等图表 |

### 请求参数（Query）

| 参数 | 类型 | 必填 | 默认 | 说明 |
|------|------|------|------|------|
| days | number | 否 | 7 | 天数，最大 30 |

### 响应格式

**成功**：
```json
{
  "code": 0,
  "data": {
    "list": [
      {
        "date": "2025-03-11",
        "salesAmount": 45000,
        "orderCount": 12
      },
      {
        "date": "2025-03-12",
        "salesAmount": 52000,
        "orderCount": 15
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
| 401 | 未登录或 token 已失效 | 鉴权失败 |
| 400 | 参数错误：days 最大 30 | 超出范围 |

---

## 4. 待处理事项统计

| 项目 | 说明 |
|------|------|
| **路径** | `GET /api/v1/manager/dashboard/pending` |
| **鉴权** | 是 |
| **说明** | 待发货订单数、待处理售后数等，用于首页待办提醒 |

### 请求参数

无。

### 响应格式

**成功**：
```json
{
  "code": 0,
  "data": {
    "pendingShipCount": 8,
    "pendingAfterSaleCount": 2,
    "lowStockCount": 3
  },
  "message": "success",
  "timestamp": 1710662400000
}
```

| 字段 | 类型 | 说明 |
|------|------|------|
| pendingShipCount | number | 待发货订单数 |
| pendingAfterSaleCount | number | 待处理售后数 |
| lowStockCount | number | 库存不足商品数（可配置阈值，如 &lt;10） |

### 错误码

| code | message | 说明 |
|------|---------|------|
| 401 | 未登录或 token 已失效 | 鉴权失败 |
