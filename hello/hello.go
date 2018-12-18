package main

import "fmt"

const englishPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

// domain code
func Hello(name, language string) string {
	if name == "" {
		name = "Beautiful"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case "spanish":
		prefix = spanishHelloPrefix
	case "french":
		prefix = frenchHelloPrefix
	default:
		prefix = englishPrefix
	}

	return
}

// side-effect
func main() {
	fmt.Println(Hello("", ""))
}
