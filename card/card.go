package card

import (
	"math/rand"
)

func PickFirstCards() []string {
	// 今はここでrandの初期化を行っているが、最終的にはmain packageで行うようにする。
	// テストのsetupでrandの初期化をしているのでここは不必要
	// rand.Seed(time.Now().UnixNano())

	hands := make([]string, 5)

	for i := 0; i < 5; {
		key := cardKeys[rand.Intn(52)]
		if cards[key] == false {
			hands[i] = key
			cards[key] = true
			i++
			continue
		}
	}

	return hands
}

func RepickCards(hands []string, changeQueryInts []int) []string {
	// 今はここでrandの初期化を行っているが、最終的にはmain packageで行うようにする。
	// テストのsetupでrandの初期化をしているのでここは不必要
	// rand.Seed(time.Now().UnixNano())
	for _, hand := range hands {
		cards[hand] = true
	}
	if len(changeQueryInts) == 0 {
		return hands
	}

	for i := 0; i < len(changeQueryInts); {
		key := cardKeys[rand.Intn(52)]
		if cards[key] == false {
			hands[changeQueryInts[i]] = key
			cards[key] = true
			i++

			continue
		}
	}

	return hands
}

var cards map[string]bool = map[string]bool{
	"Heart A":       false,
	"Heart 2":       false,
	"Heart 3":       false,
	"Heart 4":       false,
	"Heart 5":       false,
	"Heart 6":       false,
	"Heart 7":       false,
	"Heart 8":       false,
	"Heart 9":       false,
	"Heart 10":      false,
	"Heart Jack":    false,
	"Heart Queen":   false,
	"Heart King":    false,
	"Diamond A":     false,
	"Diamond 2":     false,
	"Diamond 3":     false,
	"Diamond 4":     false,
	"Diamond 5":     false,
	"Diamond 6":     false,
	"Diamond 7":     false,
	"Diamond 8":     false,
	"Diamond 9":     false,
	"Diamond 10":    false,
	"Diamond Jack":  false,
	"Diamond Queen": false,
	"Diamond King":  false,
	"Club A":        false,
	"Club 2":        false,
	"Club 3":        false,
	"Club 4":        false,
	"Club 5":        false,
	"Club 6":        false,
	"Club 7":        false,
	"Club 8":        false,
	"Club 9":        false,
	"Club 10":       false,
	"Club Jack":     false,
	"Club Queen":    false,
	"Club King":     false,
	"Spade A":       false,
	"Spade 2":       false,
	"Spade 3":       false,
	"Spade 4":       false,
	"Spade 5":       false,
	"Spade 6":       false,
	"Spade 7":       false,
	"Spade 8":       false,
	"Spade 9":       false,
	"Spade 10":      false,
	"Spade Jack":    false,
	"Spade Queen":   false,
	"Spade King":    false,
}

var cardKeys []string = []string{
	"Heart 1",
	"Heart 2",
	"Heart 3",
	"Heart 4",
	"Heart 5",
	"Heart 6",
	"Heart 7",
	"Heart 8",
	"Heart 9",
	"Heart 10",
	"Heart Jack",
	"Heart Queen",
	"Heart King",
	"Diamond 1",
	"Diamond 2",
	"Diamond 3",
	"Diamond 4",
	"Diamond 5",
	"Diamond 6",
	"Diamond 7",
	"Diamond 8",
	"Diamond 9",
	"Diamond 10",
	"Diamond Jack",
	"Diamond Queen",
	"Diamond King",
	"Club 1",
	"Club 2",
	"Club 3",
	"Club 4",
	"Club 5",
	"Club 6",
	"Club 7",
	"Club 8",
	"Club 9",
	"Club 10",
	"Club Jack",
	"Club Queen",
	"Club King",
	"Spade 1",
	"Spade 2",
	"Spade 3",
	"Spade 4",
	"Spade 5",
	"Spade 6",
	"Spade 7",
	"Spade 8",
	"Spade 9",
	"Spade 10",
	"Spade Jack",
	"Spade Queen",
	"Spade King",
}
