package v3

import (
	"adminVideos/models"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
)

func UploadImages(c *gin.Context){
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status" :400,
			"error": err,
			"data": "错误请求",
		})
		return
	}
	//文件的名称
	filename := header.Filename
	//创建空文件
	out, err := os.Create("public/upload/images/"+filename)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status" :400,
			"error": err,
			"data": "文件创建失败",
		})
		return
	}
	defer out.Close()
	// 将文件复制进空文件
	_, err = io.Copy(out, file)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status" :400,
			"error": err,
			"data": "复制文件失败",
		})
		return
	}
	//封面表添加
<<<<<<< HEAD
	id,err := models.CreatedImage("http://127.0.0.1:3000/images/"+filename)
=======
	id,err := models.CreatedImage("http://127.0.0.1:8000/images/"+filename)
>>>>>>> 1f71234280818950b8805706b9a2acdfed7fe285
	if err != nil{
		c.JSON(http.StatusOK, gin.H{
			"status" :400,
			"error": err,
			"data": "封面添加数据库失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status" :200,
		"error": nil,
		"data": id,
	})
}

func UploadVideos(c *gin.Context){
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status" :400,
			"error": err,
			"data": "错误请求",
		})
		return
	}
	//文件的名称
	filename := header.Filename
	//创建文件
	out, err := os.Create("public/upload/videos/"+filename)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status" :400,
			"error": err,
			"data": "文件创建失败",
		})
		return
	}

	defer out.Close()
	// 将文件复制进空文件
	_, err = io.Copy(out, file)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status" :400,
			"error": err,
			"data": "复制文件失败",
		})
		return
	}
	// 返回视频路径id
<<<<<<< HEAD
	id,err := models.CreatedVideoSrc("http://127.0.0.1:3000/videos/"+filename)
=======
	id,err := models.CreatedVideoSrc("http://127.0.0.1:8000/videos/"+filename)
>>>>>>> 1f71234280818950b8805706b9a2acdfed7fe285
	if err != nil{
		c.JSON(http.StatusOK, gin.H{
			"status" :400,
			"error": err,
			"data": "视频路径添加数据库失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status" :200,
		"error": nil,
		"data": id,
	})
}