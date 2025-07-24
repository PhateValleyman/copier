package config

import (
    "fmt"
    "os"
    "path/filepath"

    "gopkg.in/yaml.v3"

    "github.com/PhateValleyman/copier/config"
    "github.com/PhateValleyman/copier/worker"
)

type Entry struct {
    Name   string `yaml:"name"`
    Source string `yaml:"source"`
    Target string `yaml:"target"`
}

func LoadOrCreate() ([]Entry, error) {
    cfgDir := filepath.Join(os.Getenv("HOME"), ".config", "copier")
    cfgPath := filepath.Join(cfgDir, "dirs.yml")

    if err := os.MkdirAll(cfgDir, 0755); err != nil {
        return nil, fmt.Errorf("failed to create config dir: %w", err)
    }

    if _, err := os.Stat(cfgPath); os.IsNotExist(err) {
        defaultContent := `- name: Pohadky
  source: "` + filepath.Join(os.Getenv("HOME"), "pohadky") + `"
  target: "` + filepath.Join(os.Getenv("HOME"), "pohadky_backup") + `"`
        if err := os.WriteFile(cfgPath, []byte(defaultContent), 0644); err != nil {
            return nil, fmt.Errorf("failed to write default config: %w", err)
        }
    }

    data, err := os.ReadFile(cfgPath)
    if err != nil {
        return nil, fmt.Errorf("failed to read config: %w", err)
    }

    var entries []Entry
    if err := yaml.Unmarshal(data, &entries); err != nil {
        return nil, fmt.Errorf("failed to parse config: %w", err)
    }
    return entries, nil
}
