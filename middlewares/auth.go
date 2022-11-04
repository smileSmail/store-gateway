package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"store.getAway/utils/helper"
)

func AuthUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorization := ctx.Request.Header["Authorization"][0]
		userInfo, err := helper.ParseToken(authorization)
		if err != nil {
			fmt.Println(err)
			return
		}
		ctx.Set("UserName", userInfo.UserId)
		ctx.Set("Uid", userInfo.UserId)
		ctx.Next()
	}
}
