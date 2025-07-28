package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"gin-template/config"
	"gin-template/database"
	"gin-template/router"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

// TestDB 测试数据库连接
var TestDB *gorm.DB

// SetupTestDB 设置测试数据库
func SetupTestDB() *gorm.DB {
	cfg := config.DatabaseConfig{
		Driver: "sqlite",
		DSN:    ":memory:",
	}

	db, err := database.New(cfg)
	if err != nil {
		panic("Failed to connect to test database: " + err.Error())
	}

	TestDB = db
	return db
}

// SetupTestRouter 设置测试路由
func SetupTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	db := SetupTestDB()
	return router.New(db)
}

// MakeRequest 创建 HTTP 请求帮助函数
func MakeRequest(method, url string, body any) *http.Request {
	var bodyReader *bytes.Reader

	if body != nil {
		jsonBody, _ := json.Marshal(body)
		bodyReader = bytes.NewReader(jsonBody)
	} else {
		bodyReader = bytes.NewReader([]byte{})
	}

	req, _ := http.NewRequest(method, url, bodyReader)
	req.Header.Set("Content-Type", "application/json")
	return req
}

// AssertStatusOK 断言状态码为 200
func AssertStatusOK(t *testing.T, w *httptest.ResponseRecorder) {
	assert.Equal(t, http.StatusOK, w.Code)
}

// AssertStatusCreated 断言状态码为 201
func AssertStatusCreated(t *testing.T, w *httptest.ResponseRecorder) {
	assert.Equal(t, http.StatusCreated, w.Code)
}

// AssertStatusBadRequest 断言状态码为 400
func AssertStatusBadRequest(t *testing.T, w *httptest.ResponseRecorder) {
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

// AssertStatusNotFound 断言状态码为 404
func AssertStatusNotFound(t *testing.T, w *httptest.ResponseRecorder) {
	assert.Equal(t, http.StatusNotFound, w.Code)
}

// ParseResponseBody 解析响应体
func ParseResponseBody(t *testing.T, w *httptest.ResponseRecorder, result any) {
	err := json.Unmarshal(w.Body.Bytes(), result)
	assert.NoError(t, err)
}
