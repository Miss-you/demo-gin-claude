package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

// TestConfig 测试配置结构体
type TestConfig struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBSSLMode  string

	ServerHost string
	ServerPort string

	JWTSecret      string
	JWTExpireHours string
}

// LoadTestConfig 加载测试配置
func LoadTestConfig() (*TestConfig, error) {
	// 获取项目根目录
	rootDir, err := FindProjectRoot()
	if err != nil {
		return nil, err
	}

	// 加载 .env.test 文件
	envPath := filepath.Join(rootDir, ".env.test")
	if err := godotenv.Load(envPath); err != nil {
		return nil, fmt.Errorf("failed to load .env.test: %w", err)
	}

	config := &TestConfig{
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5433"),
		DBUser:     getEnv("DB_USER", "test_user"),
		DBPassword: getEnv("DB_PASSWORD", "test_pass"),
		DBName:     getEnv("DB_NAME", "demo_gin_test"),
		DBSSLMode:  getEnv("DB_SSLMODE", "disable"),

		ServerHost: getEnv("SERVER_HOST", "localhost"),
		ServerPort: getEnv("SERVER_PORT", "8081"),

		JWTSecret:      getEnv("JWT_SECRET", "test_jwt_secret"),
		JWTExpireHours: getEnv("JWT_EXPIRE_HOURS", "24"),
	}

	return config, nil
}

// GetDatabaseURL 获取数据库连接字符串
func (c *TestConfig) GetDatabaseURL() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		c.DBUser, c.DBPassword, c.DBHost, c.DBPort, c.DBName, c.DBSSLMode)
}

// GetDSN 获取数据库 DSN
func (c *TestConfig) GetDSN() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		c.DBHost, c.DBPort, c.DBUser, c.DBPassword, c.DBName, c.DBSSLMode)
}

// getEnv 获取环境变量，如果不存在则返回默认值
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// FindProjectRoot 查找项目根目录（包含 go.mod 的目录）
func FindProjectRoot() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	for {
		// 检查当前目录是否包含 go.mod
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir, nil
		}

		// 获取父目录
		parent := filepath.Dir(dir)
		if parent == dir {
			// 已经到达文件系统根目录
			return "", fmt.Errorf("could not find project root (go.mod)")
		}
		dir = parent
	}
}