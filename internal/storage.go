package internal

import (
	"encoding/json"
	"os"
	"path/filepath"
)

func baseDir() string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, ".fastcards")
}

func decksDir() string {
	return filepath.Join(baseDir(), "decks")
}

func EnsureDirs() error {
	return os.MkdirAll(decksDir(), 0755)
}

func deckPath(name string) string {
	return filepath.Join(decksDir(), name+".json")
}

func SaveDeck(deck Deck) error {
	if err := EnsureDirs(); err != nil {
		return err
	}

	data, err := json.MarshalIndent(deck, "", " ")
	if err != nil {
		return err
	}

	return os.WriteFile(deckPath(deck.Name), data, 0644)
}

func LoadDeck(name string) (Deck, error) {
	var deck Deck
	data, err := os.ReadFile(deckPath(name))
	if err != nil {
		return deck, err
	}

	err = json.Unmarshal(data, &deck)
	return deck, err
}

func ListDeckNames() ([]string, error) {
	dir := decksDir()

	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var names []string

	for _, e := range entries {
		if e.IsDir() {
			continue
		}

		if filepath.Ext(e.Name()) == ".json" {
			name := e.Name()[:len(e.Name())-5]
			names = append(names, name)
		}
	}

	return names, nil
}
