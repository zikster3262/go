package main

import "fmt"

const vietnameseprefix = "Xian chao, "
const czechprefix = "Ahoj, "
const englishprefix = "Hello, "

func lanprefix(language string) string {
	switch language {
	case "English":
		language = englishprefix
	case "Czech":
		language = czechprefix
	case "Vietnamese":
		language = vietnameseprefix
	default:
		language = englishprefix
	}
	return language
}

func Hello(name, language string) string {
	if name == "" {
		name = "World"
		language = "English"
	}

	fmt.Println(lanprefix(language) + name + "!")
	return lanprefix(language) + name + "!"

}

func main() {
	Hello("David", "Czech")
}
