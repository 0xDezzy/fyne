//go:build !hints

package theme

import (
	"image/color"

	"github.com/0xDezzy/fyne"
)

var (
	fallbackColor = color.Transparent
	fallbackIcon  = &fyne.StaticResource{}
)
