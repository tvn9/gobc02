package main

import (
	"fmt"
	"os"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Case Insensitive Search
//
//  Allow for case-insensitive searching
//
// EXAMPLE
//  Let's say that the user runs the program like this:
//    go run main.go LAZY
//
//  Or like this: go run main.go lAzY
//  Or like this: go run main.go lazy
//
//  For all cases above, the program should find
//  the "lazy" keyword.
// ---------------------------------------------------------

const (
	usage = `Word Finder
Program save a number of word in a wordSlice, 
you job is to enter a query word, the program will display the matching word found in memory.

Usage example: command [word]
`
	// Declare and words string
	wordStr = "" + "nhu co bac ho trong ngay vui dai thang vua ngo ra bi no go u dau nhu co bac ho"
)

func main() {

	// Read the user query from command-line
	query := os.Args[1:]
	l := len(query)

	// Verify user input syntax
	if l < 1 {
		fmt.Println(usage)
		return
	}

	// Convert words string to words lice
	wordS := strings.Fields(wordStr)
	// fmt.Printf("%T %s\n", wordS, wordS)

	// Compare user query to word slice
	// queryies:
	for _, q := range query {
		q = strings.ToLower(q)
		// fmt.Printf("%d %s\n", i, q)
		for j, w := range wordS {
			w = strings.ToLower(w)
			if q == w {
				fmt.Printf("query for %q... We found a match #%d %q\n", q, j+1, w)
				// break // queryies
				// continue // queryies
			}
		}
	}
}
