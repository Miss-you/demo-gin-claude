# Demo Gin API 接口文档

> 📖 基于 OpenAPI 3.0 规范
> 🌐 基础URL: http://localhost:8080/api/v1
> 🔒 认证方式: Bearer JWT Token

---

## 📚 接口概览

| 模块 | 接口数量 | 公开接口 | 认证接口 |
|------|----------|----------|----------|
| 认证模块 | 2 | 2 | 0 |
| 用户管理 | 4 | 0 | 4 |
| 文章管理 | 5 | 2 | 3 |
| 系统功能 | 1 | 1 | 0 |
| **总计** | **12** | **5** | **7** |

---

## 🔐 认证模块 (Authentication)

### 用户注册
```http
POST /api/v1/auth/register
Content-Type: application/json

{
  "email": "user@example.com",
  "username": "johndoe",
  "password": "password123",
  "full_name": "John Doe"
}
```

**响应示例:**
```json
{
  "message": "User registered successfully",
  "user": {
    "email": "user@example.com",
    "username": "johndoe"
  }
}
```

**状态码:**
- `201 Created` - 注册成功
- `400 Bad Request` - 请求参数错误
- `409 Conflict` - 邮箱或用户名已存在

### 用户登录
```http
POST /api/v1/auth/login
Content-Type: application/json

{
  "username": "johndoe",
  "password": "password123"
}
```

**响应示例:**
```json
{
  "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "token_type": "Bearer",
  "expires_in": 3600,
  "user": {
    "id": 1,
    "email": "user@example.com",
    "username": "johndoe",
    "full_name": "John Doe",
    "is_active": true,
    "created_at": "2024-01-01T00:00:00Z"
  }
}
```

**状态码:**
- `200 OK` - 登录成功
- `401 Unauthorized` - 凭证无效

---

## 👥 用户管理模块 (User Management)

> 🔒 所有用户管理接口都需要JWT认证

### 获取用户列表
```http
GET /api/v1/users?page=1&limit=10
Authorization: Bearer {jwt_token}
```

**查询参数:**
- `page` (int): 页码，默认1
- `limit` (int): 每页条数，范围1-100，默认10

**响应示例:**
```json
{
  "users": [
    {
      "id": 1,
      "email": "user1@example.com",
      "username": "user1",
      "full_name": "User One",
      "is_active": true,
      "created_at": "2024-01-01T00:00:00Z",
      "updated_at": "2024-01-01T00:00:00Z"
    }
  ],
  "pagination": {
    "page": 1,
    "limit": 10,
    "total": 50,
    "total_pages": 5
  }
}
```

### 获取用户详情
```http
GET /api/v1/users/{id}
Authorization: Bearer {jwt_token}
```

**响应示例:**
```json
{
  "data": {
    "id": 1,
    "email": "user@example.com",
    "username": "johndoe",
    "full_name": "John Doe",
    "is_active": true,
    "created_at": "2024-01-01T00:00:00Z",
    "updated_at": "2024-01-01T00:00:00Z"
  }
}
```

### 更新用户信息
```http
PUT /api/v1/users/{id}
Authorization: Bearer {jwt_token}
Content-Type: application/json

{
  "email": "newemail@example.com",
  "full_name": "New Full Name"
}
```

**响应示例:**
```json
{
  "message": "User updated successfully",
  "data": {
    "id": 1
  }
}
```

### 删除用户
```http
DELETE /api/v1/users/{id}
Authorization: Bearer {jwt_token}
```

**响应:**
- `204 No Content` - 删除成功，无响应体

---

## 📝 文章管理模块 (Post Management)

### 获取文章列表 (公开)
```http
GET /api/v1/posts?page=1&limit=10
```

**查询参数:**
- `page` (int): 页码，默认1
- `limit` (int): 每页条数，范围1-100，默认10

**响应示例:**
```json
{
  "posts": [
    {
      "id": 1,
      "user_id": 1,
      "title": "First Post",
      "content": "This is the first post content...",
      "status": "published",
      "published_at": "2024-01-01T00:00:00Z",
      "created_at": "2024-01-01T00:00:00Z",
      "updated_at": "2024-01-01T00:00:00Z"
    }
  ],
  "pagination": {
    "page": 1,
    "limit": 10,
    "total": 25,
    "total_pages": 3
  }
}
```

### 获取文章详情 (公开)
```http
GET /api/v1/posts/{id}
```

**响应示例:**
```json
{
  "data": {
    "id": 1,
    "user_id": 1,
    "title": "Sample Post",
    "content": "This is the complete post content...",
    "status": "published",
    "published_at": "2024-01-01T00:00:00Z",
    "created_at": "2024-01-01T00:00:00Z",
    "updated_at": "2024-01-01T00:00:00Z"
  }
}
```

### 创建文章 (需认证)
```http
POST /api/v1/posts
Authorization: Bearer {jwt_token}
Content-Type: application/json

{
  "title": "My New Post",
  "content": "This is the content of my new post...",
  "status": "draft"
}
```

**请求字段说明:**
- `title` (string): 标题，必填，长度1-255字符
- `content` (string): 内容，必填
- `status` (string): 状态，可选值: draft/published，默认draft

**响应示例:**
```json
{
  "message": "Post created successfully",
  "data": {
    "id": 1,
    "user_id": 1,
    "title": "My New Post",
    "content": "This is the content of my new post...",
    "status": "draft",
    "created_at": "2024-01-01T00:00:00Z"
  }
}
```

### 更新文章 (需认证)
```http
PUT /api/v1/posts/{id}
Authorization: Bearer {jwt_token}
Content-Type: application/json

{
  "title": "Updated Post Title",
  "content": "Updated post content...",
  "status": "published"
}
```

**请求字段说明:**
- 所有字段都是可选的
- `status` 可选值: draft/published/archived

**响应示例:**
```json
{
  "message": "Post updated successfully",
  "data": {
    "id": 1
  }
}
```

### 删除文章 (需认证)
```http
DELETE /api/v1/posts/{id}
Authorization: Bearer {jwt_token}
```

**响应:**
- `204 No Content` - 删除成功，无响应体

---

## 🏥 系统功能

### 健康检查
```http
GET /api/v1/health
```

**响应示例:**
```json
{
  "status": "ok",
  "timestamp": "2024-01-01T00:00:00Z",
  "version": "1.0.0"
}
```

---

## 🔒 认证机制

### JWT Token 使用方式
```http
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
```

### Token 信息
- **类型**: Bearer Token
- **过期时间**: 3600秒 (1小时)
- **刷新机制**: 需重新登录

---

## 📋 通用响应格式

### 成功响应结构
```json
{
  "data": {},          // 数据对象(单个资源)
  "message": "string"   // 操作消息
}
```

```json
{
  "users": [],         // 数据数组(多个资源)
  "pagination": {}     // 分页信息
}
```

### 错误响应结构
```json
{
  "error": "error_code",
  "message": "Human readable error message",
  "details": {}        // 额外错误详情
}
```

### HTTP 状态码说明
- `200 OK` - 请求成功
- `201 Created` - 资源创建成功
- `204 No Content` - 操作成功，无返回内容
- `400 Bad Request` - 请求参数错误
- `401 Unauthorized` - 未认证或认证失败
- `404 Not Found` - 资源不存在
- `409 Conflict` - 资源冲突（如重复创建）
- `500 Internal Server Error` - 服务器内部错误

---

## 📝 分页参数说明

### 查询参数
- `page`: 页码，从1开始
- `limit`: 每页条数，范围1-100

### 分页响应
```json
{
  "pagination": {
    "page": 1,           // 当前页码
    "limit": 10,         // 每页条数
    "total": 100,        // 总记录数
    "total_pages": 10    // 总页数
  }
}
```

---

## 🛠️ 接口测试示例

### 使用 curl 测试

#### 1. 用户注册
```bash
curl -X POST http://localhost:8080/api/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "email": "test@example.com",
    "username": "testuser",
    "password": "password123",
    "full_name": "Test User"
  }'
```

#### 2. 用户登录
```bash
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "username": "testuser",
    "password": "password123"
  }'
```

#### 3. 创建文章 (需要先从登录响应中获取token)
```bash
curl -X POST http://localhost:8080/api/v1/posts \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -d '{
    "title": "My First Post",
    "content": "This is my first post content",
    "status": "draft"
  }'
```

#### 4. 获取文章列表
```bash
curl -X GET "http://localhost:8080/api/v1/posts?page=1&limit=5"
```

---

*本文档基于 OpenAPI 3.0 规范生成，随项目开发进度持续更新。*