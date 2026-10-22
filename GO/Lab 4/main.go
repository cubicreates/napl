package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

type TaskResult struct {
	TaskName string
	Input    int
	Output   int
}

func computeSquare(n int, ch chan<- TaskResult, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("  [START]    Square goroutine started (n = %d)\n", n)

	time.Sleep(150 * time.Millisecond)
	res := n * n

	fmt.Printf("  [COMPLETE] Square goroutine finished -> %d^2 = %d\n", n, res)
	ch <- TaskResult{TaskName: "Square", Input: n, Output: res}
}

func computeCube(n int, ch chan<- TaskResult, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("  [START]    Cube goroutine started (n = %d)\n", n)

	time.Sleep(250 * time.Millisecond)
	res := n * n * n

	fmt.Printf("  [COMPLETE] Cube goroutine finished -> %d^3 = %d\n", n, res)
	ch <- TaskResult{TaskName: "Cube", Input: n, Output: res}
}

func computeFibonacci(n int, ch chan<- TaskResult, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("  [START]    Fibonacci goroutine started (n = %d)\n", n)

	time.Sleep(200 * time.Millisecond)
	res := fibonacci(n)

	fmt.Printf("  [COMPLETE] Fibonacci goroutine finished -> Fib(%d) = %d\n", n, res)
	ch <- TaskResult{TaskName: "Fibonacci", Input: n, Output: res}
}

func fibonacci(n int) int {
	if n <= 0 {
		return 0
	} else if n == 1 {
		return 1
	}
	a, b := 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}

func printDivider() {
	fmt.Println(strings.Repeat("-", 65))
}

func main() {
	fmt.Println("=================================================================")
	fmt.Println("  EXPERIMENT 6: Concurrent Tasks Using Goroutines and Channels   ")
	fmt.Println("=================================================================")
	fmt.Println()

	squareInput := 8
	cubeInput := 5
	fibInput := 9

	resultChan := make(chan TaskResult, 3)

	var wg sync.WaitGroup
	wg.Add(3)

	fmt.Println(">>> Step 1: Launching Goroutines Concurrently...")
	printDivider()

	go computeSquare(squareInput, resultChan, &wg)
	go computeCube(cubeInput, resultChan, &wg)
	go computeFibonacci(fibInput, resultChan, &wg)

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	fmt.Println()
	fmt.Println(">>> Step 2: Collecting Results in main() via Channel...")
	printDivider()

	for res := range resultChan {
		fmt.Printf("  [CHANNEL RECV] Task: %-10s | Input: %-3d | Result: %d\n",
			res.TaskName, res.Input, res.Output)
	}

	printDivider()
	fmt.Println("  All concurrent tasks completed and results processed in main().")
	fmt.Println("=================================================================")
}
