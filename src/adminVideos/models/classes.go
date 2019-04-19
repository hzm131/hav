package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

//第一级表 电影 动漫 电视剧 等等
type Classes struct {
	gorm.Model
	Name string `gorm:"column:name;type:varchar;"json:"name"validate:"required||string"`
}


// 第二级表 电影种类 动漫种类
type Secondary struct {
	gorm.Model
	Name string `gorm:"column:name;type:varchar;"json:"name"validate:"required||string"`
	ClassesId int `gorm:"column:classes_id;type:integer;"json:"classes_id"validate:"required||integer"`
}




// 创建第一级分类
func CreatedClass(c Classes)(id int,err error){
	cc := Classes{}
	query := Db.Raw("insert into classes(name) values(?) returning id",&c.Name).Scan(&cc)
	if err:=query.Error; err!=nil{
		fmt.Println("创建失败",err)
		return 0,err
	}
	id = cc.ID
	return
}



//创建第二级分类
func CreatedSecondary(c Secondary)(id int,err error){
	cc := Secondary{}
	query := Db.Raw("insert into secondaries(name,classes_id) values(?,?) returning id",&c.Name,&c.ClassesId).Scan(&cc)
	if err:=query.Error; err!=nil{
		fmt.Println("创建失败",err)
		return 0,err
	}
	id = cc.ID
	return
}
