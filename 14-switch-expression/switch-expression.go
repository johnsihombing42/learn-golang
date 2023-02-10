package main

import "fmt"

func main() {
	name := "John"

	switch name{
	case "John" :
		fmt.Println("Ini John")
	case "Joko":
		fmt.Println("Ini Joko")
	default:
		fmt.Println("Ini Sabrina, kuy kenalan")
	}
	
}