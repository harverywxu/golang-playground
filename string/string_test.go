package string

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"reflect"
	"strings"
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

func TestStringSplit(t *testing.T) {
	str := ""
	sl := strings.Split(str, ",")
	// len(sl) = 1, [""]
	t.Logf("sl: %+v", sl)
	t.Logf("len(sl) = %v", len(sl))
	t.Logf("sl == nil:  %v", sl == nil)
	assert.Equal(t, sl, []string{""})
}

func TestStringPrefix(t *testing.T) {
	url := "/aaa/bbb/ccc"
	if strings.HasPrefix(url, "/") {
		url = url[1:]
	}
	t.Logf("url: %v", url)
}
