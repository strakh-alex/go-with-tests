package main

import (
	"fmt"
)

const HelloEnglishPrefix = "Hello, "
const HelloSpanishPrefix = "Hola, "
const HelloFrenchPrefix = "Bonjour, "

func Hello(name, lang string) string {
	if name == "" {
		name = "world"
	}

	return greetingPrefix(lang) + name
}

func greetingPrefix(lang string) (prefix string) {
	switch lang {
	case "Spanish":
		prefix = HelloSpanishPrefix
	case "French":
		prefix = HelloFrenchPrefix
	default:
		prefix = HelloEnglishPrefix
	}

	return
}

func main() {
	fmt.Println(Hello("world", ""))
}
