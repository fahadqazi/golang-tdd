package main

import "fmt"

const englishHelloPrefix = "Hello,"
const spanishHelloPrefix = "Hola,"

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	if language == "Spanish" {
		return fmt.Sprintf("%s %s", spanishHelloPrefix, name)
	}

	return fmt.Sprintf("%s %s", englishHelloPrefix, name)
}

func main() {
	fmt.Println(Hello("Chris", "English"))
}
