package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type Role struct {
	gorm.Model
	Name string `json:"name"gorm:"column:name;type:varchar;"validate:"required||string"`
}

type RoleUser struct {
	gorm.Model
	RoleId int `gorm:"column:role_id;type:integer;"json:"role_id"validate:"required||int"`
	PersonId int `gorm:"column:person_id;type:integer;"json:"person_id"validate:"required||int"`
}

func createdRole(){
	Name := []string{
		"普通用户",
		"vip用户",
		"管理员",
	}
	for i:= 0;i<len(Name);i++{
		Db.Create(&Role{
			Name:Name[i],
		})
	}
}

func FindRole(Pid int) (role Role,err error){
	re := Db.Raw("select * from roles where name='普通用户'").Scan(&role)
	if err = re.Error; err!=nil{
		fmt.Println("查询不到",err)
		return
	}
	roleId := role.ID
	role_user := RoleUser{}
	query := Db.Raw("insert into role_users(role_id,person_id) values(?,?) returning id",roleId,Pid).Scan(&role_user)
	if err = query.Error; err!=nil{
		fmt.Println("创建失败",err)
		return
	}
	return
}