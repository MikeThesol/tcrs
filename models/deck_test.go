package models

/*
func TestNewTarotDeck(t *testing.T) {
	expectedMajorCards := GetAllMajorArcanaArray()
	expectedImagePaths := GetMajorArcanaCardImagePaths()

	actualDeck := NewTarotDeck()

	if len(actualDeck.Cards) != len(expectedMajorCards) {
		t.Errorf("Expected %d cards, got %d", len(expectedMajorCards),
			len(actualDeck.Cards))
	}

	for i, name := range expectedMajorCards {
		card := actualDeck.Cards[i]
		expectedImagePath := expectedImagePaths[name]

		if card.Name != name {
			t.Errorf("Expected card name %s, got %s", name, card.Name)
		}
		if card.Arcana != Major {
			t.Errorf("Expected arcana %v for card %s, got %v", Major, name, card.Arcana)
		}
		if card.ImagePath != expectedImagePath {
			t.Errorf("Expected image path %s for card %s, got %s",
				expectedImagePath, name, card.ImagePath)
		}

	}
}

*/
