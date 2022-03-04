package slice

import "testing"

func TestSliceAppend(t *testing.T) {
	var sli []string
	s1 := make([]int, 10)
	t.Logf("s1: %+v", s1)

	s2 := make([]int, 0, 10)
	t.Logf("s2: %+v", s2)

	sli = append(sli, "hello")
	t.Logf("sli: %v, sli length: %d, sli capacity: %d", sli, len(sli), cap(sli))
}
