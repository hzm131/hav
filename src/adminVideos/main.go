package main

import (
	"adminVideos/routers"
)




func main(){
	r := routers.InitRouter()

	r.Run(":3000")
}