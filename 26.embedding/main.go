package main

import "fmt"

//A linguagem Go não suporta herança, mas é possível incorpoorar interfaces, realizando a composição de interfaces ou incorporando as assinaturas dos métodos de outras interfaces

//interface AuthorDetails
type AuhtorDetails interface {
	details()
}

//interface AuthorAtticles
type AuthorArticles interface {
	articles()
}

//interface FinalDetails
type FinalDetails interface {
	details()
	articles()

	// AuhtorDetails
	// AuthorArticles
}

// Struct Author
type Author struct {
	name      string
	branch    string
	college   string
	year      int
	salary    int
	particles int
	tarticles int
}

//Implementando o método da interface AuthorDetails
func (a Author) details() {
	fmt.Printf("Author Name: %s", a.name)
	fmt.Printf("\nBranch: %s and passing year: %d", a.branch, a.year)
	fmt.Printf("\nCollege Name: %s", a.college)
	fmt.Printf("\nSalary: %d", a.salary)
	fmt.Printf("\nPublished articles: %d", a.particles)
}

//Implementando o método da interface AuthorArticles
func (a Author) articles() {
	pendingArticles := a.tarticles - a.particles
	fmt.Printf("\nPending articles: %d", pendingArticles)
}

func main() {
	values := Author{
		name:      "Mickey",
		branch:    "Computer science",
		college:   "XYZ",
		year:      2012,
		salary:    50000,
		particles: 209,
		tarticles: 309,
	}

	//Acessando os métodos das interfaces AuhtorDetails e AuthorArticles usando a interface FinalDetails
	var f FinalDetails = values
	f.details()
	f.articles()
}
