package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	fmt.Println("it runs for", sum)
	c <- sum // 把 sum 发送到通道 c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int, 10)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	go sum(s[len(s)/3:], c)
	// go sum(s[:len(s)/2], c)
	x := <-c // 从通道 c 中接收
	y := <-c
	z := <-c

	fmt.Println(x, y, z)
}
