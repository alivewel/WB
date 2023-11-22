package main

import (
	"fmt"
	"unsafe"
)

type one struct {
	t1 bool  // 8
	t2 int32 // 8
	t3 bool  // 8
}

type two struct {
	t1  bool
	qwe int
	t2  bool
	t3  bool
	t4  bool
	t5  bool
	t6  bool
	t7  bool
	t8  bool
}

type Foo struct {
	a int
	b bool
	c int
	d bool
}

type Bar struct {
	d bool
	b bool
	a int
	c int
}

func main() {
	first := Foo{}
	second := Bar{}
	fmt.Println(unsafe.Sizeof(first))
	fmt.Println(unsafe.Sizeof(second))
}
