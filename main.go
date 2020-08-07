package main

import "fmt"

func feature1() string {
	return "Hello from feature1!"
}

func feature2() string {
	return "Hi from feature2!"
}

func main() {
	fmt.Println("Hi from main!")
	fmt.Println(feature1())
	fmt.Println(feature2())
	fmt.Println("Bye from main!")
}
