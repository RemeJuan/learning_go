package main

import (
	"encoding/json"
	"fmt"
)

func getName() string {
	var name string
	fmt.Println("Please enter your name:")

	_, err := fmt.Scan(&name)

	if err != nil {
		return ""
	}

	return name
}

func getAddress() string {
	var address string
	fmt.Println("Please enter your address:")

	_, err := fmt.Scan(&address)

	if err != nil {
		return ""
	}

	return address
}

func main() {
	person := make(map[string]string)

	person["name"] = getName()
	person["address"] = getAddress()

	formatted, _ := json.Marshal(person)

	fmt.Println(string(formatted))
}
