package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type Users struct {
	gorm.Model
	UserName string `gorm:"column:username;type:varchar;not null;"json:"username"validate:"required||string"`
	PassWord string `gorm:"column:password;type:varchar;not null;"json:"password"validate:"required||string"`
}

type Persons struct {
	gorm.Model
	Users Users `gorm:"FOREIGNKEY:UsersId;"json:"users"`
	UsersId int `gorm:"column:users_id;type:integer;not null"json:"users_id"`
	Age int `gorm:"column:age;type:integer;"json:"age"` //年龄
	Sex string `gorm:"column:sex;type:varchar;"json:"sex"` //性别
	Birthday string `gorm:"column:birthday;type:varchar;"json:"birthday"` //生日
	NickName string  `gorm:"column:nick_name;type:varchar;"json:"nick_name"` //昵称
	Vip int `gorm:"column:vip;type:integer;"json:"vip"`  //会员
}


func FindId(user Users) (int,error) {
	//根据用户名 密码查询用户 将查询到的结果封装在user结构中
	query:= Db.Raw("select id from users where username=? and password=? limit 1",&user.UserName,&user.PassWord).Scan(&user)
	if err:=query.Error; err!=nil{
		fmt.Println("用户名或密码有问题",err)
		return 0,err
	}
	id := user.ID
	var person Persons
	per := Db.Raw("SELECT id FROM persons WHERE users_id = ?",&id).Scan(&person)
	if err:=per.Error;err!=nil{
		fmt.Println("查询失败")
		return 0,err
	}
	PersonsId := person.ID
	return PersonsId,nil
}

//创建数据
func CreateData(user Users) (int,error){
	users := Users{}
	username := Users{}
	Db.Raw("select id from users where username=?",&user.UserName).Scan(&username)
	if username.ID > 0{
		fmt.Println("用户名已存在")
		return -1,nil
	}

	query := Db.Raw("insert into users(username,password) values(?,?) returning id",&user.UserName,&user.PassWord).Scan(&users)
	if err:=query.Error; err!=nil{
		fmt.Println("创建失败",err)
		return 0,err
	}
	PersonId,err := CreatePerson(users.ID)
	if PersonId == 0 || err !=nil{
		fmt.Println("获取persons_id失败")
		return 0,nil
	}
	return PersonId,nil
}

// 创建person表 返回person_id
func CreatePerson(id int) (int,error){
	create := Db.Exec("insert into persons(users_id) values(?)",id)
	if err:=create.Error;err!=nil{
		fmt.Println("创建失败")
		return 0,err
	}

	var person Persons
	per := Db.Raw("SELECT id FROM persons WHERE users_id = ?",id).Scan(&person)
	if err:=per.Error;err!=nil{
		fmt.Println("查询失败")
		return 0,err
	}
	PersonsId := person.ID
	return PersonsId,nil
}