package router

import (
	"github.com/gin-gonic/gin"
	"store.getAway/controller"
	"store.getAway/middlewares"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middlewares.ErrorHandler())

	v1Group := r.Group("/v1")
	public := v1Group.Group("")
	{
		userController := public.Group("/user")
		userController.POST("/login", controller.UserLogin)
		userController.POST("/register", controller.UserRegister)
	}
	private := v1Group.Group("", middlewares.AuthUser())
	{
		userController := private.Group("user")
		{
			userController.GET("/loginLog", controller.LoginLogs)
			userController.PUT("/:uid/info", controller.UpdateUser)
			userController.PUT("/:uid/password", controller.ChangePassWord)
			userController.PUT("/:uid/logout", controller.LogOut)
		}
	}
	orderController := private.Group("order")
	{
		orderController.POST("", controller.CreateOrder)

	}
	// 商品路由再不验证
	goodsController := public.Group("goods")
	{
		goodsController.GET(":goodId", controller.QueryGoods)

		goodsController.POST("", controller.CreateGoods)
		goodsController.DELETE(":goodId", controller.DeleteGoods)
		goodsController.GET(":goodId/comments", controller.QueryComments)
	}
	v1Group.POST("/emailCode", controller.SendEmailCode)
	return r
}
