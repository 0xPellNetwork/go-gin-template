package middleware

import (
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
)

// Response represents the unified response structure
type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

// ErrorResponse sends an error response
func ErrorResponse(c *gin.Context, code int, message string) {
	c.JSON(code, Response{
		Code:    code,
		Message: message,
	})
}

// SuccessResponse sends a success response
func SuccessResponse(c *gin.Context, data any) {
	c.JSON(http.StatusOK, Response{
		Code:    http.StatusOK,
		Message: "success",
		Data:    data,
	})
}

// Handler 定义通用的处理器函数类型
type Handler any

// BindAndCall creates a middleware that automatically binds parameters and calls handler
func BindAndCall(handler Handler, bindTypes ...any) gin.HandlerFunc {
	return func(c *gin.Context) {
		logger := log.With().
			Str("method", c.Request.Method).
			Str("path", c.Request.URL.Path).
			Str("component", "middleware").
			Logger()

		// Prepare arguments
		args := []reflect.Value{reflect.ValueOf(c)}

		// Process each binding type
		for i, bindType := range bindTypes {
			// Create parameter instance
			paramType := reflect.TypeOf(bindType).Elem()
			paramValue := reflect.New(paramType)
			param := paramValue.Interface()

			var err error

			// Decide binding method based on HTTP method
			switch c.Request.Method {
			case "GET", "DELETE":
				err = c.ShouldBindQuery(param)
				logger.Debug().Str("binding_type", "query").Msg("Binding query parameters")
			case "POST", "PUT", "PATCH":
				if i == 0 {
					// First parameter is usually JSON body
					err = c.ShouldBindJSON(param)
					logger.Debug().Str("binding_type", "json").Msg("Binding JSON body")
				} else {
					// Other parameters might be query parameters
					err = c.ShouldBindQuery(param)
					logger.Debug().Str("binding_type", "query").Msg("Binding query parameters")
				}
			}

			if err != nil {
				logger.Error().Err(err).Msg("Parameter binding failed")
				if validationErr, ok := err.(validator.ValidationErrors); ok {
					ErrorResponse(c, http.StatusBadRequest, validationErr.Error())
				} else {
					ErrorResponse(c, http.StatusBadRequest, err.Error())
				}
				c.Abort()
				return
			}

			args = append(args, paramValue.Elem())
		}

		// Call handler
		logger.Debug().Msg("Calling handler with bound parameters")
		handlerValue := reflect.ValueOf(handler)
		handlerValue.Call(args)
	}
}

// PathParam 用于路径参数的结构体
type PathParam struct {
	ID string `uri:"id" binding:"required"`
}

// BindPathParam 绑定路径参数
func BindPathParam(c *gin.Context, param *PathParam) error {
	return c.ShouldBindUri(param)
}

// GetPathID 获取路径中的 ID 参数
func GetPathID(c *gin.Context) (string, error) {
	var param PathParam
	if err := BindPathParam(c, &param); err != nil {
		return "", err
	}
	return param.ID, nil
}
