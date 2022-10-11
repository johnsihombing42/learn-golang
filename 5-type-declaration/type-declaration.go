package main

import "fmt"

func main() {
	type KTP string
	type married bool
	var noKTPJohn KTP = "12345"
	var marriedStatus = false
	fmt.Println(noKTPJohn)
	fmt.Println(marriedStatus)
}