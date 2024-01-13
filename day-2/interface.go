package main

import "fmt"

type Geomatery interface {
	Area() float64
	Perimeter() float64
}

type Rect struct {
	Width, Height float64
}

func (r Rect) Area() float64{
	return r.Width * r.Height
}

func (r Rect) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func main(){
	// Interface

	var w Geomatery = Rect{3,4}

	fmt.Println(w)
	fmt.Println(w.Area())
	fmt.Println(w.Perimeter())
	
}
