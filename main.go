package main

func main() {
	cards := newDeck()
	cards.shuffle()

	//	cards := newDeckFromFile("my_cards")
	cards.print()
	//fmt.Println(cards.toString())
	//	hand, remainingCards := deal(cards, 5)
	//	hand.print()
	//	remainingCards.print()
	//
}
