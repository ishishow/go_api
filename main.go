package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}

	for i:=0; i<10; i++ {
		wg.Add(1) //wgをインクリメント　GoRoutineを動かす前にするのが大事
		go func(i int) {
			fmt.Println(i)
			wg.Done() //wgをデクリメント
		}(i)
	}

	wg.Wait() // wgがゼロになるまで待つ
}

