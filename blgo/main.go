package main

import (
	"${MOD_NAME}/base"

	"github.com/bit101/blgo"
)

func main() {
	base.RenderImage(800, 800, setup, render)
	// base.RenderGif(400, 400, 5, 30, setup, render)
	// base.RenderVideo(1280, 720, 10, 30, setup, render)
}

func setup(surface *blgo.Surface, width, height float64) {
	surface.ClearRGB(1, 1, 1)
}

func render(surface *blgo.Surface, width, height, percent float64) {
	surface.FillRectangle(50, 50, width - 100, height - 100)

}
