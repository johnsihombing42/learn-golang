package main

import "fmt"

func getFullName()(string, string,int) {
	return "John", "Sihombing",20
}
func main() {
	firstName, lastName,age := getFullName()
	fmt.Println(firstName,lastName,age)
}