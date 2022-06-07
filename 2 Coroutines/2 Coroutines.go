package main

import (
	"fmt"
	"sync"
	"time"
)

/* 本程序根据教程 https://www.bilibili.com/video/BV1Ju41167BL?p=1 编写， 通过利用 Go 当中信道的阻塞机制，让拿到‘球‘的协程执行代码块，
没拿到’球’的协程阻塞等待取‘球’。实现了两个协程交替输出一个字符串。
*/
func main() {

	var waiting sync.WaitGroup
	waiting.Add(2)
	str := "We Are Going Down."
	byte1 := []byte(str)
	BytesChan := make(chan byte, len(str))
	count := make(chan string, 0)
	for _, v := range byte1 {
		BytesChan <- v
	}
	close(BytesChan)

	go func() {
		defer waiting.Done()
		for {
			ball, ok1 := <-count
			if ok1 {
				print, ok2 := <-BytesChan
				if ok2 {
					fmt.Printf("routine1: %c \n", print)
				} else {
					close(count) //if there is no data taken out from stringChan, count Channel will close.
					return
				}
				count <- ball
			} else {
				return
			}

		}
	}()
	go func() {
		defer waiting.Done()
		time.Sleep(1 * time.Millisecond)
		for {
			ball, ok1 := <-count
			if ok1 {
				print, ok2 := <-BytesChan
				if ok2 {
					fmt.Printf("routine2: %c \n", print)
				} else {
					close(count) //if there is no data taken out from stringChan, count Channel will close.
					return
				}
				count <- ball
			} else {
				return
			}
		}
	}()
	count <- "ball"
	waiting.Wait()

}
