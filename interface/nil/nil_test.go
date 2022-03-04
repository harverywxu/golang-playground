package nild

import (
	"fmt"
	"testing"
	"unsafe"
)

type Coder interface {
	code()
}

type Gopher struct {
	name string
}

type st struct {
	name string
}

func (g Gopher) code() {
	fmt.Printf("%s is coding\n", g.name)
}

func TestInterfaceNil(t *testing.T)  {
	// 接口类型，动态类型和动态值都为nil，这个接口才会被认为 接口值==nil
	var c Coder
	fmt.Println(c == nil)
	fmt.Printf("c: %T, %v\n", c, c)

	// 非接口类型，只需要 值==nil
	var g *Gopher
	g = &Gopher{name:"hello world"}
	fmt.Println(g == nil)
	fmt.Printf("g: %T, %v\n", g, g)

	// 非接口类型，只需要 值==nil
	var m *Gopher
	fmt.Println(m == nil)
	fmt.Printf("m: %T, %v\n", m, m)

	// 接口类型，动态值不为nil
	var h *Gopher
	c = h
	fmt.Println(c == nil)
	fmt.Printf("c: %T, %v\n", c, c)

	// 接口类型，动态类型和动态值都不为nil
	var d Coder
	d = g
	fmt.Println(c == nil)
	fmt.Printf("d: %T, %v\n", d, d)

	fmt.Printf("hello world")
}


type MyError struct {}

// 实现error接口
func (i MyError) Error() string {
	return "MyError"
}

// 返回值，MyError转换为error
func Process() error {
	var err *MyError = nil
	return err
}

func TestError(t *testing.T)  {
	err := Process()
	// 动态值为nil
	fmt.Println(err)
	// 动态类型并不为nil
	fmt.Println(err == nil)
}

type iface struct {
	itab, data uintptr
}

// 打印接口的动态类型和值
func TestIface(t *testing.T)  {
	var a interface{} = nil

	var b interface{} = (*int)(nil)

	x := 5
	var c interface{} = (*int)(&x)

	ia := *(*iface)(unsafe.Pointer(&a))
	ib := *(*iface)(unsafe.Pointer(&b))
	ic := *(*iface)(unsafe.Pointer(&c))

	fmt.Println(ia, ib, ic)

	fmt.Println(*(*int)(unsafe.Pointer(ic.data)))
}

func TestConvert(t *testing.T)  {
	var i int = 9

	var f float64
	f = float64(i)
	fmt.Printf("%T, %v\n", f, f)

	f = 10.8
	a := int(f)
	fmt.Printf("%T, %v\n", a, a)
}

// 空接口 interface{} 没有定义任何函数，因此 Go 中所有类型都实现了空接口。当一个函数的形参是 interface{}，那么在函数中，需要对形参进行断言，从而得到它的真实类型
func TestInterfaceAssert(t *testing.T)  {
	type Student struct {
		Name string
		Age int
	}

	// 类型转换和类型断言有些相似，不同之处，在于类型断言是对接口进行的操作。
	var i interface{} = new(Student)
	s, ok := i.(Student)
	if ok {
		fmt.Println(s)
	}

}

func judge(v interface{}) {
	fmt.Printf("%p %v\n", &v, v)

	switch v := v.(type) {
	case nil:
		fmt.Printf("%p %v\n", &v, v)
		fmt.Printf("nil type[%T] %v\n", v, v)

	case Student:
		fmt.Printf("%p %v\n", &v, v)
		fmt.Printf("Student type[%T] %v\n", v, v)

	case *Student:
		fmt.Printf("%p %v\n", &v, v)
		fmt.Printf("*Student type[%T] %v\n", v, v)

	default:
		fmt.Printf("%p %v\n", &v, v)
		fmt.Printf("unknow\n")
	}
}

type Student struct {
	Name string
	Age int
}

func TestNilInterface(t *testing.T)  {
	//var i interface{} = new(Student)
	var i interface{} = (*Student)(nil)
	//var i interface{}

	fmt.Printf("%p %v\n", &i, i)

	judge(i)
}

