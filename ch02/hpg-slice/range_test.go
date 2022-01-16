package main

import (
	"fmt"
	"testing"
)

func TestRange(t *testing.T) {

	//map
	m := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	for k, v := range m {
		delete(m, "two")
		m["four"] = 4
		fmt.Printf("%v: %v\n", k, v)
	}
	//chan
	ch := make(chan string)
	go func ()  {
		ch <- "Go"
		ch <- "语言"
		ch <- "高性能"
		ch <- "编程"
		close(ch)
		
	}()
	for n := range ch {
		fmt.Println(n)
	}

}
