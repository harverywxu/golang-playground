package _range

import (
	"fmt"
	"testing"
)

// 在 range 过程中，v 是一个全局变量
func TestRange(t *testing.T) {
	arr := [2]int{1, 2}
	var resptr []*int
	var res []int
	for _, v := range arr {
		resptr = append(resptr, &v)
		res = append(res, v)
	}
	//expect: 1 2
	fmt.Println(*resptr[0], *resptr[1])
	fmt.Println(res)
	//but output: 2 2
}

//arr value:[word1 word2], ptr:0xc00000c0c0
//key :0, ptr:0xc0000162b8
//value :word1, ptr:0xc000042520
//key :1, ptr:0xc0000162b8
//value :word2, ptr:0xc000042520
func TestRangeIndex(t *testing.T) {
	arr := []string{"word1", "word2"}
	t.Logf("arr value:%v, ptr:%p", arr, arr)
	for k, v := range arr {
		t.Logf("key :%v, ptr:%v", k, &k)
		t.Logf("value :%v, ptr:%v", v, &v)
	}
}

//range_test.go:39: arr value:map[key1:value1 key2:value2], ptr:0x1400005e360
//range_test.go:41: key:key1, ptr:0x1400004a590
//range_test.go:42: value:value1, prt:0x1400004a5a0
//range_test.go:43: rv:value1, rprt:value1
//range_test.go:41: key:key2, ptr:0x1400004a590
//range_test.go:42: value:value2, prt:0x1400004a5a0
//range_test.go:43: rv:value2, rprt:value2
func TestMapIndex(t *testing.T) {
	mp := map[string]string{
		"key1": "value1",
		"key2": "value2",
	}
	t.Logf("arr value:%v, ptr:%p", mp, mp)
	for k, v := range mp {
		t.Logf("key:%v, ptr:%v", k, &k)
		t.Logf("value:%v, prt:%v", v, &v)
		t.Logf("rv:%v, rprt:%v", mp[k], mp[k])
	}
}

//in goroutine:  2 3
//in goroutine:  2 3
//slice:  [1 11 22 1 2 3]
//in map:  2 -> 3
//in map:  0 -> 3
//in map:  1 -> 3
//slice2:  [1 2 3]
//in goroutine:  2 3
func TestForRange(t *testing.T) {
	slice := []int{1, 2, 3}
	m := make(map[int]*int)
	var slice2 [3]int
	for index,value := range slice {
		slice = append(slice, value)
		go func(){
			fmt.Println("in goroutine: ",index,value)
		}()
		// time.Sleep(time.Second * 1)
		m[index] = &value
		if index == 0{
			slice[1] = 11
			slice[2] = 22
		}
		slice2[index] = value
	}
	fmt.Println("slice: ",slice)
	for key,value := range m {
		fmt.Println("in map: ",key,"->",*value)
	}
	fmt.Println("slice2: ",slice2)
	// time.Sleep(time.Second * 10)
}

// 遍历过程中 append slice 导致，slice 中存储区地址变化，新的 slice 内容变化
// 由于遍历时，副本切片 a 还是指向老的 slice，故获取的内容还是老 slice 内容

//a addr: 0x14000144030, len: 3, cap: 3
//r addr: 0x14000144048, len: 3, cap: 3
//a addr: 0x1400010c030, len: 4, cap: 6
//a addr: 0x1400010c030, len: 5, cap: 6
//a addr: 0x1400010c030, len: 6, cap: 6
//a addr: 0x1400010c030, len: 6, cap: 6
//r addr: 0x14000144048, len: 3, cap: 3
//r =  [1 2 3]
//a =  [1 12 13 1 2 3]
func TestSliceAppend(t *testing.T)  {
	var a = []int{1, 2, 3}
	var r = make([]int,3)
	fmt.Printf("a addr: %p, len: %d, cap: %d\n", a, len(a), cap(a))
	fmt.Printf("r addr: %p, len: %d, cap: %d\n", r, len(r), cap(r))
	// 此处 a 为老 slice 副本，内容不变
	for i, v := range a {
		// 磁盘新 slice 内容变化，地址、长度、容量
		a = append(a, v)
		fmt.Printf("a addr: %p, len: %d, cap: %d\n", a, len(a), cap(a))
		if i == 0 {
			// 新 slice 存储值变化
			a[1] = 12
			a[2] = 13
		}
		r[i] = v
	}
	fmt.Printf("a addr: %p, len: %d, cap: %d\n", a, len(a), cap(a))
	fmt.Printf("r addr: %p, len: %d, cap: %d\n", r, len(r), cap(r))
	fmt.Println("r = ", r)
	fmt.Println("a = ", a)
}

// 在遍历过程中 in-place 修改 slice 内容，因为 slice 拷贝副本中包含 具体内容指针，
// 且在拷贝过程中，存储内容位置未发生变化

//a addr: 0x140000160c0, len: 3, cap: 3
//r addr: 0x140000160d8, len: 3, cap: 3
//a addr: 0x140000160c0, len: 3, cap: 3
//a addr: 0x140000160c0, len: 3, cap: 3
//a addr: 0x140000160c0, len: 3, cap: 3
//a addr: 0x140000160c0, len: 3, cap: 3
//r addr: 0x140000160d8, len: 3, cap: 3
//r =  [1 12 13]
//a =  [1 12 13]
func TestSlice(t *testing.T)  {
	var a = []int{1, 2, 3}
	var r = make([]int,3)
	fmt.Printf("a addr: %p, len: %d, cap: %d\n", a, len(a), cap(a))
	fmt.Printf("r addr: %p, len: %d, cap: %d\n", r, len(r), cap(r))
	for i, v := range a {
		fmt.Printf("a addr: %p, len: %d, cap: %d\n", a, len(a), cap(a))
		if i == 0 {
			a[1] = 12
			a[2] = 13
		}
		r[i] = v
	}
	fmt.Printf("a addr: %p, len: %d, cap: %d\n", a, len(a), cap(a))
	fmt.Printf("r addr: %p, len: %d, cap: %d\n", r, len(r), cap(r))
	fmt.Println("r = ", r)
	fmt.Println("a = ", a)
}

// range 遍历数组时是数组的拷贝，修改是修改愿数组
// r =  [1 2 3]
// a =  [1 12 13]
func TestArray(t *testing.T)  {
	var a = [3]int{1, 2, 3}
	var r [3]int
	for i, v := range a {
		if i == 0 {
			a[1] = 12
			a[2] = 13
		}
		r[i] = v
	}
	fmt.Println("r = ", r)
	fmt.Println("a = ", a)
}