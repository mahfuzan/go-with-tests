package main

import "fmt"

const indo string = "Indonesia"
const spanish string = "Spanish"
const french string = "French"
const indoHelloPrefix string = "Halo, "
const spanishHelloPrefix string = "Hola, "
const frenchHelloPrefix string = "Bonjour, "

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	return greeting(language) + name
}

func greeting(language string) string {
	var prefix string

	switch language {
	case indo:
		prefix = indoHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = indoHelloPrefix
	}

	return prefix
}

func main() {
	fmt.Println(Hello("world", ""))
}
