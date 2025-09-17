# 测试目录结构

## 目录说明

```
tests/
├── integration/     # 集成测试 - 包含所有API端点的集成测试
├── fixtures/        # 测试数据 - 测试数据生成器和固定测试数据
├── helpers/         # 测试工具 - HTTP客户端、数据库工具、JWT工具等
└── config/          # 测试配置 - 测试环境配置加载器
```

## 使用说明

### 运行所有集成测试
```bash
go test -v ./tests/integration/...
```

### 运行单个测试文件
```bash
go test -v ./tests/integration/health_test.go
```

### 运行测试并生成覆盖率报告
```bash
go test -v -coverprofile=coverage.out ./...
go tool cover -html=coverage.out -o coverage.html
```

## 测试环境

测试使用独立的测试数据库，配置在 `.env.test` 文件中。
