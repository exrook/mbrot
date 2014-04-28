package mbrot

type Image struct {
  P1,P2 complex128
  W,H uint
}

func NewImage(x,y,r float64, w,h uint) Image {
  i := Image{W:w,H:h}
  if w>h { // wide
    s := float64(w)/float64(h)
    i.P1 = complex(x-(r*s),y+r)
    i.P2 = complex(x+(r*s),y-r)
  } else { // tall
    s := float64(h)/float64(w)
    i.P1 = complex(x-r,y+(r*s))
    i.P2 = complex(x+r,y-(r*s))
  }
  return i
}

func (i Image) HScale() float64 {
  return (real(i.P1)-real(i.P2))/float64(i.W)
}

func (i Image) VScale() float64 {
  return (imag(i.P1)-imag(i.P2))/float64(i.H)
}

func (i Image) Scale() (h,v float64) {
  return i.HScale(), i.VScale()
}
