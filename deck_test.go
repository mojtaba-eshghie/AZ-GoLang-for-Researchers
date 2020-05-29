package main

import (
	"fmt"
	"os"
	"testing"
)

/*
func Test(t *testing.T) {
	d := create_deck_of_cards()

	if len(d) != 9 {
		t.Errorf("Expected deck length of 20, but got %v", len(d))
	}
}
*/

func TestNewDeck(t *testing.T) {
	d := newDeck()
	failedSomething := false
	errorString := ""

	if len(d) != 9 {
		//t.Errorf("Expected deck length of 9, but got %v", len(d))
		failedSomething = true
		errorString = fmt.Sprintf("Error# Expected deck length of 9, but got %v", len(d))

	}
	if d[0] != "Ace of Spades" {
		failedSomething = true
		errorString = fmt.Sprintf("%v, \n Error# The first item is not correct", errorString)
	}

	if d[len(d)-1] != "Three of Hearts" {
		fmt.Println(d[len(d)-1])
		failedSomething = true
		errorString = fmt.Sprintf("%v, Error# The last item is not correct", errorString)
	}

	// at the end of the test, evaluate what has happened!
	if failedSomething {
		t.Errorf(errorString)
	}

}

func TestSaveToFileAndReadFromDeckFile(t *testing.T) {
	// Note: whenever you are writing a test in golang, you should take care of any cleanup
	// the go testing framework does absolutely nothing in this regard.
	// we need to test the functionality of file saving, but we also do not want to mess with the actual
	// files being created and handled by the deployed processes.
	// let's use something like _decktesting as the name of the file being created for testing (the _ here indicates this)

	// check for cleanup
	err := os.Remove("_decktesting")

	if err != nil {
		// means that we did not have a _decktesting which is fine!
		fmt.Println("No previous dirty test file found")
	} else {
		fmt.Println("Successfully deleted dirty test files")
	}

	//we need a deck to save
	dummy_deck := newDeck()
	err = dummy_deck.saveToFile("_decktesting")
	if err != nil {
		t.Errorf("Error# failed to save the _decktesting")
	}

	// now the saving is finished, we need to see if the file has been successfully stored!?
	// and also the contents are correct

	deck_content := readFromDeckFile("_decktesting")
	if len(deck_content) >= 100 {
		t.Errorf("Error# the contents of the file are bigger than expected")
	}

}
