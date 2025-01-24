package controller

import (
	"github.com/gin-gonic/gin"
	"taro-api/models"
)

func GetTarotReading(c *gin.Context) {
	arr, count := models.TarotCardReading()
	t := models.TarotRequest{[6]models.TarotCard(arr), uint8(count)}
	c.JSON(200, t)
}
