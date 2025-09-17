package helpers

import (
	"database/sql"
	"fmt"
	"path/filepath"
	"time"

	"github.com/demo/demo-gin/tests/config"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

// TestDB 测试数据库连接
type TestDB struct {
	DB     *sql.DB
	Config *config.TestConfig
}

// SetupTestDB 初始化测试数据库
func SetupTestDB() (*TestDB, error) {
	// 加载测试配置
	cfg, err := config.LoadTestConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to load test config: %w", err)
	}

	// 连接数据库
	db, err := sql.Open("postgres", cfg.GetDSN())
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	// 配置连接池
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(5 * time.Minute)

	// 测试连接
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	testDB := &TestDB{
		DB:     db,
		Config: cfg,
	}

	// 运行迁移
	if err := testDB.RunMigrations(); err != nil {
		return nil, fmt.Errorf("failed to run migrations: %w", err)
	}

	return testDB, nil
}

// RunMigrations 运行数据库迁移
func (tdb *TestDB) RunMigrations() error {
	driver, err := postgres.WithInstance(tdb.DB, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("failed to create migration driver: %w", err)
	}

	// 获取项目根目录
	rootDir, err := findProjectRoot()
	if err != nil {
		return fmt.Errorf("failed to find project root: %w", err)
	}

	// 迁移文件路径
	migrationsPath := fmt.Sprintf("file://%s", filepath.Join(rootDir, "migrations"))

	m, err := migrate.NewWithDatabaseInstance(
		migrationsPath,
		"postgres",
		driver,
	)
	if err != nil {
		return fmt.Errorf("failed to create migration instance: %w", err)
	}

	// 运行迁移
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("failed to run migrations: %w", err)
	}

	return nil
}

// Close 关闭数据库连接
func (tdb *TestDB) Close() error {
	if tdb.DB != nil {
		return tdb.DB.Close()
	}
	return nil
}

// Begin 开始事务
func (tdb *TestDB) Begin() (*sql.Tx, error) {
	return tdb.DB.Begin()
}

// Exec 执行SQL语句
func (tdb *TestDB) Exec(query string, args ...interface{}) (sql.Result, error) {
	return tdb.DB.Exec(query, args...)
}

// Query 查询数据
func (tdb *TestDB) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return tdb.DB.Query(query, args...)
}

// QueryRow 查询单行数据
func (tdb *TestDB) QueryRow(query string, args ...interface{}) *sql.Row {
	return tdb.DB.QueryRow(query, args...)
}

// findProjectRoot 查找项目根目录
func findProjectRoot() (string, error) {
	return config.FindProjectRoot()
}

// 辅助函数：获取项目根目录
func FindProjectRoot() (string, error) {
	return config.FindProjectRoot()
}