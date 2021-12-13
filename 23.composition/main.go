package main

import "fmt"

//Composição de structs
type details struct {
	genre       string
	genreRating string
	reviews     string
}

type game1 struct {
	name  string
	price string
	details
}

func (d details) showDetails() {
	fmt.Println("Genre:", d.genre)
	fmt.Println("Genre Rating:", d.genreRating)
	fmt.Println("Reviews:", d.reviews)
}

func (g game1) show() {
	fmt.Println("Name:", g.name)
	fmt.Println("Price:", g.price)
	g.showDetails()
}

//Composição de interfaces
type purchase interface {
	sell()
}

type display interface {
	show()
}

//Se todos os métodos das interfaces que compõe a interface salesman for implementado, então a interface selesman é implementada
type salesman interface {
	purchase
	display
}

type game2 struct {
	name, price    string
	gameCollection []string
}

func (t game2) sell() {
	fmt.Println("--------------------------------------")
	fmt.Println("Name:", t.name)
	fmt.Println("Price:", t.price)
	fmt.Println("--------------------------------------")
}

func (t game2) show() {
	
}

func main() {
	action := details{
		genre:       "Action",
		genreRating: "18+",
		reviews:     "mostly positive",
	}

	newGame := game1{
		name:    "XYZ",
		price:   "$125",
		details: action,
	}

	newGame.show()

}
