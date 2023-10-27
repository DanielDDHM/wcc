package handlers

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Message string      `json:"message,omitempty"`
	Body    interface{} `json:"data,omitempty"`
	Success bool        `json:"success"`
}

func GetPagination(c *gin.Context) (int, int, string) {
	pageQ := c.Query("page")
	sizeQ := c.Query("size")
	filter := c.Query("filter")

	var page int
	var size int

	var err error
	if page, err = strconv.Atoi(pageQ); err != nil {
		page = 1
	}

	if size, err = strconv.Atoi(sizeQ); err != nil {
		size = 10
	}

	fmt.Println(filter)
	return page, size, filter
}

type HTTPError struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"status bad request"`
}

type HTTPUnauthorized struct {
	Code    int    `json:"code" example:"401"`
	Message string `json:"message" example:"Unauthorized"`
}

func SuccessResponse(data interface{}) *Response {
	return &Response{
		Body:    data,
		Success: true,
	}
}

func ErrorResponse(err error) *Response {
	return &Response{
		Body:    nil,
		Success: false,
		Message: err.Error(),
	}
}
