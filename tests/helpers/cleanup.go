package helpers

import (
	"database/sql"
	"fmt"
	"strings"
)

// CleanupHelper 测试数据清理辅助工具
type CleanupHelper struct {
	DB *TestDB
}

// NewCleanupHelper 创建清理工具实例
func NewCleanupHelper(db *TestDB) *CleanupHelper {
	return &CleanupHelper{DB: db}
}

// TruncateTables 清空指定的表
func (c *CleanupHelper) TruncateTables(tables ...string) error {
	if len(tables) == 0 {
		// 如果没有指定表，则清空所有业务表
		tables = []string{"posts", "users"}
	}

	// 构建TRUNCATE语句
	query := fmt.Sprintf("TRUNCATE TABLE %s RESTART IDENTITY CASCADE", strings.Join(tables, ", "))

	// 执行清空操作
	if _, err := c.DB.Exec(query); err != nil {
		return fmt.Errorf("failed to truncate tables: %w", err)
	}

	return nil
}

// TruncateAllTables 清空所有业务表
func (c *CleanupHelper) TruncateAllTables() error {
	// 获取所有业务表（排除系统表和迁移表）
	tables, err := c.getBusinessTables()
	if err != nil {
		return fmt.Errorf("failed to get business tables: %w", err)
	}

	if len(tables) > 0 {
		return c.TruncateTables(tables...)
	}

	return nil
}

// getBusinessTables 获取所有业务表
func (c *CleanupHelper) getBusinessTables() ([]string, error) {
	query := `
		SELECT tablename
		FROM pg_tables
		WHERE schemaname = 'public'
		AND tablename NOT IN ('schema_migrations', 'pg_stat_statements')
		ORDER BY tablename
	`

	rows, err := c.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tables []string
	for rows.Next() {
		var tableName string
		if err := rows.Scan(&tableName); err != nil {
			return nil, err
		}
		tables = append(tables, tableName)
	}

	return tables, rows.Err()
}

// DeleteUser 删除指定用户及其关联数据
func (c *CleanupHelper) DeleteUser(userID int) error {
	// 开始事务
	tx, err := c.DB.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	// 删除用户的文章
	if _, err := tx.Exec("DELETE FROM posts WHERE author_id = $1", userID); err != nil {
		return fmt.Errorf("failed to delete user posts: %w", err)
	}

	// 删除用户
	if _, err := tx.Exec("DELETE FROM users WHERE id = $1", userID); err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}

	// 提交事务
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

// ResetSequences 重置所有序列
func (c *CleanupHelper) ResetSequences() error {
	query := `
		SELECT
			'ALTER SEQUENCE ' || sequence_name || ' RESTART WITH 1;' AS reset_sql
		FROM information_schema.sequences
		WHERE sequence_schema = 'public'
	`

	rows, err := c.DB.Query(query)
	if err != nil {
		return fmt.Errorf("failed to get sequences: %w", err)
	}
	defer rows.Close()

	var resetStatements []string
	for rows.Next() {
		var resetSQL string
		if err := rows.Scan(&resetSQL); err != nil {
			return err
		}
		resetStatements = append(resetStatements, resetSQL)
	}

	// 执行重置语句
	for _, stmt := range resetStatements {
		if _, err := c.DB.Exec(stmt); err != nil {
			return fmt.Errorf("failed to reset sequence: %w", err)
		}
	}

	return nil
}

// WithTransaction 在事务中执行测试
func (c *CleanupHelper) WithTransaction(fn func(*sql.Tx) error) error {
	tx, err := c.DB.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	if err := fn(tx); err != nil {
		return err
	}

	// 测试完成后回滚事务，保持数据库干净
	return tx.Rollback()
}

// TeardownTestDB 清理测试数据库连接
func TeardownTestDB(db *TestDB) error {
	if db == nil {
		return nil
	}

	// 清空所有表数据
	cleanup := NewCleanupHelper(db)
	if err := cleanup.TruncateAllTables(); err != nil {
		// 记录错误但不返回，确保连接被关闭
		fmt.Printf("Warning: failed to truncate tables: %v\n", err)
	}

	// 关闭数据库连接
	return db.Close()
}