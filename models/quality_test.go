package models

import "testing"

func TestMakeQualityForCard(t *testing.T) {
	tc := TarotCard{
		Name:      "The Fool",
		Arcana:    Major,
		ImagePath: imagePath + "The_Fool",
	}

	MakeQualityForCard(&tc)

	if tc.Quality.GetMessageOfQuality() != TheFoolGood || tc.Quality.GetMessageOfQuality() != TheFoolBad {
		t.Errorf("Expected quality %s or %s, got %s", TheFoolGood, TheFoolBad, tc.Quality.GetMessageOfQuality())
	}
}
