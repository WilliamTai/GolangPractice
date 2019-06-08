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
	for i, v := range a {
		fmt.Println(i, v)
	}
	//output:
	//Output range here:
	//0 1
	//1 22
	//2 333
	//3 4444

	fmt.Println("for loop here:")
	//This one will do the same thing above but if you miss calculating the index will causing error
	for g := 0; g <= 3; g++ {
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
	a = append(a, 123, 456)
	fmt.Println(a)

	fmt.Println("Append with two different slice object:")
	b := []int{777, 888, 999}
	a = append(a, b...) //the ... means all elements
	fmt.Println(a)

	//Remove slice element
	fmt.Println("Remove slice element")
	a = append(a[:2], a[5:]...)
	fmt.Println(a)
	//output:
	//Append slice with new element here:
	//[1 22 333 4444 123 456]
	//Append with two different slice object:
	//[1 22 333 4444 123 456 777 888 999]
	//Remove slice element
	//[1 22 456 777 888 999]

	//Multi-dimensional slice
	fmt.Println("Multi-dimensional slice:")
	a1 := []string{"aaa","bbb","ccc"}
	a2 := []string{"ddd","eee","fff"}
	a1a2 := [][]string{a1,a2}
	fmt.Println(a1a2)
	//Multi-dimensional slice:
	//[[aaa bbb ccc] [ddd eee fff]]

	//The slice will dynamic changed if slice go bigger,
	//this will causing compiler and runtime issue if do all the time.
	//To prevent the issue using make to hold the size if you already know how much the data you have.
	fmt.Println("Make slice example:")
	c := make([]int, 3, 5)
	fmt.Println(c)
	fmt.Println("len:", len(c))
	fmt.Println("cap:", cap(c))
	//Make slice example:
	//[0 0 0]
	//len: 3
	//cap: 5

	fmt.Println("Assign c[1]:")
	c[1] = 9
	fmt.Println(c)
	fmt.Println("len:", len(c))
	fmt.Println("cap:", cap(c))
	//Assign c[1]:
	//[0 9 0]
	//len: 3
	//cap: 5

	fmt.Println("Append:")
	c = append(c, 7)
	fmt.Println(c)
	fmt.Println("len:", len(c))
	fmt.Println("cap:", cap(c))
	//Append:
	//[0 9 0 7]
	//len: 4
	//cap: 5

	fmt.Println("Append one more:")
	c = append(c, 444)
	fmt.Println(c)
	fmt.Println("len:", len(c))
	fmt.Println("cap:", cap(c))
	//Append one more:
	//[0 9 0 7 444]
	//len: 5
	//cap: 5

	fmt.Println("Append one more to overhead:")
	c = append(c, 555)
	fmt.Println(c)
	fmt.Println("len:", len(c))
	fmt.Println("cap:", cap(c))
	//Append one more to overhead:
	//[0 9 0 7 444 555]
	//len: 6
	//cap: 10
}
