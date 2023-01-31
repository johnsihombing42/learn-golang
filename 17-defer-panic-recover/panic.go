package main

import "fmt"

func endApp() {
	fmt.Println("End App")
}

func runApplication(error bool){
	defer endApp()
	if(error){
		panic("Application error")
	}
	fmt.Println("Application")
}

func main() {
	runApplication(true)
}