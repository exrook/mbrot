package mbrot

import "fmt"

// Renders the function defined by fn
func Render(x,y,r float64, w,h,nMax uint, fn Fractaler) Fractal {
  //This section should be rewritten at some point
  lft :=  x-r
  rght := x+r
  tp := y+r
  bttm := y-r
  var s,xStep,yStep float64
  if h>w { //tall
    s = float64(h)/float64(w)
    tp = y+(r*s)
    bttm = y-(r*s)
  } else { //wide
    s = float64(w)/float64(h)
    rght = x+(r*s)
    lft = x-(r*s)
  }
  xStep = (rght-lft)/float64(w)
  yStep = (bttm-tp)/float64(h)
  f := NewFractal(w,h)
  fmt.Println(lft,tp)
  fmt.Println(rght,bttm)
  fmt.Println(xStep,yStep)
  for xPos := uint(0); xPos < w; xPos = xPos+1 {
    for yPos := uint(0); yPos < h; yPos = yPos+1 {
      x := lft+(float64(xPos)*xStep)
      y := bttm-(float64(yPos)*yStep)
      c := complex(x,y)
      z := complex(0,0)
      n := uint(1)
      for ;n < nMax && fn.CheckBounds(z);n=n+1 {
        z = fn.Iterate(z,c)
      }
      f.Data[xPos][yPos] = n
    }
    fmt.Printf("Column %v/%v %v%%\n", xPos,w,float32(xPos)/float32(w)*100)
  }
  return f
}
