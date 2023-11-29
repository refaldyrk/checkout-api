package http_response

import (
	"checkout-api/domain/rest"
	error_utils "checkout-api/utils/error"

	"github.com/gin-gonic/gin"
)

type ResponseWriter struct{}

type IResponseWriter interface {
	HTTPCustomErr(ctx *gin.Context, err error)
	HTTPJsonErr(ctx *gin.Context, code int, message string, detail string, data interface{})
	HTTPJson(ctx *gin.Context, data interface{}, message string)
}

func NewResponseWriter() IResponseWriter {
	return &ResponseWriter{}
}

func (r *ResponseWriter) HTTPCustomErr(ctx *gin.Context, err error) {
	customErr, ok := err.(*error_utils.CustomErr)
	if ok {
		ctx.JSON(customErr.Code, rest.BaseJSONResp{
			Error:   true,
			Message: customErr.Error(),
			Detail:  "",
			Data:    nil,
		})
		return
	}
	ctx.JSON(500, rest.BaseJSONResp{
		Error:   true,
		Message: "internal server error",
		Detail:  err.Error(),
		Data:    nil,
	})
}

func (r *ResponseWriter) HTTPJsonErr(ctx *gin.Context, code int, message string, detail string, data interface{}) {
	ctx.JSON(code, rest.BaseJSONResp{
		Error:   true,
		Message: message,
		Detail:  detail,
		Data:    data,
	})
}

func (r *ResponseWriter) HTTPJson(ctx *gin.Context, data interface{}, message string) {
	ctx.JSON(200, rest.BaseJSONResp{
		Error:   false,
		Message: message,
		Detail:  "",
		Data:    data,
	})
}
