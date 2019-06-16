package main

import "fmt"

//In Golang, we don't say class but using struct
type person struct {
	firstName string
	lastName  string
	age       int
}

//This shows how to embedded struct
type engineerProfile struct {
	person
	AwsSkillCertificated bool
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

	//Print embedded struct
	worker1 := engineerProfile{
		p1,
		true,
	}
	fmt.Println(worker1)
	//output:
	//{{William Tai 29} true}

	//Notice that the inside of struct has been promote to upper level by default
	//and you shouldn't make a same name item in upper level, something like firstName as example.
	fmt.Println(
		"First name:", worker1.firstName, //Same as worker1.person.firstName,
		"\nLast name:", worker1.lastName,
		"\nAge:", worker1.age,
		"\nAWS Skill Certificated:", worker1.AwsSkillCertificated)
	//output
	//First name: William
	//Last name: Tai
	//Age: 29
	//AWS Skill Certificated: true

}
