package chater3

import (
	"fmt"
	"log"
	"testing"
)

func TestName(t *testing.T) {
	//1.声明、赋值
	var name1 string
	name1 = "声明、赋值"
	log.Printf("%v", name1)
	//2.声明并赋值
	var name2 = "声明并赋值a"
	log.Printf("%v", name2)
	//3.声明并赋值 --->推荐,使用较多
	name3 := "声明并赋值b"
	log.Printf("%v", name3)
	//4.声明并赋值,但是指针类型
	name4 := new(string)
	log.Printf("%v", name4)

	//5.同时声明多个变量
	var (
		vname1 = "a"
		vname2 = "b"
		vname3 string
	)
	vname3 = "c"

	log.Printf("%v", vname1)
	log.Printf("%v", vname2)
	log.Printf("%v", vname3)
}

func TestName2(t *testing.T) {
	//声明数组
	array := [5]int{2, 4, 6, 8, 10}
	printArr(&array)

	//声明时初始化
	slice := []int{2, 4, 6, 8, 10}
	printSlice(slice)

	//从数组切片,对于  slice:=&array
	slice2 := array[:]
	printSlice(slice2)

	//append组合
	slice3 := append(array[1:], 1)
	printSlice(slice3)

	s2 := make([]int, 6, 8) // 使用 make 创建，指定 len 和 cap 值。
	fmt.Println(s2, len(s2), cap(s2))

	s3 := make([]int, 6) // 省略 cap，相当于 cap = len。
	fmt.Println(s3, len(s3), cap(s3))
}

func printArr(arr *[5]int) {
	fmt.Println("printArr")
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func printSlice(arr []int) {
	fmt.Println("printSlice")
	for i, v := range arr {
		fmt.Println(i, v)
	}
}
func main() {
	scoreMap := make(map[string]int)
	scoreMap["张三"] = 90
	scoreMap["小明"] = 100
	// 如果key存在ok为true,v为对应的值；不存在ok为false,v为值类型的零值
	v, ok := scoreMap["张三"]
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("查无此人")
	}
}
