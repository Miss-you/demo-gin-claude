package fixtures

import (
	"database/sql"
	"fmt"
	"math/rand"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// User 用户模型
type User struct {
	ID        int
	Username  string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Post 文章模型
type Post struct {
	ID          int
	Title       string
	Content     string
	AuthorID    int
	Status      string
	PublishedAt *time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// CreateTestUser 创建测试用户
func CreateTestUser(db *sql.DB) (*User, error) {
	username := GenerateRandomUsername()
	email := GenerateRandomEmail()
	password := "Test123456!"

	// 加密密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &User{
		Username: username,
		Email:    email,
		Password: string(hashedPassword),
	}

	// 插入数据库
	query := `
		INSERT INTO users (username, email, password, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, created_at, updated_at
	`
	now := time.Now()
	err = db.QueryRow(query, username, email, string(hashedPassword), now, now).
		Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		return nil, err
	}

	// 返回原始密码，便于测试登录
	user.Password = password
	return user, nil
}

// CreateTestUserWithData 创建指定数据的测试用户
func CreateTestUserWithData(db *sql.DB, username, email, password string) (*User, error) {
	// 加密密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &User{
		Username: username,
		Email:    email,
		Password: password, // 保存原始密码用于测试
	}

	// 插入数据库
	query := `
		INSERT INTO users (username, email, password, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, created_at, updated_at
	`
	now := time.Now()
	err = db.QueryRow(query, username, email, string(hashedPassword), now, now).
		Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt)

	return user, err
}

// CreateTestPost 创建测试文章
func CreateTestPost(db *sql.DB, userID int) (*Post, error) {
	title := GenerateRandomTitle()
	content := GenerateRandomContent()
	status := "published"
	now := time.Now()

	post := &Post{
		Title:       title,
		Content:     content,
		AuthorID:    userID,
		Status:      status,
		PublishedAt: &now,
	}

	// 插入数据库
	query := `
		INSERT INTO posts (title, content, author_id, status, published_at, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id, created_at, updated_at
	`
	err := db.QueryRow(query, title, content, userID, status, now, now, now).
		Scan(&post.ID, &post.CreatedAt, &post.UpdatedAt)

	return post, err
}

// CreateTestPostWithData 创建指定数据的测试文章
func CreateTestPostWithData(db *sql.DB, userID int, title, content, status string) (*Post, error) {
	now := time.Now()
	var publishedAt *time.Time
	if status == "published" {
		publishedAt = &now
	}

	post := &Post{
		Title:       title,
		Content:     content,
		AuthorID:    userID,
		Status:      status,
		PublishedAt: publishedAt,
	}

	// 插入数据库
	query := `
		INSERT INTO posts (title, content, author_id, status, published_at, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id, created_at, updated_at
	`
	err := db.QueryRow(query, title, content, userID, status, publishedAt, now, now).
		Scan(&post.ID, &post.CreatedAt, &post.UpdatedAt)

	return post, err
}

// GenerateRandomEmail 生成随机邮箱
func GenerateRandomEmail() string {
	return fmt.Sprintf("user%d@test.com", rand.Intn(1000000))
}

// GenerateRandomUsername 生成随机用户名
func GenerateRandomUsername() string {
	return fmt.Sprintf("user_%d", rand.Intn(1000000))
}

// GenerateRandomTitle 生成随机标题
func GenerateRandomTitle() string {
	titles := []string{
		"Introduction to Go Programming",
		"Building RESTful APIs with Gin",
		"Database Design Best Practices",
		"Microservices Architecture Guide",
		"Testing Strategies for Go Applications",
	}
	return titles[rand.Intn(len(titles))] + fmt.Sprintf(" - Part %d", rand.Intn(10)+1)
}

// GenerateRandomContent 生成随机内容
func GenerateRandomContent() string {
	return fmt.Sprintf("This is test content generated at %s. Lorem ipsum dolor sit amet, consectetur adipiscing elit. "+
		"Sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Test ID: %d",
		time.Now().Format("2006-01-02 15:04:05"), rand.Intn(10000))
}

// CreateMultipleUsers 批量创建测试用户
func CreateMultipleUsers(db *sql.DB, count int) ([]*User, error) {
	users := make([]*User, 0, count)
	for i := 0; i < count; i++ {
		user, err := CreateTestUser(db)
		if err != nil {
			return users, err
		}
		users = append(users, user)
	}
	return users, nil
}

// CreateMultiplePosts 批量创建测试文章
func CreateMultiplePosts(db *sql.DB, userID int, count int) ([]*Post, error) {
	posts := make([]*Post, 0, count)
	for i := 0; i < count; i++ {
		post, err := CreateTestPost(db, userID)
		if err != nil {
			return posts, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}