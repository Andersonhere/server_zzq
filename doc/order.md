# 订单管理 接口规格

**文档路径**：api-docs/manager-backend/order.md  
**用途**：订单列表、详情、发货、售后处理  
**适用端**：manager-miniprogram ↔ backend

---

## 1. 订单列表（分页）

| 项目 | 说明 |
|------|------|
| **路径** | `GET /api/v1/manager/orders` |
| **鉴权** | 是 |
| **说明** | 分页查询当前商家的订单列表 |

### 请求参数（Query）

| 参数 | 类型 | 必填 | 默认 | 说明 |
|------|------|------|------|------|
| page | number | 否 | 1 | 页码 |
| pageSize | number | 否 | 10 | 每页条数（最大 50） |
| status | number | 否 | - | 订单状态筛选，见下表 |
| keyword | string | 否 | - | 订单号或用户手机号搜索 |
| startDate | string | 否 | - | 开始日期 YYYY-MM-DD |
| endDate | string | 否 | - | 结束日期 YYYY-MM-DD |

**订单状态 status**：

| 值 | 含义 |
|----|------|
| 0 | 待支付 |
| 1 | 待发货 |
| 2 | 待收货 |
| 3 | 已完成 |
| 4 | 已取消 |
| 5 | 售后中 |

### 响应格式

**成功**：
```json
{
  "code": 0,
  "data": {
    "list": [
      {
        "id": 1001,
        "orderNo": "ORD202503171200001",
        "userId": 10,
        "userName": "张三",
        "userPhone": "138****8000",
        "totalAmount": 3500,
        "status": 1,
        "statusText": "待发货",
        "itemCount": 2,
        "createdAt": "2025-03-17T12:00:00.000Z"
      }
    ],
    "total": 100,
    "page": 1,
    "pageSize": 10
  },
  "message": "success",
  "timestamp": 1710662400000
}
```

| 字段 | 类型 | 说明 |
|------|------|------|
| totalAmount | number | 订单总金额，单位：分 |
| status | number | 订单状态码 |
| statusText | string | 状态中文描述 |
| itemCount | number | 商品种类数 |

### 错误码

| code | message | 说明 |
|------|---------|------|
| 401 | 未登录或 token 已失效 | 鉴权失败 |
| 400 | 参数错误 | 日期格式或分页参数不合法 |

---

## 2. 订单详情

| 项目 | 说明 |
|------|------|
| **路径** | `GET /api/v1/manager/orders/:id` |
| **鉴权** | 是 |
| **说明** | 获取订单详情及商品明细 |

### 请求参数

| 参数 | 位置 | 类型 | 必填 | 说明 |
|------|------|------|------|------|
| id | path | number | 是 | 订单 ID |

### 响应格式

**成功**：
```json
{
  "code": 0,
  "data": {
    "id": 1001,
    "orderNo": "ORD202503171200001",
    "status": 1,
    "statusText": "待发货",
    "totalAmount": 3500,
    "payAmount": 3500,
    "freightAmount": 0,
    "discountAmount": 0,
    "payTime": "2025-03-17T12:01:00.000Z",
    "createdAt": "2025-03-17T12:00:00.000Z",
    "receiver": {
      "name": "张三",
      "phone": "13800138000",
      "address": "某市某区某街道某号"
    },
    "items": [
      {
        "productId": 1,
        "productName": "新鲜苹果",
        "mainImage": "https://xxx.com/apple.jpg",
        "price": 1500,
        "quantity": 2,
        "amount": 3000,
        "specText": "500g/份"
      }
    ],
    "afterSale": null
  },
  "message": "success",
  "timestamp": 1710662400000
}
```

| 字段 | 类型 | 说明 |
|------|------|------|
| totalAmount | number | 订单总金额，单位：分 |
| payAmount | number | 实付金额，单位：分 |
| freightAmount | number | 运费，单位：分 |
| discountAmount | number | 优惠金额，单位：分 |
| afterSale | object/null | 若有售后则包含售后信息 |

### 错误码

| code | message | 说明 |
|------|---------|------|
| 401 | 未登录或 token 已失效 | 鉴权失败 |
| 404 | 订单不存在 | 订单不存在或无权限 |
| 1200 | 订单不存在 | 业务错误码 |

---

## 3. 发货

| 项目 | 说明 |
|------|------|
| **路径** | `POST /api/v1/manager/orders/:id/ship` |
| **鉴权** | 是 |
| **说明** | 商家发货，填写物流信息 |

### 请求参数

**Path**：`id`（订单 ID）

**Body JSON**：

| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| expressCompany | string | 是 | 快递公司名称 |
| expressNo | string | 是 | 快递单号 |

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
| 400 | 参数错误：expressCompany、expressNo 必填 | 参数缺失 |
| 401 | 未登录或 token 已失效 | 鉴权失败 |
| 404 | 订单不存在 | 订单不存在 |
| 1201 | 订单状态不允许发货 | 非待发货状态 |
| 1202 | 订单已发货 | 重复发货 |

---

## 4. 售后列表

| 项目 | 说明 |
|------|------|
| **路径** | `GET /api/v1/manager/after-sales` |
| **鉴权** | 是 |
| **说明** | 分页查询售后申请列表 |

### 请求参数（Query）

| 参数 | 类型 | 必填 | 默认 | 说明 |
|------|------|------|------|------|
| page | number | 否 | 1 | 页码 |
| pageSize | number | 否 | 10 | 每页条数 |
| status | number | 否 | - | 0 待处理 1 已同意 2 已拒绝 3 已完成 |

### 响应格式

**成功**：
```json
{
  "code": 0,
  "data": {
    "list": [
      {
        "id": 1,
        "orderId": 1001,
        "orderNo": "ORD202503171200001",
        "type": 1,
        "typeText": "仅退款",
        "status": 0,
        "statusText": "待处理",
        "reason": "商品质量问题",
        "refundAmount": 1500,
        "createdAt": "2025-03-17T14:00:00.000Z"
      }
    ],
    "total": 5,
    "page": 1,
    "pageSize": 10
  },
  "message": "success",
  "timestamp": 1710662400000
}
```

**售后类型 type**：1 仅退款 2 退货退款

### 错误码

| code | message | 说明 |
|------|---------|------|
| 401 | 未登录或 token 已失效 | 鉴权失败 |

---

## 5. 处理售后（同意/拒绝）

| 项目 | 说明 |
|------|------|
| **路径** | `PUT /api/v1/manager/after-sales/:id/handle` |
| **鉴权** | 是 |
| **说明** | 商家同意或拒绝售后申请 |

### 请求参数

**Path**：`id`（售后单 ID）

**Body JSON**：

| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| action | string | 是 | agree 同意 / reject 拒绝 |
| rejectReason | string | 否 | action=reject 时必填，拒绝原因 |

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
| 400 | 参数错误：action 必填 | action 缺失 |
| 400 | 拒绝时请填写拒绝原因 | reject 时 rejectReason 缺失 |
| 401 | 未登录或 token 已失效 | 鉴权失败 |
| 404 | 售后单不存在 | 售后单不存在 |
| 1203 | 售后单状态不允许操作 | 非待处理状态 |
