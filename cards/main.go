package main

//import "fmt"

func main() {
	// var card string = "Ace of Spades"
	//card := newCard() /* we use this when you are declaring your variable for one item*/

	/* for your var to be responsible for multiple items
	we make use of this 👇👇*/
	cards := deck {"Ace of diamonds", newCard()} /* we replace []string with type deck*/
			/*👆[]string (slice of string)*/
	/* How to add new element to your slice 👇👇*/
	cards = append(cards, "Ace of spades" )
	cards = append(cards, "six of spades" )

	//the range function is used to iterate inside of a slice
/*	for i, card := range cards {
		fmt.Println(i, card)
	} */ //this line of code was sent to function as a receiver 
	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}

// NB: GO IS NOT AN OBJECT ORIENTED PROGRAMMING LANGUAGE.
//SO THERE IS NO IDEA OF CLASS IN GO
