package main

import (
	"os"
	"testing"
)

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

func TestSaveToFileAndReadFromFile(t *testing.T){
	//*1 Removing any earlier files that may not have been deleted due to some error.
	err := os.RemoveAll("_decktesting")
	if err != nil {
		t.Errorf("Error in deleting earlier files, exact error is: %v",err)
		os.Exit(1)
	}

	//*2 Creating a new Deck
	deck := createDeck()

	//*3 Saving the deck to the hard-drive
	err = deck.saveToFile("_decktesting")

	if err != nil {
		t.Errorf("Error in saving the deck into a file, exact error is: %v",err)
		os.Exit(1)
	}

	//*4 Retrieving the deck from the hard-drive
	retrievedDeck := newDeckFromDrive("_decktesting")

	
	if len(retrievedDeck) != len(deck) {
		t.Errorf("Size of deck is not equal, generated Deck size is %v and retrieved Deck size is %v",len(deck),len(retrievedDeck))
		os.Exit(2)
	}

	//*5 Removing the decktesting file from hard-drive
	err = os.RemoveAll("_decktesting")
	if err != nil {
		t.Errorf("Error in deleting the current file, exact error is: %v",err)
		os.Exit(1)
	}
}