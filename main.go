package main

func main() {
	cards := deck{newCard(), "Axe of Diamonds"}
	cards = append(cards, "Six of Spades")

	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
