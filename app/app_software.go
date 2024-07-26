//go:build ci

package app

import (
	"github.com/0xDezzy/fyne"
	"github.com/0xDezzy/fyne/internal/painter/software"
	"github.com/0xDezzy/fyne/test"
)

// NewWithID returns a new app instance using the test (headless) driver.
// The ID string should be globally unique to this app.
func NewWithID(id string) fyne.App {
	return newAppWithDriver(test.NewDriverWithPainter(software.NewPainter()), id)
}
