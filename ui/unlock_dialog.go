package ui

import (
	"crypto/rand"
	"image/color"

	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"

	"passvault-fyne/internal/crypto"
)

func (a *App) ShowUnlockDialog() {
	passwordEntry := widget.NewPasswordEntry()
	passwordEntry.SetPlaceHolder("Master Password")

	errorLabel := widget.NewLabel("")
	errorLabel.Hide()

	card := canvas.NewRectangle(color.NRGBA{R: 18, G: 18, B: 22, A: 255})
	card.CornerRadius = 14
	card.StrokeColor = color.NRGBA{R: 40, G: 40, B: 46, A: 200}
	card.StrokeWidth = 1

	unlockBtn := widget.NewButton("Unlock", func() {
		password := []byte(passwordEntry.Text)
		if len(password) == 0 {
			errorLabel.SetText("Enter a master password")
			errorLabel.Show()
			return
		}

		salt := make([]byte, 16)
		if _, err := rand.Read(salt); err != nil {
			errorLabel.SetText("Failed to unlock")
			errorLabel.Show()
			return
		}

		key := crypto.DeriveKey(password, salt)
		a.State.SetMasterKey(key)
		a.ShowMainUI()
	})

	exitBtn := widget.NewButton("Exit", func() {
		a.FyneApp.Quit()
	})

	content := container.NewVBox(
		widget.NewLabel("Enter your Master Password to unlock PassVault"),
		passwordEntry,
		errorLabel,
		container.NewHBox(layout.NewSpacer(), exitBtn, unlockBtn),
	)
	content = container.NewPadded(content)
	cardWrap := container.NewMax(card, content)
	pageBg := canvas.NewRectangle(color.NRGBA{R: 12, G: 12, B: 16, A: 255})
	root := container.NewMax(pageBg, container.NewCenter(cardWrap))

	passwordEntry.OnSubmitted = func(_ string) {
		unlockBtn.OnTapped()
	}

	a.MainWindow.SetContent(root)
	a.MainWindow.Canvas().Focus(passwordEntry)
}
