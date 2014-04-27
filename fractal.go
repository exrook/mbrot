package mbrot

import (
  "image"
  "image/color"
)
type Fractal struct {
 Data [][]uint
 x,y uint
 palette color.Palette
}

func NewFractal(x,y uint) (f Fractal) {
  f = Fractal{x: x,y: y}
  f.Data = make([][]uint, x)
  data := make([]uint, x*y)
  for i := range f.Data {
    f.Data[i], data = data[:y], data[y:]
  }
  return
}

func (f *Fractal) SetPalette(c color.Palette) {
  f.palette = c
}  

func (f Fractal) X() uint {return f.x}
func (f Fractal) Y() uint {return f.y}

func (f Fractal) ColorModel() color.Model {
  return f.palette
}

func (f Fractal) Bounds() image.Rectangle {
  return image.Rectangle{image.Point{0,0},image.Point{int(f.x),int(f.y)}}
}

func (f Fractal) At(x,y int) color.Color {
  return f.palette[f.Data[uint(x)][uint(y)]]
}

func (f Fractal) ColorIndexAt(x,y int) uint8 {
  return uint8(f.Data[uint(x)][uint(y)])
}
