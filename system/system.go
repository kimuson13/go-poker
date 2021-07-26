package system

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/kimuson13/go-poker/card"
	"github.com/kimuson13/go-poker/judge"
)

type UserInput struct {
	Stdin io.Reader
}

func New() *UserInput {
	return &UserInput{
		Stdin: os.Stdin,
	}
}

const (
	royalStraightFlush = 800
	straightFlush      = 50
	fourCard           = 20
	fullHouse          = 6
	flush              = 5
	straight           = 4
	threeCard          = 3
	twoPair            = 2
	onePair            = 1
)

var chip int = 20

func Run(name string, rate int) error {
	userInput := New()

	if _, err := fmt.Printf("Welcome to go-poker, %s.\nrating is %d\nLet's start game!\n", name, rate); err != nil {
		return err
	}

	for {
		chip -= rate
		if chip <= 0 {
			if _, err := fmt.Println("you don't have enough chip!"); err != nil {
				return err
			}
			goto L
		}

		result, err := Poker(rate, userInput.Stdin)
		if err != nil {
			return fmt.Errorf("poker error: %w", err)
		}
		chip += result
		if _, err := fmt.Printf("Your current chip: %d\n", chip); err != nil {
			return err
		}

		flag, err := IsContinued(userInput.Stdin)
		if err != nil {
			return fmt.Errorf("continue error: %w", err)
		}
		if !flag {
			goto L
		}

		card.ResetCards()
	}
L:

	return nil
}

func Poker(rate int, in io.Reader) (int, error) {
	hands := card.PickFirstCards()
	if _, err := fmt.Println("Your Card:"); err != nil {
		return 0, err
	}
	for n, hand := range hands {
		num := n + 1
		if _, err := fmt.Printf("hands %d: %s\n", num, hand); err != nil {
			return 0, err
		}
	}
	if _, err := fmt.Println("Please input numbers you want to change in ascending orger.\nif you don't change cards, please input 'No'\nIf you want to change all, please input 'all'\nFor example: '1 2 3'"); err != nil {
		return 0, err
	}
	input, err := ChangeCards(in)
	if err != nil {
		return 0, err
	}

	nHands := card.RepickCards(hands, input)
	for n, nHand := range nHands {
		num := n + 1
		if _, err := fmt.Printf("hands %d: %s\n", num, nHand); err != nil {
			return 0, err
		}
	}

	result, err := Calculate(nHands, rate)
	if err != nil {
		return 0, err
	}
	if _, err := fmt.Printf("Your result: %d\n", result); err != nil {
		return 0, err
	}

	return result, nil
}

func IsContinued(in io.Reader) (bool, error) {
	if _, err := fmt.Println("If you want to continue, type 'y'. If you want to exit, type 'n'."); err != nil {
		return false, err
	}
	scanner := bufio.NewScanner(in)
	for {
		scanner.Scan()
		input := scanner.Text()
		switch input {
		case "y":
			if _, err := fmt.Println("ok, let's play next game!"); err != nil {
				return false, err
			}
			return true, nil
		case "n":
			if _, err := fmt.Println("ok, we'll be looking forward to meeting you again someday."); err != nil {
				return false, err
			}
			return false, nil
		default:
			if _, err := fmt.Println("this command is not allowed, please type again."); err != nil {
				return false, err
			}
			continue
		}
	}
}

func ChangeCards(in io.Reader) ([]int, error) {
	var input []int
	scanner := bufio.NewScanner(in)
	for {
		scanner.Scan()
		in := scanner.Text()
		switch in {
		case "No":
			if _, err := fmt.Println("change No cards."); err != nil {
				return input, err
			}
			return input, nil
		case "all":
			if _, err := fmt.Println("change all cards."); err != nil {
				return input, err
			}
			input = []int{0, 1, 2, 3, 4}
			return input, nil
		case "5":
			if _, err := fmt.Println("change hands 5"); err != nil {
				return input, err
			}
			input = []int{4}
			return input, nil
		case "4":
			if _, err := fmt.Println("change hands 4"); err != nil {
				return input, err
			}
			input = []int{3}
			return input, nil
		case "4 5":
			if _, err := fmt.Println("change hands 4 and hands 5"); err != nil {
				return input, err
			}
			input = []int{3, 4}
			return input, nil
		case "3":
			if _, err := fmt.Println("change hands 3"); err != nil {
				return input, err
			}
			input = []int{2}
			return input, nil
		case "3 5":
			if _, err := fmt.Println("change hands 3 and hands 5"); err != nil {
				return input, err
			}
			input = []int{2, 4}
			return input, nil
		case "3 4":
			if _, err := fmt.Println("change hands 3 and hands 4"); err != nil {
				return input, err
			}
			input = []int{2, 3}
			return input, nil
		case "3 4 5":
			if _, err := fmt.Println("change hands 3, hands 4 and hands 5"); err != nil {
				return input, err
			}
			input = []int{2, 3, 4}
			return input, nil
		case "2":
			if _, err := fmt.Println("change hands 2"); err != nil {
				return input, err
			}
			input = []int{1}
			return input, nil
		case "2 5":
			if _, err := fmt.Println("change hands 2 and hands 5"); err != nil {
				return input, err
			}
			input = []int{1, 4}
			return input, nil
		case "2 4":
			if _, err := fmt.Println("change hands 2 and hands 4"); err != nil {
				return input, err
			}
			input = []int{1, 3}
			return input, nil
		case "2 4 5":
			if _, err := fmt.Println("change hands 2, hands 4 hands 5"); err != nil {
				return input, err
			}
			input = []int{1, 3, 4}
			return input, nil
		case "2 3":
			if _, err := fmt.Println("change hands 2 and hands 3"); err != nil {
				return input, err
			}
			input = []int{1, 2}
			return input, nil
		case "2 3 5":
			if _, err := fmt.Println("change hands 2, hands 3 and hands 5"); err != nil {
				return input, err
			}
			input = []int{1, 2, 4}
			return input, nil
		case "2 3 4":
			if _, err := fmt.Println("change hands 2, hands 3 and hands 4"); err != nil {
				return input, err
			}
			input = []int{1, 2, 3}
			return input, nil
		case "2 3 4 5":
			if _, err := fmt.Println("change hands 2, hands 3, hands 4 and hands 5"); err != nil {
				return input, err
			}
			input = []int{1, 2, 3, 4}
			return input, nil
		case "1":
			if _, err := fmt.Println("change hands 1"); err != nil {
				return input, err
			}
			input = []int{0}
			return input, nil
		case "1 5":
			if _, err := fmt.Println("change hands 1 and hands 5"); err != nil {
				return input, err
			}
			input = []int{0, 4}
			return input, nil
		case "1 4":
			if _, err := fmt.Println("change hands 1 and hands 4"); err != nil {
				return input, err
			}
			input = []int{0, 3}
			return input, nil
		case "1 4 5":
			if _, err := fmt.Println("change hands 1, hands 4 and hands 5"); err != nil {
				return input, err
			}
			input = []int{0, 3, 4}
			return input, nil
		case "1 3":
			if _, err := fmt.Println("change hands 1 and hands 3"); err != nil {
				return input, err
			}
			input = []int{0, 2}
			return input, nil
		case "1 3 5":
			if _, err := fmt.Println("change hands 1, hands 3, hands 5"); err != nil {
				return input, err
			}
			input = []int{0, 2, 4}
			return input, nil
		case "1 3 4":
			if _, err := fmt.Println("change hands 1, hands 3 and hands 4"); err != nil {
				return input, err
			}
			input = []int{0, 2, 3}
			return input, nil
		case "1 3 4 5":
			if _, err := fmt.Println("change hands 1, hands 3, hands 4 and hands 5"); err != nil {
				return input, err
			}
			input = []int{0, 2, 3, 4}
			return input, nil
		case "1 2":
			if _, err := fmt.Println("change hands 1 and hands 2"); err != nil {
				return input, err
			}
			input = []int{0, 1}
			return input, nil
		case "1 2 5":
			if _, err := fmt.Println("change hands 1, hands 2 and hands 5"); err != nil {
				return input, err
			}
			input = []int{0, 1, 4}
			return input, nil
		case "1 2 4":
			if _, err := fmt.Println("change hands 1, hands 2 and hands 4"); err != nil {
				return input, err
			}
			input = []int{0, 1, 3}
			return input, nil
		case "1 2 4 5":
			if _, err := fmt.Println("change hands 1, hands 2, hands 4 and hands 5"); err != nil {
				return input, err
			}
			input = []int{0, 1, 3, 4}
			return input, nil
		case "1 2 3":
			if _, err := fmt.Println("change hands 1, hands 2 and hands 3"); err != nil {
				return input, err
			}
			input = []int{0, 1, 2}
			return input, nil
		case "1 2 3 5":
			if _, err := fmt.Println("change hands 1, hands 2 and hands 4"); err != nil {
				return input, err
			}
			input = []int{0, 1, 2, 4}
			return input, nil
		case "1 2 3 4":
			if _, err := fmt.Println("change hands 1, hands 2, hands 3 and hands 4"); err != nil {
				return input, err
			}
			input = []int{0, 1, 2, 3}
			return input, nil
		default:
			if _, err := fmt.Println("this command is not allowed, please type again."); err != nil {
				return input, err
			}
			continue
		}
	}
}

func Calculate(hands []string, rate int) (int, error) {
	for i := 0; i < 9; i++ {
		switch i {
		case 0:
			if judge.IsRoyalStraightFlush(hands) {
				if _, err := fmt.Println("ROYAL STRAIGHT FLUSH!!!"); err != nil {
					return 0, err
				}
				return rate * royalStraightFlush, nil
			}
		case 1:
			if judge.IsStraightFlush(hands) {
				if _, err := fmt.Println("STRAIGHT FLUSH!!!"); err != nil {
					return 0, err
				}
				return rate * straightFlush, nil
			}
		case 2:
			if judge.IsFourCard(hands) {
				if _, err := fmt.Println("FOUR CARD!!!"); err != nil {
					return 0, err
				}
				return rate * fourCard, nil
			}
		case 3:
			if judge.IsFullHouse(hands) {
				if _, err := fmt.Println("FULL HOUSE!!!"); err != nil {
					return 0, err
				}
				return rate * fullHouse, nil
			}
		case 4:
			if judge.IsFlush(hands) {
				if _, err := fmt.Println("FLUSH!!!"); err != nil {
					return 0, err
				}
				return rate * flush, nil
			}
		case 5:
			if judge.IsStraight(hands) {
				if _, err := fmt.Println("STRAIGHT!!!"); err != nil {
					return 0, err
				}
				return rate * straight, nil
			}
		case 6:
			if judge.IsThreeCard(hands) {
				if _, err := fmt.Println("THREE CARD!!!"); err != nil {
					return 0, err
				}
				return rate * threeCard, nil
			}
		case 7:
			if judge.IsTwoPair(hands) {
				if _, err := fmt.Println("TWO PAIR!!!"); err != nil {
					return 0, err
				}
				return rate * twoPair, nil
			}
		case 8:
			if judge.IsOnePair(hands) {
				if _, err := fmt.Println("ONE PAIR"); err != nil {
					return 0, err
				}
				return rate * onePair, nil
			}
		}
	}

	return 0, nil
}
