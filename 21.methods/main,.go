package main

import "fmt"

type User struct {
	Name string
	Age  uint8
}

//Método é uma função com o escopo definido para um tipo nomeado
func (user User) verifyAge() bool {
	return user.Age >= 18
}

//Um método pode ser definido para qualquer tipo nomeado, não necessariamente uma struct
type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

//Pointer Receiver
func (user *User) changeName(name string) {
	user.Name = name
}

func main() {
	user := User{Name: "Gabriel", Age: 20}
	fmt.Println(user.verifyAge())

	user.changeName("Gabr")
	fmt.Println(user)
}
