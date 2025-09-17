package integration

import (
	"net/http"
	"testing"

	"github.com/demo/demo-gin/internal/handlers"
	"github.com/demo/demo-gin/tests/helpers"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestHealthCheck(t *testing.T) {
	// 设置Gin为测试模式
	gin.SetMode(gin.TestMode)

	// 创建测试路由
	router := gin.New()
	router.GET("/health", handlers.Health)

	// 创建测试客户端
	client := helpers.NewTestClient(router)

	t.Run("health check should return 200", func(t *testing.T) {
		// 发送GET请求
		w := client.Get("/health")

		// 断言响应状态码
		assert.Equal(t, http.StatusOK, w.Code)

		// 解析响应
		var response map[string]interface{}
		err := helpers.ParseJSON(w, &response)
		assert.NoError(t, err)

		// 断言响应内容
		assert.Equal(t, "healthy", response["status"])
		assert.Equal(t, "Service is running", response["message"])
		assert.Contains(t, w.Header().Get("Content-Type"), "application/json")
	})

	t.Run("health check should not require authentication", func(t *testing.T) {
		// 不设置任何认证头，直接访问
		w := client.Get("/health")

		// 应该成功返回200
		assert.Equal(t, http.StatusOK, w.Code)

		var response map[string]interface{}
		err := helpers.ParseJSON(w, &response)
		assert.NoError(t, err)
		assert.Equal(t, "healthy", response["status"])
	})

	t.Run("health check response format should be consistent", func(t *testing.T) {
		// 多次请求验证响应格式一致性
		for i := 0; i < 3; i++ {
			w := client.Get("/health")
			assert.Equal(t, http.StatusOK, w.Code)

			var response map[string]interface{}
			err := helpers.ParseJSON(w, &response)
			assert.NoError(t, err)
			assert.Contains(t, response, "status")
			assert.Contains(t, response, "message")
		}
	})
}