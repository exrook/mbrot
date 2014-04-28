package mbrot

import (
  "runtime"
)

type point struct {
  x,y,n uint
  c complex128
}

// Renders the function defined by fn
func Render(i Image, nMax uint, fn Fractaler) Fractal {
  nCPU := runtime.NumCPU()
  runtime.GOMAXPROCS(nCPU)
  
  xStep, yStep := i.Scale()
  f := NewFractal(i.W,i.H)
  
  // points are fed through in to be computed then to d to be written to memory
  in := make(chan point,nCPU)
  d := make(chan point,nCPU)
  count := 0 // used to coordinate the closing of d
  
  // generate the points to be fed into d
  go func(i Image) {
    for x := uint(0);x < i.W;x=x+1 {
      for y := uint(0); y<i.H;y=y+1 {
        in <- point{
                x:x,
                y:y,
                c:i.P1-complex(float64(x)*xStep,float64(y)*yStep),
        }
      }
    }
    close(in)
  }(i)
  
  // Start the worker functions who perform the calculations
  for i:=0;i<nCPU;i=i+1 {
    go func() {
      for p := range in {
        z := complex128(0)
        for p.n=1;p.n<nMax && fn.CheckBounds(z);p.n=p.n+1 {
          z = fn.Iterate(z,p.c)
        }
        d<-p
      }
      if count+1 == nCPU {
        close(d)
      }
      count = count+1
    }()
  }
  
  // Enter the data into the array as it comes out
  for p := range d {
    f.Data[p.x][p.y] = p.n
  }    
    
  return f
}
