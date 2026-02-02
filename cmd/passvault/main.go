package main

import (
	"log"

	_ "passvault-fyne/internal/bootstrap"
	"passvault-fyne/ui"
)

func main() {
	app, err := ui.NewApp()
	if err != nil {
		log.Fatalf("Failed to initialize app: %v", err)
	}

	app.Start()
}
