package main

import (
	"fmt"
	"math"
)

//Definindo uma interface geometry
type geometry interface {
	Area() float64
	Perim() float64
}

//Definindo um tipo Rect
type Rect struct {
	Width, Height float64
}

//Definindo um tipo Circle
type Circle struct {
	Radius float64
}

//Uma interface é implementa impplicitamente através da implementação de seus métodos.
func (r Rect) Area() float64 {
	return r.Width * r.Height
}

func (r Rect) Perim() float64 {
	return 2*r.Width + 2*r.Height
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perim() float64 {
	return 2 * math.Pi * c.Radius
}

//Qualquer tipo que implemente a interface será aceito na função
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.Area())
	fmt.Println(g.Perim())
}

func main() {
	r := Rect{Width: 10, Height: 5}
	c := Circle{Radius: 4}

	measure(r)
	measure(c)
}
