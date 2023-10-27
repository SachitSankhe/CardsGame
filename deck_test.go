package main

import "testing"

// todo: A test for the createDeck function
func TestNewDeck(t *testing.T) {
	deck := createDeck()

	if len(deck) != 52 {
		t.Errorf("Expected length 52,got %v",len(deck))
	}
	if deck[0] != "Ace of Spades" {
		t.Errorf("Expected first card as Ace of Spades, got %v",deck[0])
	}
	if deck[len(deck)-1] != "King of Clubs" {
		t.Errorf("Expected last card as King of Clubs, got %v",deck[len(deck)-1])
	}
}