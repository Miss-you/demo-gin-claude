# Claude 开发指南

本文档为 Claude 提供项目上下文和开发规范。

## 项目概述

基于以下技术栈的 RESTful API 服务：
- **Gin** - Web 框架
- **PostgreSQL** - 数据库
- **sqlc** - 类型安全的 SQL 代码生成
- **OpenAPI 3.0** - API 规范

## 核心目录结构

```
demo-gin/
├── cmd/server/main.go     # 应用入口
├── internal/
│   ├── config/            # 配置管理
│   ├── db/
│   │   ├── queries/       # SQL 查询文件（sqlc 源文件）
│   │   └── sqlc/          # sqlc 生成的代码（不要手动修改）
│   ├── handlers/          # HTTP 处理器
│   └── middleware/        # 中间件
├── migrations/            # 数据库迁移文件
├── api/openapi.yaml      # OpenAPI 规范
└── sqlc.yaml             # sqlc 配置
```

## 开发工作流

### 数据库变更

1. **创建迁移文件**：
   ```bash
   migrate create -ext sql -dir migrations -seq <迁移名称>
   ```

2. **编写 SQL 查询**（`internal/db/queries/`）：
   - 使用 sqlc 注释格式：`-- name: GetUser :one`

3. **生成代码**：
   ```bash
   make sqlc
   ```

### 添加新接口

1. 先更新 `api/openapi.yaml`
2. 在 `internal/handlers/` 创建处理器
3. 在 `cmd/server/main.go` 注册路由
4. 生成 Swagger 文档：`make swagger`

## 代码规范

### 基本原则
- 使用结构化错误和正确的 HTTP 状态码
- 所有输入必须验证
- 遵循 RESTful 规范
- 不要提交敏感信息

### 数据库操作
- 所有查询必须通过 sqlc
- 查询文件放在 `internal/db/queries/*.sql`
- 正确处理 `sql.ErrNoRows`

### API 处理器
- 使用 Gin 的 binding 标签验证输入
- 返回一致的错误响应格式
- 列表接口必须支持分页

## 常用命令

```bash
# 开发
make run          # 启动服务
make test         # 运行测试

# 数据库
make migrate-up   # 执行迁移
make migrate-down # 回滚迁移
make sqlc         # 生成 sqlc 代码

# 构建
make build        # 构建二进制文件
make swagger      # 生成 Swagger 文档
```

## 环境变量

必需的环境变量（见 `.env.example`）：
- `DB_HOST` - 数据库主机
- `DB_PORT` - 数据库端口
- `DB_USER` - 数据库用户
- `DB_PASSWORD` - 数据库密码
- `DB_NAME` - 数据库名称
- `SERVER_PORT` - 服务端口

## 重要文件

开发前请先查看：
1. `api/openapi.yaml` - API 契约
2. `migrations/*.sql` - 数据库结构
3. `internal/db/queries/*.sql` - 已有查询
4. `internal/handlers/*.go` - 处理器模式

## 待完成事项

- [ ] 实现 JWT 认证
- [ ] 添加密码加密（bcrypt）
- [ ] 完成所有处理器的 sqlc 集成
- [ ] 添加结构化日志
- [ ] 实现优雅关闭