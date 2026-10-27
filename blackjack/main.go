package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

type card struct {
	rank int    // Ace 1/11 (depending on situation), Jack, Queen, King are all worth 10
	suit string // e.g. Spades
}

var turn string
var i int
var soft bool = false
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
	//	phv = eval(ph)
	//	dhv = eval(dh)
	//	fmt.Printf("Player hand value: %v\nDealer hand value: %v\n\n", phv, dhv)
	//
	//	fmt.Printf("Player hand: ")
	//	for _, c := range ph {
	//		fmt.Print(c.String(), " ")
	//	}
	//	fmt.Printf("(%v)\n", phv)
	//
	//	fmt.Printf("Dealer hand: ")
	//	for _, c := range dh {
	//		fmt.Print(c.String(), " ")
	//	}
	//	fmt.Printf("(%v)\n", dhv)
	//	var inp string
	//	fmt.Printf("Type h/s to hit/stand: ")
	//	fmt.Scanln(&inp)
	//	fmt.Printf("[DEBUG] INP VAR: %v, %t, %q", inp, inp, inp)
loop:
	for phv <= 21 {
		turn = "p"
		//		phv = eval(ph)
		//		dhv = eval(dh)
		//		fmt.Printf("\nPlayer hand value: %v\nDealer hand value: %v\n\n", phv, dhv)
		//
		//		fmt.Printf("Player hand: ")
		//		for _, c := range ph {
		//			fmt.Print(c.String(), " ")
		//		}
		//		fmt.Printf("(%v)\n", phv)
		//
		//		fmt.Printf("Dealer hand: ")
		//		for _, c := range dh {
		//			fmt.Print(c.String(), " ")
		//		}
		//		fmt.Printf("(%v)\n", dhv)
		var inp string
		i++
		if i == 1 {
			stats()
		}
		fmt.Printf("Type h/s to hit/stand: ")
		fmt.Scanln(&inp)
		switch inp {
		case "h", "H":
			//			fmt.Printf("\n[i] Caught input correctly\n")
			hit("ph")
			phv = eval(ph)
			clear()
			stats()
			fmt.Printf("You drew a %v\n", checkcard(ph))
			//			fmt.Printf("[i] jumped to hit\n")
			if phv > 21 {
				fmt.Printf("You busted\n")
				time.Sleep(250 * time.Millisecond)
				clear()
				stats()
			} else {
				continue
			}
			//			fmt.Printf("[i] handled eval\n")

		case "s", "S":
			break loop

		case "q", "Q":
			return
		}

		// DEALER TURN

	}
	turn = "d"
	for dhv < 18 {
		time.Sleep(1000 * time.Millisecond)
		if dhv > 21 {
			fmt.Printf("Dealer busted\n")
			break
		}
		//		fmt.Printf("[i] passed 'if'\n")
		if dhv == 17 {
			//			fmt.Printf("[i] fell into '==17'\n")
			for _, i := range dh {
				if i.rank == 1 {
					soft = true
					//					fmt.Printf("[i] soft is true\n")
				}
			}
		}
		if dhv < 17 || soft == true {
			hit("dh")
			clear()
			stats()
			dhv = eval(dh)
			fmt.Printf("\nDealer drew a %v\n", checkcard(dh))
			soft = false
			//			fmt.Printf("[i] dhv < 17 / soft == true\n")

			//			fmt.Printf("[i] should eval (to %v)\n ", dhv)
		} else {
			//			fmt.Printf("[i] hit else and broke\n")
			break
		}
	}
	if phv <= 21 && dhv <= 21 {
		if phv > dhv {
			clear()
			stats()

			fmt.Printf("You won\n")
		} else if phv < dhv {
			clear()
			stats()

			fmt.Printf("You lost\n")
		} else if phv == dhv {
			clear()
			stats()

			fmt.Printf("Tie (push)\n")
		}
	} else if phv > 21 && dhv > 21 {
		clear()
		stats()

		fmt.Printf("You both suck (push)\n")
	} else if phv > 21 && dhv <= 21 {
		clear()
		stats()

		fmt.Printf("You lost\n")
	} else if phv <= 21 && dhv > 21 {
		clear()
		stats()

		fmt.Printf("You won\n")
	}
}

func newd() []card {
	deck := make([]card, 0, 52)

	for _, suit := range suits {
		for rank := 1; rank <= 13; rank++ {
			deck = append(deck, card{rank: rank, suit: suit})
		}
	}
	for range 5 {
		shuf(deck)
	}
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
	if len(deck) == 0 {
		return
	}

	TC := deck[0]

	switch fuckhead {
	case "ph":
		ph = append(ph, TC)
	case "dh":
		dh = append(dh, TC)
	}
	deck = deck[1:]
}
func clear() {
	// \x1b[H moves cursor to top left, \x1b[2J clears the screen
	fmt.Print("\x1b[H\x1b[2J")
}

func stats() {
	phv = eval(ph)
	dhv = eval(dh)
	//	fmt.Printf("\nPlayer hand value: %v\nDealer hand value: %v\n\n", phv, dhv)

	fmt.Printf("\n\n\n")
	fmt.Printf("Player hand: ")
	for _, c := range ph {
		fmt.Print(c.String(), " ")
	}
	fmt.Printf("(%v)\n", phv)

	fmt.Printf("Dealer hand: ")
	if turn == "d" {
		for _, c := range dh {
			fmt.Print(c.String(), " ")
		}
	} else {
		for f, c := range dh {
			if f == 0 {
				fmt.Print("?? ")
				f++
			} else {
				fmt.Print(c.String(), " ")
			}

		}
	}
	if turn == "p" {
		fmt.Printf("(%v)\n", eval(dh[1:]))
	} else {
		fmt.Printf("(%v)\n", eval(dh))
	}
}

func checkcard(targ []card) card {
	lcard := targ[len(targ)-1]

	return lcard
}
