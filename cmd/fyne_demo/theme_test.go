package main

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/0xDezzy/fyne"
	"github.com/0xDezzy/fyne/test"
	"github.com/0xDezzy/fyne/theme"
)

func TestForceVariantTheme_Color(t *testing.T) {
	test.NewApp()
	dark := &forcedVariant{Theme: theme.DefaultTheme(), variant: theme.VariantDark}
	light := &forcedVariant{Theme: theme.DefaultTheme(), variant: theme.VariantLight}
	unspecified := fyne.ThemeVariant(99)

	assert.Equal(t, theme.DefaultTheme().Color(theme.ColorNameForeground, theme.VariantDark), dark.Color(theme.ColorNameForeground, unspecified))
	assert.Equal(t, theme.DefaultTheme().Color(theme.ColorNameForeground, theme.VariantLight), light.Color(theme.ColorNameForeground, unspecified))
}
