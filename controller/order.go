package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// CreateOrder
//
//	@Description: 创建订单
//	@param ctx
func CreateOrder(ctx *gin.Context) {

}

// GetOrdersByUid
//
//	@Description:  查询某个人的订单
//	@param ctx
func GetOrdersByUid(ctx *gin.Context) {
	uid, existed := ctx.Get("Uid")
	if !existed {

	}
	fmt.Println(uid)
}

func GetOrderById(ctx *gin.Context) {

}
