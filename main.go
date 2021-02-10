package main

import (
	"fmt"
	"time"
)

func rangeSum(a uint64, b uint64) uint64 {
	var sum uint64 = 0
	for i := a; i <= b; i++ {
		sum += i
	}
	return sum
}

func execAndCountRuntime(a uint64, b uint64, f func(uint64, uint64) uint64) (uint64, int64) {
	t1 := time.Now()
	ret := f(a, b)
	t2 := time.Now()
	return ret, (t2.Sub(t1)).Nanoseconds()
}

func rangeSumAsyncWorker(ch chan uint64, a uint64, b uint64) {
	ch <- rangeSum(a, b)
}

func rangeSumAsync(a uint64, b uint64) uint64 {
	var sum uint64 = 0
	interval := (b-a+1)/4 - 1
	ch := make(chan uint64, 3)
	firstArg, secondArg := a, a+interval
	for i := 0; i < 4; i++ {
		go rangeSumAsyncWorker(ch, firstArg, secondArg)
		firstArg = secondArg + 1
		secondArg = firstArg + interval
		if i == 2 {
			secondArg = b
		}
	}
	for i := 0; i < 4; i++ {
		sum += (<-ch)
	}
	return sum
}

func main() {
	fmt.Printf("Enter range: ")
	var a, b uint64 = 0, 0
	fmt.Scanf("%d %d", &a, &b)
	print, nanoSec := execAndCountRuntime(a, b, rangeSum)
	fmt.Println("Normal func")
	fmt.Println("Result:", print, "... Exec time is", nanoSec, "ns")
	print, nanoSec = execAndCountRuntime(a, b, rangeSumAsync)
	fmt.Println("Async func")
	fmt.Println("Result:", print, "... Exec time is", nanoSec, "ns")
}
