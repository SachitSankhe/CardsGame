package main

import "fmt"

func main() {

	//* A new deck generated
	cards := createDeck()

	//* Printing the deck to the terminal
	fmt.Println("Cards from the deck:- ")
	cards.print()
	fmt.Println("----------------------------------")

	//* Dealing a set of cards out of a given handsize.
	dealedCards, _ := deal(cards, 5)
	fmt.Println("Dealed cards:- ")
	dealedCards.print()
	fmt.Println("----------------------------------")


	//* Saving a deck to the hard-drive with the given filename
	fmt.Println("Saving the deck to a file")
	cards.saveToFile("my_deck")
	fmt.Println("----------------------------------")


	//* Retrieving a deck from a file named as the passed filename and printing it
	fmt.Println("Retrieving a deck from a file ")
	cardsFromDrive := newDeckFromDrive("my_deck")
	cardsFromDrive.print()
	fmt.Println("----------------------------------")


	//* Shuffling the reciever deck and printing it
	fmt.Println("Shuffling the deck:- ")
	cards.shuffle()
	cards.print()
	fmt.Println("----------------------------------")
	

}
