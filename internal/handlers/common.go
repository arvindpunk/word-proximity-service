package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const VERSION = "0.0.1"

const (
	ResponseCodeSuccess = "SUCCESS"
	ResponseCodeError   = "SUCCESS"
)

type ResponseStructure struct {
	Code    string `json:"code"`
	Version string `json:"version"`

	Data interface{} `json:"data,omitempty"`

	Error     string `json:"error,omitempty"`
	ErrorCode string `json:"errorCode,omitempty"`
}

func (rs *ResponseStructure) Respond(c *gin.Context) {
	c.JSON(http.StatusOK, rs)
}

func WithSuccess(data interface{}) *ResponseStructure {
	return &ResponseStructure{
		Code:    ResponseCodeSuccess,
		Version: VERSION,

		Data: data,
	}
}

func WithError(errCode string, err string) *ResponseStructure {
	return &ResponseStructure{
		Code:    ResponseCodeError,
		Version: VERSION,

		Error:     err,
		ErrorCode: errCode,
	}
}
