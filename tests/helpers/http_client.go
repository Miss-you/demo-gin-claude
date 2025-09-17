package helpers

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
)

// TestClient HTTP测试客户端
type TestClient struct {
	Router *gin.Engine
	Token  string
}

// NewTestClient 创建测试客户端
func NewTestClient(router *gin.Engine) *TestClient {
	return &TestClient{
		Router: router,
	}
}

// SetAuth 设置认证Token
func (c *TestClient) SetAuth(token string) {
	c.Token = token
}

// Get 发送GET请求
func (c *TestClient) Get(path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(http.MethodGet, path, nil)
	return c.doRequest(req)
}

// Post 发送POST请求
func (c *TestClient) Post(path string, body interface{}) *httptest.ResponseRecorder {
	jsonBody, _ := json.Marshal(body)
	req, _ := http.NewRequest(http.MethodPost, path, bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	return c.doRequest(req)
}

// Put 发送PUT请求
func (c *TestClient) Put(path string, body interface{}) *httptest.ResponseRecorder {
	jsonBody, _ := json.Marshal(body)
	req, _ := http.NewRequest(http.MethodPut, path, bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	return c.doRequest(req)
}

// Patch 发送PATCH请求
func (c *TestClient) Patch(path string, body interface{}) *httptest.ResponseRecorder {
	jsonBody, _ := json.Marshal(body)
	req, _ := http.NewRequest(http.MethodPatch, path, bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	return c.doRequest(req)
}

// Delete 发送DELETE请求
func (c *TestClient) Delete(path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(http.MethodDelete, path, nil)
	return c.doRequest(req)
}

// PostForm 发送表单POST请求
func (c *TestClient) PostForm(path string, values map[string]string) *httptest.ResponseRecorder {
	formData := make([]byte, 0)
	for k, v := range values {
		if len(formData) > 0 {
			formData = append(formData, '&')
		}
		formData = append(formData, []byte(k+"="+v)...)
	}

	req, _ := http.NewRequest(http.MethodPost, path, bytes.NewBuffer(formData))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	return c.doRequest(req)
}

// doRequest 执行请求
func (c *TestClient) doRequest(req *http.Request) *httptest.ResponseRecorder {
	// 添加认证头
	if c.Token != "" {
		req.Header.Set("Authorization", "Bearer "+c.Token)
	}

	// 创建响应记录器
	w := httptest.NewRecorder()

	// 执行请求
	c.Router.ServeHTTP(w, req)

	return w
}

// ParseJSON 解析JSON响应
func ParseJSON(w *httptest.ResponseRecorder, v interface{}) error {
	return json.Unmarshal(w.Body.Bytes(), v)
}

// AssertStatus 断言HTTP状态码
func AssertStatus(t interface{ Errorf(string, ...interface{}) }, w *httptest.ResponseRecorder, expectedStatus int) {
	if w.Code != expectedStatus {
		body, _ := io.ReadAll(w.Body)
		t.Errorf("Expected status %d, got %d. Response body: %s", expectedStatus, w.Code, string(body))
	}
}

// Response 通用响应结构
type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Error   string      `json:"error"`
}

// ErrorResponse 错误响应结构
type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
	Code    int    `json:"code"`
}