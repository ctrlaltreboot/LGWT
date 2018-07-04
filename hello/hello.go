package main

import "fmt"

/*
	 separation of "domain" code from "outside world"
   or "side-effects"
*/

const helloPrefix = "Hello there, "

func Hello(n string) string {
	if n == "" {
		return "Hello, World"
	} else {
		return helloPrefix + n
	}
}

func main() {
	fmt.Println(Hello("handsome"))
}
