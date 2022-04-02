package main

import "fmt"

//type deck borrow all the bahavior
//of a slice of string
//:. we can replace []string with deck

type deck []string

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
