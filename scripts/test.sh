#!/bin/bash

BASE_URL="http://115.190.178.106:9527"
TOKEN=""

# 登录获取 token
login() {
  echo "=== 登录 ==="
  response=$(curl -s -X POST "$BASE_URL/api/v1/manager/auth/login" \
    -H "Content-Type: application/json" \
    -d '{"code": "test-code", "shopName": "测试店铺"}')
  echo "$response" | jq .
  TOKEN=$(echo "$response" | jq -r '.data.token')
  echo "Token: $TOKEN"
  echo ""
}

# 获取商家信息
get_profile() {
  echo "=== 获取商家信息 ==="
  curl -s -X GET "$BASE_URL/api/v1/manager/auth/profile" \
    -H "Authorization: Bearer $TOKEN" | jq .
  echo ""
}

# 商品列表
list_products() {
  echo "=== 商品列表 ==="
  curl -s -X GET "$BASE_URL/api/v1/manager/products?page=1&pageSize=10" \
    -H "Authorization: Bearer $TOKEN" | jq .
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
    }' | jq .
  echo ""
}

# 分类列表
list_categories() {
  echo "=== 分类列表 ==="
  curl -s -X GET "$BASE_URL/api/v1/manager/categories" \
    -H "Authorization: Bearer $TOKEN" | jq .
  echo ""
}

# 创建分类
create_category() {
  echo "=== 创建分类 ==="
  curl -s -X POST "$BASE_URL/api/v1/manager/categories" \
    -H "Content-Type: application/json" \
    -H "Authorization: Bearer $TOKEN" \
    -d '{"name": "水果", "parentId": 0, "sort": 10}' | jq .
  echo ""
}

# 库存列表
list_inventory() {
  echo "=== 库存列表 ==="
  curl -s -X GET "$BASE_URL/api/v1/manager/inventory?page=1&pageSize=10" \
    -H "Authorization: Bearer $TOKEN" | jq .
  echo ""
}

# 订单列表
list_orders() {
  echo "=== 订单列表 ==="
  curl -s -X GET "$BASE_URL/api/v1/manager/orders?page=1&pageSize=10" \
    -H "Authorization: Bearer $TOKEN" | jq .
  echo ""
}

# 数据概览
dashboard_overview() {
  echo "=== 数据概览 ==="
  curl -s -X GET "$BASE_URL/api/v1/manager/dashboard/overview?period=today" \
    -H "Authorization: Bearer $TOKEN" | jq .
  echo ""
}

# 待办事项
dashboard_pending() {
  echo "=== 待办事项 ==="
  curl -s -X GET "$BASE_URL/api/v1/manager/dashboard/pending" \
    -H "Authorization: Bearer $TOKEN" | jq .
  echo ""
}

# 执行所有测试
echo "========================================="
echo "Server ZZQ API 测试脚本"
echo "========================================="
echo ""

login
get_profile
list_categories
create_category
list_products
create_product
list_inventory
list_orders
dashboard_overview
dashboard_pending

echo "========================================="
echo "测试完成!"
echo "========================================="
