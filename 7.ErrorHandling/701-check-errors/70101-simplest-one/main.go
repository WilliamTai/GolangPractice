package main

import "fmt"

func main() {
	///This is the simplest error handling.
	///Always, always, always check errors
	n, err := fmt.Println("Hello")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(n)
}
