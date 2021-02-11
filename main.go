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

type testFunc struct {
	body func(uint64, uint64) uint64
	args [2]uint64
}

func countRuntime(function testFunc) (uint64, int64) {
	t1 := time.Now()
	ret := function.body(function.args[0], function.args[1])
	t2 := time.Now()
	return ret, (t2.Sub(t1)).Nanoseconds()
}

func main() {
	fmt.Printf("Enter range: ")
	var function testFunc
	function.args = [2]uint64{}
	fmt.Scanf("%d %d", &function.args[0], &function.args[1])
	function.body = rangeSum
	print, nanoSec := countRuntime(function)
	fmt.Println("Normal func")
	fmt.Println("Result:", print, "... Exec time is", nanoSec, "ns")
	function.body = rangeSumAsync
	print, nanoSec = countRuntime(function)
	fmt.Println("Async func")
	fmt.Println("Result:", print, "... Exec time is", nanoSec, "ns")
}
