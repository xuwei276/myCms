package datasource

import (
	"fmt"
	"github.com/go-xorm/xorm"
	"myapp/model"
)

func NewMysqlEngine() *xorm.Engine{
	engine, err := xorm.NewEngine("mysql", "root:@tcp(127.0.0.1:3306)/test?charset=utf8")

	ip := engine.Ping()
	fmt.Println(ip)
	ers := engine.Sync2(new(model.Admin))
	if ers != nil{
		panic(err.Error())
	}
	engine.ShowSQL(true)
	engine.SetMaxOpenConns(10)

	return engine
}