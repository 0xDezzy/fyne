package playground

import (
	"github.com/0xDezzy/fyne/driver/software"
	"github.com/0xDezzy/fyne/test"
)

// NewSoftwareCanvas creates a new canvas in memory that can render without hardware support
func NewSoftwareCanvas() test.WindowlessCanvas {
	return software.NewCanvas()
}
