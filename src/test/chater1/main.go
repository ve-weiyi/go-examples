// 1.包名：包名一般与目录相同,相同目录只能有一个包名
package main

// 2.导包：可以使用 import "p1"  import "p2"
// 或者 import ("p1","p2" )
import (
	"log"
)

// 3.init在导入包时运行，先于main运行
func init() {
	log.SetFlags(log.LstdFlags | log.Llongfile)
	log.Println("init")
}

// 4.只有main包下的main才能运行
func main() {
	log.Println("main")
}

func Init() {

}

//输出：init   main
