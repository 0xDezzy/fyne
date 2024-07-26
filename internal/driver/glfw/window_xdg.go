//go:build linux || freebsd || openbsd || netbsd

package glfw

import "github.com/0xDezzy/fyne"

func (w *window) platformResize(canvasSize fyne.Size) {
	w.canvas.Resize(canvasSize)
}
