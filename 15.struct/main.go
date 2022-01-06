package main

import "fmt"

//Struct é uma coleção de campos de dados com tipos declarados
type User struct {
	Name string
	Age  int8
}

//Composiçao com struct
type Rectangle struct {
	Height int
	Width  int
	Color  string

	Geometry struct {
		Perimeter int
		Area      int
	}
}

func main() {
	user := User{Name: "Gabriel", Age: 18}
	user.Age = 20
	fmt.Printf("Type: %T Value: %v\n", user, user)

	//Podemos acessa os campos de um struct através de um ponteiro de struct
	rect := Rectangle{
		Height: 10,
		Width:  20,
		Color:  "Red",
	}

	p := &rect
	//Podemos simplificar a notação *p. para apenas p.
	p.Geometry.Area = p.Height * p.Width
	p.Geometry.Perimeter = 2*p.Height + 2*p.Width

	fmt.Printf("Area: %v Perimeter: %v", rect.Geometry.Area, rect.Geometry.Perimeter)
}
