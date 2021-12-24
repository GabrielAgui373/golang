package main

import "fmt"

//A composição de interface pode ser feita de duas formas. Incorporando interfaces ou incorporando as assinaturas dos métodos de outras interfaces
type AuthorDetails interface {
	details()
}

type AuthorArticles interface {
	articles()
	picked()
}

type FinalDetails interface {
	details()
	AuthorArticles
	cdeatails()
}

type author struct {
	name      string
	branch    string
	college   string
	year      int
	salary    int
	particles int
	tarticles int
	cid       int
	post      string
	pick      int
}

func (a author) details() {
	fmt.Printf("Author Name: %s", a.name)
	fmt.Printf("\nBranch: %s and passing year: %d", a.branch, a.year)
	fmt.Printf("\nCollege Name: %s", a.college)
	fmt.Printf("\nSalary: %d", a.salary)
	fmt.Printf("\nPublished articles: %d", a.particles)
}

func (a author) articles() {
	pendingArticles := a.tarticles - a.particles
	fmt.Printf("\nPending articles: %d", pendingArticles)
}

func (a author) picked() {
	fmt.Printf("\nTotal number of picked articles: %d", a.pick)
}

func (a author) cdeatails() {
	fmt.Printf("\nAuthor Id: %d", a.cid)
	fmt.Printf("\nPost: %s", a.post)
}

func main() {
	values := author{
		name:      "Mickey",
		branch:    "Computer science",
		college:   "XYZ",
		year:      2012,
		salary:    50000,
		particles: 209,
		tarticles: 309,
		cid:       3087,
		post:      "Technical content writer",
		pick:      58,
	}

	var a FinalDetails = values
	a.articles()
	a.cdeatails()
	a.details()
	a.picked()
}
