package main

import "fmt"

func main() {
	/*
	A example shows how the recover works.

	Source: https://blog.golang.org/defer-panic-and-recover
	Panic is a built-in function that stops the ordinary flow of control and begins panicking. When the function F calls panic,
	execution of F stops, any deferred functions in F are executed normally, and then F returns to its caller. To the caller,
	F then behaves like a call to panic. The process continues up the stack until all functions in the current goroutine have returned,
	at which point the program crashes. Panics can be initiated by invoking panic directly. They can also be caused by runtime errors,
	such as out-of-bounds array accesses.

	Recover is a built-in function that regains control of a panicking goroutine. Recover is only useful inside deferred functions.
	During normal execution, a call to recover will return nil and have no other effect. If the current goroutine is panicking,
	a call to recover will capture the value given to panic and resume normal execution.

	As I understand, this can be used to rollback the hole process if you fail in somewhere.
	*/
	f()
	fmt.Println("Returned normally from f.")
}

func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	fmt.Println("Calling g.")
	g(0)
	fmt.Println("Returned normally from g.")
}

func g(i int) {
	if i > 3 {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("Defer in g", i)
	fmt.Println("Printing in g", i)
	g(i + 1)
}

/*output
Calling g.
Printing in g 0
Printing in g 1
Printing in g 2
Printing in g 3
Panicking!
Defer in g 3
Defer in g 2
Defer in g 1
Defer in g 0
Recovered in f 4
Returned normally from f.

Process finished with exit code 0

 */