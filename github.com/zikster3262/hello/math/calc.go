package main

import "fmt"

func Calc(x, y int) int {
	a := x + y
	fmt.Println(a)
	return a
}

func Divide(x, y int) float64 {
	a := float64(x)
	b := float64(y)
	c := a / b
	fmt.Println(c)
	return c
}

const greetings = "Hello "

func Hello(name string) string {
	fmt.Println(greetings + name)
	return greetings + name
}

func main() {
	Calc(5, 12)
	Divide(5, 12)
	Hello("David")
}
