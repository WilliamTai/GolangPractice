package main

import "fmt"

func main() {
	//Arrays
	fmt.Println("Array here:")
	var x [5]int
	x[2] = 9
	fmt.Println(x)
	//output:
	//Array here:
	//[0 0 9 0 0]

	//Arrays is good but in Golang we have Slice which pretty cool for data structure
	//use Slice instead Arrays for effect programing

	fmt.Println("Slice here:")
	a := []int{1, 22, 333, 4444}
	fmt.Println(a)
	fmt.Println(a[1])
	fmt.Println(a[2:])
	fmt.Println(a[3:4])

	//output:
	//Slice here:
	//[1 22 333 4444]
	//22
	//[333 4444]
	//[4444]

	fmt.Println("Output range here:")
	for i, v := range a{
		fmt.Println(i,v)
	}
	//output:
	//Output range here:
	//0 1
	//1 22
	//2 333
	//3 4444

	fmt.Println("for loop here:")
	//This one will do the same thing above but if you miss calculating the index will causing error
	for  g := 0; g <= 3; g++ {
		fmt.Println(g, a[g])
	}
	//output:
	//for loop here:
	//0 1
	//1 22
	//2 333
	//3 4444

	//Append slice here.
	fmt.Println("Append slice with new element here:")
	a = append(a,123, 456)
	fmt.Println(a)

	fmt.Println("Append with two different slice object:")
	b := []int{777,888,999}
	a = append(a, b...) //the ... means all elements
	fmt.Println(a)

	//Remove slice element
	fmt.Println("Remove slice element")
	a = append(a[:2], a[5:]...)
	fmt.Println(a)

}
