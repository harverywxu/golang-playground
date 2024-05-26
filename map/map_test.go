package _map

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMapMap(t *testing.T) {
	var m = make(map[string]map[string]int)
	m["foo"] = make(map[string]int)
	m["foo"]["key1"] = 1
	x := m["foo"]
	// x before: map[key1:1]
	t.Logf("x before: %+v", x)

	m["foo"]["key1"] = 2
	m["foo"]["key2"] = 3

	t.Logf("m: %+v", m)
	// x after: map[key1:2 key2:3]
	t.Logf("x after: %+v", x)
}

func TestSetMapMember(t *testing.T) {
	var m = map[string]struct{ x, y int }{
		"foo": {2, 3},
	}

	// 1 不能直接修改 map 中成员变量的值，因为不是并发安全的
	// Cannot assign to m["foo"].x
	// m["foo"].x = 4

	fmt.Printf("result is : %+v", m)

	// 2 可以直接通过 key 取出值，修改后，覆盖原值
	tmp := m["foo"]
	tmp.x = 4
	m["foo"] = tmp
	assert.Equal(t, 4, m["foo"].x)

	// 如果值为指针，则可以直接修改
	mp := map[string]*struct{ x, y int }{
		"foo": {2, 3},
	}
	mp["foo"].x = 4
	assert.Equal(t, 4, mp["foo"].x)
}

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var (
	PtrRequiredErr     = fmt.Errorf("b must be a pointer")
	NothingToBeDoneErr = fmt.Errorf("nothing to be done")
)

func Transfer(a, b interface{}) error {
	if reflect.ValueOf(b).Kind() != reflect.Ptr {
		return PtrRequiredErr
	}
	if a == nil {
		return NothingToBeDoneErr
	}
	aBytes, err := json.Marshal(a)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(aBytes, b); err != nil {
		return err
	}
	return nil
}

func TestMap(t *testing.T) {
	p := Person{
		Name: "harvey",
		Age:  30,
	}
	target := map[string]interface{}{}
	target1 := make(map[string]interface{})

	err := Transfer(p, &target)
	if err != nil {
		t.Errorf("error: %+v", err)
	}
	t.Logf("target: %+v", target)

	err = Transfer(p, target1)
	if err != nil {
		t.Errorf("error: %+v", err)
	}
	t.Logf("target1: %+v", target1)
}

func TestMapEmpty(t *testing.T) {
	var m map[string]string
	m = nil
	t.Logf("map len: %v", len(m))
}
