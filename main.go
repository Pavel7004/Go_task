package main

import (
	"fmt"
	"sync"
	"time"
)

func rangeSum(a uint64, b uint64) uint64 {
	var sum uint64 = 0
	for i := a; i <= b; i++ {
		sum += i
	}
	return sum
}

func rangeSumAsync(a uint64, b uint64) uint64 {
	const threadCount = 4
	interval := (b-a+1)/threadCount - 1
	balance := (b - a + 1) % threadCount
	var wg sync.WaitGroup
	var sum uint64 = 0
	first := a
	second := first + interval
	wg.Add(threadCount)
	for i := 0; i < threadCount; i++ {
		go func(a uint64, b uint64) {
			sum += rangeSum(a, b)
			wg.Done()
		}(first, second)
		first = second + 1
		second += interval + 1
		if balance > 0 {
			second++
			balance--
		}
	}
	wg.Wait()
	return sum
}

type testFunc func(uint64, uint64) uint64

func countRuntime(a uint64, b uint64, function testFunc) (uint64, int64) {
	t1 := time.Now()
	ret := function(a, b)
	t2 := time.Now()
	return ret, (t2.Sub(t1)).Nanoseconds()
}

func main() {
	var a, b uint64
	fmt.Printf("Enter range: ")
	fmt.Scanf("%d %d", &a, &b)
	print, nanoSec := countRuntime(a, b, rangeSum)
	fmt.Println("Normal func")
	fmt.Println("Result:", print, "... Exec time is", nanoSec, "ns")
	print, nanoSec = countRuntime(a, b, rangeSumAsync)
	fmt.Println("Async func")
	fmt.Println("Result:", print, "... Exec time is", nanoSec, "ns")
}
