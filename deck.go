package main

import "fmt"

// Create a new type of 'deck'
// witch is a slice of strings
type deck []string

//resiver
func (d deck) print() {

	for i, card := range d {
		fmt.Println(i, card)
	}

}
