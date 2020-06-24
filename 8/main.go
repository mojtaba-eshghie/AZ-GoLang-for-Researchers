package main

import "fmt"

type bot interface {
	getGreeting() string
}
type englishBot struct{}
type spanishBot struct{}

func (eb englishBot) getGreeting() string {
	// We assume the getGreeting function's logic if very different from each other here
	return "Hello!"
}

func (sb spanishBot) getGreeting() string {
	return "Hola!"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}
