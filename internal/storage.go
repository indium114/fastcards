package internal

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func baseDir() string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, ".fastcards")
}

func DecksDir() string {
	return filepath.Join(baseDir(), "decks")
}

func ArchiveDir() string {
	return filepath.Join(baseDir(), "archive")
}

func DataDir() string {
	return filepath.Join(baseDir(), "data")
}

func EnsureDirs() error {
	return os.MkdirAll(DecksDir(), 0755)
	return os.MkdirAll(ArchiveDir(), 0755)
}

func deckPath(name string) string {
	return filepath.Join(DecksDir(), name+".json")
}

func xpPath() string {
	return filepath.Join(DataDir(), "xp.json")
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
	dir := DecksDir()

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

func CreateDeck(name string) string {
	path := deckPath(name)

	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		deck := Deck{
			ID:    NewID(),
			Name:  name,
			Cards: []Card{},
		}

		if err := SaveDeck(deck); err != nil {
			fmt.Println("Error:", err)
		}
	}

	return path
}
