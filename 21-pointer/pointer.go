package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func changeCountrytoIndonesia(address *Address){
	address.Country="Indonesia"
}

func main() {
	address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
	address2 := &address1
	address3 := &address1

	address2.City="Bandung"

	*address2= Address{"Jakarta","DKI Jakarta","Indonesia"}
	
	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)

	var alamat = Address{
		City:"Cirebon",
		Province: "Jawa Barat",
		Country:"",
	}
	changeCountrytoIndonesia(&alamat)
	fmt.Println(alamat)
}