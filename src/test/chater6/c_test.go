package chater6

import (
	"log"
	"testing"
)

func TestOverload(t *testing.T) {
	overload("a", 1, "2", 3.0)
}

func overload(name string, vs ...interface{}) {
	log.Println("name-->", name)

	//遍历
	for i, v := range vs {
		log.Printf("%v-->%v", i, v)
	}

	//判断有几个参数
	switch len(vs) {
	case 0:
	case 1:
	case 2:
	default:
		log.Println("len-->", len(vs))
	}
}
