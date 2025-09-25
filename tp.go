package main

import "fmt"

type Pokemon struct {
	Name  string
	Type  string
	Level int
	Price int
}

func main() {
	pikachu := Pokemon{Name: "Pikachu", Type: "Electric", Level: 25, Price: 400}
	bulbasaur := Pokemon{Name: "Bulbasaur", Type: "Grass/Poison", Level: 15, Price: 300}
	charmander := Pokemon{Name: "Charmander", Type: "Fire", Level: 18, Price: 500}
	lilou := Pokemon{Name: "Lilou", Type: "NRV", Level: 99, Price: 1000}
	Gateau := Pokemon{Name: "Gateau", Type: "Sucre", Level: 1, Price: 1}

	fmt.Println(pikachu)
	fmt.Println(bulbasaur)
	fmt.Println(charmander)
	fmt.Println(lilou)
	fmt.Println(Gateau)
}
