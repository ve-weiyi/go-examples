package chater2

import (
	log2 "github.com/ve-weiyi/go-examplse/src/utils/log"
	"log"
	"testing"
)

func Test1(t *testing.T) {

	//小写包内可见
	log.Printf(helloWorld())
	log.Printf(HelloWorld())
}

func Test2(t *testing.T) {
	log.Printf("init")
	log2.MyPrintln("init")

	//小写包外不可见
	//log2.println("sss")
}
