package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const chinese = "Chinese"
const english = "English"
const defaultName = "World"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const chineseHelloPrefix = "Ni hao, "

func Hello(name string, language string) string {
	if name == "" {
		name = defaultName
	}

	return grettingPrefix(language) + name
}

func grettingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	case chinese:
		prefix = chineseHelloPrefix
	case english:
		prefix = englishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("world", ""))
}
