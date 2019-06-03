package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

//第一级表 电影 动漫 电视剧 等等
type Classes struct {
	gorm.Model
	Name string `gorm:"column:name;type:varchar;"json:"name"validate:"required||string"`
	Pid string `gorm:"column:pid;type:integer;"json:"pid"validate:"required||int"`
}


// 创建菜单
func CreatedClass(c Classes)(id int,err error){
	cc := Classes{}
	query := Db.Raw("insert into classes(name,pid) values(?,?) returning id",&c.Name,&c.Pid).Scan(&cc)
	if err:=query.Error; err!=nil{
		fmt.Println("创建失败",err)
		return
	}
	id = cc.ID
	return
}



