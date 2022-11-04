package controller

import (
	"github.com/gin-gonic/gin"
	"store.getAway/utils/exception"
	"store.getAway/utils/response"
	"store.getAway/utils/rpcx"
	"store.getAway/vo"
	"strconv"
	"time"
)

// UserLogin
//
//	@Description: 用户登陆
//	@param ctx
func UserLogin(ctx *gin.Context) {
	loginForm := vo.RequestLogin{}
	err := ctx.BindJSON(&loginForm)
	exception.AssertException(ctx, exception.NewException(exception.ValidateError, err.Error()))
	loginForm.Ua = ctx.GetHeader("User-Agent")
	//var loginResult dto.ResponseLogin
	//// 调用user login 方法
	//err = rpcx.CallFuncService("User", "Login", &loginForm, &loginResult)
	//exception.AssertException(ctx, exception.NewException(exception.ValidateError, err.Error()))
	//// token生成
	//token, err := helper.GenerateToken(helper.CustomerUserInfo{
	//	UserId:   loginResult.ID,
	//	UserName: loginResult.UserName,
	//})
	//exception.AssertException(ctx, exception.NewException(exception.ValidateError, err.Error()))
	//
	//ctx.Header("Authorization", token)
	//response.WriteResponse(ctx, loginResult)
}

// UserRegister
//
//	@Description: 用户注册
//	@param ctx
func UserRegister(ctx *gin.Context) {
	reqParams := &vo.RequestRegister{}
	err := ctx.ShouldBindJSON(reqParams)
	exception.AssertException(ctx, exception.NewException(exception.ValidateError, err.Error()))
	var responseJson interface{}
	err = rpcx.CallFuncService("User", "Register", &reqParams, &responseJson)
	exception.AssertException(ctx, exception.NewException(exception.ValidateError, err.Error()))
	response.WriteResponse(ctx, responseJson)
}

func LoginLogs(ctx *gin.Context) {
	reqParams := &vo.RequestLoginLogQuery{}
	err := ctx.ShouldBindQuery(reqParams)
	exception.AssertException(ctx, exception.NewException(exception.ValidateError, err.Error()))
	//var responseJson []dto.ResponseLoginLog
	//err = rpcx.CallFuncService("User", "GetLoginLogs", &reqParams, &responseJson)
	//exception.AssertException(ctx, err)
	//response.WriteResponse(ctx, responseJson)
}

func UpdateUser(ctx *gin.Context) {
	updateJson := &vo.RequestUpdateUserInfo{}
	err := ctx.ShouldBindJSON(updateJson)
	exception.AssertException(ctx, exception.NewException(exception.ValidateError, err.Error()))
	uid, err := strconv.Atoi(ctx.Param("uid"))
	exception.AssertException(ctx, exception.NewException(exception.ValidateError, err.Error()))
	formMapData := map[string]interface{}{}
	formMapData["Uid"] = uint(uid)
	birTime, err := time.Parse("2006-01-02 15:04:05", updateJson.Birthday)
	exception.AssertException(ctx, exception.NewException(exception.ValidateError, err.Error()))

	formMapData["Birthday"] = birTime
	formMapData["Avatar"] = updateJson.Avatar
	formMapData["Phone"] = updateJson.Phone
	formMapData["Grade"] = updateJson.Grade
	var responseJson interface{}

	err = rpcx.CallFuncService("User", "UpdateUserInfo", &formMapData, &responseJson)
	exception.AssertException(ctx, err)
	response.WriteResponse(ctx, responseJson)
}
func ChangePassWord(ctx *gin.Context) {
	postBody := vo.RequestUpdatePassWord{}
	err := ctx.ShouldBindJSON(&postBody)
	exception.AssertException(ctx, exception.NewException(exception.ValidateError, err.Error()))

	uid, err := strconv.Atoi(ctx.Param("uid"))
	exception.AssertException(ctx, exception.NewException(exception.ValidateError, err.Error()))
	formMapData := map[string]interface{}{}
	formMapData["Uid"] = uid
}

// LogOut
//
//	@Description: 用户退出登陆 删除用户存在token 中的数据
//	@param ctx
func LogOut(ctx *gin.Context) {
	uid := ctx.Param("uid")
	uintUid, err := strconv.Atoi(uid)
	exception.AssertException(ctx, exception.NewException(exception.ValidateError, err.Error()))

	var delReply interface{}
	err = rpcx.CallFuncService("User", "LogOut", map[string]int{
		"Uid": uintUid,
	}, &delReply)
	exception.AssertException(ctx, exception.NewException(exception.ValidateError, err.Error()))
	response.WriteResponse(ctx, nil)
}
