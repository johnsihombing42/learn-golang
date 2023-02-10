package main

import "fmt"

func main() {
	//perulanga for biasa
	// counter:=1
	// for counter<=10{
	// 	fmt.Println("Perulangan ke",counter)
	// 	counter++
	// }

	// for counter := 1; counter <= 10; counter++ {
	// 	fmt.Println("Perulangan ke", counter)
	// }

	for i := 1; i <= 100; i++ {
		if(i%3==0 && i%5==0){
			fmt.Println("Fizz Buzz")
		}else if (i%3 == 0) {
			fmt.Println("Fizz")
		}else if (i%5== 0){
			fmt.Println("Buzz")
		}else {
			fmt.Println(i)
		}
	}

}