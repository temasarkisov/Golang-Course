/*package main

import (
	"fmt"
	"sort"
	"strconv"
)

func printSlice(slice []int) {
	fmt.Print("Slice: ")
	for _, val := range slice {
		fmt.Printf("%d ", val)
	}
	fmt.Print("\n")
}

func main() {
	var input string
	var check int = 1
	slice := []int{}

	for check != 0 {
		fmt.Print("Enter value:\n")
		check, _ = fmt.Scan(&input)
		if check == 1 {
			value, err := strconv.Atoi(input)
			if err == nil {
				slice = append(slice, value)
				sort.Ints(slice)
				printSlice(slice)
			} else if input == "X" {
				fmt.Print("Thank you, bye!\n")
				return
			}
		} else {
			fmt.Print("Incorrect input!\n")
			return
		}
	}
}
*/