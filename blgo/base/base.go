package base

import (
	"github.com/bit101/blgo"
	"github.com/bit101/blgo/anim"
	"github.com/bit101/blgo/util"
)

type RenderFunc func(surface *blgo.Surface, percent float64)

var (
	surface    *blgo.Surface
	renderFunc RenderFunc
)

const framesDir = "frames" // Must Exist!!!

func RenderImage(width, height float64, render RenderFunc) {
	outFileName := "out.png"
	surface = blgo.NewSurface(width, height)
	render(surface, 0)
	surface.WriteToPNG(outFileName)
	util.ViewImage(outFileName)
}

func RenderGif(width, height float64, seconds, fps int, render RenderFunc) {
	outFileName := "out.gif"
	renderAnim(width, height, seconds, fps, render)
	util.ConvertToGIF(framesDir, outFileName, float64(fps))
	util.ViewImage(outFileName)
}

func RenderVideo(width, height float64, seconds, fps int, render RenderFunc) {
	outFileName := "out.mp4"
	renderAnim(width, height, seconds, fps, render)
	util.ConvertToYoutube(framesDir, outFileName, fps)
	util.VLC(outFileName)
}

func renderAnim(width, height float64, seconds, fps int, render RenderFunc) {
	framesDir := "frames" // MUST EXIST!!!
	surface = blgo.NewSurface(width, height)
	renderFunc = render
	animation := anim.NewAnimation(surface, fps*seconds, framesDir)
	animation.Render(internalRender)
}

func internalRender(percent float64) {
	renderFunc(surface, percent)
}
