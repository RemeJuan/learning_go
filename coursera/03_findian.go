package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const Close = "exit"

func main() {
	var input string

	for {
		fmt.Println("Please enter any string (Type 'exit' to Quit):")
		scanner := bufio.NewScanner(os.Stdin)

		if input == Close {
			fmt.Println("Thanks, bye...")
			os.Exit(0)
			return
		} else if scanner.Scan() {
			input = scanner.Text()

			fmt.Println("Input: ", input)
			input = strings.ToLower(input)

			i := strings.HasPrefix(input, "i")
			n := strings.HasSuffix(input, "n")
			a := strings.Contains(input, "a")

			if i && n && a {
				fmt.Println("Found!")
			} else {
				fmt.Println("Not Found!")
			}
		}
	}

}
