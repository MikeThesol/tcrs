package models

type TarotRequest struct {
	Cards [6]TarotCard `json:"cards"`
	Score uint8        `json:"score"`
}
