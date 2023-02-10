package main

import "fmt"

func sayHello(firstName string, lastName string) {
	fmt.Println("Hello",firstName, lastName)
}
func main() {
	sayHello("John","Sihombing")
}