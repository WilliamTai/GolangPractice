package main

import "fmt"

//In Golang, we don't say class but using struct
type person struct {
	firstName string
	lastName  string
	age       int
}

func main() {
	p1 := person{
		firstName: "William",
		lastName:  "Tai",
		age:       29,
	}

	p2 := person{
		firstName: "John",
		lastName:  "Denny",
	}

	fmt.Println(p1)
	fmt.Println(p2)

	//output:
	//{William Tai 29}
	//{John Denny 0}

	fmt.Println("First name:", p1.firstName, "\nLast name:", p1.lastName, "\nAge:", p1.age)
	//output:
	//First name: William
	//Last name: Tai
	//Age: 29

}
