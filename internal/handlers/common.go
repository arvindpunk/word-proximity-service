package handlers

import (
	"net/http"

	"github.com/arvindpunk/word-proximity-service/internal/utils"
	"github.com/gin-gonic/gin"
)

const (
	ResponseCodeSuccess = "SUCCESS"
	ResponseCodeError   = "ERROR"
)

type ResponseStructure struct {
	Code    string `json:"code"`
	Version string `json:"version"`

	Data interface{} `json:"data,omitempty"`

	Error string `json:"error,omitempty"`
}

func (rs *ResponseStructure) Respond(c *gin.Context) {
	c.JSON(http.StatusOK, rs)
}

func WithSuccess(data interface{}) *ResponseStructure {
	return &ResponseStructure{
		Code:    ResponseCodeSuccess,
		Version: utils.VERSION,

		Data: data,
	}
}

func WithError(err error) *ResponseStructure {
	return &ResponseStructure{
		Code:    ResponseCodeError,
		Version: utils.VERSION,

		Error: err.Error(),
	}
}
