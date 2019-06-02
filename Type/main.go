package main

import "fmt"

//DECLARED
var x = 42
var y = "This is a string type variable"
var z string

//Custom type
type hotdog int
var b hotdog

func main() {
	fmt.Println(x)
	fmt.Println(y)
	//can not assign z as int because z is a string type.
	//z = 43

	//print the variable format
	fmt.Printf("%T\n", z)

	//print custom type
	b = 99
	fmt.Println(b)

	//cannot use b (type hotdog) as type int in assignment
	//x = b

	//Type conversion, as cast in C#
	x = int(b)
	fmt.Println("conversion b as int, x = ", x)
}
