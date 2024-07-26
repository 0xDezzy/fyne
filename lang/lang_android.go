package lang

import (
	"github.com/0xDezzy/fyne/internal/driver/mobile/app"

	"github.com/jeandeaual/go-locale"
)

func init() {
	locale.SetRunOnJVM(app.RunOnJVM)
}
