package main

import (
	"fmt"

	"rsc.io/quote"
)

func main(){
	fmt.Println("Hello, World! ")
	fmt.Println(quote.Go())
	greet("sam")
}

func greet(name string) {
	fmt.Println("hello", name)
}