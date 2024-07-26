//go:build wasm || test_web_driver

package glfw

import "github.com/0xDezzy/fyne"

func (d *gLDriver) SetSystemTrayMenu(m *fyne.Menu) {
	// no-op for mobile apps using this driver
}

func (d *gLDriver) catchTerm() {
}

func setDisableScreenBlank(disable bool) {
	// awaiting complete support for WakeLock
}
