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

// For accessibility announcements, we hide all output from the user and only
// send it to screen readers using a feature of the terminal called ANSI escape codes.
// This approach ensures screen reader announcements don't corrupt the UI.

// AnnounceToScreenReader outputs text specifically formatted for screen readers without breaking the UI.
func (s *Screen) AnnounceToScreenReader(announcement string) {
	if s.accessibilityEnabled {
		// Create the formatted announcement
		text := fmt.Sprintf("[SCREEN_READER_ANNOUNCEMENT] %s", announcement)
		
		// Use ANSI escape sequences to make text invisible to users but available to screen readers
		// \033[8m is the "conceal" escape code 
		// The screen reader will still read it but it won't show on screen
		invisibleText := fmt.Sprintf("\033[8m%s\033[0m", text)
		
		// Position at the end of the line and output invisibly
		// This avoids messing up the visible UI
		fmt.Fprint(os.Stderr, invisibleText)
	}
}
