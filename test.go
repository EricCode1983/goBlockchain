package main

import (
	"fmt"
	"time"
)

func main() {
	bread := make(chan int,100)
	for i:=1;i<=3;i++ {
		go produce(bread)
	}
	for i:=1;i<=2;i++ {
		go consume(bread)
	}
	time.Sleep(1*time.Second)
}

func produce(ch chan<- int) {
	for {
		ch <- 1
		fmt.Println("produce bread"+time.Now().Format("2006-01-02 15:04:05"))
		time.Sleep(100 * time.Millisecond)
	}
}

func consume(ch <-chan int) {
	for {
		<-ch
		fmt.Println("take bread"+time.Now().Format("2006-01-02 15:04:05"))
		time.Sleep(200 * time.Millisecond)
	}
}
