package controller

import (
	"github.com/gin-gonic/gin"
	"store.getAway/utils/exception"
	"store.getAway/utils/response"
	"store.getAway/utils/rpcx"
	"strconv"
)

// CreateComment
//
//	@Description: 针对商品创建评论
//	@param ctx
func CreateComment(ctx *gin.Context) {

}

// QueryComments
//
//	@Description: 查询商品评论
//	@param ctx
func QueryComments(ctx *gin.Context) {
	goodId, _ := strconv.Atoi(ctx.Param("goodId"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("page_size", "10"))
	pageNum, _ := strconv.Atoi(ctx.DefaultQuery("page_num", "1"))
	result := map[string]any{}
	err := rpcx.CallFuncService("Goods", "QueryComment", map[string]int{
		"GoodId":   goodId,
		"PageSize": pageSize,
		"PageNum":  pageNum,
	}, &result)

	exception.AssertException(ctx, err)
	response.WriteResponse(ctx, map[string]any{
		"count": result["Count"],
		"list":  result["List"],
	})
}
