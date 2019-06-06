package main

import "fmt"

func main() {
	//for INIT; CONDITION; POST{
	//	fmt.Println("Hello loop")
	//}

	//Ex1
	fmt.Println("Ex1:")
	fmt.Print("for INIT; CONDITION; POST{}:")
	for i := 0; i < 10; i++ {
		fmt.Print(i, ",")
	}
	//output:
	//Ex1:
	//simple loop:0,1,2,3,4,5,6,7,8,9,

	//there are no Do loop in Golang but you can do with for condition or for {}
	//this is useful when build a TCP listener or something else need to run all the time
	fmt.Println("\nEx2:")
	fmt.Println("for condition {}")
	fmt.Print("z % 3 == 0 output:")
	z := 0
	for z < 10 {
		if z%3 == 0 {
			fmt.Print(z, " ")
		}
		z++
	}
	//output:
	//Ex2:
	//for condition {}
	//z % 3 == 0 output:0 3 6 9

	//for {} example

	//Ex3
	fmt.Println("\nEx3:")
	fmt.Println("for {}")
	fmt.Print("x % 2 == 0 output:")
	x := 0
	for {
		x++
		//put break condition here to prevent infinity loop
		if x > 20 {
			fmt.Println("\nx = ", x, ",break!")
			break
		} else if x%2 == 0 {
			fmt.Print(x, ",")
		}

	}
	//output
	//Ex3:
	//for {}
	//x % 2 == 0 output:2,4,6,8,10,12,14,16,18,20,
	//x =  21 ,break!
}
