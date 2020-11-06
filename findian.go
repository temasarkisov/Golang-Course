/*
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Print("Enter string:\n")

	reader := bufio.NewReader(os.Stdin)
	someString, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	someString = strings.ToLower(someString[:len(someString)-1])

	if strings.HasPrefix(strings.ToLower(someString), "i") &&
		strings.HasSuffix(strings.ToLower(someString), "n") &&
		strings.Contains(strings.ToLower(someString), "a") {
		fmt.Print("Found!\n")
	} else {
		fmt.Print("Not Found!\n")
	}
}
*/