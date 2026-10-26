package main

import (
	"fmt"
	"math/rand/v2"
)

type card struct {
	rank int    // Ace 1/11 (depending on situation), Jack, Queen, King are all worth 10
	suit string // e.g. Spades
}

var suits = []string{"♠", "♥", "♦", "♣"}
var deck, ph, dh []card
var phv, dhv int // hand values

func main() {
	deck = newd()
	//	ph := []card{}
	//	dh := []card{}

	for range 2 {
		ph = append(ph, deck[0])
		deck = deck[1:]
		dh = append(dh, deck[0])
		deck = deck[1:]
	}
	//	fmt.Printf("%v\nPH: %v\nDH: %v\n", len(deck), ph, dh)
	phv = eval(ph)
	dhv = eval(dh)
	fmt.Printf("Player hand value: %v\nDealer hand value: %v\n\n", phv, dhv)

	fmt.Printf("Player hand: ")
	for _, c := range ph {
		fmt.Print(c.String(), " ")
	}
	fmt.Printf("(%v)\n", phv)

	fmt.Printf("Dealer hand: ")
	for _, c := range dh {
		fmt.Print(c.String(), " ")
	}
	fmt.Printf("(%v)\n", dhv)
}

func newd() []card {
	deck := make([]card, 0, 52)

	for _, suit := range suits {
		for rank := 1; rank <= 13; rank++ {
			deck = append(deck, card{rank: rank, suit: suit})
		}
	}
	shuf(deck)
	return deck
}

func shuf(deck []card) {
	rand.Shuffle(len(deck), func(i, j int) {
		deck[i], deck[j] = deck[j], deck[i]
	})
}

func eval(shit []card) int {
	bullshit := 0
	aces := 0
	for _, c := range shit {
		//		fmt.Printf("Card %v: %v%v\n", c, c.suit, c.rank)
		if c.rank == 1 {
			bullshit += 11
			aces++
		} else if c.rank > 10 {
			bullshit += 10
		} else {
			bullshit += c.rank
		}
		for aces != 0 && bullshit > 21 {
			bullshit -= 10
			aces--
		}

		//		fmt.Printf("bullshit: %v\n", bullshit)
	}
	return bullshit

}

func (c card) String() string {
	rank := fmt.Sprint(c.rank)

	switch c.rank {
	case 1:
		rank = "A"
	case 11:
		rank = "J"
	case 12:
		rank = "Q"
	case 13:
		rank = "K"
	}
	return c.suit + rank
}

func hit(fuckhead string) {
	fuckhead = append(fuckhead, deck[0])
	deck = deck[1:]
}
