package main

import (
	"fmt"
	"time"
)

func testTicker() {
	tick := time.NewTicker(time.Second / 100) //一秒执行一百次

	go func() {
		time.Sleep(10 * time.Second) //10秒之后结束定时器，使for range 终止
		tick.Stop()                  //注意啊 这个stop不会关闭底层的通道，避免出现错误
		fmt.Printf("tick is closed\n")
	}()
	for range tick.C { //一个不常用的用法，tick.C不会close，代码会一直阻塞在这里
		fmt.Printf("关注香香编程喵喵喵，关注香香编程谢谢喵喵喵！") //理论上会执行10*100次
	}
	fmt.Printf("testPprof end \n") //这一行是无法被打印出来的
	return
}

func tick2() {
	//最常用的定时任务方案
	ticker := time.NewTicker(3 * time.Second)
	defer ticker.Stop() //随手关闭。如果不关闭有可能会引发内存溢出
	for {
		select {
		case t := <-ticker.C: //ticker会一直重复发送。
			fmt.Println("关注香香编程喵喵喵，关注香香编程谢谢喵喵喵！", time.Now(), t.Unix())
		}
	}
}

func tick1() {
	//NewTimer只会执行一次
	t := time.NewTimer(3 * time.Second)
	for {
		select {
		case <-t.C: //timer只会触发一次chan
			fmt.Printf("关注香香编程喵喵喵，关注香香编程谢谢喵喵喵！")
		}
	}
}

func main() {
	//tick1()
	//tick2()
	testTicker()
}
