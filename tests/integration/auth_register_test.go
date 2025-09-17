package integration

import (
	"net/http"
	"testing"

	"github.com/demo/demo-gin/internal/handlers"
	"github.com/demo/demo-gin/tests/helpers"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestUserRegister(t *testing.T) {
	// 设置Gin为测试模式
	gin.SetMode(gin.TestMode)

	// 创建测试路由
	router := gin.New()
	authHandler := handlers.NewAuthHandler(nil) // 暂时使用nil，因为当前实现还没有真正使用数据库
	router.POST("/auth/register", authHandler.Register)

	// 创建测试客户端
	client := helpers.NewTestClient(router)

	t.Run("successful registration with valid data", func(t *testing.T) {
		// 准备测试数据
		requestBody := map[string]interface{}{
			"email":     "test@example.com",
			"username":  "testuser",
			"password":  "Test123456!",
			"full_name": "Test User",
		}

		// 发送POST请求
		w := client.Post("/auth/register", requestBody)

		// 断言响应状态码
		assert.Equal(t, http.StatusCreated, w.Code)

		// 解析响应
		var response map[string]interface{}
		err := helpers.ParseJSON(w, &response)
		assert.NoError(t, err)

		// 断言响应内容
		assert.Equal(t, "User registered successfully", response["message"])
		assert.Contains(t, response, "user")

		user := response["user"].(map[string]interface{})
		assert.Equal(t, "test@example.com", user["email"])
		assert.Equal(t, "testuser", user["username"])

		// 确保密码不在响应中
		assert.NotContains(t, user, "password")
	})

	t.Run("registration fails with invalid email", func(t *testing.T) {
		requestBody := map[string]interface{}{
			"email":    "invalid-email",
			"username": "testuser",
			"password": "Test123456!",
		}

		w := client.Post("/auth/register", requestBody)

		// 应该返回400错误
		assert.Equal(t, http.StatusBadRequest, w.Code)

		var response map[string]interface{}
		err := helpers.ParseJSON(w, &response)
		assert.NoError(t, err)
		assert.Contains(t, response, "error")
	})

	t.Run("registration fails with short username", func(t *testing.T) {
		requestBody := map[string]interface{}{
			"email":    "test@example.com",
			"username": "ab", // 少于3个字符
			"password": "Test123456!",
		}

		w := client.Post("/auth/register", requestBody)

		assert.Equal(t, http.StatusBadRequest, w.Code)

		var response map[string]interface{}
		err := helpers.ParseJSON(w, &response)
		assert.NoError(t, err)
		assert.Contains(t, response, "error")
	})

	t.Run("registration fails with short password", func(t *testing.T) {
		requestBody := map[string]interface{}{
			"email":    "test@example.com",
			"username": "testuser",
			"password": "1234567", // 少于8个字符
		}

		w := client.Post("/auth/register", requestBody)

		assert.Equal(t, http.StatusBadRequest, w.Code)

		var response map[string]interface{}
		err := helpers.ParseJSON(w, &response)
		assert.NoError(t, err)
		assert.Contains(t, response, "error")
	})

	t.Run("registration fails with missing required fields", func(t *testing.T) {
		requestBody := map[string]interface{}{
			"username": "testuser",
			// 缺少 email 和 password
		}

		w := client.Post("/auth/register", requestBody)

		assert.Equal(t, http.StatusBadRequest, w.Code)

		var response map[string]interface{}
		err := helpers.ParseJSON(w, &response)
		assert.NoError(t, err)
		assert.Contains(t, response, "error")
	})

	t.Run("registration fails with empty request body", func(t *testing.T) {
		w := client.Post("/auth/register", map[string]interface{}{})

		assert.Equal(t, http.StatusBadRequest, w.Code)

		var response map[string]interface{}
		err := helpers.ParseJSON(w, &response)
		assert.NoError(t, err)
		assert.Contains(t, response, "error")
	})

	t.Run("registration handles malformed JSON", func(t *testing.T) {
		// 发送非JSON数据
		w := client.PostForm("/auth/register", map[string]string{
			"email": "test@example.com",
		})

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}