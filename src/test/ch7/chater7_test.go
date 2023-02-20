package ch7

import (
	"errors"
	"log"
	"testing"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Llongfile)
	log.Println("invoke")
}

// 1.defer语句在函数运行结束后才执行
func TestDefer(t *testing.T) {
	log.Println("defer begin")
	// 将defer放入延迟调用栈
	defer log.Println(1)
	defer log.Println(2)
	// 最后一个放入, 位于栈顶, 最先调用
	defer log.Println(3)
	log.Println("defer end")
}

// 3.panic异常，会使程序崩溃
func TestPanic(t *testing.T) {
	panic("I'm a panic")
	//不可达
	log.Printf("hello world:%v", "panic")
}

// 4.recover恢复，捕获panic
func TestErr(t *testing.T) {
	rec := func() {
		if err := recover(); err != nil {
			v, ok := err.(string)
			if ok {
				log.Println(v) // 将 interface{} 转型为具体类型。
			} else {
				log.Printf("类型断言失败:%T %v", err, err)
			}
		}
	}

	defer rec()
	if str, err := getError(); err != nil {
		panic(err)
	} else {
		log.Println(str)
	}
}

func getError() (string, error) {
	return "", errors.New("this is a error")
	//return "this is a str", nil
}
