package main

import "fmt"

func main() {
	person := map[string]string{
		"name":   "John",
		"adress": "Medan",
	}
	person["title"]="programer"
	fmt.Println(person)
	fmt.Println(person["name"])
}