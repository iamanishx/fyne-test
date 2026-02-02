package ui

import (
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"

	"passvault-fyne/internal/database"
)

func (a *App) makeSidebar() fyne.CanvasObject {
	a.Search = widget.NewEntry()
	a.Search.SetPlaceHolder("Search")
	a.Search.OnChanged = func(_ string) {
		a.applyFilter()
	}

	a.List = widget.NewList(
		func() int {
			return len(a.Filtered)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(a.Filtered[i].Name)
		},
	)

	a.List.OnSelected = func(id widget.ListItemID) {
		if id >= 0 && id < len(a.Filtered) {
			a.showSecretDetails(a.Filtered[id].ID)
		}
	}

	addBtn := widget.NewButton("Add Secret", func() {
		a.ShowAddSecretDialog()
	})

	return container.NewBorder(container.NewVBox(a.Search), addBtn, nil, nil, a.List)
}

func (a *App) refreshSecretList() {
	secrets, err := a.DB.GetSecrets()
	if err != nil {
		dialog.ShowError(err, a.MainWindow)
		return
	}

	a.Secrets = secrets
	a.applyFilter()
}

func (a *App) applyFilter() {
	query := strings.ToLower(strings.TrimSpace(a.Search.Text))
	if query == "" {
		a.Filtered = a.Secrets
	} else {
		filtered := make([]database.SecretEntry, 0, len(a.Secrets))
		for _, s := range a.Secrets {
			if strings.Contains(strings.ToLower(s.Name), query) {
				filtered = append(filtered, s)
			}
		}
		a.Filtered = filtered
	}

	if a.List != nil {
		a.List.Refresh()
	}
}
