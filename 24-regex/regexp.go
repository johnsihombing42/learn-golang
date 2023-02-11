package main

import (
	"fmt"
	"regexp"
)

func main() {
    var text = "banana burger soup"
    var regex, err = regexp.Compile(`[a-z]+`)

    if err != nil {
        fmt.Println(err.Error())
    }

    var res1 = regex.FindAllString(text, 2)
    fmt.Printf("%#v \n", res1)


    var res2 = regex.FindAllString(text, -1)
    fmt.Printf("%#v \n", res2)

	//tes
}