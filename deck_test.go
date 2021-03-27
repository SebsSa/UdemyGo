package main

import "testing"

//Test always start with a UPPERCASE - Test
func TestNewDeck(t *testing.T) {

	d := newDeck()
	if len(d) != 32 {
		t.Errorf("Error: Expected deck length 32 but got %v", len(d))
	}

}
