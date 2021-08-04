package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
)

const Quit = "x"

func main() {
	var input string
	items := make([]int, 0, 3)

	for {
		fmt.Println("Please enter an integer value (Type 'x' to Quit):")
		_, err := fmt.Scan(&input)

		if err != nil {
			return
		}

		if input == Quit {
			fmt.Println("Thanks, bye...")
			os.Exit(0)
			return
		} else {
			parsed, _ := strconv.ParseInt(input, 10, 64)
			items = append(items, int(parsed))
			sort.Ints(items)
			fmt.Println(items)
		}
	}

}
