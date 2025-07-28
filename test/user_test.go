package test

import (
	"encoding/json"
	"net/http/httptest"
	"testing"

	"gin-template/middleware"
	"gin-template/models"

	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	router := SetupTestRouter()

	// 测试用例
	testCases := []struct {
		name           string
		payload        any
		expectedStatus int
	}{
		{
			name: "Valid user creation",
			payload: models.CreateUserRequest{
				Name:  "John Doe",
				Email: "john@example.com",
				Age:   25,
				Phone: "1234567890",
			},
			expectedStatus: 200,
		},
		{
			name: "Invalid email",
			payload: models.CreateUserRequest{
				Name:  "Jane Doe",
				Email: "invalid-email",
				Age:   30,
			},
			expectedStatus: 400,
		},
		{
			name: "Missing required field",
			payload: models.CreateUserRequest{
				Email: "test@example.com",
				Age:   25,
			},
			expectedStatus: 400,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req := MakeRequest("POST", "/api/v1/users", tc.payload)
			w := httptest.NewRecorder()

			router.ServeHTTP(w, req)

			assert.Equal(t, tc.expectedStatus, w.Code)

			if tc.expectedStatus == 200 {
				var response middleware.Response
				ParseResponseBody(t, w, &response)
				assert.Equal(t, 200, response.Code)
				assert.Equal(t, "success", response.Message)
				assert.NotNil(t, response.Data)
			}
		})
	}
}

func TestGetUser(t *testing.T) {
	router := SetupTestRouter()

	// 先创建一个用户
	createReq := MakeRequest("POST", "/api/v1/users", models.CreateUserRequest{
		Name:  "Test User",
		Email: "test@example.com",
		Age:   25,
	})
	createW := httptest.NewRecorder()
	router.ServeHTTP(createW, createReq)

	var createResponse middleware.Response
	ParseResponseBody(t, createW, &createResponse)

	userData, _ := json.Marshal(createResponse.Data)
	var user models.User
	json.Unmarshal(userData, &user)

	// 测试获取用户
	req := MakeRequest("GET", "/api/v1/users/1", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	AssertStatusOK(t, w)

	var response middleware.Response
	ParseResponseBody(t, w, &response)
	assert.Equal(t, "success", response.Message)
}

func TestGetUsers(t *testing.T) {
	router := SetupTestRouter()

	// 创建几个用户
	users := []models.CreateUserRequest{
		{Name: "User 1", Email: "user1@example.com", Age: 25},
		{Name: "User 2", Email: "user2@example.com", Age: 30},
	}

	for _, user := range users {
		req := MakeRequest("POST", "/api/v1/users", user)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
	}

	// 测试获取用户列表
	req := MakeRequest("GET", "/api/v1/users?page=1&page_size=10", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	AssertStatusOK(t, w)

	var response middleware.Response
	ParseResponseBody(t, w, &response)
	assert.Equal(t, "success", response.Message)
	assert.NotNil(t, response.Data)
}

func TestUpdateUser(t *testing.T) {
	router := SetupTestRouter()

	// 先创建一个用户
	createReq := MakeRequest("POST", "/api/v1/users", models.CreateUserRequest{
		Name:  "Original Name",
		Email: "original@example.com",
		Age:   25,
	})
	createW := httptest.NewRecorder()
	router.ServeHTTP(createW, createReq)

	// 更新用户
	newName := "Updated Name"
	updateReq := MakeRequest("PUT", "/api/v1/users/1", models.UpdateUserRequest{
		Name: &newName,
	})
	w := httptest.NewRecorder()

	router.ServeHTTP(w, updateReq)

	AssertStatusOK(t, w)

	var response middleware.Response
	ParseResponseBody(t, w, &response)
	assert.Equal(t, "success", response.Message)
}

func TestDeleteUser(t *testing.T) {
	router := SetupTestRouter()

	// 先创建一个用户
	createReq := MakeRequest("POST", "/api/v1/users", models.CreateUserRequest{
		Name:  "To Be Deleted",
		Email: "delete@example.com",
		Age:   25,
	})
	createW := httptest.NewRecorder()
	router.ServeHTTP(createW, createReq)

	// 删除用户
	req := MakeRequest("DELETE", "/api/v1/users/1", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	AssertStatusOK(t, w)

	var response middleware.Response
	ParseResponseBody(t, w, &response)
	assert.Equal(t, "success", response.Message)
}
