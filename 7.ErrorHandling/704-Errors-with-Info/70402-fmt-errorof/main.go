package main

import (
	"fmt"
	"log"
)

func main() {
	_, err := sqrt(-10)
	if err != nil{
		log.Fatal(err)
	}
}

func sqrt(f float64) (float64, error){
	if f < 0 {
		errorMsg := fmt.Errorf("negative number is not allow. input: %v", f)
		//you can put more information with string format which is good for debugging
		return 0, errorMsg
	}

	return 42, nil
}

/* output
2019/07/28 14:52:40 negative number is not allow. input: -10
Process finished with exit code 1
*/