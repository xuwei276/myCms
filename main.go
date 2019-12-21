package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"
	"myapp/config"
	"myapp/controller"
	"myapp/datasource"
	"myapp/service"
	"time"
)

//主入口
func main() {
	app := newApp()
	
	configation(app)
	mvcHand(app)

	confi := config.InitConfi()
	addr := ":"+confi.Port

	app.Run(
		iris.Addr(addr),
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithOptimizations,
	)
}

func newApp() *iris.Application{
	app := iris.New()

	//设置日志级别
	app.Logger().SetLevel("debug")

	app.StaticWeb("/static","./static")
	app.StaticWeb("/manage/static","./static")

	app.RegisterView(iris.HTML("./static",".html"))
	app.Get("/", func(c context.Context) {
		c.View("index.html")
	})
	return app
}
func configation(app *iris.Application)  {
	app.Configure(iris.WithConfiguration(iris.Configuration{
		Charset:"UTF-8",
	}))

	app.OnErrorCode(iris.StatusFound, func(c context.Context) {
		c.JSON(iris.Map{
			"errmsg":iris.StatusFound,
			"msg":"not found",
			"data":iris.Map{},
		})
	})
	app.OnErrorCode(iris.StatusInternalServerError, func(c context.Context) {
		c.JSON(iris.Map{
			"errmsg":iris.StatusInternalServerError,
			"msg":"interal error",
			"data":iris.Map{},
		})
	})
}
func mvcHand(app *iris.Application){
	sessManager := sessions.New(sessions.Config{Cookie: "sessincookie", Expires: 24 * time.Hour})

	engine := datasource.NewMysqlEngine() //映射数据库
	adminService := service.NewAdminService(engine) //数据库操作引擎

	admin := mvc.New(app.Party("/admin"))
	admin.Register(adminService,sessManager.Start)
	admin.Handle(new(controller.AdminController))

}
