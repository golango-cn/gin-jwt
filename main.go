package main

import (
	"github.com/gin-gonic/gin"
	"golango.cn/cc-gin/router"
	"log"
)

func main() {

	gin.SetMode(gin.DebugMode)
	r := router.InitRouter()
	if err := r.Run(":8000"); err != nil {
		log.Fatal(err)
	}

}
