package main

import "fmt"

func endApp() {
	message:=recover()
	if message != nil{
		fmt.Println("Error occured :",message)
	}
	fmt.Println("Application End")	
}

func runApplication(error bool){
	defer endApp()
	if(error){
		panic("Application error")
	}
	fmt.Println("Start Application")
}

func main() {
	runApplication(true)
	fmt.Println("John")
}