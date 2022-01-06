package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Client struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Cpf   int    `json:"cpf"`
}

func (c Client) display() {
	fmt.Println(c.Cpf)
	fmt.Println(c.Email)
	fmt.Println(c.Name)
}

type ClientInternationa struct {
	Client
	Country string `json:"country"`
}

func main() {
	cliente := Client{
		Name:  "João",
		Email: "joao@gmail.com",
		Cpf:   12345678900,
	}

	cliente2 := ClientInternationa{
		Client: Client{
			Name:  "Maria",
			Email: "maria@outl ook.com",
			Cpf:   12345678900,
		},
		Country: "EUA",
	}

	cliente.display()
	cliente2.display()

	clientJson, err := json.Marshal(cliente)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(string(clientJson))

	jsonClient := `{"name":"João","email":"joao@gmail.com","cpf":12345678900}`
	client3 := Client{}

	json.Unmarshal([]byte(jsonClient), &client3)
	fmt.Println(client3)

}
