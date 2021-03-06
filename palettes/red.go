package palettes

import "image/color"

func Red(size int) color.Palette {
  c := make(color.Palette, size)
  for i := range c {
    c[i] = color.RGBA64{uint16((float64(i)/float64(size)) * 65535),0,0,65535}
  }
  return c
}
