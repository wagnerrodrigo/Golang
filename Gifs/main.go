package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"time"
)

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0 // primeira cor da paleta
	blackIndex = 1 // próxima cor da paleta
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycle   = 5     // número de revoluções completas do oscilador x
		res     = 0.001 // resolução angular
		size    = 100   // canvas da imagem cobre de  [-size ... +size]
		nframes = 64    // números de quadros da animação
		delay   = 8     // tempo entre quadros em unidades de 10 ms
	)
	freq := rand.Float64() * 3.0 // frequência relativa do oscilador y
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // diferença de fase
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycle*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
