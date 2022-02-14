// package main

// import "fmt"

// 在 切片(slice)性能及陷阱 这篇文章中，我们也提到了切片由三个值构成：

// *ptr 指向底层数组的指针
// len 长度
// cap 容量
// 因此，将切片作为参数时，拷贝了一个新切片，即拷贝了构成切片的三个值，包括底层数组的指针。对切片中某个元素的修改，实际上是修改了底层数组中的值，因此原切片也发生了改变。

// func foo(a []int) {
// 	//切片操作并不复制切片指向的元素，创建一个新的切片会复用原来切片的底层数组，
// 	a[0] = 200
// }

// func main() {
// 	a := []int{1, 2}
// 	foo(a)
// 	fmt.Println(a)
// }
