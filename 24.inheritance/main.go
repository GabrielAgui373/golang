package main

import "fmt"

//Golang não suporta classes, então a herança é feita através de composição
type Comic struct {
	Universe string
}

type Marvel struct {
	Comic
}

type DC struct {
	Comic
}

func (c Comic) ComicUniverse() string {
	return c.Universe
}

//Multipla herança
type First struct {
	base_one string
}

type Second struct {
	base_two string
}

func (f First) printBase1() string {
	return f.base_one
}

func (s Second) printBase2() string {
	return s.base_two
}

type Child struct {
	First
	Second
}

func main() {
	c1 := Marvel{
		Comic{
			Universe: "MCU",
		},
	}

	c2 := DC{
		Comic{
			Universe: "DC",
		},
	}

	fmt.Println("Universe is: ", c1.ComicUniverse())
	fmt.Println("Universe is: ", c2.ComicUniverse())

	c := Child{
		First{base_one: "In base struct 1."},
		Second{base_two: "\nIn base struct 2.\n"},
	}

	fmt.Println(c.printBase1())
	fmt.Println(c.printBase2())
}
