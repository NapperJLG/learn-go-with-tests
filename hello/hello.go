package main

import "fmt"

//Hello stupid linter
func Hello(name string) string {
	return "Hello " + name
}

func main() {
	fmt.Println(Hello("Jake"))
}
