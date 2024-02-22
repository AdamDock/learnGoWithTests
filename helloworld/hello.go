package main

import "fmt"
const (
	spanish = "Spanish"
	french = "French"
	italian = "Italian"

	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix = "Bonjour, "
	italianHelloPrefix = "Ciao, "
)


func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	
	return greetingPrefix(language) + name
}
func greetingPrefix(language string) (prefix string){
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	case italian:
		prefix = italianHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("world", "Italian"))
	fmt.Println(Hello("world", "French"))
	fmt.Println(Hello("world", "Spanish"))
	fmt.Println(Hello("world", ""))
	fmt.Println(Hello("world", "English"))
}