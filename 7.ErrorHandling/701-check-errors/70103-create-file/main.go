package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	f, err := os.Create("70103-output.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	r := strings.NewReader("Hello!")
	io.Copy(f, r)
}
