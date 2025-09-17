# Demo-Gin: TDD-Driven API Framework

一个基于 **测试驱动开发(TDD)** 理念构建的 Go API 框架示例，专为 **AI 自动化开发** 场景设计。

## 🎯 项目目标

本项目旨在创建一个标准化的、可被 AI 工具理解和扩展的后端服务框架，具备以下特点：

- **TDD 优先**：所有功能开发遵循"先写测试，再写代码"的原则
- **AI 友好**：清晰的代码结构和注释，便于 AI 理解和自动生成代码
- **类型安全**：使用 sqlc 生成类型安全的数据库代码
- **API 规范**：基于 OpenAPI 3.0 的契约优先开发
- **自动化测试**：完整的单元测试、集成测试和 E2E 测试体系

## 🛠 技术栈

- **Gin** - 高性能 Web 框架
- **PostgreSQL** - 主数据库
- **sqlc** - 类型安全的 SQL 代码生成
- **OpenAPI 3.0** - API 规范定义
- **testify** - 测试断言库
- **gomock** - Mock 框架
- **Docker** - 容器化部署

## 📁 项目结构

```
demo-gin/
├── api/                    # API 契约定义
│   └── openapi.yaml       # OpenAPI 3.0 规范（API 优先设计）
├── cmd/                   # 应用入口
│   └── server/
│       └── main.go        # 主程序入口
├── internal/              # 内部应用代码
│   ├── config/           # 配置管理
│   ├── db/               # 数据库相关
│   │   ├── queries/      # SQL 查询文件（sqlc 源文件）
│   │   └── sqlc/         # sqlc 生成的类型安全代码
│   ├── handlers/         # HTTP 请求处理器
│   ├── middleware/       # HTTP 中间件
│   ├── models/           # 领域模型
│   ├── services/         # 业务逻辑层
│   └── utils/            # 工具函数
├── tests/                 # 测试套件（TDD 核心）
│   ├── unit/            # 单元测试
│   ├── integration/     # 集成测试
│   ├── e2e/            # 端到端测试
│   └── fixtures/        # 测试数据
├── migrations/            # 数据库迁移脚本
├── pkg/                   # 公共包
│   ├── database/         # 数据库连接
│   └── logger/           # 日志工具
├── docs/                  # 生成的 API 文档
├── docker/               # Docker 相关配置
│   └── docker-compose.yml # 本地开发环境
├── .github/              # GitHub Actions CI/CD
├── CLAUDE.md             # AI 开发指南
├── Makefile              # 自动化脚本
├── sqlc.yaml             # sqlc 配置
└── README.md             # 项目说明

```

## 🚀 TDD 开发流程

### 1. 编写测试优先
```bash
# 创建测试文件
make test-new feature=user_profile

# 运行测试（会失败）
make test

# 实现功能直到测试通过
make watch-test  # 自动监控测试
```

### 2. AI 辅助开发
本项目针对 AI 工具（如 Claude、GitHub Copilot）优化：
- 清晰的函数签名和接口定义
- 完整的测试用例作为行为规范
- 标准化的错误处理模式
- CLAUDE.md 文件提供 AI 上下文

### 3. 测试金字塔
```
        /\      E2E 测试 (10%)
       /  \     - 完整业务流程
      /    \    - Docker 环境
     /      \
    /--------\  集成测试 (30%)
   /          \ - API 端点测试
  /            \- 数据库交互
 /              \
/________________\ 单元测试 (60%)
                   - 业务逻辑
                   - 工具函数
```

## 📋 环境要求

- Go 1.22+
- PostgreSQL 14+
- Docker & Docker Compose
- [sqlc](https://sqlc.dev/)
- [golang-migrate](https://github.com/golang-migrate/migrate)
- [swag](https://github.com/swaggo/swag) (Swagger 文档生成)

## ⚡ 快速开始

### 使用 Docker（推荐）
```bash
# 1. 克隆项目
git clone <repository-url>
cd demo-gin

# 2. 启动所有服务（数据库、API、测试环境）
make docker-up

# 3. 运行测试套件
make test-all
```

### 本地开发
```bash
# 1. 安装依赖
go mod download

# 2. 配置环境变量
cp .env.example .env
# 编辑 .env 文件

# 3. 启动 PostgreSQL（使用 Docker）
docker-compose up -d postgres

# 4. 初始化数据库
make db-setup   # 创建数据库、运行迁移、生成代码

# 5. 运行测试
make test

# 6. 启动服务
make run
```

## 🧪 测试驱动开发

### 运行测试
```bash
# 所有测试
make test-all

# 单元测试
make test-unit

# 集成测试
make test-integration

# E2E 测试
make test-e2e

# 测试覆盖率报告
make test-coverage

# 持续测试（文件变更自动运行）
make watch-test
```

### 开发命令
```bash
# 启动开发服务器（热重载）
make dev

# 构建生产版本
make build

# 代码检查
make lint

# 格式化代码
make fmt
```

## 📖 API 文档

- **OpenAPI Specification**: `api/openapi.yaml`
- **Swagger UI**: After running the server, visit `http://localhost:8080/swagger/index.html`

## 🔌 API 端点

### Authentication
- `POST /api/v1/auth/register` - Register a new user
- `POST /api/v1/auth/login` - Login user

### Users (Protected)
- `GET /api/v1/users` - List users
- `GET /api/v1/users/:id` - Get user by ID
- `PUT /api/v1/users/:id` - Update user
- `DELETE /api/v1/users/:id` - Delete user

### Posts
- `GET /api/v1/posts` - List posts (public)
- `GET /api/v1/posts/:id` - Get post by ID (public)
- `POST /api/v1/posts` - Create post (protected)
- `PUT /api/v1/posts/:id` - Update post (protected)
- `DELETE /api/v1/posts/:id` - Delete post (protected)

### Health
- `GET /api/v1/health` - Health check

## 💾 数据库架构

The application includes two main tables:
- **users**: User accounts with authentication
- **posts**: Content posts linked to users

See `migrations/000001_init_schema.up.sql` for the complete schema.

## 🔧 Makefile 命令

### 基础命令
```bash
make help         # 显示所有可用命令
make run          # 运行应用
make dev          # 开发模式（热重载）
make build        # 构建二进制文件
make clean        # 清理构建产物
```

### 测试命令
```bash
make test         # 运行所有测试
make test-unit    # 仅单元测试
make test-integration # 仅集成测试
make test-e2e     # 仅 E2E 测试
make test-coverage # 生成覆盖率报告
make watch-test   # 监控模式测试
```

### 数据库命令
```bash
make db-setup     # 初始化数据库
make migrate-up   # 执行迁移
make migrate-down # 回滚迁移
make sqlc         # 生成 sqlc 代码
make db-seed      # 填充测试数据
```

### Docker 命令
```bash
make docker-up    # 启动所有容器
make docker-down  # 停止所有容器
make docker-build # 构建镜像
make docker-test  # 容器中运行测试
```

## 🎓 AI 自动化开发指南

### 为 AI 工具准备的特性

1. **CLAUDE.md 文件**
   - 项目上下文和规范
   - AI 可读的开发指南
   - 代码生成模板

2. **标准化的测试模式**
   ```go
   // 测试文件命名：*_test.go
   // 测试函数命名：Test<功能>_<场景>_<预期结果>
   func TestCreateUser_ValidInput_Success(t *testing.T) {...}
   func TestCreateUser_DuplicateEmail_ReturnsError(t *testing.T) {...}
   ```

3. **清晰的接口定义**
   - OpenAPI 优先设计
   - 类型安全的 sqlc
   - 明确的错误类型

### 示例：使用 AI 添加新功能

```bash
# 1. 定义需求（给 AI 的提示）
"基于现有的用户模块，添加一个用户头像上传功能，
要求：
- 先写测试
- 支持 JPG/PNG
- 最大 5MB
- 存储到 S3"

# 2. AI 生成测试
make test  # 运行失败的测试

# 3. AI 实现功能
make test  # 测试通过

# 4. AI 生成文档
make swagger  # 更新 API 文档
```

## 🚧 开发路线图

### Phase 1: 基础框架 ✅
- [x] 项目结构搭建
- [x] 基本 CRUD 示例
- [x] 数据库集成
- [x] API 文档生成

### Phase 2: TDD 体系 🚧
- [ ] 完整的测试套件
- [ ] Mock 和 Stub 框架
- [ ] 测试数据工厂
- [ ] 性能基准测试

### Phase 3: 认证与安全
- [ ] JWT 认证实现
- [ ] 密码加密 (bcrypt)
- [ ] 权限管理 (RBAC)
- [ ] API 限流

### Phase 4: 生产就绪
- [ ] 结构化日志
- [ ] 分布式追踪
- [ ] 健康检查
- [ ] 优雅关闭
- [ ] 配置热重载

### Phase 5: AI 增强
- [ ] AI 代码生成模板
- [ ] 自动化测试生成
- [ ] AI 驱动的代码审查
- [ ] 智能错误诊断

## 📝 贡献指南

欢迎贡献！请遵循以下原则：

1. **TDD 优先**：先写测试，再写代码
2. **保持简单**：代码应该易于 AI 理解
3. **文档完善**：每个功能都要有清晰的文档
4. **遵循规范**：使用项目的代码风格和命名约定

## 📄 许可证

MIT License - 详见 [LICENSE](LICENSE) 文件

## 🤝 联系方式

- Issues: [GitHub Issues](https://github.com/yourusername/demo-gin/issues)
- Discussions: [GitHub Discussions](https://github.com/yourusername/demo-gin/discussions)