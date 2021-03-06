package main

import (
	"fmt"
	"log"
	"os"
	"runtime/trace"
)

//**What is the largest prime factor of the number 600851475143?**
//Takes way too long (~47min). Attempted pipline model.
func main() {
	file, err := os.Create("trace.out")
	if err != nil {
		log.Fatalf("failed to create trace output file: %v", err)
	}
	defer file.Close()

	if err := trace.Start(file); err != nil {
		log.Fatalf("failed to start trace: %v", err)
	}
	defer trace.Stop()
	f := factorsOf(600851475143)
	pr := isPrime(f)
	fmt.Println(<-pr)
}

func factorsOf(n int64) chan int64 {
	out := make(chan int64)
	go func() {
		for i := int64(3); i < n; i += 2 {
			if n%i == 0 {
				out <- i
			}
		}
		close(out)
	}()
	return out
}

func isPrime(ch chan int64) chan int64 {
	out := make(chan int64)
	var lrg int64
	go func() {
		for i := range ch {
			for j := int64(3); j < i; j += 2 {
				if i%j == 0 {
					i = 0
				}
			}
			if i > lrg {
				lrg = i
			}
		}
		out <- lrg
		close(out)
	}()
	return out
}
