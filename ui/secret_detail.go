package ui

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func (a *App) showSecretDetails(id string) {
	secret, err := a.DB.GetSecret(id)
	if err != nil {
		dialog.ShowError(err, a.MainWindow)
		return
	}

	if secret == nil {
		return
	}

	nameLabel := widget.NewLabelWithStyle(secret.Name, fyne.TextAlignLeading, fyne.TextStyle{Bold: true})
	fieldsContainer := container.NewVBox()
	if secret.Format == "json" || secret.Format == "text" {
		contentLabel := widget.NewLabel(secret.Content)
		contentLabel.Wrapping = fyne.TextWrapWord
		copyBtn := widget.NewButton("Copy", func() {
			a.Clipboard.CopyWithAutoClear(secret.Content, 30*time.Second)
		})
		row := container.NewHBox(widget.NewLabel("Content:"), layout.NewSpacer(), copyBtn)
		fieldsContainer.Add(row)
		fieldsContainer.Add(contentLabel)
	} else {
		for _, field := range secret.Fields {
			valStr := string(field.Value)
			if field.IsSensitive {
				valStr = "********"
			}

			valLabel := widget.NewLabel(valStr)
			keyLabel := widget.NewLabelWithStyle(field.Key+":", fyne.TextAlignTrailing, fyne.TextStyle{Bold: true})

			copyBtn := widget.NewButton("Copy", func() {
				a.Clipboard.CopyWithAutoClear(string(field.Value), 30*time.Second)
			})

			row := container.NewHBox(keyLabel, valLabel, layout.NewSpacer(), copyBtn)
			fieldsContainer.Add(row)
		}
	}

	editBtn := widget.NewButton("Edit", func() {
		a.ShowEditSecretDialog(secret)
	})

	deleteBtn := widget.NewButton("Delete", func() {
		dialog.ShowConfirm("Delete Secret", "Are you sure?", func(ok bool) {
			if ok {
				if err := a.DB.DeleteSecret(secret.ID); err != nil {
					dialog.ShowError(err, a.MainWindow)
				} else {
					a.refreshSecretList()
					empty := widget.NewLabel("Select a secret")

					split := a.MainWindow.Content().(*container.Split)
					split.Trailing = empty
					split.Refresh()
				}
			}
		}, a.MainWindow)
	})

	topBar := container.NewBorder(nil, nil, nameLabel, container.NewHBox(editBtn, deleteBtn), nil)

	detailView := container.NewVBox(topBar, widget.NewSeparator(), fieldsContainer)

	split := a.MainWindow.Content().(*container.Split)
	split.Trailing = container.NewPadded(detailView)
	split.Refresh()
}
