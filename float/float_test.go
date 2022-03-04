package float

import (
	"fmt"
	"testing"
)

func TestFloat(t *testing.T)  {
	var f1 float64
	var f2 float64

	if f1 > 0 {
		fmt.Println("default float value bigger than 0")
	} else if f1 < 0 {
		fmt.Println("default float value smaller than 0")
	} else {
		fmt.Println("default float value equal than 0")
	}

	if f2 == 0 {
		fmt.Println("f2 default float value equal than 0")
	}

	fmt.Println("default float value: ", f1)
}
