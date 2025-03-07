package tui

import (
	"fmt"
	"os"

	"github.com/rivo/tview"
)

// Screen is a shell screen.
type Screen struct {
	*tview.Application
	accessibilityEnabled bool
}

// NewScreen creates a new screen.
func NewScreen() *Screen {
	app := tview.NewApplication()
	// Check if accessibility is enabled via env var
	_, accessibilityEnabled := os.LookupEnv("JIRA_ACCESSIBILITY_MODE")
	return &Screen{
		Application:          app,
		accessibilityEnabled: accessibilityEnabled,
	}
}

// Paint paints UI to the screen.
func (s *Screen) Paint(root tview.Primitive) error {
	return s.SetRoot(root, true).SetFocus(root).Run()
}

// AnnounceToScreenReader outputs text specifically formatted for screen readers.
func (s *Screen) AnnounceToScreenReader(announcement string) {
	if s.accessibilityEnabled {
		// This prints directly to stderr which screen readers typically monitor
		// The special marker helps screen readers identify this as an announcement
		fmt.Fprintf(os.Stderr, "\n[SCREEN_READER_ANNOUNCEMENT] %s\n", announcement)
	}
}
