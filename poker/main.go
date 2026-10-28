package main

import (
	"fmt"
	"math/rand/v2"
	"sort"
	"strings"
)

type card struct {
	rank int    // Ace 1/11 (depending on situation), Jack, Queen, King are all worth 10
	suit string // e.g. Spades
}

const (
	highcard      = iota // 0
	onepair              // 1
	twopair              // 2
	threekind            // 3
	straight             // 4
	flush                // 5
	fullhouse            // 6
	fourkind             // 7
	straightflush        // 8
)

var turn string
var i int
var soft bool = false
var q bool = false
var suits = []string{"♠", "♥", "♦", "♣"}
var deck, ph, dh []card
var phv, dhv []int // hand values
var b string = "\x1b[1m"
var n int = 0

func main() {

	for q == false {
		clear()
		if n == 0 {
			deck = newd()
			ph, dh = nil, nil
			n++
			for range 5 {
				ph = append(ph, deck[0])
				deck = deck[1:]
				dh = append(dh, deck[0])
				deck = deck[1:]
			}
		}
		fmt.Printf("\n")
		phv = eval(ph)
		dhv = eval(dh)
		fmt.Printf("Your hand:\n\n")
		//		fmt.Printf("%v\n\n\n", ph)
		prhand(ph, false)
		//		fmt.Printf("%d\n", phv)
		fmt.Printf("Dealer hand:\n\n")
		//		fmt.Printf("%v\n\n\n", dh)
		prhand(dh, false) //  prints the hand
		fmt.Printf("\n")
		//		fmt.Printf("%d\n", dhv)

		var inp string
		fmt.Printf("(R)edraw, (e)nd turn, (f)old, or (q)uit: ")
		fmt.Scanln(&inp)
		fmt.Printf("\n")
		switch inp {
		case "r", "R":
			var shit string
			fmt.Printf(": ")
			fmt.Scanln(&shit)
			redraw(shit)

		case "e", "E":
			continue
		case "q", "Q":
			q = true
		case "f", "F":
			dealer()
			n = 0
			clear()
			fmt.Printf("\n")
			phv = eval(ph)
			dhv = eval(dh)
			fmt.Printf("Your hand:\n\n")
			//		fmt.Printf("%v\n\n\n", ph)
			prhand(ph, false)
			//		fmt.Printf("%d\n", phv)
			fmt.Printf("Dealer hand:\n\n")
			//		fmt.Printf("%v\n\n\n", dh)
			prhand(dh, false) //  prints the hand
			fmt.Printf("\n")

			//			fmt.Printf("%v\n%v\n", phv[0], dhv[0])
			var inp byte
			// only pussies fold
			switch {
			case phv[0] > dhv[0]:
				fmt.Printf("You win\nClick enter to continue")
				fmt.Scanln(&inp)
				if inp != 0 {
					continue
				}
			case phv[0] < dhv[0]:
				fmt.Printf("Dealer wins\nClick enter to continue")
				fmt.Scanln(&inp)
				if inp != 0 {
					continue
				}
			default:
				win := compare(phv, dhv)
				switch win {
				case 1:
					fmt.Printf("You win\n")
				case -1:
					fmt.Printf("Dealer wins\n")
				default:
					fmt.Printf("Tie\n")
				}
				fmt.Printf("Click enter to continue")
				fmt.Scanln(&inp)
				if inp != 0 {
					continue
				}
			}

		}
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

func eval(user []card) []int {
	cunts := make(map[int]int)
	suits := make(map[string]int)
	ranks := make([]int, 0, len(user))
	for _, i := range user {
		cunts[i.rank]++
		suits[i.suit]++
		ranks = append(ranks, i.rank)

	}
	sort.Ints(ranks)

	// flush
	hasFlush := false
	for _, v := range suits {
		if v == 5 {
			hasFlush = true
			break
		}
	}

	hasStraight := false
	if len(ranks) == 5 {
		hasStraight = true
		for i := 1; i < len(ranks); i++ {
			if ranks[i] != ranks[i-1]+1 {
				hasStraight = false
				break
			}
		}
	}

	pairs := 0
	three := false
	four := false

	for _, count := range cunts {
		switch count {
		case 2:
			pairs++
		case 3:
			three = true
		case 4:
			four = true
		}
	}
	switch {
	case hasFlush && hasStraight:
		return []int{straightflush}
	case four:
		return []int{fourkind}
	case three && pairs == 1:
		return []int{fullhouse}
	case hasFlush:
		return []int{flush}
	case hasStraight:
		return []int{straight}

	case three:
		return []int{threekind}
	case pairs == 2:
		return []int{twopair}
	case pairs == 1:
		for rank, count := range cunts {
			if count == 2 {
				return []int{onepair, rank}
			}
		}
	case true:
		return []int{highcard}
	}
	return []int{highcard}

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

func clear() {
	// \x1b[H moves cursor to top left, \x1b[2J clears the screen
	fmt.Print("\x1b[H\x1b[2J")
}

func stats() {

}

func checkcard(targ []card) card {
	lcard := targ[len(targ)-1]

	return lcard
}

func (c card) lines() [3]string {
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

	width := 3

	leftp := (width - len(rank)) / 2
	center := strings.Repeat(" ", leftp) + rank + strings.Repeat(" ", width-len(rank)-leftp)

	return [3]string{
		c.suit + strings.Repeat(" ", width-1), center, strings.Repeat(" ", width-1) + c.suit,
	}
}

func prhand(hand []card, hidden bool) {
	back := [4]string{"?  ", " ? ", "  ?"}

	for row := range 3 {
		for _, c := range hand {
			color := "\x1b[47;30m" // black default
			if hidden {
				color := "\x1b[47;30m" // black on white for middle row
				if row == 0 || row == 2 {
					color = "\x1b[47;31m" // red on white background
				}
				fmt.Printf("%s%s%s\x1b[0m ", b, color, back[row])
				continue
			}
			if c.suit == "♥" || c.suit == "♦" {
				color = "\x1b[47;31m"
			}
			fmt.Printf("%s%s%s\x1b[0m ", b, color, c.lines()[row])
		}
		fmt.Printf("\n")
	}
}

func redraw(shit string) {
	for _, char := range shit {
		if char < '1' || char > '5' {
			fmt.Printf("Input integers 1-5\n")
			return
		} else {
			i := int(char - '1')
			ph[i] = deck[0]
			deck = deck[1:]
		}

	}
}

func compare(a, b []int) int {
	// a and b are from eval
	if len(a) == 0 || len(b) == 0 {
		return 0
	}
	//compare hand catagory

	if a[0] != b[0] {
		if a[0] > b[0] {
			return 1
		}
		return -1
	}

	// same catagory. ties
	for i := 1; i < len(a) && i < len(b); i++ {
		if a[i] != b[i] {
			if a[i] > b[i] {
				return 1
			}
			return -1
		}
	}

	return 0
}

func dealer() {
	for range 3 {
		score := eval(dh)

		if score[0] == highcard {
			// look for near flush or straight
			if hfd(dh) || hsd(dh) {
				ddraw(1)
				return
			}
			ddraw(3)
			return
		}
		switch score[0] {
		case highcard:
			ddraw(3)
			continue

		case onepair:
			pairr := score[1]
			if pairr <= 6 {
				ddraw(2)
				continue
			} else if pairr <= 10 {
				ddraw(1)
			} else {
				return
			}
			continue
		case twopair:
			return
		case threekind:
			return
		case straight:
			return
		case flush:
			return
		case fullhouse:
			return
		case fourkind:
			return
		case straightflush:
			return
		default:
			// two pair tripple, straight, flush, etc
			return
		}
	}
}

func ddraw(count int) {
	// pick the shittiest cards and redraw them
	for range count {
		weak := 0
		for i := 1; i < len(dh); i++ {
			if dh[i].rank < dh[weak].rank {
				weak = i
			}
		}

		dh[weak] = deck[0]
		deck = deck[1:]
	}
}

func hfd(hand []card) bool {
	scunts := map[string]int{}
	for _, aksfdhklasjddhksfkljsdhksffslahsakjldffgyusalhdkjffjskjdldfhsjsahsjdkfsaadfgjhahsdfhk := range hand {
		scunts[aksfdhklasjddhksfkljsdhksffslahsakjldffgyusalhdkjffjskjdldfhsjsahsjdkfsaadfgjhahsdfhk.suit]++
	}
	for _, v := range scunts {
		if v >= 4 {
			return true
		}
	}
	return false
}

func hsd(hand []card) bool {
	ranks := make([]int, len(hand))
	for i, c := range hand {
		ranks[i] = c.rank
	}
	sort.Ints(ranks)

	for i := 0; i <= len(ranks)-4; i++ {
		window := ranks[i : i+4]

		if window[3]-window[0] == 3 {
			return true
		}
	}
	return false
}
