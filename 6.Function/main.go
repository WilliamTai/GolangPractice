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

	printVariadicParameter(2, 3, 4, 5, 6)
	//output:
	//	[2 3 4 5 6]
	//[]int
	//Loop, now is:  0 Adding:  2 sum:  2
	//Loop, now is:  1 Adding:  3 sum:  5
	//Loop, now is:  2 Adding:  4 sum:  9
	//Loop, now is:  3 Adding:  5 sum:  14
	//Loop, now is:  4 Adding:  6 sum:  20
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

func printVariadicParameter(x ...int){
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	sum := 0
	for i, v := range x {
		sum += v
		fmt.Println("Loop, now is: ", i, "Adding: ", v, "sum: ", sum)
	}
}