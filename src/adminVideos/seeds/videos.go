package main

import (
	"github.com/jinzhu/gorm"
)

type Videos struct {
	gorm.Model
	Name string `gorm:"column:name;type:varchar;"json:"name"` //视频名称

	Date string  `gorm:"column:date;type:varchar;"json:"date"` //上映时间

	VideoPersons VideoPersons `gorm:"FOREIGNKEY:VideoPersonsId;"json:"video_persons_id"` //演员表
	VideoPersonsId int `gorm:"column:video_persons_id;type:integer;"json:"video_persons_id"`

	Introduce string  `gorm:"column:introduce;type:varchar;"json:"introduce"` //视频简介

	CommentId int `gorm:"column:comment_id;type:integer;"json:"comment_id"` // 评论 一对多
	Comment []Comment `gorm:"FOREIGNKEY:VideosId;"json:"actor_id"`

	ImageId int `gorm:"column:image_id;type:integer;"json:"image_id"` //引入视频封面表
	Image Image `gorm:"FOREIGNKEY:ImageId;"json:"image"` //封面 belongs to 视频

	SrcId int `gorm:"column:src_id;type:integer;"json:"src_id"` //引入视频路径表
	Src Src `gorm:"FOREIGNKEY:SrcId;"json:"src_id"` //视频路径

	ClassesId int `gorm:"column:classes_id;type:integer;"json:"classes_id"`  //分类表的id
	Classes Classes `gorm:"FOREIGNKEY:ClassesId;"json:"classes"`

}

type Image struct {
	gorm.Model
	Path string
}
//视频路径
type Src struct {
	gorm.Model
	Patch string `gorm:"column:path;type:varchar;"json:"path"`
}

func ForVideoImage(){
	ImagePath := []string{
		"http://127.0.0.1:8000/static/images/akali.jpg",
		"http://127.0.0.1:8000/static/images/jie.jpg",
		"http://127.0.0.1:8000/static/images/yasuo.jpg",
		"http://127.0.0.1:8000/static/images/daomei.jpg",
	}

	for i:=0;i< len(ImagePath);i++{
		db.Create(&Image{
			Path: ImagePath[i],
		})
	}

}

func ForVideos(){
	Name := []string{
		"山村老尸",
		"复仇者联盟",
		"人在囧途",
		"功夫",
		"死亡笔记",
		"杀人回忆",
	}
	Date := []string{
		"1999-12-12",
		"2018-12-09",
		"2016-05-15",
		"2012-02-23",
		"2006-03-03",
		"2007-11-13",
	}
	VideoPersonId := []int{
		1,2,3,4,5,6,
	}
	Introduce := []string{
		"讲述的是小明（海俊杰）的四个朋友在黄山村玩过一次“招魂游戏”后先后离奇身亡，因直觉事有蹊跷，小明向当记者的姐姐CISSY（黎姿）求助，为获独家新闻，...",
		"“复仇者联盟”应运而生。他们各显神通，团结一心，终于战胜了邪恶势力，保证了地球的安全",
		"讲述了玩具集团老板李成功（徐峥饰）和讨债的挤奶工牛耿（王宝强饰）前往长沙的旅程故事...",
		"讲述了一个街头混混变成一代武术家，一个本性善良的黑帮流氓改邪归正的故事",
		"天才与天才的较量",
		"讲述了韩国的一个小镇上，出现了一系列的连环杀人案，小镇警察和从汉城调来的苏警察共同办案的故事",
	}
	ClassesId := []int{
		1,2,3,4,5,6,
	}
	for i:=0;i<6;i++{
		db.Create(&Videos{
			Name:Name[i],
			Date:Date[i],
			VideoPersonsId:VideoPersonId[i],
			Introduce:Introduce[i],
			ClassesId:ClassesId[i],
		})
	}
}


