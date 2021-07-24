package judge

import (
	"strings"
)

func IsRoyalStraightFlush(hands []string) bool {
	cardMarks, cardNums := classifyHands(hands)
	if len(cardMarks) == 1 {
		_, isA := cardNums["A"]
		_, is10 := cardNums["10"]
		_, isJack := cardNums["Jack"]
		_, isQueen := cardNums["Queen"]
		_, isKing := cardNums["King"]

		if isA && is10 && isJack && isQueen && isKing {
			return true
		}
	}

	return false
}

func IsStraightFlush(hands []string) bool {
	cardMarks, cardNums := classifyHands(hands)
	if len(cardMarks) == 1 {
		flag := isNumsContinued(cardNums)
		if flag {
			return true
		}
	}

	return false
}

func IsFourCard(hands []string) bool {
	cardMarks, cardNums := classifyHands(hands)
	if len(cardMarks) == 4 {
		for _, n := range cardNums {
			if n == 4 {
				return true
			}
		}
	}

	return false
}

func IsFullHouse(hands []string) bool {
	_, cardNums := classifyHands(hands)
	if len(cardNums) == 2 {
		var count int
		for _, v := range cardNums {
			if v == 2 || v == 3 {
				count++
			}
		}

		if count == 2 {
			return true
		}
	}

	return false
}

func IsFlush(hands []string) bool {
	cardMarks, _ := classifyHands(hands)
	if len(cardMarks) == 1 {
		return true
	}

	return false
}

func IsStraight(hands []string) bool {
	_, cardNums := classifyHands(hands)
	flag := isNumsContinued(cardNums)
	if flag {
		return true
	}

	return false
}

func IsThreeCard(hands []string) bool {
	_, cardNums := classifyHands(hands)
	if len(cardNums) != 3 {
		return false
	}
	for _, n := range cardNums {
		if n == 3 {
			return true
		}
	}

	return false
}

func IsTwoPair(hands []string) bool {
	_, cardNums := classifyHands(hands)
	if len(cardNums) != 3 {
		return false
	}

	var count int
	for _, n := range cardNums {
		if n == 2 {
			count++
		}
	}

	if count == 2 {
		return true
	}

	return false
}

func IsOnePair(hands []string) bool {
	_, cardNums := classifyHands(hands)
	if len(cardNums) != 4 {
		return false
	}

	var count int
	for _, n := range cardNums {
		if n == 2 {
			count++
		}
	}

	if count == 1 {
		return true
	}

	return false
}

func isNumsContinued(cardNums map[string]int) bool {
	if len(cardNums) != 5 {
		return false
	}

	_, isA := cardNums["A"]
	_, is2 := cardNums["2"]
	_, is3 := cardNums["3"]
	_, is4 := cardNums["4"]
	_, is5 := cardNums["5"]
	_, is6 := cardNums["6"]
	_, is7 := cardNums["7"]
	_, is8 := cardNums["8"]
	_, is9 := cardNums["9"]
	_, is10 := cardNums["10"]
	_, isJack := cardNums["Jack"]
	_, isQueen := cardNums["Queen"]
	_, isKing := cardNums["King"]

	switch {
	case isA && is2 && is3 && is4 && is5:
		return true
	case is2 && is3 && is4 && is5 && is6:
		return true
	case is3 && is4 && is5 && is6 && is7:
		return true
	case is4 && is5 && is6 && is7 && is8:
		return true
	case is5 && is6 && is7 && is8 && is9:
		return true
	case is6 && is7 && is8 && is9 && is10:
		return true
	case is7 && is8 && is9 && is10 && isJack:
		return true
	case is8 && is9 && is10 && isJack && isQueen:
		return true
	case is9 && is10 && isJack && isQueen && isKing:
		return true
	}

	return false
}

func classifyHands(hands []string) (map[string]int, map[string]int) {
	cardMarks := make(map[string]int)
	cardNums := make(map[string]int)

	for _, hand := range hands {
		markAndNum := strings.Split(hand, " ")
		if _, ok := cardMarks[markAndNum[0]]; ok {
			cardMarks[markAndNum[0]]++
		} else {
			cardMarks[markAndNum[0]] = 1
		}

		if _, ok := cardNums[markAndNum[1]]; ok {
			cardNums[markAndNum[1]]++
		} else {
			cardNums[markAndNum[1]] = 1
		}
	}

	return cardMarks, cardNums
}
