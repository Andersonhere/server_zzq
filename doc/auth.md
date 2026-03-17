# 商家登录/认证 接口规格

**文档路径**：api-docs/manager-backend/auth.md  
**用途**：商家管理端登录、token 刷新、退出登录  
**适用端**：manager-miniprogram ↔ backend

---

## 1. 商家微信小程序登录

商家通过微信小程序 `wx.login` 获取 code，后端用 code 换 session_key 并生成 token。

| 项目 | 说明 |
|------|------|
| **路径** | `POST /api/v1/manager/auth/login` |
| **鉴权** | 否 |
| **说明** | 使用微信 code 换取商家 token |

### 请求参数（Body JSON）

| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| code | string | 是 | 微信 wx.login 返回的临时 code |
| shopName | string | 否 | 首次登录时的店铺名称（注册用） |

### 响应格式

**成功**：
```json
{
  "code": 0,
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "expireIn": 7200,
    "shop": {
      "id": 1,
      "shopName": "水果铺子",
      "status": 1
    }
  },
  "message": "success",
  "timestamp": 1710662400000
}
```

| 字段 | 类型 | 说明 |
|------|------|------|
| token | string | JWT token，后续请求需携带 |
| expireIn | number | token 有效期（秒） |
| shop.id | number | 商家/店铺 ID |
| shop.shopName | string | 店铺名称 |
| shop.status | number | 1 正常 0 禁用 |

### 错误码

| code | message | 说明 |
|------|---------|------|
| 400 | 参数错误：code 必填 | code 缺失 |
| 1400 | 微信登录失败 | 微信 code 无效或过期 |
| 1401 | 商家未注册，请先完善店铺信息 | 新商家需先注册 |
| 1402 | 商家已被禁用 | 店铺状态异常 |

---

## 2. 刷新 Token

token 即将过期时调用，获取新 token。

| 项目 | 说明 |
|------|------|
| **路径** | `POST /api/v1/manager/auth/refresh` |
| **鉴权** | 是（需携带当前 token 或 refresh token） |
| **说明** | 刷新 token 延长登录态 |

### 请求参数

无 body，依赖请求头 `Authorization: Bearer {token}`。

### 响应格式

**成功**：
```json
{
  "code": 0,
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "expireIn": 7200
  },
  "message": "success",
  "timestamp": 1710662400000
}
```

### 错误码

| code | message | 说明 |
|------|---------|------|
| 401 | 未登录或 token 已失效 | token 无效或过期 |
| 1403 | 刷新失败，请重新登录 | refresh 失败 |

---

## 3. 退出登录

商家主动退出，后端使当前 token 失效（如加入黑名单）。

| 项目 | 说明 |
|------|------|
| **路径** | `POST /api/v1/manager/auth/logout` |
| **鉴权** | 是 |
| **说明** | 退出登录，token 失效 |

### 请求参数

无 body。

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
| 401 | 未登录或 token 已失效 | 未携带有效 token |

---

## 4. 获取当前商家信息

用于页面初始化或权限校验时获取当前商家/店铺信息。

| 项目 | 说明 |
|------|------|
| **路径** | `GET /api/v1/manager/auth/profile` |
| **鉴权** | 是 |
| **说明** | 获取当前登录商家详情 |

### 请求参数

无。

### 响应格式

**成功**：
```json
{
  "code": 0,
  "data": {
    "id": 1,
    "shopName": "水果铺子",
    "logo": "https://xxx.com/logo.png",
    "contactPhone": "13800138000",
    "status": 1,
    "createdAt": "2025-01-01T00:00:00.000Z"
  },
  "message": "success",
  "timestamp": 1710662400000
}
```

### 错误码

| code | message | 说明 |
|------|---------|------|
| 401 | 未登录或 token 已失效 | token 无效 |
| 404 | 商家不存在 | 商家已被删除 |
