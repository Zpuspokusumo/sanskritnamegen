package main

import (
	"example/sanskritgen/services"
	"image/color"
	"math/rand"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

/*type mytheme struct{}

func (m myTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	if name == theme.ColorNameBackground {
		if variant == theme.VariantLight {
			return color.White
		}
		return color.Black
	}

	return theme.DefaultTheme().Color(name, variant)
}*/

//if you are looking for the API code, go to servermain.go
//this project went from an API to a GUI App

// go get fyne.io/fyne/v2@latest
func main() {
	a := app.New()
	win := a.NewWindow("Sanskrit Name Generator")
	syz := fyne.NewSize(1000, 600)
	win.Resize(syz)

	title := canvas.NewText("Sanskrit Name Generator", color.Black)
	title.TextStyle = fyne.TextStyle{
		Bold: true,
	}
	title.Alignment = fyne.TextAlignCenter
	title.TextSize = 24

	attri := canvas.NewText("By Zpuspokusumo", color.Black)
	attri.TextStyle = fyne.TextStyle{
		Bold: true,
	}
	attri.Alignment = fyne.TextAlignCenter
	attri.TextSize = 20

	//a.Settings().SetTheme(&myTheme{})
	//var _ fyne.Theme = (*mytheme)(nil)

	names := services.Sktnamegen10()
	handom := rand.NewSource(time.Now().UnixNano())
	name := services.Sktnamegen(handom)

	namebox1 := widget.NewLabel(name)
	namebox1.Wrapping = fyne.TextWrapWord
	namebox1.Alignment = fyne.TextAlignCenter

	//list widget for 10 names
	list10 := widget.NewList(
		func() int {
			return len(names)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("Names?")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(names[i])
			o.(*widget.Label).TextStyle = fyne.TextStyle{Bold: true}
		})

	button1 := widget.NewButton("Generate 1 name", func() {
		handom := rand.NewSource(time.Now().UnixNano())
		name := services.Sktnamegen(handom)
		namebox1.SetText(name)
		namebox1.Refresh()
	})
	//updates list10
	button10 := widget.NewButton("Generate 10 names",
		func() {
			names = services.Sktnamegen10()
			list10.Refresh()
		})

	hBox := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), button1, layout.NewSpacer(), button10, layout.NewSpacer())
	vBox := container.New(layout.NewVBoxLayout(), title, attri, hBox)
	nBox := container.New(layout.NewVBoxLayout(), vBox, namebox1, layout.NewSpacer())
	lBox := container.New(layout.NewMaxLayout(), list10)

	win.SetContent(
		container.NewBorder(
			nBox,
			nil,
			nil,
			nil,
			lBox,
		),
	)

	/*win.SetContent(
		container.NewVSplit(
			vBox,
			list10,
		),
	)*/
	//vBox.Resize(namebox1.MinSize())
	win.ShowAndRun()
}
