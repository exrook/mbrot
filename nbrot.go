package mbrot

import "math/cmplx"

type Nbrot struct {
	N complex128
}

func (n Nbrot) Iterate(z, c complex128) complex128 {
	return cmplx.Pow(z, n.N) + c
}

func (n Nbrot) CheckBounds(z complex128) bool {
	return cmplx.Abs(z) < 2
}
