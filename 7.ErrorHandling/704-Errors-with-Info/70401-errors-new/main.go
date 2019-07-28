package main

import (
	"github.com/pkg/errors"
	"log"
)

var errorMsg = errors.New("negative number is not allow")

func main() {
	_, err := sqrt(-10)
	if err != nil{
		log.Fatal(err)
	}
}

func sqrt(f float64) (float64, error){
	if f < 0 {
		return 0, errorMsg
	}

	return 42, nil
}

/* output
2019/07/28 14:41:38 negative number is not allow
Process finished with exit code 1
 */