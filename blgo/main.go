package main

import (
	"github.com/bit101/blgo"
	"github.com/bit101/blgo/base"
)

func main() {
	sketch := base.NewSketch(setup, render)
	sketch.RenderImage(800, 800, 1)
	// sketch.RenderGif(400, 400, 5, 30)
	// sketch.RenderVideo(1280, 720, 10, 30)
}

func setup(surface *blgo.Surface, width, height float64) {
	surface.ClearRGB(1, 1, 1)
}

func render(surface *blgo.Surface, width, height, percent float64) {
	surface.SetSourceRGB(1, 0, 0)
	surface.FillRectangle(0, 0, width*percent, height)
}
