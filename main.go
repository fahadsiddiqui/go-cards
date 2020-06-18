package main

func main() {
	// cards := newDeck()

	// hand, remainingCards := deal(cards, 44)

	// hand.print()
	// fmt.Println("-------------")
	// remainingCards.print()

	// fmt.Println("-------------")
	// fmt.Println(remainingCards.toString())

	// remainingCards.saveToFile("./myRemainingCards.txt")
	deck1 := newDeckFromFile("./myRemainingCards.txt")

	deck1.print()
	deck1.shuffle()

	deck1.print()
}
