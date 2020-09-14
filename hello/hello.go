package main

import "fmt"

const englishHelloPrefix = "Hello "

//Hello stupid linter
func Hello(name string) string {
	if name == "" {
		name = "Stranger"
	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("Jake"))
}
