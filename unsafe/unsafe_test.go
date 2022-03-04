package unsafe

import (
	"fmt"
	"testing"
	"unsafe"
)

//type slice struct {
//	array unsafe.Pointer // 元素指针
//	len   int // 长度
//	cap   int // 容量
//}

func TestSlice(t *testing.T)  {
	s := make([]int, 9, 20)
	var Len = *(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&s)) + uintptr(8)))
	fmt.Println(Len, len(s)) // 9 9

	var Cap = *(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&s)) + uintptr(16)))
	fmt.Println(Cap, cap(s)) // 20 20
}

//type hmap struct {
//	count     int
//	flags     uint8
//	B         uint8
//	noverflow uint16
//	hash0     uint32
//
//	buckets    unsafe.Pointer
//	oldbuckets unsafe.Pointer
//	nevacuate  uintptr
//
//	extra *mapextra
//}

func TestMap(t *testing.T)  {
	mp := make(map[string]int)
	mp["qcrao"] = 100
	mp["stefno"] = 18

	count := **(**int)(unsafe.Pointer(&mp))
	fmt.Println(count, len(mp)) // 2 2
}


