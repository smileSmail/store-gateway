package controller

import (
	"github.com/gin-gonic/gin"
	"store.getAway/utils/exception"
	"store.getAway/utils/response"
	"store.getAway/utils/rpcx"
)

type RequestSendEmail struct {
	Email string `json:"email" binding:"required"`
}

// SendEmailCode
//
//	@Description: 请求发送邮箱验证码
//	@param ctx
func SendEmailCode(ctx *gin.Context) {
	postForm := RequestSendEmail{}
	err := ctx.ShouldBind(&postForm)
	exception.AssertException(ctx, exception.NewException(exception.ValidateError, err.Error()))

	reply := map[string]interface{}{}
	err = rpcx.CallFuncService("User", "SendEmailCode", &postForm, &reply)
	exception.AssertException(ctx, err)
	response.WriteResponse(ctx, nil)
}
