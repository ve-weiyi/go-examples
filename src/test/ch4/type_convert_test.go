package ch4

import (
	"fmt"
	"log"
	"testing"
)

// 格式化输出：https://blog.csdn.net/shn111/article/details/127877743
// %T输出类型，%v输出值
func TestCovert(t *testing.T) {
	// number类型只能转换number类型  int、float
	num := 100
	flo := float32(num)
	num64 := int64(num)

	log.Printf("%T-->%v", num, num)
	log.Printf("%T-->%v", flo, flo)
	log.Printf("%T-->%v", num64, num64)

	//#不能把int类型赋值给int64
	//num64 == num

	//不能把int类型转换为string，需要使用fmt.Sprint()
	//str := string(num)
	str := string(fmt.Sprint(num))
	log.Printf("%T-->%v", str, str)
}
