package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type Videos struct {
	gorm.Model
	Name string `gorm:"column:name;type:varchar;"json:"name"validate:"required||string"` //视频名字

	VideoSrcId int `gorm:"column:video_src_id;type:integer;"json:"video_src_id"validate:"required||integer"` //视频路径

	ImageSrcId int `gorm:"column:image_src_id;type:integer;"json:"image_src_id"validate:"required||integer"` //封面路径

	SecondaryId int `gorm:"column:secondary_id;type:integer;"json:"secondary_id"validate:"required||integer"` //父类id

	Area string  `gorm:"column:area;type:varchar;"json:"area"validate:"string"` //产地

	Language string  `gorm:"column:language;type:varchar;"json:"language"validate:"string"` //语言

	Released string `gorm:"column:released;type:varchar;"json:"released"validate:"string"` //上映日期

	Updated string `gorm:"column:updated;type:varchar;"json:"updated"validate:"string"` //更新日期

	Director string `gorm:"column:director;type:varchar;"json:"director"validate:"string"` //导演

	Actor string  `gorm:"column:actor;type:varchar;"json:"actor"validate:"string"` //演员

	Score string `gorm:"column:score;type:varchar;"json:"score"validate:"string"` // 评分

	Plot string  `gorm:"column:plot;type:varchar;"json:"plot"validate:"string"` // 剧情简介

	Screenshot string  `gorm:"column:screenshot;type:varchar;"json:"screenshot"validate:"string"` // 截图
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


//上传视频封面返回id
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
//上传视频返回id
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


//创建视频
func CreatedVideo(v Videos)(videoId int,err error){
	video := Videos{}
	create := Db.Raw(`insert into videos(name,video_src_id,image_src_id,
secondary_id,area,language,released,updated,director,actor,score,plot,screenshot) values(?,?,?,?,?,?,?,?,?,?,?,?,?) returning id`,
&v.Name,
&v.VideoSrcId,
&v.ImageSrcId,
&v.SecondaryId,
&v.Area,
&v.Language,
&v.Released,
&v.Updated,
&v.Director,
&v.Actor,
&v.Score,
&v.Plot,
&v.Screenshot,
).Scan(&video)
	if err:= create.Error; err!=nil{
		fmt.Println("创建失败",err)
		return 0,err
	}
	videoId = video.ID
	return
}

//更新视频
func UpdatedVideo(id int,v Videos) bool{
	fmt.Println("id:",id)
	if id == 0 {
		fmt.Println("id没传进来")
		return false
	}
	update := Db.Exec(`update videos set name=?,video_src_id=?,image_src_id=?,
secondary_id=?,area=?,language=?,released=?,updated=?,director=?,actor=?,score=?,plot=?,screenshot=? where id=?`,
		&v.Name,
		&v.VideoSrcId,
		&v.ImageSrcId,
		&v.SecondaryId,
		&v.Area,
		&v.Language,
		&v.Released,
		&v.Updated,
		&v.Director,
		&v.Actor,
		&v.Score,
		&v.Plot,
		&v.Screenshot,
		id,
	)
	if err:= update.Error; err!=nil{
		fmt.Println("更新有问题",err)
		return false
	}
	return true
}