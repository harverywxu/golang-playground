package string

import (
	"fmt"
	"reflect"
	"testing"
)

func TestString(t *testing.T) {
	// int32
	var x = 'C'
	// uint8
	var y byte = 'C'
	// int32, (rune, 默认类型，可忽略)
	var z rune = 'C'
	fmt.Println("type x", reflect.TypeOf(x))
	fmt.Println("type y", reflect.TypeOf(y))
	fmt.Println("type z", reflect.TypeOf(z))
}
