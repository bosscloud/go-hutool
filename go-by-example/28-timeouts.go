package main

import "time"
import "fmt"

func main() {
	//在我们的例子中，假如我们执行一个外部调用，并在 2 秒后 通过通道 c1 返回它的执行结果。

	c1 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		c1 <- "result 1"
	}()
	//这里是使用 select 实现一个超时操作。 res := <- c1 等待结果，<-Time.After 等待超时 时间 1 秒后发送的值。由于 select 默认处理第一个 已准备好的接收操作，如果这个操作超过了允许的 1 秒 的话，将会执行超时 case。

	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(time.Second * 1):
		fmt.Println("timeout 1")
	}
	//如果我允许一个长一点的超时时间 3 秒，将会成功的从 c2 接收到值，并且打印出结果。

	c2 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "result 2"
	}()
	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(time.Second * 3):
		fmt.Println("timeout 2")
	}
}
