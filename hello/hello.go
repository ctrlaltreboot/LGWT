package main

import "fmt"

/*
	 separation of "domain" code from "outside world"
   or "side-effects"
*/

func Hello() string {
	return "Hello, World... of... tests..."
}

func main() {
	fmt.Println(Hello())
}
