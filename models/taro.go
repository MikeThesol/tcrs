package models

import (
	"fmt"
)

type ArcanaType int

const (
	Major ArcanaType = iota
	Minor
)

const imagePath = "/images/"

type TarotCard struct {
	Name      string
	Arcana    ArcanaType
	ImagePath string
	Quality
	IsGood bool
}

func GetAllMajorArcanaArray() []string {
	majorArcanaNames := []string{
		"The Fool", "The Magician", "The High Priestess", "The Empress", "The Emperor",
		"The Hierophant", "The Lovers", "The Chariot", "Strength", "The Hermit",
		"Wheel of Fortune", "Justice", "The Hanged Man", "Death", "Temperance",
		"The Devil", "The Tower", "The Star", "The Moon", "The Sun",
		"Judgement", "The World",
	}

	return majorArcanaNames
}

func GetAllMinorArcanaArray() []string {
	minorArcanaNames := []string{
		// Масть Жезлов
		"Ace of Wands", "Two of Wands", "Three of Wands", "Four of Wands", "Five of Wands",
		"Six of Wands", "Seven of Wands", "Eight of Wands", "Nine of Wands", "Ten of Wands",
		"Page of Wands", "Knight of Wands", "Queen of Wands", "King of Wands",

		// Масть Мечей
		"Ace of Swords", "Two of Swords", "Three of Swords", "Four of Swords", "Five of Swords",
		"Six of Swords", "Seven of Swords", "Eight of Swords", "Nine of Swords", "Ten of Swords",
		"Page of Swords", "Knight of Swords", "Queen of Swords", "King of Swords",

		// Масть Кубков
		"Ace of Cups", "Two of Cups", "Three of Cups", "Four of Cups", "Five of Cups",
		"Six of Cups", "Seven of Cups", "Eight of Cups", "Nine of Cups", "Ten of Cups",
		"Page of Cups", "Knight of Cups", "Queen of Cups", "King of Cups",

		// Масть Пентаклей
		"Ace of Pentacles", "Two of Pentacles", "Three of Pentacles", "Four of Pentacles", "Five of Pentacles",
		"Six of Pentacles", "Seven of Pentacles", "Eight of Pentacles", "Nine of Pentacles", "Ten of Pentacles",
		"Page of Pentacles", "Knight of Pentacles", "Queen of Pentacles", "King of Pentacles",
	}

	return minorArcanaNames
}

func GetMajorArcanaCardImagePaths() map[string]string {
	cards := make(map[string]string)
	cardsArr := GetAllMajorArcanaArray()

	for i := 0; i < len(cardsArr); i++ {
		cards[cardsArr[i]] = imagePath + cardsArr[i] + ".jpg"
	}
	return cards
}

func (card TarotCard) Print() {
	fmt.Printf("Name: %s\nArcana: %v\nImage: %s\n", card.Name, card.Arcana, card.ImagePath)
}
