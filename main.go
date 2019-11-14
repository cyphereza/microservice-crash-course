package main

import "fmt"

type Hello struct {
	Name string
	age  int
}

func main() {
	newstruct := new(Hello)
	newstruct.Name = "asdasda"
	newstruct.age = 10

	fmt.Printf("%+v", newstruct)

	var helloWorld1 string
	helloWorld1 = "Hello world 1"

	helloWorld2 := "Hello world 2"

	var helloWorld3 = "Hello world 3"

	var helloWorld4 string = "Hello world 4"

	var helloWorld5, helloWorld6 string
	helloWorld5 = "Hello world 5"
	helloWorld6 = "Hello world 6"

	var helloWorld7 = string("Hello world 7")

	fmt.Println(newstruct.Name)
	fmt.Println(helloWorld2)
	fmt.Println(helloWorld3)
	fmt.Println(helloWorld4)
	fmt.Println(helloWorld5)
	fmt.Println(helloWorld6)
	fmt.Println(helloWorld7)
}
