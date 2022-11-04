package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type CustomResponse struct {
	Code    uint64      `json:"code"`
	Data    interface{} `json:"data"`
	Success bool        `json:"success"`
	Msg     string      `json:"msg"`
}

// WriteResponse
//
//	@Description: 重写返回
func WriteResponse(ctx *gin.Context, respData interface{}) {
	response := CustomResponse{
		Code:    200,
		Msg:     "成功",
		Success: true,
		Data:    respData,
	}
	ctx.JSON(http.StatusOK, response)
}
