package digit_paint

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"image"
	"image/color"
)

var clorBlack = color.NRGBA{R: 0, G: 0, B: 0, A: 255}

func NewForm() fyne.Window {
	paintApp := app.New()

	w := paintApp.NewWindow("digit paint")
	w.Resize(fyne.Size{
		Width:  800,
		Height: 605,
	})

	vbox := container.NewVBox()
	vbox.Resize(fyne.Size{
		Width: 200,
	})

	iig := image.NewRGBA(image.Rectangle{Min: image.Point{X: 0, Y: 0}, Max: image.Point{X: 28, Y: 28}})
	pw := NewPaintWidget(iig)
	pw.Resize(fyne.NewSize(600, 600))
	pw.Refresh()

	vbox.Add(widget.NewButton("Очистить", func() {
		pw.Clear().Refresh()
	}))

	pbs := make([]*widget.ProgressBar, 0, 10)

	for i := 0; i < 10; i++ {
		lb := widget.NewLabel(fmt.Sprint(i))
		pb := widget.NewProgressBar()
		vbox.Add(container.NewHBox(lb, pb))

		pbs = append(pbs, pb)
	}

	grw := container.NewGridWrap(fyne.NewSize(600, 600), pw)
	cn := container.NewHBox(grw, vbox)

	w.SetContent(cn)

	//w.Resize(fyne.NewSize(100, 100))
	return w
}
