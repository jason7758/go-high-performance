package main

import (
	"testing"
	"time"
)

// func BenchmarkFib(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		fib(30)
// 	}
// }
func BenchmarkFib(b *testing.B) {
	time.Sleep(time.Second * 3) // 模拟耗时准备任务
	b.ResetTimer()              // 重置定时器;
	for n := 0; n < b.N; n++ {
		fib(30) // run fib(30) b.N times
	}
}
