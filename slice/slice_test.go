package slice

import (
	"fmt"
	"testing"
)

func TestSliceAppend(t *testing.T) {
	var sli []string
	s1 := make([]int, 10)
	t.Logf("s1: %+v", s1)

	s2 := make([]int, 0, 10)
	t.Logf("s2: %+v", s2)

	sli = append(sli, "hello")
	t.Logf("sli: %v, sli length: %d, sli capacity: %d", sli, len(sli), cap(sli))
}

func appendToSlice(s *[]string, vals ...string) {
	*s = append(*s, vals...)
}

func TestSlicePara(t *testing.T) {
	sli := make([]string, 0)

	appendToSlice(&sli, "a")
	appendToSlice(&sli, "b")
	appendToSlice(&sli, "c")
	appendToSlice(&sli, "c")
	appendToSlice(&sli, "c")
	appendToSlice(&sli, "c")
	appendToSlice(&sli, "c")
	appendToSlice(&sli, "c")
	appendToSlice(&sli, "c", "d", "e", "f", "g")

	t.Logf("sli out function: %v", sli)
}

func updateSlice(s []int) {
	for i := range s {
		s[i] *= 2
	}
}

func TestSlicePara2(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}
	updateSlice(s)
	fmt.Println(s)
}
