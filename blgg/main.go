package main

import (
	"github.com/bit101/blgg/blgg"
	"github.com/bit101/blgg/blmath"
	"github.com/bit101/blgg/render"
	"github.com/bit101/blgg/util"
)

func main() {
	target := render.Gif

	switch target {
	case render.Image:
		render.RenderImage(800, 800, "out.png", renderFrame, 0.5)
		util.ViewImage("out.png")
		break

	case render.Gif:
		render.RenderFrames(400, 400, 60, "frames", renderFrame)
		util.MakeGIF("ffmpeg", "frames", "out.gif", 30)
		util.ViewImage("out.gif")
		break

	case render.Video:
		render.RenderFrames(1280, 800, 60, "frames", renderFrame)
		util.ConvertToYoutube("frames", "out.mp4", 60)
		util.VLC("out.mp4", true)
		break
	}
}

func renderFrame(context *blgg.Context, width, height, percent float64) {
	context.BlackOnWhite()
	x := blmath.LerpSin(percent, 0, width-50)
	context.FillRectangle(x, height/2-25, 50, 50)
}
