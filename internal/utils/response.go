package utils

import "time"

// Response 统一响应结构
type Response struct {
	Code      int         `json:"code"`
	Data      interface{} `json:"data"`
	Message   string      `json:"message"`
	Timestamp int64       `json:"timestamp"`
}

// Success 成功响应
func Success(data interface{}) *Response {
	return &Response{
		Code:      0,
		Data:      data,
		Message:   "success",
		Timestamp: timeNow(),
	}
}

// SuccessWithMessage 带消息的成功响应
func SuccessWithMessage(data interface{}, message string) *Response {
	return &Response{
		Code:      0,
		Data:      data,
		Message:   message,
		Timestamp: timeNow(),
	}
}

// Error 错误响应
func Error(code int, message string) *Response {
	return &Response{
		Code:      code,
		Data:      nil,
		Message:   message,
		Timestamp: timeNow(),
	}
}

// BadRequest 参数错误
func BadRequest(message string) *Response {
	return &Response{
		Code:      400,
		Data:      nil,
		Message:   message,
		Timestamp: timeNow(),
	}
}

// Unauthorized 未授权
func Unauthorized(message string) *Response {
	return &Response{
		Code:      401,
		Data:      nil,
		Message:   message,
		Timestamp: timeNow(),
	}
}

// NotFound 未找到
func NotFound(message string) *Response {
	return &Response{
		Code:      404,
		Data:      nil,
		Message:   message,
		Timestamp: timeNow(),
	}
}

// InternalError 服务器内部错误
func InternalError(message string) *Response {
	return &Response{
		Code:      500,
		Data:      nil,
		Message:   message,
		Timestamp: timeNow(),
	}
}

// PageResponse 分页响应
type PageResponse struct {
	List     interface{} `json:"list"`
	Total    int64       `json:"total"`
	Page     int         `json:"page"`
	PageSize int         `json:"pageSize"`
}

func timeNow() int64 {
	return time.Now().UnixMilli()
}
