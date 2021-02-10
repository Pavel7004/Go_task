package main

import (
	"fmt"
	"sync"
	"time"
)

func createMultiInterval(a uint64, b uint64, count uint64) [][2]uint64 {
	interval := (b-a+1)/count - 1
	balance := (b - a + 1) % count
	var ret [][2]uint64 = make([][2]uint64, int(count))
	first := a
	second := a + interval
	for i := 0; i < int(count); i++ {
		ret[i][0] = first
		ret[i][1] = second
		first = second + 1
		second += interval + 1
		if balance > 0 {
			second++
			balance--
		}
	}
	return ret
}

func rangeSum(a uint64, b uint64) uint64 {
	var sum uint64 = 0
	for i := a; i <= b; i++ {
		sum += i
	}
	return sum
}

func rangeSumAsync(a uint64, b uint64) uint64 {
	const threadCount = 4
	intervals := createMultiInterval(a, b, threadCount)
	var wg sync.WaitGroup
	var sum uint64 = 0
	wg.Add(threadCount)
	for i := 0; i < threadCount; i++ {
		go func(id int) {
			sum += rangeSum(intervals[id][0], intervals[id][1])
			wg.Done()
		}(i)
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
