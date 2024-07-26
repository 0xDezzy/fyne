package common

import (
	"github.com/0xDezzy/fyne"
	"github.com/0xDezzy/fyne/internal/cache"
)

// CanvasForObject returns the canvas for the specified object.
func CanvasForObject(obj fyne.CanvasObject) fyne.Canvas {
	return cache.GetCanvasForObject(obj)
}
