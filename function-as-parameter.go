package main

import "fmt"

func sayHelloWithFilter(name string, filter func(string) string) {
	nameFiltered := filter(name)
	fmt.Println("Hello",nameFiltered)
}

func spamFilter(name string)string{
	if name == "Anjing"{
		return "..."
	}else{
		return name
	}
}

func main() {
	sayHelloWithFilter("John",spamFilter)
}