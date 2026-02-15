package internal

import "time"

func interval(state int) int {
	switch state {
	case 1:
		return 1
	case 2:
		return 3
	case 3:
		return 7
	default:
		return 14
	}
}

func IsDue(card Card) bool {
	if card.LastReviewed == nil {
		return true
	}

	next := card.LastReviewed.AddDate(0, 0, interval(card.State))
	return time.Now().After(next)
}

func Promote(card *Card) {
	if card.State < 4 {
		card.State++
	}
	now := time.Now()
	card.LastReviewed = &now
}

func Reset(card *Card) {
	card.State = 1
	now := time.Now()
	card.LastReviewed = &now
}
