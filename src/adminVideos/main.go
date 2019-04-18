package main

import (
	"adminVideos/routers"
)




func main(){
	r := routers.InitRouter()

	r.Run("192.168.2.219:8000")
}