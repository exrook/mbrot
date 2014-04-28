package mbrot

import "math/cmplx"

type Mandelbrot struct{}

func (m Mandelbrot) Iterate(z, c complex128) complex128 {
	return (z * z) + c
}

func (m Mandelbrot) CheckBounds(z complex128) bool {
	return cmplx.Abs(z) < 2
}
