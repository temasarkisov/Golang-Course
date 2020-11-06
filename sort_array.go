/*
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

func InputSlice() []int {
	intSlice := []int{} // Slice containing entered integers
	fmt.Print("Enter series of integers: ")
	reader := bufio.NewReader(os.Stdin)
	inputString, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	inputString = (inputString[:len(inputString)-1])
	if inputString[len(inputString)-1] == ' ' { // Preventing space as last symbol
		inputString = (inputString[:len(inputString)-1])
	}
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

func Min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func PartitionSlice(intSlice []int) [][]int {
	sliceOfSLices := [][]int{}
	limit := 0
	if len(intSlice)%4 == 0 {
		limit = len(intSlice) / 4
	} else {
		limit = (len(intSlice) / 4) + 1
	}
	for i := 0; i < len(intSlice); i += limit {
		subSlice := intSlice[i:Min(i+limit, len(intSlice))]
		sliceOfSLices = append(sliceOfSLices, subSlice)
	}
	return sliceOfSLices
}

func Swap(intSlice []int, i int) {
	temp := intSlice[i]
	intSlice[i] = intSlice[i+1]
	intSlice[i+1] = temp
}

func Sort(intSlice []int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Print("Sorting subarray: ")
	PrintSlice(intSlice)
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

func MergeSubSlices(sliceOfSLices [][]int) []int {
	intSlice := []int{}
	for _, slice := range sliceOfSLices {
		for _, value := range slice {
			intSlice = append(intSlice, value)
		}
	}
	return intSlice
}

func main() {
	wg := &sync.WaitGroup{}

	intSlice := InputSlice()
	fmt.Print("\n")
	sliceOfSlices := PartitionSlice(intSlice)
	wg.Add(4)
	go Sort(sliceOfSlices[0], wg)
	go Sort(sliceOfSlices[1], wg)
	go Sort(sliceOfSlices[2], wg)
	go Sort(sliceOfSlices[3], wg)
	wg.Wait()
	intSlice = MergeSubSlices(sliceOfSlices)
	fmt.Print("Entire sorted list : ")
	PrintSlice(intSlice)
	fmt.Print("\n")
}
*/