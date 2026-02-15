package internal

import "time"

type Card struct {
	Front        string     `json:"front"`
	Back         string     `json:"back"`
	State        int        `json:"state"`
	LastReviewed *time.Time `json:"lastReviewed,omitempty"`
}

type Deck struct {
	Name  string `json:"name"`
	Cards []Card `json:"cards"`
}
