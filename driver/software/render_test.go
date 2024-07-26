package software

import (
	"image/color"
	"testing"

	"github.com/0xDezzy/fyne"
	"github.com/0xDezzy/fyne/canvas"
	"github.com/0xDezzy/fyne/container"
	"github.com/0xDezzy/fyne/internal/painter"
	"github.com/0xDezzy/fyne/test"

	"github.com/0xDezzy/fyne/theme"
	"github.com/0xDezzy/fyne/widget"
)

func TestRender(t *testing.T) {
	obj := widget.NewLabel("Hi")
	test.AssertImageMatches(t, "label.png", Render(obj, test.Theme()))
	painter.ClearFontCache() // avoid side effects of the cause of #4937
	test.AssertImageMatches(t, "label_ugly_theme.png", Render(obj, test.NewTheme()))
	painter.ClearFontCache() // avoid side effects of the cause of #4937
}

func TestRender_State(t *testing.T) {
	obj := widget.NewButtonWithIcon("Cancel", theme.CancelIcon(), func() {})
	test.AssertImageMatches(t, "button.png", Render(obj, test.Theme()))

	obj.Importance = widget.HighImportance
	obj.Refresh()
	test.AssertImageMatches(t, "button_important.png", Render(obj, test.Theme()))
}

func TestRender_Focus(t *testing.T) {
	obj := widget.NewEntry()
	test.AssertImageMatches(t, "entry.png", Render(obj, test.Theme()))

	obj.FocusGained()
	test.AssertImageMatches(t, "entry_focus.png", Render(obj, test.Theme()))
}

func TestRenderCanvas(t *testing.T) {
	obj := container.NewAppTabs(
		container.NewTabItem("Tab 1", container.NewVBox(
			widget.NewLabel("Label"),
			widget.NewButton("Button", func() {}),
		)))

	c := NewCanvas()
	c.SetContent(obj)

	if fyne.CurrentDevice().IsMobile() {
		test.AssertImageMatches(t, "canvas_mobile.png", RenderCanvas(c, test.Theme()))
	} else {
		test.AssertImageMatches(t, "canvas.png", RenderCanvas(c, test.Theme()))
	}
}

func TestRender_ImageSize(t *testing.T) {
	image := canvas.NewImageFromFile("../../theme/icons/fyne.png")
	image.FillMode = canvas.ImageFillOriginal
	bg := canvas.NewCircle(color.NRGBA{255, 0, 0, 128})
	bg.StrokeColor = color.White
	bg.StrokeWidth = 5

	c := container.NewStack(image, bg)

	test.AssertImageMatches(t, "image_size.png", Render(c, test.Theme()))
}
