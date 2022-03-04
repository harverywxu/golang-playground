package panicc

import (
	"fmt"
	"os"
	"testing"
	"time"
)

func TestPanic(t *testing.T) {
	var user = os.Getenv("USER_")
	go func() {
		defer func() {
			fmt.Println("defer here")
		}()

		if user == "" {
			panic("should set user env.")
		}
	}()

	time.Sleep(1 * time.Second)
	fmt.Printf("get result %s\r\n", user)

}

func TestPanicRecovery(t *testing.T) {
	defer fmt.Println("defer main") // will this be called when panic?
	var user = os.Getenv("USER_")
	go func() {
		defer func() {
			fmt.Println("defer caller")
			if err := recover(); err != nil {
				fmt.Println("recover success.")
			}
		}()
		func() {
			defer func() {
				fmt.Println("defer here")
			}()

			if user == "" {
				panic("should set user env.")
			}
			fmt.Println("after panic")
		}()
	}()

	time.Sleep(1 * time.Second)
	fmt.Printf("get result %s\r\n", user)
}

func TestPanicDemo(t *testing.T) {
	// 错误的 recover 调用示例
	recover()         // 什么都不会捕捉
	panic("not good") // 发生 panic，主程序退出
	recover()         // 不会被执行
	println("ok")
}

func TestPanicRecoveryDemo(t *testing.T) {
	// 正确的 recover 调用示例
	defer func() {
		fmt.Println("recovered: ", recover())
	}()
	panic("not good")
}
