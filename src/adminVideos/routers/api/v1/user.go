package v1

import (
	"adminVideos/middleware/jwt"
	"adminVideos/models"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/smokezl/govalidators"
	"io/ioutil"
	"net/http"
)

// 登录
func Login(c *gin.Context){
	validator := govalidators.New()
	user := models.Users{}
	value,err := ioutil.ReadAll(c.Request.Body)
	if err != nil{
		return
	}
	json.Unmarshal(value,&user)
	errList := validator.Validate(&user)
	if errList != nil {
		for _, err := range errList {
			c.JSON(http.StatusOK, gin.H{
				"status" :400,
				"error": err,
				"data":"json数据类型不匹配",
			})
		}
		return
	}
	PersonsId,Role,err := models.FindId(user)
	if err != nil || PersonsId == 0 || Role.Name == "" {
		fmt.Errorf("查询id失败",err)
		return
	}

	str,err := jwt.CreateJWT(PersonsId,Role)
	if err != nil {
		fmt.Errorf("失败")
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status" :200,
		"error": nil,
		"data": str,
		"role":Role.Name,
	})
}

//注册
func Registered(c *gin.Context){
	validator := govalidators.New()
	user := models.Users{}
	value,err := ioutil.ReadAll(c.Request.Body)
	if err != nil{
		return
	}
	json.Unmarshal(value,&user)
	errList := validator.Validate(&user)
	if errList != nil {
		for _, err := range errList {
			c.JSON(http.StatusOK, gin.H{
				"status" :400,
				"error": err,
				"data":"json数据类型不匹配",
			})
		}
		return
	}
	PersonsId,err := models.CreateData(user)
	if err != nil || PersonsId == 0{
		fmt.Errorf("查询返回id失败")
		return
	}
	if PersonsId == -1{
		c.JSON(http.StatusOK, gin.H{
			"status" :400,
			"error": "用户名已经存在",
		})
		return
	}
	fmt.Println(PersonsId)

	role,err := models.FindRole(PersonsId)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status" :400,
			"error": err,
			"data":"生成角色失败",
		})
		return
	}
	str,err := jwt.CreateJWT(PersonsId,role) //返回完整token
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status" :200,
			"error": err,
			"data":"jwt生成失败",
		})
		return
	}
	fmt.Println("打印完整的token:",str) //打印token
	c.JSON(http.StatusOK, gin.H{
		"status" :200,
		"error": nil,
		"data": str, //jwt
		"role":role.Name, //用户角色
	})
}