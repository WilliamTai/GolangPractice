package main

import "fmt"

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
}
