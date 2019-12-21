package service

import (
	"github.com/go-xorm/xorm"
	"myapp/model"
)

type AdminServic interface {

	GetByAdminNamePassword(username,password string)(model.Admin,bool)

}

func NewAdminService(db *xorm.Engine)AdminServic  {
	return &adminServic{engine:db,}
}

type adminServic struct {
	engine *xorm.Engine
}

func (ac *adminServic) GetByAdminNamePassword(username,password string)(model.Admin,bool) {
	var admin model.Admin
	ac.engine.Where("admin_name = ? and pwd = ?",username,password).Get(&admin)
	return  admin,admin.AdminId != 0
}