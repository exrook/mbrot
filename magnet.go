package mbrot

import (
	"math/cmplx"
)

type Magnent1 struct{}

func (m Magnent1) Iterate(z, c complex128) complex128 {
	t := (((z * z) + (c - 1)) / ((2 * z) + (c - 2)))
	return t * t
}

func (m Magnent1) CheckBounds(z complex128) bool {
	return cmplx.Abs(z) < 2
}

/*
type Magnent2 struct{}

func (m Magnent2) Iterate(z,c complex128) complex128 {
  t := (((z*z)+(c-1))/((2*z)+(c-2)))
  return t*t
}

func (m Magnent2) CheckBounds(z complex128) bool {
  return cmplx.Abs(z) < 2
}
*/
