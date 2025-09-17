package integration

import (
	"net/http"
	"testing"

	"github.com/demo/demo-gin/internal/handlers"
	"github.com/demo/demo-gin/tests/helpers"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestUserLogin(t *testing.T) {
	// 设置Gin为测试模式
	gin.SetMode(gin.TestMode)

	// 创建测试路由
	router := gin.New()
	authHandler := handlers.NewAuthHandler(nil) // 暂时使用nil，因为当前实现还没有真正使用数据库
	router.POST("/auth/login", authHandler.Login)

	// 创建测试客户端
	client := helpers.NewTestClient(router)

	t.Run("successful login with valid credentials", func(t *testing.T) {
		// 准备测试数据
		requestBody := map[string]interface{}{
			"username": "testuser",
			"password": "Test123456!",
		}

		// 发送POST请求
		w := client.Post("/auth/login", requestBody)

		// 断言响应状态码
		assert.Equal(t, http.StatusOK, w.Code)

		// 解析响应
		var response map[string]interface{}
		err := helpers.ParseJSON(w, &response)
		assert.NoError(t, err)

		// 断言响应结构
		assert.Contains(t, response, "access_token")
		assert.Contains(t, response, "token_type")
		assert.Contains(t, response, "expires_in")
		assert.Contains(t, response, "user")

		// 断言具体值
		assert.Equal(t, "mock_jwt_token", response["access_token"])
		assert.Equal(t, "Bearer", response["token_type"])
		assert.Equal(t, float64(3600), response["expires_in"])

		user := response["user"].(map[string]interface{})
		assert.Equal(t, "testuser", user["username"])

		// 确保密码不在响应中
		assert.NotContains(t, user, "password")
	})

	t.Run("login fails with missing username", func(t *testing.T) {
		requestBody := map[string]interface{}{
			"password": "Test123456!",
			// 缺少 username
		}

		w := client.Post("/auth/login", requestBody)

		// 应该返回400错误
		assert.Equal(t, http.StatusBadRequest, w.Code)

		var response map[string]interface{}
		err := helpers.ParseJSON(w, &response)
		assert.NoError(t, err)
		assert.Contains(t, response, "error")
	})

	t.Run("login fails with missing password", func(t *testing.T) {
		requestBody := map[string]interface{}{
			"username": "testuser",
			// 缺少 password
		}

		w := client.Post("/auth/login", requestBody)

		assert.Equal(t, http.StatusBadRequest, w.Code)

		var response map[string]interface{}
		err := helpers.ParseJSON(w, &response)
		assert.NoError(t, err)
		assert.Contains(t, response, "error")
	})

	t.Run("login fails with empty username", func(t *testing.T) {
		requestBody := map[string]interface{}{
			"username": "",
			"password": "Test123456!",
		}

		w := client.Post("/auth/login", requestBody)

		assert.Equal(t, http.StatusBadRequest, w.Code)

		var response map[string]interface{}
		err := helpers.ParseJSON(w, &response)
		assert.NoError(t, err)
		assert.Contains(t, response, "error")
	})

	t.Run("login fails with empty password", func(t *testing.T) {
		requestBody := map[string]interface{}{
			"username": "testuser",
			"password": "",
		}

		w := client.Post("/auth/login", requestBody)

		assert.Equal(t, http.StatusBadRequest, w.Code)

		var response map[string]interface{}
		err := helpers.ParseJSON(w, &response)
		assert.NoError(t, err)
		assert.Contains(t, response, "error")
	})

	t.Run("login fails with empty request body", func(t *testing.T) {
		w := client.Post("/auth/login", map[string]interface{}{})

		assert.Equal(t, http.StatusBadRequest, w.Code)

		var response map[string]interface{}
		err := helpers.ParseJSON(w, &response)
		assert.NoError(t, err)
		assert.Contains(t, response, "error")
	})

	t.Run("login handles malformed JSON", func(t *testing.T) {
		// 发送非JSON数据
		w := client.PostForm("/auth/login", map[string]string{
			"username": "testuser",
		})

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("login response format is consistent", func(t *testing.T) {
		// 多次登录请求验证响应格式一致性
		for i := 0; i < 3; i++ {
			requestBody := map[string]interface{}{
				"username": "testuser",
				"password": "Test123456!",
			}

			w := client.Post("/auth/login", requestBody)
			assert.Equal(t, http.StatusOK, w.Code)

			var response map[string]interface{}
			err := helpers.ParseJSON(w, &response)
			assert.NoError(t, err)

			// 确保必需字段都存在
			assert.Contains(t, response, "access_token")
			assert.Contains(t, response, "token_type")
			assert.Contains(t, response, "expires_in")
			assert.Contains(t, response, "user")
		}
	})
}