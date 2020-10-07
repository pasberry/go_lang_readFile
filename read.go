//Write a program which reads information from a file and represents it in a slice of structs.
//Assume that there is a text file which contains a series of names.
//Each line of the text file has a first name and a last name, in that order, separated by a single space on the line.
//Your program will define a name struct which has two fields, fname for the first name, and lname for the last name.
//Each field will be a string of size 20 (characters).
//Your program should prompt the user for the name of the text file. Your program will successively read each line
//of the text file and create a struct which contains the first and last names found in the file.
//Each struct created will be added to a slice, and after all lines have been read from the file,
//your program will have a slice containing one struct for each line in the file.
//After reading all lines from the file, your program should iterate through your slice of structs and print the first
//and last names found in each struct.

//Submit your source code for the program, “read.go”.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Name struct {
	fname string
	lname string
}

func main() {

	fmt.Println("Please enter the name of the file")

	inputScanner := bufio.NewScanner(os.Stdin)
	inputScanner.Scan()

	dataFileName := inputScanner.Text()
	
	dataFile, err := os.Open(dataFileName)
	
	if err != nil {
		fmt.Println("The file does not exist, exiting the program")
		os.Exit(0)
	}

	fileScanner := bufio.NewScanner(dataFile)

	names := make([]Name, 0)

	for fileScanner.Scan() {

		tokens := strings.Split(fileScanner.Text(), " ")

		names = append(names , Name{fname:tokens[0] , lname:tokens[1]})

	}
	dataFile.Close()

	for i := range names {

		fmt.Println("First Name: " + names[i].fname + "\nLast Name: " + names[i].lname + "\n")
	}
}
