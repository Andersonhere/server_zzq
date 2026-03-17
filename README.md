# Server ZZQ - 商家后台管理系统

这是一个基于 Go 语言开发的商家后台管理系统 API 服务，服务于生鲜电商/零售场景。

## 功能模块

- **认证模块** - 微信小程序登录、Token 管理、商家信息
- **商品管理** - 商品 CRUD、上下架、打折/预售、轮播推荐
- **分类管理** - 树形分类结构
- **库存管理** - 库存查询、调整、低库存预警
- **订单管理** - 订单列表/详情、发货、售后处理
- **数据看板** - 销售统计、趋势分析、待办事项

## 技术栈

- **框架**: Gin
- **ORM**: GORM
- **数据库**: MySQL 8.0+
- **缓存**: Redis (可选)
- **认证**: JWT
- **配置**: Viper
- **日志**: Zap

## 项目结构

```
server_zzq/
├── cmd/
│   └── server/
│       └── main.go          # 程序入口
├── internal/
│   ├── config/              # 配置加载
│   ├── handlers/            # HTTP 处理器
│   ├── models/              # 数据模型
│   ├── middleware/          # 中间件
│   ├── dto/                 # 数据传输对象
│   │   ├── request/
│   │   └── response/
│   └── utils/               # 工具函数
├── configs/
│   └── config.yaml          # 配置文件
├── migrations/              # 数据库迁移
└── Makefile
```

## 快速开始

### 环境要求

- Go 1.21+
- MySQL 8.0+
- Redis (可选)

### 安装依赖

```bash
go mod download
```

### 配置

编辑 `configs/config.yaml` 文件，配置数据库和微信等信息：

```yaml
server:
  port: 8080
  mode: debug

database:
  host: localhost
  port: 3306
  user: root
  password: your_password
  dbname: server_zzq

jwt:
  secret: your-secret-key

wechat:
  appid: your-wechat-appid
  secret: your-wechat-secret
```

### 运行

```bash
# 开发模式
go run cmd/server/main.go

# 编译后运行
go build -o bin/server ./cmd/server
./bin/server
```

### 使用 Makefile

```bash
make run      # 运行开发服务器
make build    # 编译
make test     # 运行测试
make clean    # 清理
```

## API 文档

### 认证接口

| 方法 | 路径 | 说明 | 鉴权 |
|------|------|------|------|
| POST | /api/v1/manager/auth/login | 商家登录 | 否 |
| POST | /api/v1/manager/auth/refresh | 刷新 Token | 是 |
| POST | /api/v1/manager/auth/logout | 退出登录 | 是 |
| GET | /api/v1/manager/auth/profile | 获取商家信息 | 是 |

### 商品接口

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /api/v1/manager/products | 商品列表 |
| POST | /api/v1/manager/products | 创建商品 |
| PUT | /api/v1/manager/products/:id | 更新商品 |
| DELETE | /api/v1/manager/products/:id | 删除商品 |
| PUT | /api/v1/manager/products/:id/status | 修改上下架 |
| PUT | /api/v1/manager/products/:id/discount | 设置打折 |
| PUT | /api/v1/manager/products/:id/presale | 设置预售 |

### 分类接口

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /api/v1/manager/categories | 分类列表 |
| POST | /api/v1/manager/categories | 创建分类 |
| PUT | /api/v1/manager/categories/:id | 更新分类 |
| DELETE | /api/v1/manager/categories/:id | 删除分类 |

### 库存接口

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /api/v1/manager/inventory | 库存列表 |
| PUT | /api/v1/manager/inventory/:productId/adjust | 调整库存 |

### 订单接口

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /api/v1/manager/orders | 订单列表 |
| GET | /api/v1/manager/orders/:id | 订单详情 |
| POST | /api/v1/manager/orders/:id/ship | 订单发货 |

### 数据看板

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /api/v1/manager/dashboard/overview | 数据概览 |
| GET | /api/v1/manager/dashboard/compare | 多周期对比 |
| GET | /api/v1/manager/dashboard/trend | 销售趋势 |
| GET | /api/v1/manager/dashboard/pending | 待办事项 |

## 响应格式

```json
{
  "code": 0,
  "data": {},
  "message": "success",
  "timestamp": 1710662400000
}
```

### 错误码

| Code | 说明 |
|------|------|
| 0 | 成功 |
| 400 | 参数错误 |
| 401 | 未授权 |
| 404 | 未找到 |
| 500 | 服务器错误 |

## 开发计划

- [ ] 实现微信登录逻辑
- [ ] 实现商品管理业务逻辑
- [ ] 实现订单管理业务逻辑
- [ ] 添加 Redis 缓存支持
- [ ] 添加单元测试
- [ ] 添加 API 文档生成

## License

MIT
