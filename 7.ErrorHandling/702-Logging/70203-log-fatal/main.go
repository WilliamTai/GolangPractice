package main

import (
	"log"
	"os"
)

func main() {
	_, err := os.Open("no-file.txt")
	if err != nil {
		log.Fatalln(err)
		//Fatal will call os.exit(1) which means shutdown the program with error code immediately
	}

	/*
	2019/07/27 19:49:49 open no-file.txt: no such file or directory
	Process finished with exit code 1
	*/
}
