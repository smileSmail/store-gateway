package exception

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"store.getAway/utils/response"
)

type Exception struct {
	Msg  string `json:"msg"`
	Code uint64 `json:"code"`
}

func NewException(code uint64, msg string) *Exception {
	return &Exception{Msg: msg, Code: code}
}

func AssertException(ctx *gin.Context, respData interface{}) {
	// 如果是Exception 则代表是自定义错误
	customResponse := response.CustomResponse{}
	if v, ok := respData.(Exception); ok {
		customResponse.Success = false
		customResponse.Msg = v.Msg
		customResponse.Code = v.Code
		ctx.JSON(http.StatusOK, customResponse)
		return
	}
	// 如果是其他错误需要记录日志
	if v, ok := respData.(error); ok {
		customResponse.Success = false
		customResponse.Msg = v.Error()
		customResponse.Code = http.StatusInternalServerError
		ctx.JSON(http.StatusInternalServerError, customResponse)
		return
	}

}
