//go:build !debug && !release

package build

import "github.com/0xDezzy/fyne"

// Mode is the application's build mode.
const Mode = fyne.BuildStandard
