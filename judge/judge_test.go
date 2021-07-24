package judge_test

import (
	"testing"

	"github.com/kimuson13/go-poker/judge"
)

func TestIsRoyalStraightFlush(t *testing.T) {
	cases := map[string]struct {
		hands []string
		want  bool
	}{
		"RoyalStraightFlush_with_Heart": {
			hands: []string{"Heart A", "Heart 10", "Heart Jack", "Heart Queen", "Heart King"},
			want:  true,
		},
		"RoyalStraightFlush_with_Diamond": {
			hands: []string{"Diamond A", "Diamond 10", "Diamond Jack", "Diamond Queen", "Diamond King"},
			want:  true,
		},
		"RoyalStraightFlush_with_Club": {
			hands: []string{"Club A", "Club 10", "Club Jack", "Club Queen", "Club King"},
			want:  true,
		},
		"RoyalStraightFlush_with_Spade": {
			hands: []string{"Spade A", "Spade 10", "Spade Jack", "Spade Queen", "Spade King"},
			want:  true,
		},
		"Not_RoyalStraightFlush_because_of_mark_does_not_same": {
			hands: []string{"Heart A", "Heart, 10", "Heart Jack", "Heart Queen", "Club King"},
			want:  false,
		},
		"Not_RoyalStraightFlush_because_of_numders_are_not_continued": {
			hands: []string{"Heart A", "Heart, 10", "Heart Jack", "Heart Queen", "Heart 6"},
			want:  false,
		},
	}

	for name, c := range cases {
		got := judge.IsRoyalStraightFlush(c.hands)
		if c.want != got {
			t.Errorf("Test failed with %s: this test want %v, but got %v", name, c.want, got)
		}
	}
}

func TestIsStraightFlush(t *testing.T) {
	cases := map[string]struct {
		hands []string
		want  bool
	}{
		"StraihgtFlush_with_start_ace_and_Heart": {
			hands: []string{"Heart A", "Heart 2", "Heart 3", "Heart 4", "Heart 5"},
			want:  true,
		},
		"StraihgtFlush_with_start_2_and_Heart": {
			hands: []string{"Heart 2", "Heart 3", "Heart 4", "Heart 5", "Heart 6"},
			want:  true,
		},
		"StraihgtFlush_with_start_3_and_Heart": {
			hands: []string{"Heart 3", "Heart 4", "Heart 5", "Heart 6", "Heart 7"},
			want:  true,
		},
		"StraihgtFlush_with_start_4_and_Heart": {
			hands: []string{"Heart 4", "Heart 5", "Heart 6", "Heart 7", "Heart 8"},
			want:  true,
		},
		"StraihgtFlush_with_start_5_and_Heart": {
			hands: []string{"Heart 5", "Heart 6", "Heart 7", "Heart 8", "Heart 9"},
			want:  true,
		},
		"StraihgtFlush_with_start_6_and_Heart": {
			hands: []string{"Heart 6", "Heart 7", "Heart 8", "Heart 9", "Heart 10"},
			want:  true,
		},
		"StraihgtFlush_with_start_7_and_Heart": {
			hands: []string{"Heart 7", "Heart 8", "Heart 9", "Heart 10", "Heart Jack"},
			want:  true,
		},
		"StraihgtFlush_with_start_8_and_Heart": {
			hands: []string{"Heart 8", "Heart 9", "Heart 10", "Heart Jack", "Heart Queen"},
			want:  true,
		},
		"StraihgtFlush_with_start_ace_and_Diamond": {
			hands: []string{"Diamond A", "Diamond 2", "Diamond 3", "Diamond 4", "Diamond 5"},
			want:  true,
		},
		"StraihgtFlush_with_start_2_and_Diamond": {
			hands: []string{"Diamond 2", "Diamond 3", "Diamond 4", "Diamond 5", "Diamond 6"},
			want:  true,
		},
		"StraihgtFlush_with_start_3_and_Diamond": {
			hands: []string{"Diamond 3", "Diamond 4", "Diamond 5", "Diamond 6", "Diamond 7"},
			want:  true,
		},
		"StraihgtFlush_with_start_4_and_Diamond": {
			hands: []string{"Diamond 4", "Diamond 5", "Diamond 6", "Diamond 7", "Diamond 8"},
			want:  true,
		},
		"StraihgtFlush_with_start_5_and_Diamond": {
			hands: []string{"Diamond 5", "Diamond 6", "Diamond 7", "Diamond 8", "Diamond 9"},
			want:  true,
		},
		"StraihgtFlush_with_start_6_and_Diamond": {
			hands: []string{"Diamond 6", "Diamond 7", "Diamond 8", "Diamond 9", "Diamond 10"},
			want:  true,
		},
		"StraihgtFlush_with_start_7_and_Diamond": {
			hands: []string{"Diamond 7", "Diamond 8", "Diamond 9", "Diamond 10", "Diamond Jack"},
			want:  true,
		},
		"StraihgtFlush_with_start_8_and_Diamond": {
			hands: []string{"Diamond 8", "Diamond 9", "Diamond 10", "Diamond Jack", "Diamond Queen"},
			want:  true,
		},
		"StraihgtFlush_with_start_ace_and_Club": {
			hands: []string{"Club A", "Club 2", "Club 3", "Club 4", "Club 5"},
			want:  true,
		},
		"StraihgtFlush_with_start_2_and_Club": {
			hands: []string{"Club 2", "Club 3", "Club 4", "Club 5", "Club 6"},
			want:  true,
		},
		"StraihgtFlush_with_start_3_and_Club": {
			hands: []string{"Club 3", "Club 4", "Club 5", "Club 6", "Club 7"},
			want:  true,
		},
		"StraihgtFlush_with_start_4_and_Club": {
			hands: []string{"Club 4", "Club 5", "Club 6", "Club 7", "Club 8"},
			want:  true,
		},
		"StraihgtFlush_with_start_5_and_Club": {
			hands: []string{"Club 5", "Club 6", "Club 7", "Club 8", "Club 9"},
			want:  true,
		},
		"StraihgtFlush_with_start_6_and_Club": {
			hands: []string{"Club 6", "Club 7", "Club 8", "Club 9", "Club 10"},
			want:  true,
		},
		"StraihgtFlush_with_start_7_and_Club": {
			hands: []string{"Club 7", "Club 8", "Club 9", "Club 10", "Club Jack"},
			want:  true,
		},
		"StraihgtFlush_with_start_8_and_Club": {
			hands: []string{"Club 8", "Club 9", "Club 10", "Club Jack", "Club Queen"},
			want:  true,
		},
		"StraihgtFlush_with_start_ace_and_Spade": {
			hands: []string{"Spade A", "Spade 2", "Spade 3", "Spade 4", "Spade 5"},
			want:  true,
		},
		"StraihgtFlush_with_start_2_and_Spade": {
			hands: []string{"Spade 2", "Spade 3", "Spade 4", "Spade 5", "Spade 6"},
			want:  true,
		},
		"StraihgtFlush_with_start_3_and_Spade": {
			hands: []string{"Spade 3", "Spade 4", "Spade 5", "Spade 6", "Spade 7"},
			want:  true,
		},
		"StraihgtFlush_with_start_4_and_Spade": {
			hands: []string{"Spade 4", "Spade 5", "Spade 6", "Spade 7", "Spade 8"},
			want:  true,
		},
		"StraihgtFlush_with_start_5_and_Spade": {
			hands: []string{"Spade 5", "Spade 6", "Spade 7", "Spade 8", "Spade 9"},
			want:  true,
		},
		"StraihgtFlush_with_start_6_and_Spade": {
			hands: []string{"Spade 6", "Spade 7", "Spade 8", "Spade 9", "Spade 10"},
			want:  true,
		},
		"StraihgtFlush_with_start_7_and_Spade": {
			hands: []string{"Spade 7", "Spade 8", "Spade 9", "Spade 10", "Spade Jack"},
			want:  true,
		},
		"StraihgtFlush_with_start_8_and_Spade": {
			hands: []string{"Spade 8", "Spade 9", "Spade 10", "Spade Jack", "Spade Queen"},
			want:  true,
		},
		"Not_StraightFlush_because_of_not_continued": {
			hands: []string{"Spade A", "Spade 3", "Spade 4", "Spade 5", "Spade 6"},
			want:  false,
		},
		"Not_StraightFlush_because_of_not_same_mark": {
			hands: []string{"Spade A", "Heart 2", "Heart 3", "Heart 4", "Heart 5"},
			want:  false,
		},
	}

	for name, c := range cases {
		got := judge.IsStraightFlush(c.hands)
		if c.want != got {
			t.Errorf("Test failed with %s: this test want %v, but got %v", name, c.want, got)
		}
	}
}

func TestIsFourCard(t *testing.T) {
	cases := map[string]struct {
		hands []string
		want  bool
	}{
		"Fourcard_with_ace": {
			hands: []string{"Heart A", "Spade A", "Diamond A", "Club A", "Club 3"},
			want:  true,
		},
		"Fourcard_with_2": {
			hands: []string{"Heart 2", "Spade 2", "Diamond 2", "Club 2", "Club 3"},
			want:  true,
		},
		"Fourcard_with_3": {
			hands: []string{"Heart 3", "Spade 3", "Diamond 3", "Club A", "Club 3"},
			want:  true,
		},
		"Fourcard_with_4": {
			hands: []string{"Heart 4", "Spade 4", "Diamond 4", "Club 4", "Club 3"},
			want:  true,
		},
		"Fourcard_with_5": {
			hands: []string{"Heart 5", "Spade 5", "Diamond 5", "Club 5", "Club 3"},
			want:  true,
		},
		"Fourcard_with_6": {
			hands: []string{"Heart 6", "Spade 6", "Diamond 6", "Club 6", "Club 3"},
			want:  true,
		},
		"Fourcard_with_7": {
			hands: []string{"Heart 7", "Spade 7", "Diamond 7", "Club 7", "Club 3"},
			want:  true,
		},
		"Fourcard_with_8": {
			hands: []string{"Heart 8", "Spade 8", "Diamond 8", "Club 8", "Club 3"},
			want:  true,
		},
		"Fourcard_with_9": {
			hands: []string{"Heart 9", "Spade 9", "Diamond 9", "Club 9", "Club 3"},
			want:  true,
		},
		"Fourcard_with_10": {
			hands: []string{"Heart 10", "Spade 10", "Diamond 10", "Club 10", "Club 3"},
			want:  true,
		},
		"Fourcard_with_Jack": {
			hands: []string{"Heart Jack", "Spade Jack", "Diamond Jack", "Club Jack", "Club 3"},
			want:  true,
		},
		"Fourcard_with_Queen": {
			hands: []string{"Heart Queen", "Spade Queen", "Diamond Queen", "Club Queen", "Club 3"},
			want:  true,
		},
		"Fourcard_with_King": {
			hands: []string{"Heart King", "Spade King", "Diamond King", "Club King", "Club 3"},
			want:  true,
		},
		"Not_Fourcard": {
			hands: []string{"Heart 2", "Heart 5", "Spade 2", "Club 2", "Spade 3"},
			want:  false,
		},
	}

	for name, c := range cases {
		got := judge.IsFourCard(c.hands)
		if c.want != got {
			t.Errorf("Test failed with %s: this test want %v, but got %v", name, c.want, got)
		}
	}
}

func TestIsFullHouse(t *testing.T) {
	cases := map[string]struct {
		hands []string
		want  bool
	}{
		"FullHouse_ok_case1": {
			hands: []string{"Heart 3", "Spade 3", "Club 3", "Diamond 2", "Heart 2"},
			want:  true,
		},
		"FullHouse_ok_case2": {
			hands: []string{"Heart 5", "Spade 5", "Club 3", "Diamond 5", "Heart 3"},
			want:  true,
		},
		"FullHouse_ok_case3": {
			hands: []string{"Heart Jack", "Spade Jack", "Club Jack", "Diamond Queen", "Heart Queen"},
			want:  true,
		},
		"Not_FullHouse_because_it_is_FourCard": {
			hands: []string{"Heart 5", "Spade 5", "Club 5", "Diamond 5", "Heart 3"},
			want:  false,
		},
		"Not_FullHouse": {
			hands: []string{"Heart 4", "Spade 5", "Club 5", "Diamond 5", "Heart 3"},
			want:  false,
		},
	}

	for name, c := range cases {
		got := judge.IsFullHouse(c.hands)
		if c.want != got {
			t.Errorf("Test failed with %s: this test want %v, but got %v", name, c.want, got)
		}
	}
}

func TestIsFlush(t *testing.T) {
	cases := map[string]struct {
		hands []string
		want  bool
	}{
		"Flush_with_Heart": {
			hands: []string{"Heart 3", "Heart 4", "Heart Jack", "Heart 8", "Heart 2"},
			want:  true,
		},
		"Flush_with_Diamond": {
			hands: []string{"Diamond 3", "Diamond 4", "Diamond Jack", "Diamond 8", "Diamond 2"},
			want:  true,
		},
		"Flush_with_Spade": {
			hands: []string{"Spade 3", "Spade 4", "Spade Jack", "Spade 8", "Spade 2"},
			want:  true,
		},
		"Flush_with_Club": {
			hands: []string{"Club 3", "Club 4", "Club Jack", "Club 8", "Club 2"},
			want:  true,
		},
		"Not_Flush_because_of_not_same_mark": {
			hands: []string{"Heart A", "Heart 4", "Heart 6", "Heart 8", "Spade 10"},
			want:  false,
		},
	}

	for name, c := range cases {
		got := judge.IsFlush(c.hands)
		if c.want != got {
			t.Errorf("Test failed with %s: this test want %v, but got %v", name, c.want, got)
		}
	}
}

func TestIsStraight(t *testing.T) {
	cases := map[string]struct {
		hands []string
		want  bool
	}{
		"Straight_start_ace": {
			hands: []string{"Heart A", "Spade 2", "Club, 3", "Diamond 4", "Heart 5"},
			want:  true,
		},
		"Straihgt_start_2": {
			hands: []string{"Heart 2", "Diamond 3", "Spade 4", "Heart 5", "Heart 6"},
			want:  true,
		},
		"Straihgt_with_start_3": {
			hands: []string{"Spade 3", "Club 4", "Heart 5", "Heart 6", "Heart 7"},
			want:  true,
		},
		"Straihgt_with_start_4": {
			hands: []string{"Club 4", "Spade 5", "Heart 6", "Diamond 7", "Heart 8"},
			want:  true,
		},
		"Straihgt_with_start_5": {
			hands: []string{"Diamond 5", "Heart 6", "Club 7", "Heart 8", "Diamond 9"},
			want:  true,
		},
		"Straihgt_with_start_6": {
			hands: []string{"Heart 6", "Diamond 7", "Heart 8", "Spade 9", "Heart 10"},
			want:  true,
		},
		"Straihgt_with_start_7": {
			hands: []string{"Spade 7", "Heart 8", "Heart 9", "Heart 10", "Heart Jack"},
			want:  true,
		},
		"Straihgt_with_start_8": {
			hands: []string{"Heart 8", "Diamond 9", "Heart 10", "Heart Jack", "Heart Queen"},
			want:  true,
		},
		"Straihgt_with_start_9": {
			hands: []string{"Diamond 9", "Diamond 10", "Heart Jack", "Heart King", "Heart Queen"},
			want:  true,
		},
		"Not_Straight_because_of_not_continued_numbers": {
			hands: []string{"Diamond 2", "Heart 5", "Spade 6", "Diamond 7", "Club 8"},
			want:  false,
		},
	}

	for name, c := range cases {
		got := judge.IsStraight(c.hands)
		if c.want != got {
			t.Errorf("Test failed with %s: this test want %v, but got %v", name, c.want, got)
		}
	}
}

func TestIsThreeCard(t *testing.T) {
	cases := map[string]struct {
		hands []string
		want  bool
	}{
		"ThreeCard_with_Ace": {
			hands: []string{"Heart A", "Diamond A", "Spade A", "Diamond 2", "Spade 3"},
			want:  true,
		},
		"ThreeCard_with_2": {
			hands: []string{"Heart 2", "Diamond 2", "Spade 2", "Diamond 4", "Spade 3"},
			want:  true,
		},
		"ThreeCard_with_3": {
			hands: []string{"Heart 3", "Diamond 3", "Spade A", "Diamond 2", "Spade 3"},
			want:  true,
		},
		"ThreeCard_with_4": {
			hands: []string{"Heart 4", "Diamond 2", "Spade 4", "Diamond 4", "Spade 3"},
			want:  true,
		},
		"ThreeCard_with_5": {
			hands: []string{"Heart 5", "Diamond A", "Spade 2", "Diamond 5", "Spade 5"},
			want:  true,
		},
		"ThreeCard_with_6": {
			hands: []string{"Heart 2", "Diamond 2", "Spade 2", "Diamond 4", "Spade 3"},
			want:  true,
		},
		"ThreeCard_with_7": {
			hands: []string{"Heart 7", "Diamond 7", "Spade 7", "Diamond 2", "Spade 3"},
			want:  true,
		},
		"ThreeCard_with_8": {
			hands: []string{"Heart 8", "Diamond 8", "Spade 8", "Diamond 4", "Spade 3"},
			want:  true,
		},
		"ThreeCard_with_9": {
			hands: []string{"Heart 9", "Diamond 9", "Spade 9", "Diamond 2", "Spade 3"},
			want:  true,
		},
		"ThreeCard_with_10": {
			hands: []string{"Heart 10", "Diamond 10", "Spade 10", "Diamond 4", "Spade 3"},
			want:  true,
		},
		"ThreeCard_with_Jack": {
			hands: []string{"Heart Jack", "Diamond Jack", "Spade Jack", "Diamond 2", "Spade 3"},
			want:  true,
		},
		"ThreeCard_with_Queen": {
			hands: []string{"Heart Queen", "Diamond A", "Spade 4", "Diamond Queen", "Spade Queen"},
			want:  true,
		},
		"ThreeCard_with_King": {
			hands: []string{"Heart King", "Spade King", "Club King", "Spade 1", "Club 10"},
			want:  true,
		},
		"Not_ThreeCard_because_it_is_one_pair": {
			hands: []string{"Heart A", "Diamond A", "Spade King", "Dimond Queen", "Club 4"},
			want:  false,
		},
	}

	for name, c := range cases {
		got := judge.IsThreeCard(c.hands)
		if c.want != got {
			t.Errorf("Test failed with %s: this test want %v, but got %v", name, c.want, got)
		}
	}
}

func TestIsTwoPair(t *testing.T) {
	cases := map[string]struct {
		hands []string
		want  bool
	}{
		"TwoPair_case1": {
			hands: []string{"Heart A", "Diamond A", "Diamond King", "Club King", "Spade 3"},
			want:  true,
		},
		"TwoPair_case2": {
			hands: []string{"Heart 4", "Diamond 4", "Diamond Jack", "Club King", "Spade Jack"},
			want:  true,
		},
		"TwoPair_case3": {
			hands: []string{"Heart 1", "Diamond 1", "Diamond 2", "Club 3", "Spade 3"},
			want:  true,
		},
		"Not_TwoPair_because_it_is_ThreeCard": {
			hands: []string{"Heart King", "Diamond King", "Club King", "Heart 4", "Diamond 3"},
			want:  false,
		},
		"Not_TwoPari_because_it_is_OnePari": {
			hands: []string{"Heart King", "Spade King", "Diamond 5", "Diamond 7", "Club 2"},
			want:  false,
		},
	}

	for name, c := range cases {
		got := judge.IsTwoPair(c.hands)
		if c.want != got {
			t.Errorf("Test failed with %s: this test want %v, but got %v", name, c.want, got)
		}
	}
}

func TestIsOnePair(t *testing.T) {
	cases := map[string]struct {
		hands []string
		want  bool
	}{
		"OnePair_case1": {
			hands: []string{"Spade 2", "Diamond 2", "Spade 9", "Dimaond 10", "Club Jack"},
			want:  true,
		},
		"OnePair_case2": {
			hands: []string{"Spade King", "Diamond 2", "Spade 9", "Dimaond Jack", "Club Jack"},
			want:  true,
		},
		"OnePair_case3": {
			hands: []string{"Spade 5", "Diamond 2", "Spade 9", "Dimaond 5", "Club Jack"},
			want:  true,
		},
		"Not_OnePair_because_it_is_TwoPair": {
			hands: []string{"Heart A", "Diamond A", "Diamond King", "Club King", "Spade 3"},
			want:  false,
		},
		"Not_OnePair_because_it_is_ThreeCard": {
			hands: []string{"Heart King", "Diamond King", "Club King", "Heart 4", "Diamond 3"},
			want:  false,
		},
	}

	for name, c := range cases {
		got := judge.IsOnePair(c.hands)
		if c.want != got {
			t.Errorf("Test failed with %s: this test want %v, but got %v", name, c.want, got)
		}
	}
}
