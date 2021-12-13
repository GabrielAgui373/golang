package main

import "fmt"

//Como a linguagem Go não suporta classes, não é possível objter o polimorfismo através dessa estrutura, mas o polimorfismo pode ser obtido através das interfaces

//interface employee
type employee interface {
	develop() int
	name() string
}

//struct team1
type team1 struct {
	totalapp_1 int
	name_1     string
}

//struct team1 implementa a interface employee
func (t1 team1) develop() int {
	return t1.totalapp_1
}

func (t1 team1) name() string {
	return t1.name_1
}

//struct team2
type team2 struct {
	totalapp_2 int
	name_2     string
}

//struct team2 implementa a interface employee
func (t2 team2) develop() int {
	return t2.totalapp_2
}

func (t2 team2) name() string {
	return t2.name_2
}

func finalDevelop(i []employee) {
	totalProject := 0

	for _, el := range i {
		fmt.Printf("\nProject environment = %s\n", el.name())
		fmt.Printf("Total number of project %d\n", el.develop())
		totalProject += el.develop()
	}

	fmt.Printf("\nTotal projects completed by the company = %d", totalProject)
}

func main() {
	res1 := team1{totalapp_1: 20, name_1: "Android"}
	res2 := team2{totalapp_2: 35, name_2: "IOS"}

	final := []employee{res1, res2}
	finalDevelop(final)
}
