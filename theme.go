package main

import (
    "fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	"image/color"
)
type myTheme struct {}
var _ fyne.Theme = (*myTheme)(nil)
func (m myTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	if name == theme.IconNameHome {
		return fyne.NewStaticResource("myHome", []byte{})
	}
	
	return theme.DefaultTheme().Icon(name)
}
func (m myTheme) Font(style fyne.TextStyle) fyne.Resource {
	return theme.DefaultTheme().Font(style)
}

func (m myTheme) Size(name fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(name)
}

func (m myTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	if name == theme.ColorNameBackground {
		return color.Black
	}

	if name == theme.ColorNameButton {
		return color.RGBA{0,255,0,255}
	}

	if name == theme.ColorNameBackground {
		return color.RGBA{0,255,0,255}
	}

	return theme.DefaultTheme().Color(name, variant)
}