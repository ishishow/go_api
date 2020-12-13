package main

import (
	"time"
	"fmt"
)

func main() {
	//  無名関数(クロージャ)
	go func() {
		time.Sleep(3*time.Second)
		fmt.Println("実行終了！")
	} ()
	fmt.Println("実行開始")
	time.Sleep(10*time.Second)
}