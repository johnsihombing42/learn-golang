package main

import "fmt"

func main() {
	var months = [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	fmt.Println("Slice 1")
	var slice1 = months[4:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	fmt.Println("Slice 2")
	var slice2=months[10:]
	fmt.Println(slice2)

	fmt.Println("Append Slice 2")
	var slice3=append(slice2,"John")
	fmt.Println(slice3)


	var fruits = []string{"apple", "grape", "banana", "melon","orange","pinneaple"}
	fmt.Println(len(fruits)) 
}