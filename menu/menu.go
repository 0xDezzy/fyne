package menu

import (
	"image/color"

	"github.com/0xDezzy/fyne"
	"github.com/0xDezzy/fyne/canvas"
)

type systemTrayDriver interface {
	fyne.Driver
	SetSystemTrayMenu(*Menu)
	SystemTrayMenu() *Menu
}

// Menu stores the information required for a standard menu.
// A menu can pop down from a MainMenu or could be a pop out menu.
type Menu struct {
	Label           string
	Items           []*MenuItem
	backgroundColor color.Color
	backgroundRect  *canvas.Rectangle
}

// NewMenu creates a new menu given the specified label (to show in a MainMenu) and list of items to display.
func NewMenu(label string, items ...*MenuItem) *Menu {
	return &Menu{Label: label, Items: items}
}

// SetBackgroundColor sets the background color for the menu.
func (m *Menu) SetBackgroundColor(color color.Color) {
	m.backgroundColor = color
	m.applyBackgroundColor()
}

// applyBackgroundColor applies the background color to all menu items.
func (m *Menu) applyBackgroundColor() {
	if m.backgroundRect == nil {
		m.backgroundRect = canvas.NewRectangle(m.backgroundColor)
	} else {
		m.backgroundRect.FillColor = m.backgroundColor
	}
	canvas.Refresh(m.backgroundRect)

	for _, item := range m.Items {
		item.SetBackgroundColor(m.backgroundColor)
	}
}

// Refresh will instruct this menu to update its display.
//
// Since: 2.2
func (m *Menu) Refresh() {
	for _, w := range fyne.CurrentApp().Driver().AllWindows() {
		main := w.MainMenu()
		if main != nil {
			for _, menu := range main.Items {
				if menu == m {
					w.SetMainMenu(main)
					break
				}
			}
		}
	}

	if d, ok := fyne.CurrentApp().Driver().(systemTrayDriver); ok {
		if m == d.SystemTrayMenu() {
			d.SetSystemTrayMenu(m)
		}
	}
}

// MenuItem is a single item within any menu, it contains a display Label and Action function that is called when tapped.
type MenuItem struct {
	ChildMenu       *Menu
	IsQuit          bool
	IsSeparator     bool
	Label           string
	Action          func()
	Disabled        bool
	Checked         bool
	Icon            fyne.Resource
	Shortcut        fyne.Shortcut
	backgroundColor color.Color
	backgroundRect  *canvas.Rectangle
}

// NewMenuItem creates a new menu item from the passed label and action parameters.
func NewMenuItem(label string, action func()) *MenuItem {
	return &MenuItem{Label: label, Action: action}
}

// NewMenuItemSeparator creates a menu item that is to be used as a separator.
func NewMenuItemSeparator() *MenuItem {
	return &MenuItem{IsSeparator: true, Action: func() {}}
}

// SetBackgroundColor sets the background color for the menu item.
func (m *MenuItem) SetBackgroundColor(color color.Color) {
	m.backgroundColor = color
	m.applyBackgroundColor()
}

// applyBackgroundColor applies the background color to the menu item.
func (m *MenuItem) applyBackgroundColor() {
	if m.backgroundRect == nil {
		m.backgroundRect = canvas.NewRectangle(m.backgroundColor)
	} else {
		m.backgroundRect.FillColor = m.backgroundColor
	}
	canvas.Refresh(m.backgroundRect)
}

// MainMenu defines the data required to show a menu bar (desktop) or other appropriate top level menu.
type MainMenu struct {
	Items           []*Menu
	backgroundColor color.Color
	backgroundRect  *canvas.Rectangle
}

// NewMainMenu creates a top level menu structure used by fyne.Window for displaying a menubar
// (or appropriate equivalent).
func NewMainMenu(items ...*Menu) *MainMenu {
	return &MainMenu{Items: items}
}

// SetBackgroundColor sets the background color for the main menu.
func (m *MainMenu) SetBackgroundColor(color color.Color) {
	m.backgroundColor = color
	m.applyBackgroundColor()
}

// applyBackgroundColor applies the background color to the main menu and all its items.
func (m *MainMenu) applyBackgroundColor() {
	if m.backgroundRect == nil {
		m.backgroundRect = canvas.NewRectangle(m.backgroundColor)
	} else {
		m.backgroundRect.FillColor = m.backgroundColor
	}
	canvas.Refresh(m.backgroundRect)

	for _, menu := range m.Items {
		menu.SetBackgroundColor(m.backgroundColor)
	}
}

// Refresh will instruct any rendered menus using this struct to update their display.
//
// Since: 2.2
func (m *MainMenu) Refresh() {
	for _, w := range fyne.CurrentApp().Driver().AllWindows() {
		menu := w.MainMenu()
		if menu != nil && menu == m {
			w.SetMainMenu(m)
		}
	}
}
