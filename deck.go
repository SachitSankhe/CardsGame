package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

/*
*Create a new type of 'deck' which is a slice of strings
 */
type deck []string

// todo : A createDeck function that creates and returns a new deck
func createDeck() deck {
	cards := deck{}

	//* Creating 2 slices such that one contains suits and other contains cards
	cardsSuits := [4]string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardsNumbers := [13]string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight",
		"Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardsSuits {
		for _, number := range cardsNumbers {
			cards = append(cards, number+" of "+suit)
		}
	}

	return cards
}

// todo: Always keep the reciever name as 1 or 2 letter word (Standard)
func (deckInstance deck) print() {
	for index, card := range deckInstance {
		fmt.Println(index+1, card)
	}
}

// todo : Create a deal function that deals a deck of a specified handSize and return it along with the remaining cards deck
func deal(currentDeck deck, handSize int) (deck, deck) {
	return currentDeck[:handSize], currentDeck[handSize:]
}

// todo : Convert a deck to a single string
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

// todo: Save a string or technically a byte slice to a file on hard-drive
func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

// todo: Retrieve a deck from a file stored on hard-drive
func newDeckFromDrive(filename string) deck {
	cardsBytes, err := readDataFromFile(filename)

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	return deck(strings.Split(string(cardsBytes), ","))

}
func readDataFromFile(filename string) ([]byte, error) {
	return os.ReadFile(filename)
}

func (d deck) shuffle() {
	size := len(d)
	cards := []string(d)

	for index, card := range cards {
		cards[index] = cards[rand.Intn(size-1)]
		cards[rand.Intn(size-1)] = card
	}
}
