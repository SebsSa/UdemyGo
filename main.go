package main

import "fmt"

func main() {

	cards := newDeck()
	fmt.Println(cards.toString())
	cards.saveToFile("my_Cards")

}
