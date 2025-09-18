package save

import (
	"Projet-Red_GoonLegend/character"
	"encoding/json"
	"os"
)

func SaveCharacter(hero *character.Character, filename string) error {
	data, err := json.MarshalIndent(hero, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}

func LoadCharacter(filename string) (*character.Character, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var hero character.Character
	if err := json.Unmarshal(data, &hero); err != nil {
		return nil, err
	}
	return &hero, nil
}
