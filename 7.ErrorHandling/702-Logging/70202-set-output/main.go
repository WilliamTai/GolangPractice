package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Create("70202-log.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	log.SetOutput(f)
	///log.SetOutput will open a writer and get ready for output log

	f2, err := os.Open("70202-no-file.txt")
	if err != nil {
		log.Println("Error! ", err)
		///The error message will be dump to the log file.
	}
	defer f2.Close()

	fmt.Println("Hey! take a look the log file!")
}
