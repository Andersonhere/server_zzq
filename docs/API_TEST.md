# API 接口测试指南

## 使用工具

你可以使用以下工具测试 API：
- **curl** - 命令行工具
- **Postman** - 图形化 API 测试工具
- **Apifox** - 国产 API 调试工具
- **HTTPie** - 简化版 curl

## 接口列表

### 1. 商家登录

**请求**
```bash
POST http://localhost:8080/api/v1/manager/auth/login
Content-Type: application/json

{
  "code": "wechat-login-code",
  "shopName": "测试店铺"
}
```

**curl 示例**
```bash
curl -X POST http://localhost:8080/api/v1/manager/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "code": "test-code",
    "shopName": "测试店铺"
  }'
```

**响应**
```json
{
  "code": 0,
  "data": {
    "token": "mock-token",
    "expireIn": 7200,
    "shop": {
      "id": 1,
      "shopName": "测试店铺",
      "logo": "",
      "contactPhone": "",
      "status": 1
    }
  },
  "message": "success",
  "timestamp": 1710662400000
}
```

---

### 2. 获取商家信息（需要鉴权）

**请求**
```bash
GET http://localhost:8080/api/v1/manager/auth/profile
Authorization: Bearer {token}
```

**curl 示例**
```bash
curl -X GET http://localhost:8080/api/v1/manager/auth/profile \
  -H "Authorization: Bearer mock-token"
```

---

### 3. 商品列表（需要鉴权）

**请求**
```bash
GET http://localhost:8080/api/v1/manager/products?page=1&pageSize=10&keyword=苹果
Authorization: Bearer {token}
```

**curl 示例**
```bash
curl -X GET "http://localhost:8080/api/v1/manager/products?page=1&pageSize=10" \
  -H "Authorization: Bearer mock-token"
```

---

### 4. 创建商品（需要鉴权）

**请求**
```bash
POST http://localhost:8080/api/v1/manager/products
Content-Type: application/json
Authorization: Bearer {token}

{
  "name": "新鲜苹果",
  "mainImage": "https://example.com/apple.jpg",
  "images": ["https://example.com/apple1.jpg", "https://example.com/apple2.jpg"],
  "categoryId": 1,
  "price": 1500,
  "originalPrice": 2000,
  "stock": 100,
  "description": "新鲜采摘的苹果",
  "status": 1
}
```

**curl 示例**
```bash
curl -X POST http://localhost:8080/api/v1/manager/products \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer mock-token" \
  -d '{
    "name": "新鲜苹果",
    "mainImage": "https://example.com/apple.jpg",
    "categoryId": 1,
    "price": 1500,
    "stock": 100
  }'
```

---

### 5. 分类列表（需要鉴权）

**请求**
```bash
GET http://localhost:8080/api/v1/manager/categories?tree=0
Authorization: Bearer {token}
```

**curl 示例**
```bash
curl -X GET "http://localhost:8080/api/v1/manager/categories?tree=0" \
  -H "Authorization: Bearer mock-token"
```

---

### 6. 创建分类（需要鉴权）

**请求**
```bash
POST http://localhost:8080/api/v1/manager/categories
Content-Type: application/json
Authorization: Bearer {token}

{
  "name": "水果",
  "parentId": 0,
  "sort": 10,
  "status": 1
}
```

---

### 7. 库存列表（需要鉴权）

**请求**
```bash
GET http://localhost:8080/api/v1/manager/inventory?page=1&pageSize=10
Authorization: Bearer {token}
```

---

### 8. 库存调整（需要鉴权）

**请求**
```bash
PUT http://localhost:8080/api/v1/manager/inventory/1/adjust
Content-Type: application/json
Authorization: Bearer {token}

{
  "delta": 50,
  "reason": "补货"
}
```

---

### 9. 订单列表（需要鉴权）

**请求**
```bash
GET http://localhost:8080/api/v1/manager/orders?page=1&pageSize=10&status=1
Authorization: Bearer {token}
```

---

### 10. 订单发货（需要鉴权）

**请求**
```bash
POST http://localhost:8080/api/v1/manager/orders/1/ship
Content-Type: application/json
Authorization: Bearer {token}

{
  "expressCompany": "顺丰快递",
  "expressNo": "SF1234567890"
}
```

---

### 11. 数据概览（需要鉴权）

**请求**
```bash
GET http://localhost:8080/api/v1/manager/dashboard/overview?period=today
Authorization: Bearer {token}
```

**period 参数可选值**: today, yesterday, 7d, 30d

---

### 12. 待办事项（需要鉴权）

**请求**
```bash
GET http://localhost:8080/api/v1/manager/dashboard/pending
Authorization: Bearer {token}
```

## Postman 集合

导入以下 JSON 到 Postman：

```json
{
  "info": {
    "name": "Server ZZQ API",
    "schema": "https://schema.getpostman.com/json/collection/v2.1.0/collection.json"
  },
  "variable": [
    {
      "key": "baseUrl",
      "value": "http://localhost:8080"
    },
    {
      "key": "token",
      "value": ""
    }
  ],
  "item": [
    {
      "name": "Auth",
      "item": [
        {
          "name": "Login",
          "request": {
            "method": "POST",
            "header": [{"key": "Content-Type", "value": "application/json"}],
            "url": "{{baseUrl}}/api/v1/manager/auth/login",
            "body": {
              "mode": "raw",
              "raw": "{\n  \"code\": \"test-code\",\n  \"shopName\": \"测试店铺\"\n}"
            }
          }
        },
        {
          "name": "Get Profile",
          "request": {
            "method": "GET",
            "header": [{"key": "Authorization", "value": "Bearer {{token}}"}],
            "url": "{{baseUrl}}/api/v1/manager/auth/profile"
          }
        }
      ]
    },
    {
      "name": "Products",
      "item": [
        {
          "name": "List Products",
          "request": {
            "method": "GET",
            "header": [{"key": "Authorization", "value": "Bearer {{token}}"}],
            "url": "{{baseUrl}}/api/v1/manager/products?page=1&pageSize=10"
          }
        },
        {
          "name": "Create Product",
          "request": {
            "method": "POST",
            "header": [
              {"key": "Authorization", "value": "Bearer {{token}}"},
              {"key": "Content-Type", "value": "application/json"}
            ],
            "url": "{{baseUrl}}/api/v1/manager/products",
            "body": {
              "mode": "raw",
              "raw": "{\n  \"name\": \"新鲜苹果\",\n  \"mainImage\": \"https://example.com/apple.jpg\",\n  \"categoryId\": 1,\n  \"price\": 1500,\n  \"stock\": 100\n}"
            }
          }
        }
      ]
    },
    {
      "name": "Categories",
      "item": [
        {
          "name": "List Categories",
          "request": {
            "method": "GET",
            "header": [{"key": "Authorization", "value": "Bearer {{token}}"}],
            "url": "{{baseUrl}}/api/v1/manager/categories"
          }
        },
        {
          "name": "Create Category",
          "request": {
            "method": "POST",
            "header": [
              {"key": "Authorization", "value": "Bearer {{token}}"},
              {"key": "Content-Type", "value": "application/json"}
            ],
            "url": "{{baseUrl}}/api/v1/manager/categories",
            "body": {
              "mode": "raw",
              "raw": "{\n  \"name\": \"水果\",\n  \"parentId\": 0,\n  \"sort\": 10\n}"
            }
          }
        }
      ]
    },
    {
      "name": "Inventory",
      "item": [
        {
          "name": "List Inventory",
          "request": {
            "method": "GET",
            "header": [{"key": "Authorization", "value": "Bearer {{token}}"}],
            "url": "{{baseUrl}}/api/v1/manager/inventory?page=1&pageSize=10"
          }
        },
        {
          "name": "Adjust Inventory",
          "request": {
            "method": "PUT",
            "header": [
              {"key": "Authorization", "value": "Bearer {{token}}"},
              {"key": "Content-Type", "value": "application/json"}
            ],
            "url": "{{baseUrl}}/api/v1/manager/inventory/1/adjust",
            "body": {
              "mode": "raw",
              "raw": "{\n  \"delta\": 50,\n  \"reason\": \"补货\"\n}"
            }
          }
        }
      ]
    },
    {
      "name": "Orders",
      "item": [
        {
          "name": "List Orders",
          "request": {
            "method": "GET",
            "header": [{"key": "Authorization", "value": "Bearer {{token}}"}],
            "url": "{{baseUrl}}/api/v1/manager/orders?page=1&pageSize=10"
          }
        },
        {
          "name": "Ship Order",
          "request": {
            "method": "POST",
            "header": [
              {"key": "Authorization", "value": "Bearer {{token}}"},
              {"key": "Content-Type", "value": "application/json"}
            ],
            "url": "{{baseUrl}}/api/v1/manager/orders/1/ship",
            "body": {
              "mode": "raw",
              "raw": "{\n  \"expressCompany\": \"顺丰快递\",\n  \"expressNo\": \"SF1234567890\"\n}"
            }
          }
        }
      ]
    },
    {
      "name": "Dashboard",
      "item": [
        {
          "name": "Overview",
          "request": {
            "method": "GET",
            "header": [{"key": "Authorization", "value": "Bearer {{token}}"}],
            "url": "{{baseUrl}}/api/v1/manager/dashboard/overview?period=today"
          }
        },
        {
          "name": "Pending",
          "request": {
            "method": "GET",
            "header": [{"key": "Authorization", "value": "Bearer {{token}}"}],
            "url": "{{baseUrl}}/api/v1/manager/dashboard/pending"
          }
        }
      ]
    }
  ]
}
```

## 快速测试脚本

创建一个测试脚本 `scripts/test.sh`:

```bash
#!/bin/bash

BASE_URL="http://localhost:8080"
TOKEN=""

# 登录获取 token
login() {
  echo "=== 登录 ==="
  response=$(curl -s -X POST "$BASE_URL/api/v1/manager/auth/login" \
    -H "Content-Type: application/json" \
    -d '{"code": "test-code", "shopName": "测试店铺"}')
  echo "$response"
  TOKEN=$(echo "$response" | grep -o '"token":"[^"]*"' | cut -d'"' -f4)
  echo "Token: $TOKEN"
}

# 获取商家信息
get_profile() {
  echo "=== 获取商家信息 ==="
  curl -s -X GET "$BASE_URL/api/v1/manager/auth/profile" \
    -H "Authorization: Bearer $TOKEN"
  echo ""
}

# 商品列表
list_products() {
  echo "=== 商品列表 ==="
  curl -s -X GET "$BASE_URL/api/v1/manager/products?page=1&pageSize=10" \
    -H "Authorization: Bearer $TOKEN"
  echo ""
}

# 创建商品
create_product() {
  echo "=== 创建商品 ==="
  curl -s -X POST "$BASE_URL/api/v1/manager/products" \
    -H "Content-Type: application/json" \
    -H "Authorization: Bearer $TOKEN" \
    -d '{
      "name": "新鲜苹果",
      "mainImage": "https://example.com/apple.jpg",
      "categoryId": 1,
      "price": 1500,
      "stock": 100
    }'
  echo ""
}

# 分类列表
list_categories() {
  echo "=== 分类列表 ==="
  curl -s -X GET "$BASE_URL/api/v1/manager/categories" \
    -H "Authorization: Bearer $TOKEN"
  echo ""
}

# 创建分类
create_category() {
  echo "=== 创建分类 ==="
  curl -s -X POST "$BASE_URL/api/v1/manager/categories" \
    -H "Content-Type: application/json" \
    -H "Authorization: Bearer $TOKEN" \
    -d '{"name": "水果", "parentId": 0, "sort": 10}'
  echo ""
}

# 数据概览
dashboard_overview() {
  echo "=== 数据概览 ==="
  curl -s -X GET "$BASE_URL/api/v1/manager/dashboard/overview?period=today" \
    -H "Authorization: Bearer $TOKEN"
  echo ""
}

# 执行测试
login
get_profile
list_categories
create_category
list_products
create_product
dashboard_overview
```

运行测试：
```bash
chmod +x scripts/test.sh
./scripts/test.sh
