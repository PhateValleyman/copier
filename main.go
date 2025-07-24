package main

import (
    "log"

	"github.com/PhateValleyman/copier/config"
	"github.com/PhateValleyman/copier/gui"
	"github.com/PhateValleyman/copier/worker"
)

func main() {
    entries, err := config.LoadOrCreate()
    if err != nil {
        log.Fatalf("Config error: %v", err)
    }
    gui.Run(entries)
}
