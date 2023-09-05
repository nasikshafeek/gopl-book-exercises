// lissajous generates GIF animations of random Lissajous figures.
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

var pallete = []color.Color{color.White, color.RGBA{R: 0xff, A: 255}, color.RGBA{G: 0xff, A: 255}, color.RGBA{B: 0xff, A: 255}} // Green color lines instead of black (Passing named parameters)
// var pallete = []color.Color{color.White, color.RGBA{0x00, 0xff, 0x00, 0xff}} // Green color lines instead of black (Passing positional parameters)

const (
	whiteIndex = 0 // first color in pallete
	blackIndex = 1 // next color in pallete
)

func main() {
	lassojous(os.Stdout)
}

func lassojous(writer io.Writer) {
	const (
		nFrames = 64
		size = 100
		res = 0.001
		cycles = 5
		delay = 8
	)

	frequency :=  rand.Float64() * 3.5 // 
	phase := 0.0;
	anim := gif.GIF{LoopCount: nFrames}

	for eachFrame := 0; eachFrame < nFrames; eachFrame ++ {
		imageFrame := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(imageFrame, pallete)
		
		// color index here will update the color for each frame
		// colorIndex := uint8(uint32(oscillator) % 3) + 1

		for oscillator := 0.0; oscillator < 2*math.Pi*cycles; oscillator += res {
			x := math.Sin(oscillator)
			y := math.Sin(oscillator * frequency + phase)
			colorIndex := uint8(uint32(oscillator) % 3) + 1 // Playing around with color index and oscillator to produce same color consistenlty for a few stretch
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), colorIndex)
		}

		phase += 0.01
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}

	gif.EncodeAll(writer, &anim)
}