package reflect

import (
	"fmt"
	"io"
	"os"
	"reflect"
	"testing"
)

func TestReflect(t *testing.T)  {
	type Reader interface {
		Read(p []byte) (n int, err error)
	}

	type Writer interface {
		Write(p []byte) (n int, err error)
	}

	var r io.Reader
	tty, err := os.OpenFile("../interface_test.go", os.O_RDWR, 0)
	if err != nil {
		t.Error(err.Error())
	}
	r = tty
	t.Logf("reader:%v", &r)

	var w io.Writer
	w = r.(io.Writer)
	t.Logf("reader:%v", &w)

}

type Rect struct {
	Width int
	Height int
}

func SetRectAttr(r *Rect, name string, value int) {
	var v = reflect.ValueOf(r)
	var field = v.Elem().FieldByName(name)
	field.SetInt(int64(value))
}

func TestReflectElem(t *testing.T)  {
	var r = Rect{50, 100}
	SetRectAttr(&r, "Width", 100)
	SetRectAttr(&r, "Height", 200)
	fmt.Println(r)
}

