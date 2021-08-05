package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

/*
Write a program which reads information from a file and represents it in a slice of structs.
Assume that there is a text file which contains a series of names. Each line of the text file
has a first name and a last name, in that order, separated by a single space on the line.

Your program will define a name struct which has two fields, fname for the first name, and
lname for the last name. Each field will be a string of size 20 (characters).

Your program should prompt the user for the name of the text file. Your program will
successively read each line of the text file and create a struct which contains the first
and last names found in the file. Each struct created will be added to a slice, and after
ll lines have been read from the file, your program will have a slice containing one struct
for each line in the file. After reading all lines from the file, your program should iterate
through your slice of structs and print the first and last names found in each struct.
*/

type Name struct {
	fname string
	lname string
}

func getFileName() string {
	var fileName string
	fmt.Println("Enter the name of the text file:")
	_, err := fmt.Scan(&fileName)

	if err != nil {
		return ""
	}

	return fileName
}

func main() {
	fileName := getFileName()

	data, _ := ioutil.ReadFile(fileName)
	namesArr := strings.Split(string(data), "\n")
	names := make([]Name, 0)

	for i := 0; i < len(namesArr); i++ {
		if len(namesArr[i]) > 0 {
			var name Name
			n := strings.Split(namesArr[i], " ")

			name.fname, name.lname = n[0], n[1]
			names = append(names, name)
		}
	}

	for _, v := range names {
		fmt.Println("First Name:", v.fname, "Last Name:", v.lname)
	}
}
