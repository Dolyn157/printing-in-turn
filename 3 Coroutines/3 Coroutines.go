package main

import (
	"fmt"
	"sync"
)

var waiting sync.WaitGroup

func main() {
	//引入了回合控制机制

	waiting.Add(3)
	str := "We Are Going Down."
	byte1 := []byte(str)
	BytesChan := make(chan byte, len(str))
	count := make(chan int)
	for _, v := range byte1 {
		BytesChan <- v
	}
	close(BytesChan)

	go Routine1(BytesChan, count)
	go Routine2(BytesChan, count)
	go Routine3(BytesChan, count)

	count <- 1
	waiting.Wait()
	/*启动三个协程， 取到了‘球’的协程通过判断球号 ball 来判断是否是自己的回合，如果球号显示是自己的回合，则执行代码块，之后把‘球’放回‘Count’信道。
	如果球号表示不是自己的回合，也把球放回信道，进入下一回合。
	没有取到球的协程阻塞暂停在取‘球’命令 ball := <-count 处。
	*/
}

func Routine1(bChan chan byte, count chan int) {
	defer waiting.Done()
	for {
		ball, ok1 := <-count
		if ok1 {
			if ball == 1 {
				print, ok2 := <-bChan
				if ok2 {
					fmt.Printf("routine1: %c \n", print)
				} else {
					close(count) //if there is no data taken out from stringChan, count Channel will close.
					return
				}
				ball = 2
			}
			count <- ball
		} else {
			return
		}
	}
}

func Routine2(bChan chan byte, count chan int) {
	defer waiting.Done()
	for {
		ball, ok1 := <-count
		if ok1 {
			if ball == 2 {
				print, ok2 := <-bChan
				if ok2 {
					fmt.Printf("routine2: %c \n", print)
				} else {
					close(count) //if there is no data taken out from stringChan, count Channel will close.
					return
				}
				ball = 3
			}
			count <- ball
		} else {
			return
		}
	}
}

func Routine3(bChan chan byte, count chan int) {
	defer waiting.Done()
	for {
		ball, ok1 := <-count
		if ok1 {
			if ball == 3 {
				print, ok2 := <-bChan
				if ok2 {
					fmt.Printf("routine3: %c \n", print)
				} else {
					close(count) //if there is no data taken out from stringChan, count Channel will close.
					return
				}
				ball = 1
			}
			count <- ball
		} else {
			return
		}
	}
}
