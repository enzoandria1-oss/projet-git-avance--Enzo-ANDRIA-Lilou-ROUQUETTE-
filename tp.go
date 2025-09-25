package main

import "fmt"

type Pokemon struct {
	Name  string
	Type  string
	Level int
}

func main() {
	pikachu := Pokemon{Name: "Pikachu", Type: "Electric", Level: 25}
	bulbasaur := Pokemon{Name: "Bulbasaur", Type: "Grass/Poison", Level: 15}
	charmander := Pokemon{Name: "Charmander", Type: "Fire", Level: 18}

	fmt.Println(pikachu)
	fmt.Println(bulbasaur)
	fmt.Println(charmander)
}
