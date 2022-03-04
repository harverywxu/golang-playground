package value_point

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// [关于 Go 中 Map 类型和 Slice 类型的传递](https://www.cnblogs.com/snowInPluto/p/7477365.html)

func TestSliceAppend(t *testing.T) {
	sl := []string{"hello", "world"}
	f := func(s []string) {
		s = append(s, "add")
		t.Logf("func, add slice: %+v", s)
		assert.Equal(t, []string{"hello", "world", "add"}, s)
	}
	f(sl)
	t.Logf("add slice: %+v", sl)
	assert.Equal(t, []string{"hello", "world"}, sl)
}

func TestSliceSet(t *testing.T) {
	sl := []string{"hello", "world"}
	f := func(s []string) {
		s[0] = "11"
		s[1] = "22"
		assert.Equal(t, []string{"11", "22"}, s)
	}
	f(sl)
	t.Logf("add slice: %+v", sl)
	assert.Equal(t, []string{"11", "22"}, sl)
}

func TestSlicePoint(t *testing.T) {
	sl := []string{"hello", "world"}
	f := func(s *[]string) {
		*s = append(*s, "add")
		assert.Equal(t, []string{"hello", "world", "add"}, *s)
	}
	f(&sl)
	assert.Equal(t, []string{"hello", "world", "add"}, sl)
}

func TestMap(t *testing.T)  {
	m := make(map[string]string)
	m["hello"] = "hello"
	m["world"] = "world"

	f := func(mt map[string]string) {
		mt["add"] = "add"
		assert.Equal(t, map[string]string{
			"hello": "hello",
			"world": "world",
			"add": "add",
		}, mt)
	}

	f(m)
	assert.Equal(t, map[string]string{
		"hello": "hello",
		"world": "world",
		"add": "add",
	}, m)
}