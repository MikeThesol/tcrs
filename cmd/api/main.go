package main

import (
	"fmt"
	"taro-api/models"
)

func init() {

}

func main() {
	arr := models.TarotCardReading()
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

}
