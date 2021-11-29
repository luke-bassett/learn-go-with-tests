package main

import "fmt"

const spanish = "Spanish"
const defaultName = "World"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "

func Hello(name string, language string) string {
	if name == "" {
		name = defaultName
	}
	if language == spanish {
		return spanishHelloPrefix + name
	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("world", ""))
}
