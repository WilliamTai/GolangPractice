package main

import (
	"log"
	"os"
)

func main() {
	_, err := os.Open("no-file.txt")
	if err != nil {
		log.Panicln(err)
		//Call panic and log
		//This allows a program to manage behavior of a panicking goroutine.
		//executing a call to recover inside a deferred function (but not any function called by it)
		//read the go doc for more detail
	}

	/*
	2019/07/27 19:55:59 open no-file.txt: no such file or directory
	panic: open no-file.txt: no such file or directory


	goroutine 1 [running]:
	log.Panicln(0xc00006af78, 0x1, 0x1)
	        /usr/local/go/src/log/log.go:340 +0xc0
	main.main()
	        /THE_PATH/main.go:11 +0x7d

	Process finished with exit code 2

	*/
}
