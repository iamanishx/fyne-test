package ui

import (
	"bytes"
	"encoding/json"
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
		contentText := secret.Content
		if secret.Format == "json" {
			var pretty bytes.Buffer
			if err := json.Indent(&pretty, []byte(secret.Content), "", "  "); err == nil {
				contentText = pretty.String()
			}
		}
		contentEntry := widget.NewMultiLineEntry()
		contentEntry.SetText(contentText)
		contentEntry.Wrapping = fyne.TextWrapWord
		contentEntry.Scroll = container.ScrollVerticalOnly
		contentEntry.SetMinRowsVisible(18)
		copyBtn := widget.NewButton("Copy", func() {
			selected := contentEntry.SelectedText()
			if selected == "" {
				selected = contentText
			}
			a.Clipboard.CopyWithAutoClear(selected, 30*time.Second)
		})
		row := container.NewHBox(widget.NewLabel("Content:"), layout.NewSpacer(), copyBtn)
		fieldsContainer.Add(row)
		fieldsContainer.Add(contentEntry)
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
