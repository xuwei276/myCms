package controller

import (
	"encoding/json"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"
	"myapp/service"
)

type  AdminController struct {
	Ctx iris.Context
	Servic service.AdminServic
	Session *sessions.Session
}

const (
	ADMIN = "admin"
)

type AdminLogin struct {
	UserName string `form:"user_name"`
	PassWord string `form:"password"`
}


func (ac *AdminController)PostLogin(context iris.Context)mvc.Result{
	iris.New().Logger().Info("admin login")

	var adminLogin AdminLogin
	ac.Ctx.ReadForm(&adminLogin)


	if adminLogin.UserName == "" || adminLogin.PassWord == ""{
		return  mvc.Response{
			Object:map[string]interface{}{
				"status":"8",
				"success":"登陆失败",
				"message":"用户名密码为空,请重新填写后尝试登陆",
			},
		}
	}

	admin,exist := ac.Servic.GetByAdminNamePassword(adminLogin.UserName, adminLogin.PassWord)

	if !exist {
		return mvc.Response{
			Object:map[string] interface{}{
				"status":"0",
				"success":"登陆失败",
				"message":"用户名密码错误",
			},
		}
	}

	userByte,_ := json.Marshal(admin)
	ac.Session.Set(ADMIN,userByte)

	return mvc.Response{
		Object:map[string] interface{}{
			"status":"1",
			"success":"登陆成功",
			"message":"管理员登陆成功",
		},
	}
}

