package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type Videos struct {
	gorm.Model
	Name string `gorm:"column:name;type:varchar;"json:"name"validate:"required||string"` //视频名字
	VideoSrcId int `gorm:"column:video_src_id;type:integer;"json:"video_src_id"validate:"required||integer"` //视频路径
	VideoSrc VideoSrc  `gorm:"FOREIGNKEY:SrcId"json:"video_src"`

	ImageSrcId int `gorm:"column:image_src_id;type:integer;"json:"image_src_id"validate:"required||integer"`
	ImageSrc ImageSrc `gorm:"FOREIGNKEY:ImageSrcId"json:"image_src"`

	SecondaryId int `gorm:"column:secondary_id;type:integer;"json:"secondary_id"validate:"required||integer"`
	Secondary Secondary `gorm:"FOREIGNKEY:SecondaryId"json:"secondary"`
}

//上传视频封面
type ImageSrc struct {
	gorm.Model
	SrcPath string `gorm:"column:src_path;type:varchar;"json:"src_path"validate:"required||string"`
}

// 上传视频路径
type VideoSrc struct {
	gorm.Model
	SrcPath string `gorm:"column:src_path;type:varchar;"json:"src_path"validate:"required||string"`
}


//创建视频封面返回id
func CreatedImage(str string) (imgId int,err error){
	img_id := ImageSrc{}
	find := Db.Raw("insert into image_srcs(src_path) values(?) returning id",str).Scan(&img_id)
	if err:= find.Error; err!=nil{
		fmt.Println("创建失败",err)
		return 0,err
	}
	imgId = img_id.ID
	return
}

func CreatedVideoSrc(str string)(v_src_id int,err error){
	video_src := VideoSrc{}
	find := Db.Raw("insert into video_srcs(src_path) values(?) returning id",str).Scan(&video_src)
	if err:=find.Error; err!=nil{
		fmt.Println("创建失败",err)
		return 0,err
	}
	//创建成功后返回id
	v_src_id = video_src.ID  //拿到id
	return
}