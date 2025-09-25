package main 

import ("fmt")

type Name struct{
	Name string
	Age int
	Class string
}


func main () {
	enzo := Initname("enzo","B1C",17)
	lilou := Initname("lilou","B1C",18)
	fmt.Println(enzo)
	fmt.Println(lilou)
}

func Initname (name, class string, age int)Name{
	return Name {
		Name: name,
		Age: age,
		Class: class,
	}
}
