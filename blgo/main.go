package main

import (
	"${MOD_NAME}/base"

	"github.com/bit101/blgo"
)

const mode = "image"

func main() {
	switch mode {
	case "image":
		base.RenderImage(800, 800, render)
	case "gif":
		base.RenderGif(400, 400, 5, 30, render)
	case "video":
		base.RenderVideo(1280, 720, 10, 30, render)
	}
}

func render(surface *blgo.Surface, percent float64) {
	surface.ClearRGB(percent, 0, 1-percent)

}
