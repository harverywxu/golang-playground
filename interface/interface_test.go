package _interface

import (
	"fmt"
	"testing"

	"github.com/golang-collections/lib.go/assert"
)

// Interface, method implement
type IGreeting interface {
	sayHello()
}

func sayHello(i IGreeting) {
	i.sayHello()
}

type Go struct{}

func (g Go) sayHello() {
	fmt.Println("Hi, I am GO!")
}

type PHP struct{}

func (p PHP) sayHello() {
	fmt.Println("Hi, I am PHP!")
}

func TestInterface(t *testing.T) {
	golang := Go{}
	php := PHP{}

	sayHello(golang)
	sayHello(php)
}

// Interface, method acceptance
type Person struct {
	age int
}

func (p Person) howOld() int {
	return p.age
}

func (p *Person) growUp() {
	p.age += 1
}

func TestMethodAcceptance(t *testing.T) {
	ast := assert.New(t)
	// qcrao 是值类型
	qcrao := Person{age: 18}

	// 值类型 调用接收者也是值类型的方法
	fmt.Println(qcrao.howOld())

	// 值类型 调用接收者是指针类型的方法
	qcrao.growUp()
	fmt.Println(qcrao.howOld())

	ast.Equal(qcrao.howOld(), 19)

	// ----------------------

	// stefno 是指针类型
	stefno := &Person{age: 100}

	// 指针类型 调用接收者是值类型的方法
	fmt.Println(stefno.howOld())

	// 指针类型 调用接收者也是指针类型的方法
	stefno.growUp()
	fmt.Println(stefno.howOld())
	ast.Equal(stefno.howOld(), 101)
}

// Interface implement, acceptance 2
type coder interface {
	code()
	debug()
}

type Gopher struct {
	language string
}

// Gopher/*Gopher（隐式）类型都实现了该方法
// 对于接收者是值类型的方法，在方法中不会对接收者本身产生影响，故默认其指针类型隐式实现该方法
func (p Gopher) code() {
	fmt.Printf("I am coding %s language\n", p.language)
}

// 仅*Gopher类型都实现了该方法
// 接收者是指针类型的方法，很可能在方法中会对接收者的属性进行更改操作，从而影响接收者
func (p *Gopher) debug() {
	fmt.Printf("I am debuging %s language\n", p.language)
}

func TestInterfaceCoder(t *testing.T) {
	var c coder = &Gopher{"GO"}
	c.code()
	c.debug()
}

func TestMapInterfacePrint(t *testing.T) {
	m := map[string]interface{}{
		"key1": "value1",
	}
	v1 := fmt.Sprint(m["key1"])
	var v2 string
	if _, ok := m["key2"]; !ok {
		v2 = ""
	} else {
		v2 = fmt.Sprint(m["key2"])
	}
	assert.Equal(t, v1, "value1")
	assert.Equal(t, v2, "")
}
