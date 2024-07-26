package mobile

import (
	"testing"

	_ "github.com/0xDezzy/fyne/test"
)

func TestDevice_HideVirtualKeyboard_BeforeRun(t *testing.T) {
	hideVirtualKeyboard() // should not crash!
}
