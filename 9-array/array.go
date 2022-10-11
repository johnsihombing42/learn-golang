package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "John"
	names[1] = "Tri"
	names[2] = "Putra"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])
	fmt.Println(len(names))

	//array
	var values=[3]int{
		80,
		90,
		100,
	}
	fmt.Println(len(values))
	fmt.Println(values)
}
