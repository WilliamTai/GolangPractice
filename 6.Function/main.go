package main

import "fmt"

// func (r receiver) identifier(parameter(s)) (return(s)) { code }

type person struct {
	firstName string
	lastName  string
}

type engineer struct {
	person
	awsCertificated bool
}

//When you have a receiver, it is going to attach the function to the type
func (s engineer) coding() {
	fmt.Println("write some code.")
}

func (s engineer) speaking() {
	fmt.Println("I'm ", s.firstName, s.lastName)
}

func main() {
	printHelloWorld()
	//output:
	//	Hello World

	printTheText("Hello function")
	//output:
	//	From printTheText: Hello function

	longString := setFullFillString("Sending string here")
	println(longString)
	//output:
	//	From fullFillString: Sending string here

	name, flag := getTwoReturnValue("William", "Tai")
	fmt.Println("Name: ", name, "Flag: ", flag)
	//output:
	//Name:  WilliamTai Flag:  true

	printVariadicParameter(2, 3, 4, 5, 6)
	//output:
	//	[2 3 4 5 6]
	//[]int
	//Loop, now is:  0 Adding:  2 sum:  2
	//Loop, now is:  1 Adding:  3 sum:  5
	//Loop, now is:  2 Adding:  4 sum:  9
	//Loop, now is:  3 Adding:  5 sum:  14
	//Loop, now is:  4 Adding:  6 sum:  20

	//This shows how to attach methods to type struct
	eng1 := engineer{
		person: person{
			firstName: "William",
			lastName:  "Tai",
		},
		awsCertificated: true,
	}
	eng1.speaking()
	eng1.coding()
	//output:
	//I'm  William Tai
	//write some code.

	//Anonymous Function
	func (x int){
		fmt.Println("Print from Anonymous function", x)
	}(999)
	//output:
	//Print from Anonymous function 999

	//How to assign a function to a variable
	//I love this, it's easy to read!
	func1 := func(x int){
		fmt.Println("Assigned function print:", x)
	}

	func1(8888)
	//output:
	//Assigned function print: 8888
}

func printHelloWorld() {
	//The simplest function.
	fmt.Println("Hello World")
}

func printTheText(s string) {
	//Print the parameter
	fmt.Println("From printTheText: ", s)
}

func setFullFillString(s string) string {
	return fmt.Sprint("From fullFillString: ", s)
}

func getTwoReturnValue(firstName, lastName string) (string, bool) {
	return fmt.Sprint(firstName, lastName), true
}

func printVariadicParameter(x ...int) {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	sum := 0
	for i, v := range x {
		sum += v
		fmt.Println("Loop, now is: ", i, "Adding: ", v, "sum: ", sum)
	}
}
