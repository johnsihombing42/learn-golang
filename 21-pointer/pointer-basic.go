package main

import "fmt"

func main() {
	// numberA := 5
	// numberB := &numberA
	// fmt.Println(numberA)
	// fmt.Println(numberB)
	// fmt.Println(*numberB)

	// *numberB=10
	// fmt.Println(numberA)
	// fmt.Println(*numberB)

	// var numberA int = 5
	// var numberB *int = &numberA

	// fmt.Println(numberA)
	// fmt.Println(numberB)
	// fmt.Println(*numberB)
	// fmt.Println()

	// numberA=20
	// fmt.Println(numberA)
	// fmt.Println(numberB)
	// fmt.Println(*numberB)

	number := 5
	fmt.Println("Alamat memori", &number)
	fmt.Println("Nilai awal :", number)

	change(&number, 100)

	fmt.Println("Nilai akhir :", number)
	fmt.Println("Alamat memori", &number)

}

func change(old *int, new int) {
	*old = new
	fmt.Println("Nilai function:",*old)
	fmt.Println("Alamat memori old", old)
}
