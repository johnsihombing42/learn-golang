package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	fmt.Println(rand.Int())
	fmt.Println(rand.Int())
	fmt.Println(rand.Int())
}