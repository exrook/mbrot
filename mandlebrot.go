package mbrot

import "math/cmplx"

type Mandlebrot struct {}

func (m Mandlebrot) Iterate(z,c complex128) complex128 {
  return (z*z)+c
}

func (m Mandlebrot) CheckBounds(z complex128) bool {
  return cmplx.Abs(z) < 2
}
