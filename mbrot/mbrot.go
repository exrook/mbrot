package main

import (
	"flag"
	"fmt"
	"github.com/exrook/mbrot"
	"github.com/exrook/mbrot/palettes"
	"image/color"
	"image/png"
	"os"
	"runtime"
)

var paletteList = map[string]func(int) color.Palette{
	"gray":  palettes.Gray,
	"red":   palettes.Red,
	"green": palettes.Green,
	"blue":  palettes.Blue,
}

var fractalList = map[string]mbrot.Fractaler{
	"mbrot": mbrot.Mandelbrot{},
	"bship": mbrot.BurningShip{},
	"nbrot": mbrot.Nbrot{3},
	"mag1":  mbrot.Magnent1{},
}

func main() {
	p := flag.Bool("p", false, "Whether or not to print to the command-line")
	o := flag.String("o", "", "The name of a file to write a PNG image to")
	x := flag.Float64("x", 0, "X coord")
	y := flag.Float64("y", 0, "Y coord")
	r := flag.Float64("r", 2, "Radius")
	w := flag.Uint("w", 35, "Horizontal Resolution")
	h := flag.Uint("h", 30, "Vertical Resolution")
	n := flag.Uint("n", 20, "Maximum iterations")
	P := flag.String("P", "gray", "Which color palette to use for PNG output")
	t := flag.String("t", "mbrot", "mbrot, bship, or nbrot")
	c := flag.Int("c", 0, "number of real threads to use, 0 for autodetect")
	flag.Parse()
	nCPU := *c
	if *c < 1 {
    nCPU = runtime.NumCPU()
  }
  runtime.GOMAXPROCS(nCPU)
	m := fractalList[*t]
	if m == nil {
		m = mbrot.Mandelbrot{}
	}
	f := mbrot.Render(mbrot.NewImage(*x, *y, *r, *w, *h), *n, m)
	if c := paletteList[*P]; c != nil {
		f.SetPalette(c(int(*n)))
	} else {
		f.SetPalette(palettes.Gray(int(*n)))
	}
	if *p {
		draw(f, 20)
	}
	if *o != "" {
		out, err := os.Create(*o)
		if err != nil {
			fmt.Println(err)
			return
		}
		if err := png.Encode(out, f); err != nil {
		  fmt.Println(err)
		  return
		}
	}
}

func draw(f mbrot.Fractal, maxN int) {
	sep := "@"
	for i := 0; i < int(f.X())*2; i = i + 1 {
		sep = sep + "-"
	}
	fmt.Println(sep + "@")
	for y := 0; uint(y) < f.Y(); y = y + 1 {
		out := ""
		for x := 0; uint(x) < f.X(); x = x + 1 {
			if f.Data[x][y] >= uint(maxN) {
				out = out + "# "
				//out = out + fmt.Sprint(vv)
			} else {
				out = out + "  "
				//out = out + fmt.Sprint(vv)
			}
		}
		fmt.Println("|" + out + "|")
	}
	fmt.Println(sep + "@")
}
