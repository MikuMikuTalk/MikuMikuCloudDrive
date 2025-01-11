package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code    int    `json:"code"`
	Data    any    `json:"data"`
	Message string `json:"message"`
}

func (r *Response) SuccessResponse(c *gin.Context, data any, message ...string) {
	r.Code = http.StatusOK
	if len(message) > 0 {
		r.Message = message[0]
	}
	r.Data = data
	c.JSON(r.Code, r)
}
func (r *Response) ErrorResponse(c *gin.Context, code int, message ...string) {
	r.Code = code
	r.Data = nil
	if len(message) > 0 {
		r.Message = message[0]
	}
	c.JSON(r.Code, r)
}

func (r *Response) Response(c *gin.Context, code int, data any, message string) {
	r.Code = code
	r.Data = data
	r.Message = message
	c.JSON(r.Code, r)
}
