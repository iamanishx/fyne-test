package ui

import (
	"image/color"

	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func (a *App) ShowMainUI() {
	sidebar := a.makeSidebar()
	label := widget.NewLabel("Select a secret to view details")

	bg := canvas.NewRectangle(color.NRGBA{R: 12, G: 12, B: 16, A: 255})
	content := container.NewMax(bg, container.NewPadded(label))

	split := container.NewHSplit(sidebar, content)
	split.SetOffset(0.3)

	a.MainWindow.SetContent(split)
	a.refreshSecretList()
}
