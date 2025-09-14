# API 测试 CURL 命令

## 基础信息
- **Base URL**: `http://localhost:8080/api/v1`
- **Content-Type**: `application/json`

## 1. 健康检查

### 检查服务状态
```bash
curl http://localhost:8080/api/v1/health
```

## 2. 认证相关

### 用户注册
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

### 用户登录
```bash
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "username": "testuser",
    "password": "password123"
  }'
```

## 3. 文章管理（公开接口）

### 获取文章列表
```bash
# 默认分页
curl http://localhost:8080/api/v1/posts

# 指定分页参数
curl "http://localhost:8080/api/v1/posts?page=1&limit=10"
```

### 获取单个文章
```bash
curl http://localhost:8080/api/v1/posts/1
```

## 4. 文章管理（需要认证）

### 创建文章
```bash
curl -X POST http://localhost:8080/api/v1/posts \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer mock_jwt_token" \
  -d '{
    "title": "My New Post",
    "content": "This is the content of my new post",
    "status": "draft"
  }'
```

### 更新文章
```bash
curl -X PUT http://localhost:8080/api/v1/posts/1 \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer mock_jwt_token" \
  -d '{
    "title": "Updated Post Title",
    "content": "Updated content",
    "status": "published"
  }'
```

### 删除文章
```bash
curl -X DELETE http://localhost:8080/api/v1/posts/1 \
  -H "Authorization: Bearer mock_jwt_token"
```

## 5. 用户管理（需要认证）

### 获取用户列表
```bash
# 默认分页
curl http://localhost:8080/api/v1/users \
  -H "Authorization: Bearer mock_jwt_token"

# 指定分页参数
curl "http://localhost:8080/api/v1/users?page=1&limit=10" \
  -H "Authorization: Bearer mock_jwt_token"
```

### 获取用户详情
```bash
curl http://localhost:8080/api/v1/users/1 \
  -H "Authorization: Bearer mock_jwt_token"
```

### 更新用户信息
```bash
curl -X PUT http://localhost:8080/api/v1/users/1 \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer mock_jwt_token" \
  -d '{
    "email": "newemail@example.com",
    "username": "newusername",
    "full_name": "New Name"
  }'
```

### 删除用户
```bash
curl -X DELETE http://localhost:8080/api/v1/users/1 \
  -H "Authorization: Bearer mock_jwt_token"
```

## 使用说明

### 格式化输出
如果想要格式化 JSON 输出，可以使用 `jq`：
```bash
curl http://localhost:8080/api/v1/posts | jq '.'
```

### 查看请求详情
添加 `-v` 参数查看详细的请求和响应头：
```bash
curl -v http://localhost:8080/api/v1/health
```

### 保存响应到文件
```bash
curl http://localhost:8080/api/v1/posts -o posts.json
```

### 查看响应头
```bash
curl -I http://localhost:8080/api/v1/health
```

## 注意事项

1. **认证令牌**: 当前使用的是模拟令牌 `mock_jwt_token`，实际环境中需要从登录接口获取真实的 JWT token
2. **数据持久化**: 当前版本没有连接数据库，所有数据都是模拟的，重启服务后数据会丢失
3. **错误处理**: 如果遇到 401 错误，请检查 Authorization header 是否正确
4. **CORS**: 服务器已配置 CORS，支持跨域请求

## 常见错误

### 401 Unauthorized
```json
{
  "error": "Authorization header required"
}
```
**解决方案**: 添加 `Authorization: Bearer mock_jwt_token` 头

### 400 Bad Request
```json
{
  "error": "Invalid request body"
}
```
**解决方案**: 检查 JSON 格式是否正确，必填字段是否完整

### 404 Not Found
```json
{
  "error": "Resource not found"
}
```
**解决方案**: 检查 URL 路径和资源 ID 是否正确