/*package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func InputSlice() []int {
	intSlice := []int{}
	fmt.Print("Enter a sequence of up to 10 integers\n")
	reader := bufio.NewReader(os.Stdin)
	inputString, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	inputString = (inputString[:len(inputString)-1])
	stringSlice := strings.Split(inputString, " ")
	for _, strValue := range stringSlice {
		intValue, err := strconv.Atoi(strValue)
		if err != nil {
			panic(err)
		}
		intSlice = append(intSlice, intValue)
	}

	return intSlice
}

func Swap(intSlice []int, i int) {
	temp := intSlice[i]
	intSlice[i] = intSlice[i+1]
	intSlice[i+1] = temp
}

func BubbleSort(intSlice []int) {
	for i := len(intSlice); i > 0; i-- {
		for j := 0; j < i-1; j++ {
			if intSlice[j] > intSlice[j+1] {
				Swap(intSlice, j)
			}
		}
	}
}

func PrintSlice(intSlice []int) {
	for _, value := range intSlice {
		fmt.Printf("%d ", value)
	}
	fmt.Print("\n")
}

func main() {
	intSlice := InputSlice()
	BubbleSort(intSlice)
	PrintSlice(intSlice)
}
*/