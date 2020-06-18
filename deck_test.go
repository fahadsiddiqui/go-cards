package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected length 52, we got %v", len(d))
	}

	if d[0] != "Ace of Diamonds" {
		t.Errorf("Expected first element to be Ace of Diamonds, we got %v", d[0])
	}

	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("Expected last element to be King of Clubs, we got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_deck_testing")

	deck := newDeck()
	deck.saveToFile("_deck_testing")

	loadedDeck := newDeckFromFile("_deck_testing")

	if len(loadedDeck) != 52 {
		t.Errorf("Expected length 52, we got %v", len(loadedDeck))
	}

	os.Remove("_deck_testing")
}
