package unmarshal

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func TestJsonUnmarshal(t *testing.T) {
	js1 := "{\"name\":\"test\"}"
	js2 := "{\"Name\":\"test\"}"

	c1 := &Person{}
	c2 := &Person{}
	_ = json.Unmarshal([]byte(js1), c1)
	_ = json.Unmarshal([]byte(js2), c2)

	assert.Equal(t, c1, c2)
}
