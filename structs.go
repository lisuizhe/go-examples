package main

import (
	"fmt"
)

// Person is
type Person struct {
	name string
	age  int
}

// NewPerson do
func NewPerson(name string) *Person {
	p := Person{name: name}
	p.age = 42
	return &p
}

func main() {
	fmt.Println(Person{"Bob", 20})

	fmt.Println(Person{name: "Alice", age: 30})

	fmt.Println(Person{name: "Fred"})

	fmt.Println(&Person{name: "Ann", age: 40})

	fmt.Println(NewPerson("Jon"))

	s := Person{name: "Sean", age: 50}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 51
	fmt.Println(sp.age)
}
