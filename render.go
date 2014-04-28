package mbrot

import (
	"runtime"
)

type point struct {
	x, y, n uint
	c       complex128
}

// Renders the function defined by fn
func Render(i Image, nMax uint, fn Fractaler) Fractal {
	nCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(nCPU)

	xStep, yStep := i.Scale()
	f := NewFractal(i.W, i.H)

	// points are fed through in to be computed then to d to be written to memory
	prein := make(chan point, nCPU)
	in := make(chan point, nCPU)
	d := make(chan point, nCPU)
	pcount, icount, count := 0, 0, 0 // used to coordinate the closing of d

	// create the raw points
	go func() {
		for x := uint(0); x < i.W; x = x + 1 {
			go func(x uint) {
				for y := uint(0); y < i.H; y = y + 1 {
					prein <- point{x: x, y: y}
				}
				if pcount+1 == int(i.W) {
					close(prein)
				}
				pcount = pcount + 1
			}(x)
		}
	}()

	//Create the complex numbers
	for j := 0; j < nCPU; j = j + 1 {
		go func() {
			for p := range prein {
				p.c = i.P1 - complex(float64(p.x)*xStep, float64(p.y)*yStep)
				in <- p
			}
			if icount+1 == nCPU {
				close(in)
			}
			icount = icount + 1
		}()
	}

	// Start the worker functions who perform the calculations
	for i := 0; i < nCPU; i = i + 1 {
		go func() {
			for p := range in {
				z := complex128(0)
				for p.n = 1; p.n < nMax && fn.CheckBounds(z); p.n = p.n + 1 {
					z = fn.Iterate(z, p.c)
				}
				d <- p
			}
			if count+1 == nCPU {
				close(d)
			}
			count = count + 1
		}()
	}

	// Enter the data into the array as it comes out
	go func() {
		for p := range d {
			f.Data[p.x][p.y] = p.n
		}
	}()
	for p := range d {
		f.Data[p.x][p.y] = p.n
	}

	return f
}
