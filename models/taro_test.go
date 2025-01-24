package models

import "testing"

func TestGetAllMajorArcanaArray(t *testing.T) {
	expected := []string{
		"The Fool", "The Magician", "The High Priestess", "The Empress", "The Emperor",
		"The Hierophant", "The Lovers", "The Chariot", "Strength", "The Hermit",
		"Wheel of Fortune", "Justice", "The Hanged Man", "Death", "Temperance",
		"The Devil", "The Tower", "The Star", "The Moon", "The Sun",
		"Judgement", "The World",
	}

	result := GetAllMajorArcanaArray()

	if len(expected) != len(result) {
		t.Errorf("Expected length %d, got %d", len(expected), len(result))
	}

	for i, v := range expected {
		if result[i] != v {
			t.Errorf("Expected %s at index %d, got %s", expected[i], i, result[i])
		}
	}

}

func TestGetMajorArcanaCardImagePaths(t *testing.T) {
	expected := map[string]string{
		"The Fool":           imagePath + "The Fool",
		"The Magician":       imagePath + "The Magician",
		"The High Priestess": imagePath + "The High Priestess",
		"The Empress":        imagePath + "The Empress",
		"The Emperor":        imagePath + "The Emperor",
		"The Hierophant":     imagePath + "The Hierophant",
		"The Lovers":         imagePath + "The Lovers",
		"The Chariot":        imagePath + "The Chariot",
		"Strength":           imagePath + "Strength",
		"The Hermit":         imagePath + "The Hermit",
		"Wheel of Fortune":   imagePath + "Wheel of Fortune",
		"Justice":            imagePath + "Justice",
		"The Hanged Man":     imagePath + "The Hanged Man",
		"Death":              imagePath + "Death",
		"Temperance":         imagePath + "Temperance",
		"The Devil":          imagePath + "The Devil",
		"The Tower":          imagePath + "The Tower",
		"The Star":           imagePath + "The Star",
		"The Moon":           imagePath + "The Moon",
		"The Sun":            imagePath + "The Sun",
		"Judgement":          imagePath + "Judgement",
		"The World":          imagePath + "The World",
	}

	result := GetMajorArcanaCardImagePaths()

	if len(expected) != len(result) {
		t.Errorf("Expected length %d, got %d", len(expected), len(result))
	}

	for card, path := range expected {
		if result[card] != path {
			t.Errorf("Expected path for %s to be %s, got %s", card, path, result[card])
		}
	}
}
