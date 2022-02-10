package main

import (
	"fmt"
	"unsafe"
)

type demo3 struct {
	c int32
	a struct{}
}

type demo4 struct {
	a struct{}
	c int32
}

func main() {
	fmt.Println(unsafe.Sizeof(demo3{})) //当 struct{} 作为其他 struct 最后一个字段时，需要填充额外的内存保证安全.即额外填充了 4 字节的空间
	fmt.Println(unsafe.Sizeof(demo4{}))
}
