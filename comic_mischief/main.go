package main

import (
	"fmt"
	"math"
	t "time"
)

func main() {
	var publisher, writer, artist, title string
	var year, pageNumber int
	var grade float64

	title = "Mr. GoToSleep"
	writer = "Tracey Hatchet"
	artist = "Jewel Tampson"
	publisher = "DizzyBooks Publishing Inc."
	year = 1997
	pageNumber = 14
	grade = 6.5

	printBookDetails(title, writer, artist, publisher, year, pageNumber, grade)

	title = "Epic Vol. 1"
	writer = "Ryan N. Shawn"
	artist = "Phoebe Paperclips"
	year = 2013
	pageNumber = 160
	grade = 9.0

	printBookDetails(title, writer, artist, publisher, year, pageNumber, grade)
}

func printBookDetails(
	title string, writer string, artist string, publisher string, year int, pageNumber int, grade float64,
) {
	var cost, age float64
	now := t.Now().Year()
	basePrice := 1.75

	age = float64(now - year)
	cost = basePrice * (age / grade)

	fmt.Println(title, "written by ", writer, " and drawn by ", artist)
	fmt.Println("Published by ", publisher, "in Year ", year)
	fmt.Println("Grading: ", grade)
	fmt.Println("Pages ", pageNumber)
	fmt.Println("Price: ", math.Round(cost*100)/100)
	fmt.Println()
}
