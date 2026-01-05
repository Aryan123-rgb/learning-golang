package main

import "fmt"

func Hello(name string) string {
	if name == "" {
		return "Hello World"
	}
	return "Hello " + name
}

func main() {
	fmt.Println(Hello("Chris"))
}