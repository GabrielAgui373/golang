package main

import "fmt"

type Vertex struct {
	Lat, Log float64
}

func main() {
	//Map é uma estrutura de dados não ordenada de chave-valor
	//Map Literal: são como struct literais, mas com chaves obrigatórias (map[key-type]var=type)
	m := map[string]Vertex{
		"Bell Labs": Vertex{40.344, -23.434},
		"Google":    Vertex{30.344, -23.4344},
	}

	//É possível omitir o tipo do valor
	m = map[string]Vertex{
		"Bell Labs": {40.68433, -74.39967},
		"Google":    {37.42202, -122.08408},
	}
	fmt.Println(m)

	//O valor zero de um map é nil. Não é possível adicionar valores a um map com valor nil. Para criar um map vazio é interessante utilizar a função make para alocar um map
	m1 := make(map[string]Vertex)
	m1["Bell Labs"] = Vertex{40.3232, -223.2323}

	fmt.Println(m1)

	//iterando sobre um map com range. É retornado a chave e o valor de cada elemento
	for key, value := range m {
		fmt.Printf("key: %v value: %v\n", key, value)
	}

	//Mutação de maps
	m2 := make(map[string]int)

	//Inserir um elemento
	m2["Answer"] = 42
	fmt.Println("The value:", m2["Answer"])

	//Atualizar um elemento
	m2["Answer"] = 40
	fmt.Println("The value:", m2["Answer"])

	//Recuperar um elemento
	el := m2["Answer"]
	fmt.Println("The value:", el)

	//Exluir um elemento
	delete(m2, "Answer")

	//Verificar se o elemento está presente
	elem, ok := m2["Answer"]
	fmt.Println("The value:", elem, "Present?", ok)

}
