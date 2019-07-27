package main

import (
	"log"
	"os"
)

func main() {
	_, err := os.Open("no-file.txt")
	if err != nil {
		log.Println("Error happend:", err)
	}

	/*output
	2019/07/27 17:53:33 Error happend: open no-file.txt: no such file or directory
	*/
	///log.Println calls Output to print to the standard logger.
}
