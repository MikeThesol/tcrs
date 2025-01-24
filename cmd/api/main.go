package main

import (
	"github.com/gin-gonic/gin"
	"taro-api/controller"
)

func main() {
	r := gin.Default()
	r.Static("/images", "./images")
	r.GET("/tarot", controller.GetTarotReading)
	r.Run(":3000")
}
