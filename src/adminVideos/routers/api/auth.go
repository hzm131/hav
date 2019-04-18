package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"adminVideos/middleware/jwt"
)

func GetAuth(c *gin.Context){
	h := c.Request.Header.Get("authorization")
	if h == "" {
		fmt.Println("authorization不能为空")
		c.JSON(http.StatusOK, gin.H{
			"status" :401,
			"error": nil,
			"data": "authorization不能为空",
		})
		c.Abort()
		return
	}
	//还需要进一步验证token格式

	str := h[7:len(h)] //截取token
	//验证token
	jwt.ParseToken(c,str)
}
