#!/bin/bash

# HTTPS POST 请求测试脚本
# 使用自签名证书需要添加 --insecure 参数或使用 -k 参数

# 支持自定义 URL，默认本地
BASE_URL="${1:-https://localhost:8080}"
API_PREFIX="/api/v1/manager"

# 颜色输出
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

echo "=========================================="
echo "HTTPS POST API 测试脚本"
echo "=========================================="
echo "基础 URL: ${BASE_URL}"
echo "=========================================="
echo ""

# 测试登录接口
# 注意：登录需要微信授权 code 参数，这里使用模拟的 code 进行测试
echo -e "${YELLOW}[测试 1] POST 登录接口${NC}"
echo "URL: ${BASE_URL}${API_PREFIX}/auth/login"
LOGIN_RESPONSE=$(curl -k -s -X POST "${BASE_URL}${API_PREFIX}/auth/login" \
    -H "Content-Type: application/json" \
    -d '{
        "code": "mock-auth-code-12345",
        "shopName": "测试店铺"
    }')
echo "响应：${LOGIN_RESPONSE}"
echo ""

# 从登录响应中提取 token（支持有空格和无空格格式）
# 尝试多种模式匹配
TOKEN=$(echo "$LOGIN_RESPONSE" | grep -oE '"token"\s*:\s*"[^"]*"' | head -1 | sed 's/.*"\([^"]*\)"$/\1/')
if [ -z "$TOKEN" ]; then
    TOKEN="your-token-here"
    echo -e "${RED}未获取到 token，使用默认值${NC}"
else
    echo -e "${GREEN}获取到 token: ${TOKEN:0:20}...${NC}"
fi
echo ""

# 测试获取个人信息（需要 token）
echo -e "${YELLOW}[测试 2] POST 获取个人信息 (需要 token)${NC}"
echo "URL: ${BASE_URL}${API_PREFIX}/auth/profile"
curl -k -s -X POST "${BASE_URL}${API_PREFIX}/auth/profile" \
    -H "Content-Type: application/json" \
    -H "Authorization: Bearer ${TOKEN}"
echo ""
echo ""

# 测试获取商品列表
# 注意：使用 URL 参数传递分页参数（后端使用 form 标签验证）
echo -e "${YELLOW}[测试 3] POST 获取商品列表 (需要 token)${NC}"
echo "URL: ${BASE_URL}${API_PREFIX}/products?page=1&pageSize=10"
curl -k -s -X POST "${BASE_URL}${API_PREFIX}/products?page=1&pageSize=10" \
    -H "Content-Type: application/json" \
    -H "Authorization: Bearer ${TOKEN}"
echo ""
echo ""

# 测试获取分类列表
echo -e "${YELLOW}[测试 4] POST 获取分类列表 (需要 token)${NC}"
echo "URL: ${BASE_URL}${API_PREFIX}/categories"
curl -k -s -X POST "${BASE_URL}${API_PREFIX}/categories" \
    -H "Content-Type: application/json" \
    -H "Authorization: Bearer ${TOKEN}"
echo ""
echo ""

# 测试获取订单列表
# 注意：使用 URL 参数传递分页参数（后端使用 form 标签验证）
echo -e "${YELLOW}[测试 5] POST 获取订单列表 (需要 token)${NC}"
echo "URL: ${BASE_URL}${API_PREFIX}/orders?page=1&pageSize=10"
curl -k -s -X POST "${BASE_URL}${API_PREFIX}/orders?page=1&pageSize=10" \
    -H "Content-Type: application/json" \
    -H "Authorization: Bearer ${TOKEN}"
echo ""
echo ""

# 测试数据看板
echo -e "${YELLOW}[测试 6] POST 获取数据看板 (需要 token)${NC}"
echo "URL: ${BASE_URL}${API_PREFIX}/dashboard/overview"
curl -k -s -X POST "${BASE_URL}${API_PREFIX}/dashboard/overview" \
    -H "Content-Type: application/json" \
    -H "Authorization: Bearer ${TOKEN}"
echo ""
echo ""

echo -e "${GREEN}=========================================="
echo "测试完成"
echo "==========================================${NC}"
echo ""
echo "提示："
echo "1. 使用 -k 或 --insecure 参数忽略自签名证书验证"
echo "2. 将 TOKEN 替换为实际的 JWT token 进行测试"
echo "3. 生产环境请使用正式 SSL 证书"
echo ""
echo "用法："
echo "  ./test_https.sh                          # 测试本地 (localhost:8080)"
echo "  ./test_https.sh https://your-server:8080 # 测试远程服务器"
