package main

func main() {
	//card := newCard() // 型推論
	//fmt.Println(card)

	//cards := deck{"Ace of Diamonds", newCard()}
	//cards = append(cards, "Six of Spades") // 破壊的変更ではない

	//cards := newDeck()
	//hand, remainingCards := deals(cards, 5)
	//hand.print()
	//remainingCards.print()

	//greeting := "Hi there!"
	//fmt.Println([]byte(greeting))

	//cards := newDeck()
	//cards.saveToFile("my_cards")

	cards := newDeckFromFile("my_cards")
	cards.shuffle()
	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
