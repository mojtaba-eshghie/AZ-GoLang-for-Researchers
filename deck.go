package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

//let's define a receiver for this deck type
func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

func (d deck) d_append() {
	d = append(d, "WElcom!")
	d.print()
}

func (d deck) return_string() string {
	return "hello world!"
}

func newDeck() deck {
	deck_of_cards := deck{}

	card_suits := []string{"Spades", "Diamonds", "Hearts"}
	card_values := []string{"Ace", "Two", "Three"}

	for _, suit := range card_suits {
		for _, value := range card_values {
			deck_of_cards = append(deck_of_cards, value+" of "+suit)
		}
	}

	return deck_of_cards
}

func to_string(d deck) string {
	return strings.Join(d, ", ")
}

func (d deck) saveToFile(filename string) error {
	my_string := to_string(d)
	my_byte := []byte(my_string)
	return ioutil.WriteFile(filename, my_byte, 0666)
}

func readFromDeckFile(filename string) deck {
	byte_slice, err := ioutil.ReadFile(filename)
	if err != nil {
		//other than just returning, we could also do some logging stuff here.
		//could quit the entire program
		//could just put the value of the err code to the console output (and the log) then return an empty deck
		//could silently do nothing (which is the counterintuitive approach to computer science!)
		/*
			fmt.Println("Error happened in retrieving the file contents, err is: ", err)
			return deck{}
		*/
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	my_string := string(byte_slice)
	my_slice := strings.Split(my_string, ", ")
	return my_slice
}

func (d deck) shuffle() deck {

	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	length := len(d)
	for i := 0; i < length; i++ {
		// j := rand.Intn(length)
		j := r.Intn(length)

		//do the swapping operation here
		/*
			temp := d[j]
			d[j] = d[i]
			d[i] = temp
		*/

		d[i], d[j] = d[j], d[i]
	}

	return d
}
