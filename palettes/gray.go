package palettes

import "image/color"

func Gray(size int) color.Palette {
  c := make(color.Palette, size)
  for i := range c {
    c[i] = color.Gray16{uint16((float64(i)/float64(size)) * 65535)}
  }
  return c
}
