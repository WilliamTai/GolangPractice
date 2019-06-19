package main

import "fmt"

// func (r receiver) identifier(parameter(s)) (return(s)) { code }

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
