package main 

import ("fmt")

type Name struct{
	Name string
	Age int
}


func main () {
	enzo := Initname("enzo",17)
	lilou := Initname("lilou",18)
	fmt.Println(enzo)
	fmt.Println(lilou)
}

func Initname (name string, age int)Name{
	return Name {
		Name: name,
		Age: age,
	}
}
