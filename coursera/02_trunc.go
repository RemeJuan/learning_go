package main

import (
	"fmt"
	"os"
	"strconv"
)

const Exit = "exit"

func main() {
	var input string

	for {
		fmt.Println("Please enter a float value (Type 'exit' to Quit):")
		_, err := fmt.Scan(&input)

		if err != nil {
			return
		}

		if input == Exit {
			fmt.Println("Thanks, bye...")
			os.Exit(0)
			return
		} else {
			float, _ := strconv.ParseFloat(input, 32)
			intVal := int32(float)
			fmt.Println(intVal)
		}
	}

}
