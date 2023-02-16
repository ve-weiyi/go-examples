package log

import (
	"fmt"
)

func init() {
	fmt.Println("my log init ")
}

func MyPrintln(str string) {
	fmt.Println(str)
}

func println(str string) {
	fmt.Println(str)
}
