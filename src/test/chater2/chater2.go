package chater2

import "log"

func init() {
	log.SetFlags(log.LstdFlags | log.Llongfile)
	log.Println("sys log init ")
}

func helloWorld() string {
	return "hello world"
}

func HelloWorld() string {
	return "hello world"
}
