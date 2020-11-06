/*
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type person struct {
	fname string
	lname string
}

func main() {
	personsSlice := []person{}

	f, err := os.Open("test.txt")
	if err == nil {
		defer f.Close()
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			personInfo := strings.Split(scanner.Text(), " ")
			newPerson := person{fname: personInfo[0], lname: personInfo[1]}

			personsSlice = append(personsSlice, newPerson)
		}
		fmt.Print("Persons list:\n")
		for i, person := range personsSlice {
			fmt.Printf("Person's (%d) first name - %s last name - %s", i, person.fname, person.lname)
			fmt.Print("\n")
		}
	}
}
*/