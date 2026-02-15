package internal

import (
	"encoding/json"
	"os"
)

func LoadXP() (int, error) {
	path := xpPath()

	f, err := os.Open(path)
	if err != nil {
		if os.IsNotExist(err) {
			return 0, nil
		}
		return 0, err
	}
	defer f.Close()

	var data struct {
		XP int `json:"xp"`
	}

	err = json.NewDecoder(f).Decode(&data)
	if err != nil {
		return 0, err
	}

	return data.XP, nil
}

// SaveXP writes the current XP to disk.
func SaveXP(xp int) error {
	path := xpPath()

	data := struct {
		XP int `json:"xp"`
	}{
		XP: xp,
	}

	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	return json.NewEncoder(f).Encode(data)
}
