/*
package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var name string
	var address string
	m := make(map[string]string)
	//mBack := make(map[string]string)

	fmt.Print("Please, enter name\n")
	_, err := fmt.Scan(&name)
	if err == nil {
		fmt.Print("Please, enter address\n")
		_, err := fmt.Scan(&address)
		if err == nil {
			m["name"] = name
			m["address"] = address
			mJson, err := json.Marshal(m)
			if err == nil {
				fmt.Println(string(mJson))
				fmt.Print("\n")
				//err := json.Unmarshal(mJson, &mBack)
				//if err == nil {
				//fmt.Print(mBack)
				//fmt.Print("\n")
				//}
			}
		}
	}
}
*/