package main

import (
	"fmt"
	"log"
)

type MathError struct {
	lat string
	long string
	err error
}

func main() {
	_, err := sqrt(-10)
	if err != nil{
		log.Println(err)
	}
}

func (n MathError) Error() string {
	return fmt.Sprintf("Error occured: %v %v %v", n.lat, n.long, n.err)
}

func sqrt(f float64) (float64, error)  {
	if f < 0{
		errMsg := fmt.Errorf("Negative message not allow %v", f)
		return 0, MathError{"XXXX", "YYYY", errMsg}
	}
	return 42, nil
}
/*output
2019/07/28 15:33:10 Error occured: XXXX YYYY Negative message not allow -10
Process finished with exit code 0
 */