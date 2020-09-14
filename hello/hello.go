package main

import "fmt"

const englishHelloPrefix = "Hello "
const spanishHelloPrefix = "Hola "
const frenchHelloPrefix = "Bonjour "

//Hello stupid linter
func Hello(name string, language string) string {
	if name == "" {
		name = "Stranger"
	}

	return greetingPrefix(language) + name
}

//greetingPrefix sets the greeting language
func greetingPrefix(language string) (prefix string) {
	switch language {
	case "Spanish":
		prefix = spanishHelloPrefix
	case "French":
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("Jake", ""))
}
