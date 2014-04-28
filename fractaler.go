package mbrot

// Defines the functions required by any iterative fractal function
type Fractaler interface {
	Iterate(z complex128, c complex128) complex128
	CheckBounds(z complex128) bool
}
