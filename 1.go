package main

import "fmt"

func main() {
	// cards := []string{newCard(), newCard(), newCard()}
	// let's replace the []string with our newly created type which is deck!
	cards := deck{newCard(), newCard(), newCard()}
	cards = append(cards, "Six of Spades")

	cards.print()

	C := deck{"one"}

	C.d_append()

	new_deck := newDeck()
	fmt.Println(new_deck)

	d1, d2 := return_the_two_decks(deck{"A", "C"}, deck{"B", "D"})
	d1.print()
	d2.print()

	new_deck.saveToFile("temp")
	newly_read_deck := readFromDeckFile("temp")
	fmt.Println("=======================")
	fmt.Println(newly_read_deck)

	shuffled_deck := new_deck.shuffle()
	fmt.Println("**** Shuffled ****")
	fmt.Println(shuffled_deck)
}

func newCard() string {
	return "Five of diamonds"
}

func return_the_two_decks(d1 deck, d2 deck) (deck, deck) {
	return d1, d2
}
