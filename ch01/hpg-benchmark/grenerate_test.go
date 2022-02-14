package main

import (
	"math/rand"
	"testing"
	"time"
)

func generateWithCap(n int) []int {
	rand.Seed(time.Now().UnixNano())
	//创建切片时，将切片的容量(capacity)设置为 n，
	//这样切片就会一次性申请 n 个整数所需的内存。
	nums := make([]int, 0, n)
	for i := 0; i < n; i++ {
		nums = append(nums, rand.Int())
	}
	return nums
}

func generate(n int) []int {
	rand.Seed(time.Now().UnixNano())
	nums := make([]int, 0)
	for i := 0; i < n; i++ {
		nums = append(nums, rand.Int())
	}
	return nums
}

// func BenchmarkGenerateWithCap(b *testing.B) {
// 	for n := 0; n < b.N; n++ {
// 		generateWithCap(1000000)
// 	}
// }

// func BenchmarkGenerate(b *testing.B) {
// 	for n := 0; n < b.N; n++ {
// 		generate(1000000)
// 	}
// }

func benchmarkGenerate(i int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		generate(i)
	}
}

func BenchmarkGenerate1000(b *testing.B)    { benchmarkGenerate(1000, b) }
func BenchmarkGenerate10000(b *testing.B)   { benchmarkGenerate(10000, b) }
func BenchmarkGenerate100000(b *testing.B)  { benchmarkGenerate(100000, b) }
func BenchmarkGenerate1000000(b *testing.B) { benchmarkGenerate(1000000, b) }
