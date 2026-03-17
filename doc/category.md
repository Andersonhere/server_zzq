# 分类管理 接口规格

**文档路径**：api-docs/manager-backend/category.md  
**用途**：商品分类增删改查  
**适用端**：manager-miniprogram ↔ backend

---

## 1. 分类列表

| 项目 | 说明 |
|------|------|
| **路径** | `GET /api/v1/manager/categories` |
| **鉴权** | 是 |
| **说明** | 获取当前商家的全部分类（树形或平铺） |

### 请求参数（Query）

| 参数 | 类型 | 必填 | 默认 | 说明 |
|------|------|------|------|------|
| tree | number | 否 | 0 | 1 返回树形结构，0 平铺列表 |
| status | number | 否 | - | 0 禁用 1 启用，不传查全部 |

### 响应格式（平铺 list，tree=0）

**成功**：
```json
{
  "code": 0,
  "data": {
    "list": [
      {
        "id": 1,
        "name": "水果",
        "parentId": 0,
        "sort": 10,
        "status": 1,
        "productCount": 5,
        "createdAt": "2025-01-01T00:00:00.000Z"
      }
    ]
  },
  "message": "success",
  "timestamp": 1710662400000
}
```

### 响应格式（树形，tree=1）

**成功**：
```json
{
  "code": 0,
  "data": {
    "list": [
      {
        "id": 1,
        "name": "水果",
        "parentId": 0,
        "sort": 10,
        "status": 1,
        "productCount": 5,
        "children": [
          {
            "id": 2,
            "name": "苹果",
            "parentId": 1,
            "sort": 1,
            "status": 1,
            "productCount": 2,
            "children": []
          }
        ]
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

---

## 2. 分类详情

| 项目 | 说明 |
|------|------|
| **路径** | `GET /api/v1/manager/categories/:id` |
| **鉴权** | 是 |
| **说明** | 根据 ID 获取分类详情 |

### 请求参数

| 参数 | 位置 | 类型 | 必填 | 说明 |
|------|------|------|------|------|
| id | path | number | 是 | 分类 ID |

### 响应格式

**成功**：
```json
{
  "code": 0,
  "data": {
    "id": 1,
    "name": "水果",
    "parentId": 0,
    "sort": 10,
    "status": 1,
    "productCount": 5,
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
| 404 | 分类不存在 | 分类不存在或无权限 |

---

## 3. 新增分类

| 项目 | 说明 |
|------|------|
| **路径** | `POST /api/v1/manager/categories` |
| **鉴权** | 是 |
| **说明** | 创建新分类 |

### 请求参数（Body JSON）

| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| name | string | 是 | 分类名称，1-20 字 |
| parentId | number | 否 | 父分类 ID，0 表示一级分类，默认 0 |
| sort | number | 否 | 排序值，越大越靠前，默认 0 |
| status | number | 否 | 0 禁用 1 启用，默认 1 |

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
| 400 | 参数错误：name 必填 | 必填缺失 |
| 401 | 未登录或 token 已失效 | 鉴权失败 |
| 1105 | 父分类不存在 | parentId 无效 |
| 1106 | 同级下分类名称不能重复 | 重名 |

---

## 4. 更新分类

| 项目 | 说明 |
|------|------|
| **路径** | `PUT /api/v1/manager/categories/:id` |
| **鉴权** | 是 |
| **说明** | 修改分类信息 |

### 请求参数

**Path**：`id`（分类 ID）

**Body JSON**（均为可选）：

| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| name | string | 否 | 分类名称 |
| parentId | number | 否 | 父分类 ID |
| sort | number | 否 | 排序值 |
| status | number | 否 | 0 禁用 1 启用 |

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
| 404 | 分类不存在 | 分类不存在或无权限 |
| 1107 | 不能将父分类设为自己或子级 | 循环引用 |
| 1106 | 同级下分类名称不能重复 | 重名 |

---

## 5. 删除分类

| 项目 | 说明 |
|------|------|
| **路径** | `DELETE /api/v1/manager/categories/:id` |
| **鉴权** | 是 |
| **说明** | 软删除分类，若有子分类或关联商品需先处理 |

### 请求参数

| 参数 | 位置 | 类型 | 必填 | 说明 |
|------|------|------|------|------|
| id | path | number | 是 | 分类 ID |

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
| 404 | 分类不存在 | 分类不存在 |
| 1108 | 该分类下存在子分类，无法删除 | 需先删除子分类 |
| 1109 | 该分类下存在商品，无法删除 | 需先解除商品关联或删除商品 |
