package models

import
(
	"github.com/jinzhu/gorm"
	_"github.com/lib/pq"
)

/*
	此包的其他文件 去定义表 以及表关系 操作表的方法封装
*/
var Db *gorm.DB

func init(){
	var err error
	Db,err = gorm.Open("postgres","user=postgres dbname=postgres password=123456 sslmode=disable")
	if err != nil {
		panic(err)
	}
	//创建表 自动迁移
	Db.AutoMigrate(&Users{},&Persons{})
}

func CloseDB() {
	defer Db.Close()
}
