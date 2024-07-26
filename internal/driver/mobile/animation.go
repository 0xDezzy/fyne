package mobile

import "github.com/0xDezzy/fyne"

func (d *driver) StartAnimation(a *fyne.Animation) {
	d.animation.Start(a)
}

func (d *driver) StopAnimation(a *fyne.Animation) {
	d.animation.Stop(a)
}
