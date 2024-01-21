package digit_paint

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"image"
	"image/color"
)

var clorBlack = color.NRGBA{R: 0, G: 0, B: 0, A: 255}

func NewForm() fyne.Window {
	paintApp := app.New()

	w := paintApp.NewWindow("digit paint")
	w.Resize(fyne.Size{
		Width:  2000,
		Height: 1000,
	})

	vbox := container.NewVBox()
	vbox.Resize(fyne.Size{
		Width: 200,
	})

	iig := image.NewRGBA(image.Rectangle{Min: image.Point{X: 0, Y: 0}, Max: image.Point{X: 1800, Y: 800}})
	pw := NewPaintWidget(iig)
	pw.Resize(fyne.NewSize(1800, 800))
	pw.Refresh()
	grw := container.NewGridWrap(fyne.NewSize(1800, 800), pw)
	cn := container.NewHBox(grw, vbox)

	w.SetContent(cn)

	//w.Resize(fyne.NewSize(100, 100))
	return w
}
