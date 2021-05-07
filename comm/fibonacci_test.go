package comm

import (
	"fmt"
	"testing"
)

func TestFibonacci(t *testing.T) {
	c := make(chan int, 10)
	Fibonacci(c)
	for n := range c {
		fmt.Println(n)
	}
}
