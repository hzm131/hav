package main

import ("github.com/jinzhu/gorm"
_"github.com/lib/pq")
/*
	这是添加测试数据的包
*/
var db *gorm.DB
func main(){
	var err error
	db,err = gorm.Open("postgres","user=postgres dbname=postgres password=123456 sslmode=disable")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	ForUsers() //创建用户
	//ForVideoImage() //创建上传视频的封面
	ForVideos() //创建上传视频
	ForClass()

	ForVideoPerson() //添加演员
}

