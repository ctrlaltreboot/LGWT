package main

import "fmt"

/*
	 separation of "domain" code from "outside world"
   or "side-effects"
*/

func Hello(n string) string {
	return "Hello there, " + n
}

func main() {
	fmt.Println(Hello("handsome"))
}
