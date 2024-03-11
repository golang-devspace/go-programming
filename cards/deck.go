package main

import "fmt"

// Create a new type of 'deck'
// which is a slice of strings
type deck []string

// Function with receivers
// example receiver on function also called custom type
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
