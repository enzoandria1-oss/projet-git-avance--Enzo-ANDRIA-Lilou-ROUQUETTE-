package main 

import ("fmt")

type Name struct{
	Name string
}


func main () {
	enzo := Initname("enzo")
	lilou := Initname("lilou")
	fmt.Println(enzo)
	fmt.Println(lilou)
}

func Initname (name1 string)Name{
	return Name {
		Name: name1,
	}
}
