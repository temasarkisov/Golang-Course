/*package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	foodEaten        string
	locomotionMethod string
	spokenSound      string
}

func (animal *Animal) AnimalInit(foodEaten string, locomotionMethod string, spokenSound string) {
	animal.foodEaten = foodEaten
	animal.locomotionMethod = locomotionMethod
	animal.spokenSound = spokenSound
}

func (animal Animal) Eat() {
	fmt.Printf("%s\n", animal.foodEaten)
}

func (animal Animal) Move() {
	fmt.Printf("%s\n", animal.locomotionMethod)
}

func (animal Animal) Speak() {
	fmt.Printf("%s\n", animal.spokenSound)
}

func main() {
	var cow, bird, snake Animal
	animalProc := make(map[string]func())

	// Initialization
	cow.AnimalInit("grass", "walk", "moo")
	bird.AnimalInit("worms", "fly", "peep")
	snake.AnimalInit("mice", "slither", "hsss")
	animalProc["cow eat"] = cow.Eat
	animalProc["cow move"] = cow.Move
	animalProc["cow speak"] = cow.Speak
	animalProc["bird eat"] = bird.Eat
	animalProc["bird move"] = bird.Move
	animalProc["bird speak"] = bird.Speak
	animalProc["snake eat"] = snake.Eat
	animalProc["snake move"] = snake.Move
	animalProc["snake speak"] = snake.Speak

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("%s", ">")
		requestString, err := reader.ReadString('\n')
		if err != nil {
			return
		}
		requestString = strings.ToLower(requestString[:len(requestString)-1])
		animalProc[requestString]()
	}
}
*/