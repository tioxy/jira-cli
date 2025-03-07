package main

import (
	"fmt"
	"os"

	"github.com/ankitpokhrel/jira-cli/internal/cmd/root"
)

func main() {
	// Check if accessibility help was requested
	for i, arg := range os.Args {
		if arg == "--accessibility-help" || arg == "-a-help" {
			displayAccessibilityHelp()
			return
		}

		// Enable accessibility mode if requested
		if arg == "--accessibility" || arg == "-a" {
			if err := os.Setenv("JIRA_ACCESSIBILITY_MODE", "1"); err != nil {
				fmt.Fprintf(os.Stderr, "Error setting environment variable: %s\n", err)
			}

			// Remove the flag from arguments to prevent it from causing issues
			if i < len(os.Args)-1 {
				os.Args = append(os.Args[:i], os.Args[i+1:]...)
			} else {
				os.Args = os.Args[:i]
			}
			break
		}
	}

	rootCmd := root.NewCmdRoot()
	if _, err := rootCmd.ExecuteC(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func displayAccessibilityHelp() {
	help := `
Jira CLI Accessibility Features
===============================

Jira CLI includes accessibility features for screen readers.

Usage: 
  jira [command] --accessibility    Enable accessibility features
  jira --accessibility-help         Display this help message

When accessibility mode is enabled:
- All UI interactions will be announced for screen readers
- Navigation changes will be verbalized
- Additional key commands will be available

Keyboard shortcuts in accessibility mode:
- Ctrl+S: Speak the current selection
- Ctrl+A: Hear accessibility help message
- Arrow keys: Navigate between items
- Tab: Move between UI sections
- Enter: Select the current item

To enable accessibility mode permanently, set the environment variable:
  export JIRA_ACCESSIBILITY_MODE=1
`
	fmt.Println(help)
}
