/*package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type AnimalT struct {
	foodEaten        string
	locomotionMethod string
	spokenSound      string
}

type AnimalI interface {
	Eat()
	Move()
	Speak()
	AnimalInit()
}

func (animal AnimalT) Eat() {
	fmt.Printf("%s\n", animal.foodEaten)
}

func (animal AnimalT) Move() {
	fmt.Printf("%s\n", animal.locomotionMethod)
}

func (animal AnimalT) Speak() {
	fmt.Printf("%s\n", animal.spokenSound)
}

func QueryProc(animal AnimalT, action string) {
	switch action {
	case "eat":
		animal.Eat()
	case "move":
		animal.Move()
	case "speak":
		animal.Speak()
	default:
		fmt.Println("Incorrect request! Valid actions include: 'eat', 'move', 'speak'")
	}
}

func main() {
	animals := make(map[string]AnimalT)
	animals["cow"] = AnimalT{"grass", "walk", "moo"}
	animals["bird"] = AnimalT{"worms", "fly", "peep"}
	animals["snake"] = AnimalT{"mice", "slither", "hsss"}

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(">")
		requestString, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		requestString = strings.ToLower(requestString[:len(requestString)-1])
		requestSlice := strings.Split(requestString, " ")

		switch requestSlice[0] {
		case "newanimal":
			animals[requestSlice[1]] = animals[requestSlice[2]]
			fmt.Println("Created it!")
		case "query":
			QueryProc(animals[requestSlice[1]], requestSlice[2])
		default:
			fmt.Println("Incorrect request!")
		}
	}
}
*/