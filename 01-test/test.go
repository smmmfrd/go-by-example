package main

import (
	"fmt"
)

const (
	englishHelloPrefix = "Hello, "

	spanish            = "Spanish"
	spanishHelloPrefix = "Hola, "

	french            = "French"
	frenchHelloPrefix = "Bonjour, "
)

func Hello(name string, language string) string {
	// Public functions start with a capital!
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	// Lowercase functions are internal/private.

	// In the return type, we create a variable that will exist in our function!
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	// Due to our named return value, this actually means = return prefix
	return
}

func main() {
	fmt.Println(Hello("world", ""))
}
