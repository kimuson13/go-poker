package system_test

import (
	"bytes"
	"database/sql"
	"fmt"
	"os"
	"testing"

	highscoreDB "github.com/kimuson13/go-poker/db"
	"github.com/kimuson13/go-poker/system"
)

func setUpTestDB(t *testing.T, name string) (*sql.DB, func()) {
	t.Helper()
	testDB, err := highscoreDB.GetDB(name)
	dbName := fmt.Sprintf("%s.db", name)
	if err != nil {
		t.Fatal("Error: ", err)
	}

	return testDB, func() {
		if err := testDB.Close(); err != nil {
			t.Fatal("Error: ", err)
		}
		if err := os.Remove(dbName); err != nil {
			t.Fatal("Error ", err)
		}
	}
}

func TestShowHighScorewithNoRecord(t *testing.T) {
	testDB, cleanFunc := setUpTestDB(t, "testsys1")
	defer cleanFunc()

	want := struct {
		score int
		is    bool
	}{
		score: 0,
		is:    false,
	}

	gotIs, gotScore, err := system.ShowHighScore(testDB)
	if err != nil {
		t.Fatal(err)
	}

	switch {
	case gotIs != want.is:
		t.Errorf("want bool is %v, but got %v", want.is, gotIs)
	case gotScore != want.score:
		t.Errorf("want score is %d, but got %d", want.score, gotScore)
	}
}

func TestShowHighScoreWithAlreadyRecordRxists(t *testing.T) {
	testDB, cleanFunc := setUpTestDB(t, "testsys2")
	defer cleanFunc()

	model := highscoreDB.HighScore{
		Name:  "test",
		Score: 100,
	}
	if err := highscoreDB.CreateHighScore(testDB, model); err != nil {
		t.Fatal(err)
	}

	want := struct {
		score int
		is    bool
	}{
		score: 100,
		is:    true,
	}

	gotIs, gotScore, err := system.ShowHighScore(testDB)
	if err != nil {
		t.Fatal(err)
	}
	switch {
	case gotIs != want.is:
		t.Errorf("want bool is %v, but got %v", want.is, gotIs)
	case gotScore != want.score:
		t.Errorf("want score is %d, but got %d", want.score, gotScore)
	}
}

func TestAddHighScoreWithNoRecord(t *testing.T) {
	testDB, cleanFunc := setUpTestDB(t, "testsys3")
	defer cleanFunc()

	input := struct {
		chip             int
		currentHighScore int
		name             string
		flag             bool
	}{
		chip:             100,
		currentHighScore: 10,
		name:             "test2",
		flag:             false,
	}

	err := system.AddHighScore(testDB, input.flag, input.chip, input.currentHighScore, input.name)
	if err != nil {
		t.Fatal(err)
	}
}

func TestAddHighScoreWithAlreadyRecordExists(t *testing.T) {
	testDB, cleanFunc := setUpTestDB(t, "testsys4")
	defer cleanFunc()
	model := highscoreDB.HighScore{
		Name:  "test",
		Score: 100,
	}
	if err := highscoreDB.CreateHighScore(testDB, model); err != nil {
		t.Fatal(err)
	}

	cases := map[string]struct {
		chip             int
		currentHighScore int
		name             string
		flag             bool
	}{
		"it_is_under_score": {
			chip:             50,
			currentHighScore: model.Score,
			name:             "test",
			flag:             true,
		},
		"it_is_over_score": {
			chip:             200,
			currentHighScore: model.Score,
			name:             "test",
			flag:             true,
		},
	}

	for name, c := range cases {
		c := c
		t.Run(name, func(t *testing.T) {
			err := system.AddHighScore(testDB, c.flag, c.chip, c.currentHighScore, c.name)
			if err != nil {
				t.Fatal(err)
			}
		})
	}
}

func TestReadytoStart(t *testing.T) {
	cases := map[string]string{
		"answer_yes":               "y",
		"answer_default_after_yes": "aaa\ny",
	}

	for name, c := range cases {
		c := c
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			testInput := bytes.NewBufferString(c)
			in := &system.UserInput{
				Stdin: testInput,
			}
			err := system.ReadyToStart(in.Stdin, "test", 0)
			if err != nil {
				t.Error(err)
			}
		})
	}
}

func TestPoker(t *testing.T) {
	input := "No"
	testInput := bytes.NewBufferString(input + "\n")
	in := &system.UserInput{
		Stdin: testInput,
	}

	got, err := system.Poker(1, in.Stdin)
	if err != nil {
		t.Fatal(err)
	}
	if got < 0 {
		t.Error("it is not ok")
	}
}

func TestIsContinued(t *testing.T) {
	cases := map[string]struct {
		input string
		want  bool
	}{
		"answer_yes": {
			input: "y",
			want:  true,
		},
		"answer_no": {
			input: "n",
			want:  false,
		},
		"answer_others_last_true": {
			input: "aaa\ny",
			want:  true,
		},
		"answer_others_last_false": {
			input: "iii\nn",
			want:  false,
		},
	}

	for name, c := range cases {
		c := c
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			testInput := bytes.NewBufferString(c.input + "\n")
			in := &system.UserInput{
				Stdin: testInput,
			}

			got, err := system.IsContinued(in.Stdin)
			if err != nil {
				t.Fatal(err)
			}
			if got != c.want {
				t.Errorf("got: %v, want: %v", got, c.want)
			}
		})
	}
}

func TestIsChangeCardsWithDefault(t *testing.T) {
	input1 := "aiueo"
	input2 := "all"
	testInput := bytes.NewBufferString(input1 + "\n" + input2 + "\n")
	in := &system.UserInput{
		Stdin: testInput,
	}
	want := []int{0, 1, 2, 3, 4}
	got, err := system.ChangeCards(in.Stdin)
	if err != nil {
		t.Fatal(err)
	}

	for i := 0; i < len(want); i++ {
		w := want[i]
		g := got[i]
		if w != g {
			t.Errorf("slice number %d: want value is %d, but got %d", i, w, g)
		}
	}
}

func TestChangeCards(t *testing.T) {
	cases := map[string]struct {
		input string
		want  []int
	}{
		"answer_no": {
			input: "No",
			want:  []int{},
		},
		"answer_all": {
			input: "all",
			want:  []int{0, 1, 2, 3, 4},
		},
		"answer_5": {
			input: "5",
			want:  []int{4},
		},
		"answer_4": {
			input: "4",
			want:  []int{3},
		},
		"answer_4_and_5": {
			input: "4 5",
			want:  []int{3, 4},
		},
		"answer_3": {
			input: "3",
			want:  []int{2},
		},
		"answer_3_and_5": {
			input: "3 5",
			want:  []int{2, 4},
		},
		"answer_3_and_4": {
			input: "3 4",
			want:  []int{2, 3},
		},
		"answer_3_4_5": {
			input: "3 4 5",
			want:  []int{2, 3, 4},
		},
		"answer_2": {
			input: "2",
			want:  []int{1},
		},
		"answer_2_and_5": {
			input: "2 5",
			want:  []int{1, 4},
		},
		"answer_2_and_4": {
			input: "2 4",
			want:  []int{1, 3},
		},
		"answer_2_4_5": {
			input: "2 4 5",
			want:  []int{1, 3, 4},
		},
		"answer_2_3": {
			input: "2 3",
			want:  []int{1, 2},
		},
		"answer_2_3_5": {
			input: "2 3 5",
			want:  []int{1, 2, 4},
		},
		"answer_2_3_4": {
			input: "2 3 4",
			want:  []int{1, 2, 3},
		},
		"answer_2_3_4_5": {
			input: "2 3 4 5",
			want:  []int{1, 2, 3, 4},
		},
		"answer_1": {
			input: "1",
			want:  []int{0},
		},
		"answer_1_5": {
			input: "1 5",
			want:  []int{0, 4},
		},
		"answer_1_4": {
			input: "1 4",
			want:  []int{0, 3},
		},
		"answer_1_4_5": {
			input: "1 4 5",
			want:  []int{0, 3, 4},
		},
		"answer_1_3": {
			input: "1 3",
			want:  []int{0, 2},
		},
		"answer_1_3_5": {
			input: "1 3 5",
			want:  []int{0, 2, 4},
		},
		"answer_1_3_4": {
			input: "1 3 4",
			want:  []int{0, 2, 3},
		},
		"answer_1_3_4_5": {
			input: "1 3 4 5",
			want:  []int{0, 2, 3, 4},
		},
		"answer_1_2": {
			input: "1 2",
			want:  []int{0, 1},
		},
		"answer_1_2_5": {
			input: "1 2 5",
			want:  []int{0, 1, 4},
		},
		"answer_1_2_4": {
			input: "1 2 4",
			want:  []int{0, 1, 3},
		},
		"answer_1_2_4_5": {
			input: "1 2 4 5",
			want:  []int{0, 1, 3, 4},
		},
		"answer_1_2_3": {
			input: "1 2 3",
			want:  []int{0, 1, 2},
		},
		"answer_1_2_3_5": {
			input: "1 2 3 5",
			want:  []int{0, 1, 2, 4},
		},
		"answer_1_2_3_4": {
			input: "1 2 3 4",
			want:  []int{0, 1, 2, 3},
		},
	}

	for name, c := range cases {
		c := c
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			testInput := bytes.NewBufferString(c.input + "\n")
			in := &system.UserInput{
				Stdin: testInput,
			}

			got, err := system.ChangeCards(in.Stdin)
			if err != nil {
				t.Fatal(err)
			}

			if len(got) != len(c.want) {
				t.Errorf("length want %d, but got %d", len(c.want), len(got))
			}

			if len(c.want) > 0 {
				for i := 0; i < len(c.want); i++ {
					w := c.want[i]
					g := got[i]
					if w != g {
						t.Errorf("slice number %d: want value is %d, but got %d", i, w, g)
					}
				}
			}
		})
	}
}

func TestCaculate(t *testing.T) {
	cases := map[string]struct {
		rate  int
		want  int
		hands []string
	}{
		"got_royalStraight_with_rate_1": {
			rate:  1,
			want:  800,
			hands: []string{"Heart A", "Heart 10", "Heart Jack", "Heart Queen", "Heart King"},
		},
		"got_StraightFlush_with_rate_1": {
			rate:  1,
			want:  50,
			hands: []string{"Heart A", "Heart 2", "Heart 3", "Heart 4", "Heart 5"},
		},
		"got_FourCard_wiht_rate_1": {
			rate:  1,
			want:  20,
			hands: []string{"Heart A", "Spade A", "Diamond A", "Club A", "Club 3"},
		},
		"got_FullHouse_with_rate_1": {
			rate:  1,
			want:  6,
			hands: []string{"Heart 3", "Spade 3", "Club 3", "Diamond 2", "Heart 2"},
		},
		"got_Flush_with_rate_1": {
			rate:  1,
			want:  5,
			hands: []string{"Heart 3", "Heart 4", "Heart Jack", "Heart 8", "Heart 2"},
		},
		"got_Straight_with_rate_1": {
			rate:  1,
			want:  4,
			hands: []string{"Heart A", "Spade 2", "Club, 3", "Diamond 4", "Heart 5"},
		},
		"got_ThreeCard_with_rate_1": {
			rate:  1,
			want:  3,
			hands: []string{"Heart A", "Diamond A", "Spade A", "Diamond 2", "Spade 3"},
		},
		"got_TwoPair_with_rate_1": {
			rate:  1,
			want:  2,
			hands: []string{"Heart A", "Diamond A", "Diamond King", "Club King", "Spade 3"},
		},
		"got_OnePair_with_rate_1": {
			rate:  1,
			want:  1,
			hands: []string{"Spade 2", "Diamond 2", "Spade 9", "Dimaond 10", "Club Jack"},
		},
		"got_royalStraight_with_Other_rate": {
			rate:  5,
			want:  4000,
			hands: []string{"Heart A", "Heart 10", "Heart Jack", "Heart Queen", "Heart King"},
		},
		"got_StraightFlush_with_Other_rate": {
			rate:  5,
			want:  250,
			hands: []string{"Heart A", "Heart 2", "Heart 3", "Heart 4", "Heart 5"},
		},
		"got_FourCard_wiht_other_rate": {
			rate:  5,
			want:  100,
			hands: []string{"Heart A", "Spade A", "Diamond A", "Club A", "Club 3"},
		},
		"got_FullHouse_with_other_rate": {
			rate:  5,
			want:  30,
			hands: []string{"Heart 3", "Spade 3", "Club 3", "Diamond 2", "Heart 2"},
		},
		"got_Flush_with_other_rate": {
			rate:  5,
			want:  25,
			hands: []string{"Heart 3", "Heart 4", "Heart Jack", "Heart 8", "Heart 2"},
		},
		"got_Straight_with_other_rate": {
			rate:  5,
			want:  20,
			hands: []string{"Heart A", "Spade 2", "Club, 3", "Diamond 4", "Heart 5"},
		},
		"got_ThreeCard_with_other_rate": {
			rate:  5,
			want:  15,
			hands: []string{"Heart A", "Diamond A", "Spade A", "Diamond 2", "Spade 3"},
		},
		"got_TwoPair_with_other_rate": {
			rate:  5,
			want:  10,
			hands: []string{"Heart A", "Diamond A", "Diamond King", "Club King", "Spade 3"},
		},
		"got_OnePair_with_other_rate": {
			rate:  5,
			want:  5,
			hands: []string{"Spade 2", "Diamond 2", "Spade 9", "Dimaond 10", "Club Jack"},
		},
	}

	for name, c := range cases {
		c := c
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got, err := system.Calculate(c.hands, c.rate)
			if err != nil {
				t.Fatal(err)
			}
			if got != c.want {
				t.Errorf("want Calculate(%v, %d) = %d, but got %d", c.hands, c.rate, c.want, got)
			}
		})
	}
}
