package card_test

import (
	"fmt"
	"math/rand"
	"os"
	"testing"
	"time"

	"github.com/kimuson13/go-poker/card"
)

func setUp() {
	var t *testing.T
	if _, err := fmt.Println("setup"); err != nil {
		t.Fatal(err)
	}
	rand.Seed(time.Now().UnixNano())
}

func TestMain(m *testing.M) {
	setUp()
	ret := m.Run()
	os.Exit(ret)
}

// ランダムでカードを引く関数なので、カードがダブっていないかだけを確認するテスト
func TestPickFirstCards(t *testing.T) {
	cases := map[string]struct {
		mp   map[string]int
		want int
	}{
		"test_case_1": {
			mp:   make(map[string]int),
			want: 5,
		},
		"test_case_2": {
			mp:   make(map[string]int),
			want: 5,
		},
		"test_case_3": {
			mp:   make(map[string]int),
			want: 5,
		},
		"test_case_4": {
			mp:   make(map[string]int),
			want: 5,
		},
		"test_case_5": {
			mp:   make(map[string]int),
			want: 5,
		},
	}

	for name, c := range cases {
		c := c
		t.Run(name, func(t *testing.T) {
			// t.Parallel()
			hands := card.PickFirstCards()
			for _, hand := range hands {
				if _, ok := c.mp[hand]; ok {
					c.mp[hand]++
				} else {
					c.mp[hand] = 1
				}
			}
			if len(c.mp) != c.want {
				t.Errorf("PickFirst card want len(%d) map, but got len(%d) map", c.want, len(c.mp))
			}
		})
		// hands := card.PickFirstCards()
		// for _, hand := range hands {
		// 	if _, ok := c.mp[hand]; ok {
		// 		c.mp[hand]++
		// 	} else {
		// 		c.mp[hand] = 1
		// 	}
		// }

		// if len(c.mp) != c.want {
		// 	t.Errorf("Error in %s: this card is not unique. want length = %d, but bot %d", name, c.want, len(c.mp))
		// }
	}
}

func TestRepickCards(t *testing.T) {
	cases := map[string]struct {
		hands   []string
		handmap map[string]int
		input   []int
		want    int
	}{
		"Repick_one_card": {
			hands:   []string{"Heart 1", "Heart 2", "Heart 3", "Heart 4", "Heart 5"},
			handmap: map[string]int{"Heart 1": 1, "Heart 2": 1, "Heart 3": 1, "Heart 4": 1, "Heart 5": 1},
			input:   []int{0},
			want:    6,
		},
		"Repick_two_cards": {
			hands:   []string{"Heart 1", "Heart 2", "Heart 3", "Heart 4", "Heart 5"},
			handmap: map[string]int{"Heart 1": 1, "Heart 2": 1, "Heart 3": 1, "Heart 4": 1, "Heart 5": 1},
			input:   []int{0, 1},
			want:    7,
		},
		"Repick_three_cards": {
			hands:   []string{"Heart 1", "Heart 2", "Heart 3", "Heart 4", "Heart 5"},
			handmap: map[string]int{"Heart 1": 1, "Heart 2": 1, "Heart 3": 1, "Heart 4": 1, "Heart 5": 1},
			input:   []int{0, 1, 2},
			want:    8,
		},
		"Repick_four_cards": {
			hands:   []string{"Heart 1", "Heart 2", "Heart 3", "Heart 4", "Heart 5"},
			handmap: map[string]int{"Heart 1": 1, "Heart 2": 1, "Heart 3": 1, "Heart 4": 1, "Heart 5": 1},
			input:   []int{0, 1, 2, 3},
			want:    9,
		},
		"Repick_all_cards": {
			hands:   []string{"Heart 1", "Heart 2", "Heart 3", "Heart 4", "Heart 5"},
			handmap: map[string]int{"Heart 1": 1, "Heart 2": 1, "Heart 3": 1, "Heart 4": 1, "Heart 5": 1},
			input:   []int{0, 1, 2, 3, 4},
			want:    10,
		},
		"Repick_no_cards": {
			hands:   []string{"Heart 1", "Heart 2", "Heart 3", "Heart 4", "Heart 5"},
			handmap: map[string]int{"Heart 1": 1, "Heart 2": 1, "Heart 3": 1, "Heart 4": 1, "Heart 5": 1},
			input:   []int{},
			want:    5,
		},
	}

	for name, c := range cases {
		c := c
		t.Run(name, func(t *testing.T) {
			// t.Parallel()
			repickHands := card.RepickCards(c.hands, c.input)
			for _, rHand := range repickHands {
				if _, ok := c.handmap[rHand]; ok {
					c.handmap[rHand]++
				} else {
					c.handmap[rHand] = 1
				}
			}

			if len(c.handmap) != c.want {
				t.Errorf("RepickCards want len(%d) map, but got len(%d) map", c.want, len(c.handmap))
			}
		})
	}
}
