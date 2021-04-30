package main

import "github.com/gotk3/gotk3/gtk"

func main() {
	gtk.Init(nil)

	win, _ := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	win.SetSizeRequest(800, 800)
	win.SetTitle("Hello Gotk3!")
	win.Connect("destroy", gtk.MainQuit)

	win.ShowAll()
	gtk.Main()
}
