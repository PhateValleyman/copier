// Package config obsahuje konfigurační struktury a funkce pro načítání konfigurace.
package config

import (
	"encoding/json" // pro načítání JSON souboru
	"os"            // pro práci se soubory
)

// Config definuje strukturu konfiguračního souboru.
type Config struct {
	SourceDir string `json:"source_dir"` // adresář odkud kopírovat
	TargetDir string `json:"target_dir"` // cílový adresář
}

// Load načte konfiguraci ze zadaného souboru a vrátí strukturu Config.
func Load(filename string) (*Config, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	var cfg Config
	err = decoder.Decode(&cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
