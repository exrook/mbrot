package mbrot

import (
  "math"
  "math/cmplx"
)

type BurningShip struct {}

func (b BurningShip) Iterate(z,c complex128) complex128 {
  zz := complex(math.Abs(real(z)),math.Abs(imag(z)))
  return (zz*zz)+c
}

func (b BurningShip) CheckBounds(z complex128) bool {
  return cmplx.Abs(z) < 2
}
