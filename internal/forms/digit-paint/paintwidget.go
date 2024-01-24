package digit_paint

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
	"image"
	"image/color"
)

var _ desktop.Hoverable = (*PaintWidget)(nil)

type PaintWidget struct {
	widget.BaseWidget
	Raster *canvas.Raster
	Img    *image.RGBA

	Width  float32
	Height float32
}

func (w *PaintWidget) MouseOut() {
	fmt.Println("OUT")
}

func NewPaintWidget(img *image.RGBA) *PaintWidget {
	sz := img.Bounds()
	for i := 0; i <= sz.Max.X; i++ {
		for j := 0; j <= sz.Max.Y; j++ {
			img.Set(j, i, color.Black)
		}
	}

	w := &PaintWidget{}
	w.Img = img
	w.Raster = canvas.NewRasterWithPixels(func(x, y, w, h int) color.Color {
		return color.Black
	})
	w.Raster.Generator = func(_, _ int) image.Image {
		return img
	}
	w.Resize(fyne.NewSize(float32(sz.Max.X), float32(sz.Max.X)))
	w.Raster.Refresh()
	w.ExtendBaseWidget(w)
	return w
}

func (w *PaintWidget) CreateRenderer() fyne.WidgetRenderer {
	objects := []fyne.CanvasObject{w.Raster}
	return &PaintWidgetRenderer{w.Raster, objects, w}
}

func (w *PaintWidget) commonDraw(ev *desktop.MouseEvent) {
	if ev.Button == 0 {
		return
	}

	//fmt.Println("Mouse Down")
	drawColor := color.White

	if ev.Button == desktop.MouseButtonSecondary {
		drawColor = color.Black
	}

	x := ev.Position.X * float32(w.Img.Rect.Max.X) / w.Width
	y := ev.Position.Y * float32(w.Img.Rect.Max.Y) / w.Height

	drawCircle(w.Img, int(x), int(y), 1, drawColor)
	w.Raster.Generator = func(_, _ int) image.Image {
		return w.Img
	}

	canvas.Refresh(w.Raster)
}

func (w *PaintWidget) MouseDown(ev *desktop.MouseEvent) {
	w.commonDraw(ev)
}

func (w *PaintWidget) MouseUp(ev *desktop.MouseEvent) {
	//fmt.Println("Mouse Up")
}

func (w *PaintWidget) MouseIn(ev *desktop.MouseEvent) {
	//fmt.Println("Mouse In")
}

func (w *PaintWidget) MouseMoved(ev *desktop.MouseEvent) {
	w.commonDraw(ev)
}

func (w *PaintWidget) Resize(size fyne.Size) {
	w.Height = size.Height
	w.Width = size.Width
	w.BaseWidget.Resize(size)
}

func (w *PaintWidget) Clear() *PaintWidget {
	for x := 0; x < w.Img.Rect.Max.X; x++ {
		for y := 0; y < w.Img.Rect.Max.Y; y++ {
			w.Img.Set(x, y, color.Black)
		}
	}

	return w
}

type PaintWidgetRenderer struct {
	raster  *canvas.Raster
	objects []fyne.CanvasObject
	parent  *PaintWidget
}

func (r *PaintWidgetRenderer) Destroy() {
}

func (r *PaintWidgetRenderer) Layout(size fyne.Size) {
	r.raster.Resize(size)
}

func (r *PaintWidgetRenderer) MinSize() fyne.Size {
	return r.raster.MinSize()
}

func (r *PaintWidgetRenderer) Objects() []fyne.CanvasObject {
	return r.objects
}

func (r *PaintWidgetRenderer) Refresh() {
	r.raster = r.parent.Raster
	r.raster.Generator = func(w, h int) image.Image {
		return r.parent.Img
	}
	r.raster.Refresh()
}

//func tttte() {
//	myApp := app.New()
//	label := NewTestWidget("Hello World!")
//	container.NewVBox(label)
//	myWindow := myApp.NewWindow("Label Widget")
//	myWindow.SetContent(container.NewVBox(label))
//	myWindow.ShowAndRun()
//}

func drawCircle(img *image.RGBA, x, y, radius int, colorI color.Color) {
	for i := x - radius; i <= x+radius; i++ {
		for j := y - radius; j <= y+radius; j++ {
			if (i-x)*(i-x)+(j-y)*(j-y) <= radius*radius {
				img.Set(i, j, colorI)
			}
		}
	}
}
