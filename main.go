package main

import (
	"fmt"
	"log"
//	"os"

	"github.com/PhateValleyman/copier/config"
	"github.com/PhateValleyman/copier/gui"
	"github.com/PhateValleyman/copier/worker"
)

func main() {
	// Načteme config
	cfg, err := config.Load("config.json")
	if err != nil {
		fmt.Println("❌ Chyba při načítání konfigurace:", err)
		// alternativně nabídneme ruční zadání
		cfg = &config.Config{
			SourceDir: gui.AskInput("Zadej cestu ke zdroji: "),
			TargetDir: gui.AskInput("Zadej cestu k cíli: "),
		}
	}

	// Spustíme kopírování
	fmt.Println("🔄 Kopíruji soubory...")
	err = worker.CopyDirectory(cfg.SourceDir, cfg.TargetDir)
	if err != nil {
		log.Fatalf("❌ Kopírování selhalo: %v", err)
	}

	fmt.Println("✅ Hotovo!")
}
