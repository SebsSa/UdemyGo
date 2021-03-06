package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// Create a new type of 'deck'
// witch is a slice of strings
type deck []string

func newDeck() deck {

	cards := deck{}

	cardSuits := []string{"Pik", "Herz", "Kreuz", "Karo"}
	cardValues := []string{"Ass", "König", "Dame", "Bube", "Zehn", "Neun", "Acht", "Sieben"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {

			cards = append(cards, suit+" "+value)
		}
	}

	return cards

}

//resiver
func (d deck) print() {

	for i, card := range d {
		fmt.Println(i, card)
	}

}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {

	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)

}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile((filename))
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	return deck(strings.Split(string(bs), ","))
}
