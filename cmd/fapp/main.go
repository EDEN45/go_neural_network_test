package main

import digit_paint "github.com/EDEN45/go_neural_network_test/internal/forms/digit-paint"

func main() {
	//a := appcli.New()
	//w := a.NewWindow("Hello")
	//w.Resize(fyne.Size{
	//	Width:  800,
	//	Height: 600,
	//})
	//
	//hello := widget.NewLabel("Hello Fyne!")
	//widget.NewCard()
	//w.SetContent(container.NewVBox(
	//	hello,
	//	widget.NewButton("Hi!", func() {
	//		hello.SetText("Welcome :)")
	//	}),
	//))
	//
	//w.ShowAndRun()

	form := digit_paint.NewForm()
	form.ShowAndRun()
}

//func main() {
//	err := glfw.Init(nil)
//	if err != nil {
//		panic(err)
//	}
//	defer glfw.Terminate()
//
//	fmt.Println("Library initialized.")
//
//	window, err := glfw.CreateWindow(640, 480, "Event Linter", nil, nil)
//	if err != nil {
//		panic(err)
//	}
//	//windowIds[window] = 1 // First (and only) window has id 1.
//	//
//	//window.SetPosCallback(PosCallback)
//	//window.SetSizeCallback(SizeCallback)
//	//window.SetFramebufferSizeCallback(FramebufferSizeCallback)
//	//window.SetCloseCallback(CloseCallback)
//	//window.SetRefreshCallback(RefreshCallback)
//	//window.SetFocusCallback(FocusCallback)
//	//window.SetIconifyCallback(IconifyCallback)
//	//window.SetMouseButtonCallback(MouseButtonCallback)
//	//window.SetCursorPosCallback(CursorPosCallback)
//	//window.SetCursorEnterCallback(CursorEnterCallback)
//	//window.SetScrollCallback(ScrollCallback)
//	//window.SetKeyCallback(KeyCallback)
//	//window.SetCharCallback(CharCallback)
//	//window.SetCharModsCallback(CharModsCallback)
//	//window.SetDropCallback(DropCallback)
//
//	fmt.Println("Main loop starting.")
//
//	for !window.ShouldClose() {
//		glfw.WaitEvents()
//	}
//}
