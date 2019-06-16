package main

import (
	"fmt"
)

func main() {
	person := map[string]int{
		"John":  30,
		"Jenny": 24,
	}

	fmt.Println(person)
	//output:
	//map[John:30 Jenny:24]

	fmt.Println(person["John"])
	//output:
	//30

	//This won't causing error but print 0
	fmt.Println(person["Kitty"])
	//output:
	//0

	//This will get the value and it is existing or not
	v, ok := person["Kitty"]
	fmt.Println(v)
	fmt.Println(ok)
	//output:
	//0
	//false

	//If person[KEY] is exist
	if v, ok := person["Jenny"]; ok{
		fmt.Println("Jenny's age:", v)
	}
	//output:
	//Jenny's age: 24

	//Add element and print all with for loop
	person["William"] = 29
	for k, v := range person {
		fmt.Println(k, v)
	}
	//output:
	//William 29
	//John 30
	//Jenny 24

	//Delete element
	delete(person,"Jenny")
	fmt.Println(person)
	//output:
	//map[John:30 William:29]

	//If delete with item which didn't exist, it won't be an error.
	//so the code below shows how to verified then delete
	if v,ok := person["William"]; ok {
		fmt.Println("Value:", v)
		delete(person, "William")
		fmt.Println(person)
	}
	//output:
	//Value: 29
	//map[John:30]
}
