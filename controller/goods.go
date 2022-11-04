package controller

import (
	"github.com/gin-gonic/gin"
	"store.getAway/utils/exception"
	"store.getAway/utils/response"
	"store.getAway/utils/rpcx"
	"store.getAway/vo"
	"strconv"
)

func QueryGoods(ctx *gin.Context) {

}

// CreateGoods
//
//	@Description: 新增商品
//	@param ctx
func CreateGoods(ctx *gin.Context) {
	goods := new(vo.RequestCreateGoods)
	err := ctx.BindJSON(goods)
	if err != nil {
		exception.AssertException(ctx, exception.NewException(exception.ValidateError, err.Error()))
		return
	}
	resp := map[string]interface{}{}
	//  如果resp传递nil，微服务报错也将不会得到err
	err = rpcx.CallFuncService("Goods", "CreateGoods", goods, &resp)
	if err != nil {
		exception.AssertException(ctx, err)
		return
	}
	response.WriteResponse(ctx, nil)
}

func DeleteGoods(ctx *gin.Context) {
	goodId, err := strconv.Atoi(ctx.Param("goodId"))
	if err != nil {
		exception.AssertException(ctx, exception.ValidateError)
		return
	}
	err = rpcx.CallFuncService("Goods", "QueryComment", goodId, nil)
	if err != nil {
		exception.AssertException(ctx, err)
	}
	response.WriteResponse(ctx, nil)
}

// PublishGoods
//
//	@Description: 发布商品
//	@param ctx
func PublishGoods(ctx *gin.Context) {

}
