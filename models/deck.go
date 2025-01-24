package models

import (
	"math/rand"
)

type deck struct {
	Cards []TarotCard
}

const maxValueForReading = 6

func newTarotDeck() deck {
	majorCardsArr := GetAllMajorArcanaArray()
	mapOfImagesPath := GetMajorArcanaCardImagePaths()

	var cards []TarotCard

	for _, name := range majorCardsArr {
		cards = append(cards, TarotCard{
			Name:      name,
			Arcana:    Major,
			ImagePath: mapOfImagesPath[name],
		})
	}

	for i := 0; i < len(cards); i++ {
		MakeQualityForCard(&cards[i])
	}

	shuffle(&cards)

	return deck{cards}
}

func shuffle(cards *[]TarotCard) {
	rand.Shuffle(len(*cards), func(i, j int) {
		(*cards)[i], (*cards)[j] = (*cards)[j], (*cards)[i]
	})
}

func TarotCardReading() []TarotCard {
	dk := newTarotDeck().Cards

	newDeck := make([]TarotCard, maxValueForReading)

	for i := 0; i < maxValueForReading; i++ {
		newDeck[i] = dk[i]
	}
	return newDeck
}
