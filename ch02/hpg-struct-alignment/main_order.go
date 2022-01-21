package main

import (
	"fmt"
	"unsafe"
)

type demo1 struct {
	a int8  //1
	b int16 //2 开始的位置必须是2的整数倍如从，2、4、6、8的位置开始
	c int32 //a0bbcccc //开始的位置必须是4的整数倍如从，4、8的位置开始
	//最终放置是a0bbcccc
}
type demo2 struct {
	a int8
	c int32 //4开始前边空三个
	b int16 //放置的位置从2的整数倍开始
	//unsafe.Alignof 的规则, struct 结构体类型的变量 x, 计算 x 每一个字段 f 的,等于其中的最大值。
	//demo2 的对齐倍数由 c 的对齐倍数决定, 是4
	//最终放置是a000ccccbb00

}

func main() {
	fmt.Println(unsafe.Sizeof(demo1{}))
	fmt.Println(unsafe.Sizeof(demo2{}))

}
