package clipboard

import (
	"time"

	"fyne.io/fyne/v2"
)

type Manager struct {
	clipboard fyne.Clipboard
}

func NewManager(c fyne.Clipboard) *Manager {
	return &Manager{clipboard: c}
}

func (m *Manager) CopyWithAutoClear(content string, duration time.Duration) {
	m.clipboard.SetContent(content)

	go func() {
		time.Sleep(duration)
		if m.clipboard.Content() == content {
			m.clipboard.SetContent("")
		}
	}()
}
