package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const enHelloPrefix = "Hello, "
const frHelloPrefix = "Bonjour, "
const esHelloPrefix = "Hola, "

// Hello somthing
func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = esHelloPrefix
	case french:
		prefix = frHelloPrefix
	default:
		prefix = enHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("world", "English"))
}
