package main

import (
	"os"
	"testing"
)

//Test always start with a UPPERCASE - Test
func TestNewDeck(t *testing.T) {

	d := newDeck()
	if len(d) != 32 {
		t.Errorf("Error: Expected deck length 32 but got %v", len(d))
	}

	if d[0] != "Pik Ass" {

		t.Errorf("Expected Pik Ass but got %v", d[0])
	}

	if d[len(d)-1] != "Karo Sieben" {

		t.Errorf("Expected Karo Sieben but got %v", d[len(d)-1])
	}

}

// Long Name because its easy to find the test for a function (Strg + F search Name of Function)
func TestSafeToDeckAndNewDeckTestFromFile(t *testing.T) {

	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadDeck := newDeckFromFile("_decktesting")

	if len(loadDeck) != 32 {

		t.Errorf("Expected 32 cards in deck, got %v", len(loadDeck))

	}

}
