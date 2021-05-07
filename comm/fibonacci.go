package comm

// 将斐波那契数列项存入通道
func Fibonacci(c chan int) {
	x, y := 0, 1
	for i := 0; i < cap(c); i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}
