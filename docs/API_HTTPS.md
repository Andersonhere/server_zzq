# HTTPS POST API 文档

## 概述

本项目所有 API 端点现已改为通过 HTTPS 协议使用 POST 请求方法。这提供了更好的安全性和一致性。

## 配置说明

### HTTPS 配置

在 `configs/config.yaml` 中配置 HTTPS：

```yaml
server:
  port: 8080
  mode: debug
  https:
    enabled: true
    cert_file: certs/server.crt
    key_file: certs/server.key
```

### 生成自签名证书

```bash
# 运行证书生成脚本
bash scripts/generate_certs.sh
```

## API 端点列表

### 基础 URL
```
https://localhost:8080/api/v1/manager
```

### 认证相关

| 端点 | 方法 | 描述 | 鉴权 |
|------|------|------|------|
| `/auth/login` | POST | 用户登录 | 否 |
| `/auth/profile` | POST | 获取当前用户信息 | 是 |

### 商品管理

| 端点 | 方法 | 描述 | 鉴权 |
|------|------|------|------|
| `/products` | POST | 获取商品列表 | 是 |
| `/products/:id` | POST | 获取单个商品 | 是 |
| `/products` | POST | 创建商品 | 是 |
| `/products/:id` | POST | 更新商品 | 是 |
| `/products/:id/delete` | POST | 删除商品 | 是 |
| `/products/:id/status` | POST | 更新商品状态 | 是 |
| `/products/:id/discount` | POST | 设置折扣 | 是 |
| `/products/:id/presale` | POST | 设置预售 | 是 |
| `/products/discount/batch` | POST | 批量设置折扣 | 是 |
| `/products/presale/batch` | POST | 批量设置预售 | 是 |

### 分类管理

| 端点 | 方法 | 描述 | 鉴权 |
|------|------|------|------|
| `/categories` | POST | 获取分类列表 | 是 |
| `/categories/:id` | POST | 获取单个分类 | 是 |
| `/categories` | POST | 创建分类 | 是 |
| `/categories/:id` | POST | 更新分类 | 是 |
| `/categories/:id/delete` | POST | 删除分类 | 是 |

### 库存管理

| 端点 | 方法 | 描述 | 鉴权 |
|------|------|------|------|
| `/inventory` | POST | 获取库存列表 | 是 |
| `/inventory/:productId` | POST | 获取单个库存 | 是 |
| `/inventory/:productId/adjust` | POST | 调整库存 | 是 |
| `/inventory/batch-adjust` | POST | 批量调整库存 | 是 |

### 订单管理

| 端点 | 方法 | 描述 | 鉴权 |
|------|------|------|------|
| `/orders` | POST | 获取订单列表 | 是 |
| `/orders/:id` | POST | 获取单个订单 | 是 |
| `/orders/:id/ship` | POST | 订单发货 | 是 |

### 售后管理

| 端点 | 方法 | 描述 | 鉴权 |
|------|------|------|------|
| `/after-sales` | POST | 获取售后列表 | 是 |
| `/after-sales/:id/handle` | POST | 处理售后 | 是 |

### 数据看板

| 端点 | 方法 | 描述 | 鉴权 |
|------|------|------|------|
| `/dashboard/overview` | POST | 数据概览 | 是 |
| `/dashboard/compare` | POST | 数据对比 | 是 |
| `/dashboard/trend` | POST | 数据趋势 | 是 |
| `/dashboard/pending` | POST | 待处理事项 | 是 |

## 请求示例

### 登录请求

```bash
curl -k -X POST "https://localhost:8080/api/v1/manager/auth/login" \
    -H "Content-Type: application/json" \
    -d '{
        "username": "admin",
        "password": "admin123"
    }'
```

### 获取商品列表（需要 token）

```bash
curl -k -X POST "https://localhost:8080/api/v1/manager/products" \
    -H "Content-Type: application/json" \
    -H "Authorization: Bearer YOUR_TOKEN" \
    -d '{
        "page": 1,
        "page_size": 10
    }'
```

## 测试脚本

使用提供的测试脚本快速验证 API：

```bash
chmod +x scripts/test_https.sh
bash scripts/test_https.sh
```

## 注意事项

1. **自签名证书警告**：使用自签名证书时，浏览器和客户端会显示安全警告，这是正常现象。
   - 使用 curl 测试时添加 `-k` 或 `--insecure` 参数
   - 生产环境请使用正式的 SSL 证书

2. **请求方法统一**：所有 API 端点现在都使用 POST 方法，包括传统的 GET 操作（如列表查询、获取详情等）

3. **请求参数**：原本通过 URL 查询参数传递的数据，现在可以通过 POST body 传递

## 启动服务器

```bash
# 构建
go build -o bin/server ./cmd/server/main.go

# 运行
./bin/server
```

服务器启动后会显示：
```
HTTPS enabled, cert=certs/server.crt, key=certs/server.key
Server starting, addr= :8080, https= true
