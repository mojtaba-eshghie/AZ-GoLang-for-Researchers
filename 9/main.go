package main

import "fmt"

type Animal interface {
	makeSound() string
}

type Dog struct{}
type Cat struct{}

func (dog Dog) makeSound() string {
	return "Wooof!"
}

func (cat Cat) makeSound() string {
	return "Meeeew!"
}

func playTheSound(animal Animal) {
	fmt.Println(animal.makeSound())
}

func main() {
	/*
		Now, I can find out about the exact kind of animal at the runtime, then give
		it to the playTheSound method which will result in playing the particular sound
		we found about its type in runtime (not hardcoded into the struct itself)
	*/

	animalArray := make([]Animal, 10)

	animalArray[0] = Dog{}
	animalArray[1] = Cat{}

	for _, animal := range animalArray {
		if animal != nil {
			playTheSound(animal)
		}
	}
}
