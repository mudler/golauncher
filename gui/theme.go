package gui

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

// See also https://github.com/lusingander/fyne-theme-generator for generating colors
var _ fyne.Theme = &myTheme{}

type myTheme struct {
	ColorNameBackground      color.NRGBA
	ColorNameButton          color.NRGBA
	ColorNameDisabledButton  color.NRGBA
	ColorNameDisabled        color.NRGBA
	ColorNameError           color.NRGBA
	ColorNameFocus           color.NRGBA
	ColorNameForeground      color.NRGBA
	ColorNameHover           color.NRGBA
	ColorNameInputBackground color.NRGBA
	ColorNamePlaceHolder     color.NRGBA
	ColorNamePressed         color.NRGBA
	ColorNamePrimary         color.NRGBA
	ColorNameScrollBar       color.NRGBA
	ColorNameShadow          color.NRGBA
	DefaultColor             color.NRGBA
}

func NewTheme() *myTheme {
	return &myTheme{
		ColorNameButton:          color.NRGBA{R: 0xff, G: 0xff, B: 0xff, A: 0xf},
		ColorNameBackground:      color.NRGBA{R: 0x30, G: 0x30, B: 0x30, A: 0xff},
		ColorNameDisabledButton:  color.NRGBA{R: 0x26, G: 0x26, B: 0x26, A: 0xff},
		ColorNameDisabled:        color.NRGBA{R: 0xff, G: 0xff, B: 0xff, A: 0x42},
		ColorNameError:           color.NRGBA{R: 0xf4, G: 0x43, B: 0x36, A: 0xff},
		ColorNameFocus:           color.NRGBA{R: 0x21, G: 0x96, B: 0xf3, A: 0x7f},
		ColorNameForeground:      color.NRGBA{R: 0xff, G: 0xff, B: 0xff, A: 0xff},
		ColorNameHover:           color.NRGBA{R: 0xff, G: 0xff, B: 0xff, A: 0xf},
		ColorNameInputBackground: color.NRGBA{R: 46, G: 46, B: 46, A: 254},
		ColorNamePlaceHolder:     color.NRGBA{R: 0xb2, G: 0xb2, B: 0xb2, A: 0xff},
		ColorNamePressed:         color.NRGBA{R: 0xff, G: 0xff, B: 0xff, A: 0x66},
		ColorNamePrimary:         color.NRGBA{R: 0x21, G: 0x96, B: 0xf3, A: 0xff},
		ColorNameScrollBar:       color.NRGBA{R: 0x0, G: 0x0, B: 0x0, A: 0x99},
		ColorNameShadow:          color.NRGBA{R: 0x0, G: 0x0, B: 0x0, A: 0x66},
		DefaultColor:             color.NRGBA{R: 0x21, G: 0x96, B: 0xf3, A: 0xff},
	}
}

func (m *myTheme) Color(c fyne.ThemeColorName, v fyne.ThemeVariant) color.Color {
	switch c {
	case theme.ColorNameBackground:
		return m.ColorNameBackground
	case theme.ColorNameButton:
		return m.ColorNameButton
	case theme.ColorNameDisabledButton:
		return m.ColorNameDisabledButton
	case theme.ColorNameDisabled:
		return m.ColorNameDisabled
	case theme.ColorNameError:
		return m.ColorNameError
	case theme.ColorNameFocus:
		return m.ColorNameFocus
	case theme.ColorNameForeground:
		return m.ColorNameForeground
	case theme.ColorNameHover:
		return m.ColorNameHover
	case theme.ColorNameInputBackground:
		return m.ColorNameInputBackground
	case theme.ColorNamePlaceHolder:
		return m.ColorNamePlaceHolder
	case theme.ColorNamePressed:
		return m.ColorNamePressed
	case theme.ColorNamePrimary:
		return m.ColorNamePrimary
	case theme.ColorNameScrollBar:
		return m.ColorNameScrollBar
	case theme.ColorNameShadow:
		return m.ColorNameShadow
	default:
		return m.DefaultColor
	}
}

func (myTheme) Font(s fyne.TextStyle) fyne.Resource {
	if s.Monospace {
		return theme.DefaultTheme().Font(s)
	}
	if s.Bold {
		if s.Italic {
			return theme.DefaultTheme().Font(s)
		}
		return theme.DefaultTheme().Font(s)
	}
	if s.Italic {
		return theme.DefaultTheme().Font(s)
	}
	return theme.DefaultTheme().Font(s)
}

func (myTheme) Icon(n fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(n)
}

func (myTheme) Size(s fyne.ThemeSizeName) float32 {
	switch s {
	case theme.SizeNameCaptionText:
		return 11
	case theme.SizeNameInlineIcon:
		return 20
	case theme.SizeNamePadding:
		return 4
	case theme.SizeNameScrollBar:
		return 16
	case theme.SizeNameScrollBarSmall:
		return 3
	case theme.SizeNameSeparatorThickness:
		return 1
	case theme.SizeNameText:
		return 14
	case theme.SizeNameInputBorder:
		return 2
	default:
		return theme.DefaultTheme().Size(s)
	}
}
