package main

func main() {
	// cards := newDeck()

	// //multiple return values from one function
	// hand, remainingCards := deal(cards, 5)

	// hand.print()
	// remainingCards.print()

	cards := newDeck()
	cards.shuffle()
	cards.print()

	// //type conversion
	// greeting := "Hi there!"
	// fmt.Println([]byte(greeting))
}

//Arrays in Go are at a fixed length

//Slices however are arrays that can grow or shrink
