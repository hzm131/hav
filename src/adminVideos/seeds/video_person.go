package main

import "github.com/jinzhu/gorm"

//影视人员
type VideoPersons struct {
	gorm.Model
	Director string `gorm:"column:director;type:varchar;not null;"json:"director"` //导演
	Actor []Actor `gorm:"FOREIGNKEY:ActorId;"json:"actor_id"`
	ActorId int `gorm:"column:actor_id;type:integer;not null;"json:"actor_id"`
}

//演员
type Actor struct {
	gorm.Model
	Starring1 string `gorm:"column:starring1;type:varchar;"json:"starring1"` //主演
	Starring2 string `gorm:"column:starring2;type:varchar;"json:"starring2"`
	Starring3 string `gorm:"column:starring3;type:varchar;"json:"starring3"`
	Lead1 string `gorm:"column:lead1;type:varchar;"json:"lead1"` //领衔主演
	Lead2 string `gorm:"column:lead2;type:varchar;"json:"lead2"`
	Lead3 string `gorm:"column:lead3;type:varchar;"json:"lead3"`
}

func ForVideoPerson(){
	Director:=[]string{
		"梁鸿华","乔斯·韦登","徐峥","周星驰","大场鸫","奉俊昊",
	}
	Starring1 := []string{
		"吴镇宇","小罗伯特·唐尼","徐峥","周星驰","松山健一","宋康昊",
	}
	Starring2 := []string{
		"黎姿","史蒂芬·罗杰斯","王宝强","梁小龙","藤原龙也","金相庆",
	}
	Starring3 := []string{
		"海俊杰","克里斯.埃文斯","黄渤","黄圣依","户田惠梨香","金雷河",
	}
	ActorId := []int{
		1,2,3,4,5,6,
	}
	for i:=0;i<6 ;i++  {
		db.Create(&Actor{
			Starring1:Starring1[i],
			Starring2:Starring2[i],
			Starring3:Starring3[i],
		})
		db.Create(&VideoPersons{
			Director:Director[i],
			ActorId:ActorId[i],
		})
	}
}