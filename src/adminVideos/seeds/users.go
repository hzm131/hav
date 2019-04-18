package main

import (
	"github.com/jinzhu/gorm"
)

type Users struct {
	gorm.Model
	UserName string `gorm:"column:username;type:varchar;not null;"json:"username"`
	PassWord string `gorm:"column:password;type:varchar;not null;"json:"password"`
}
type Persons struct {
	gorm.Model
	Users Users `gorm:"FOREIGNKEY:UsersId;"json:"users_id"`
	UsersId uint `gorm:"column:users_id;type:integer;not null"json:"users_id"`
	Age int `gorm:"column:age;type:integer;"json:"age"` //年龄
	Sex string `gorm:"column:sex;type:varchar;"json:"sex"` //性别
	Birthday string `gorm:"column:birthday;type:varchar;"json:"birthday"` //生日
	NickName string  `gorm:"column:nick_name;type:varchar;"json:"nick_name"` //昵称
	Vip int `gorm:"column:vip;type:integer;"json:"vip"`  //会员
}
func CreatedUsers(user Users,person Persons){
	db.Create(&user).First(&user)
	userId := user.ID
	db.Exec("insert into persons(users_id,age,sex,birthday,nick_name,vip) values(?,?,?,?,?,?)",userId,person.Age,person.Sex,person.Birthday,person.NickName,person.Vip)
}
func ForUsers(){
	username := []string{
		"13122179125","891453178@qq.com","914152174@qq.com","13166049681","78456874@qq.com","461841521@qq.com","4145212@qq.com","56415414@qq.com",
	}
	password := []string{
		"123456","544654","4545564","7552389","511245","542241","54252","785563",
	}


	age := []int{
		18,20,25,40,62,41,25,15,
	}
	sex := []string{
		"男","女","男","女","男","女","男","女",
	}
	birthday := []string{
		"1995-11-23","1996-10-13","1998-01-03","1985-01-02","1996-8-17","1975-04-29","1987-12-20","1981-02-21",
	}
	nickname := []string{
		"雷瑟守备","班德尔城","德玛西亚","洛克萨斯","艾欧尼亚","黑色玫瑰","皮尔特沃夫","巨神峰",
	}

	for i:=0;i<len(username);i++{
		CreatedUsers(Users{
			UserName: username[i],
			PassWord:password[i],
		},Persons{
			Age: age[i],
			Sex:sex[i],
			Birthday:birthday[i],
			NickName:nickname[i],
			Vip: 0,
		})
	}
}