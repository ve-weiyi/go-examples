package consumer

import (
	"fmt"
	"log"
	"testing"
	"time"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Llongfile)
	log.Println("invoke")
}

func Test7(t *testing.T) {
	// 传参的时候显式类型像隐式类型转换，双向管道向单向管道转换
	ch := make(chan string) //无缓冲channel
	go producer(ch)         // 子go程作为生产者
	consumer(ch)            // 主go程作为消费者
}

func producer(out chan<- string) {
	for i := 0; i < 10; i++ {
		data := time.Now().String()
		time.Sleep(1 * time.Second)
		fmt.Println("生产者生产数据:", data)
		out <- data // 缓冲区写入数据
	}
	close(out) //写完关闭管道
}

func consumer(in <-chan string) {
	// 无需同步机制，先做后做
	// 没有数据就阻塞等
	for data := range in {
		fmt.Println("消费者得到数据：", data)
	}

}
