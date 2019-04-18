package main

import (
	"github.com/jinzhu/gorm"
)

//视频类型表
type Classes struct {
	gorm.Model
	Name  string  `gorm:"column:name;type:varchar;"json:"name"` // 类名
	ParentId int `gorm:"column:parent_id;type:integer;DEFAULT:0;"json:"parent_id"` // 父母分类的id  默认值是0
	ParentList  string  `gorm:"column:parent_list;type:varchar;DEFAULT:0;"json:"parent_list"` // 分类的层级关系，从最高级到自己 默认值是0
	Depth int  `gorm:"column:depth;type:integer;DEFAULT:1;"json:"depth"` // 深度  假设1,2,3  那么就是3层  默认值是1
	Status int  `gorm:"column:status;type:integer;DEFAULT:1;"json:"status"`  // 0是禁用  1是启用  默认值是1
	Priority string `gorm:"column:priority;type:varchar;DEFAULT:0;"json:"priority"`  //Priority  默认值是0
}
/*
	1 电影   2动画   3原创
*/
func ForClass(){
	name := []string{
		"恐怖","科幻","喜剧","动作","悬疑","犯罪",
	}
	ParentId := []int{
		1,1,1,1,1,1,
	}
	ParentList := []string{
		"1,1","1,2","1,3","1,4","1,5","1,6",
	}

	Depth := []int{
		2,2,2,2,2,2,
	}
	for i:=0;i<6;i++{
		db.Create(&Classes{
			Name: name[i],
			ParentId:ParentId[i],
			ParentList:ParentList[i],
			Depth:Depth[i],
		})
	}
}