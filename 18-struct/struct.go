package main

import "fmt"

type Customer struct {
	Name, Address, Phone string
	Age                  int
}

func (customer Customer)sayHi(name string){
	fmt.Println("Hello",name,"My name is",customer.Name)
}

func main() {
	var john Customer
	john.Address = "Medan"
	john.Phone = "085338383"
	john.Age = 21
	fmt.Println(john)
	
	//struct literal
	joko:= Customer{
		Name: "Joko",
		Address: "Cirebon",
		Age: 17,
	}
	fmt.Println(joko)

	//struct method
	joko.sayHi("John Sihombing")
}