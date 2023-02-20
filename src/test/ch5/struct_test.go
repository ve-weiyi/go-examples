package ch5

import (
	"github.com/ve-weiyi/go-examplse/src/utils/jsonconv"
	"log"
	"testing"
)

func TestName(t *testing.T) {
	log.SetFlags(log.LstdFlags | log.Llongfile)

	people := People{
		Name:  "jack",
		Sex:   "1",
		Age:   0,
		phone: "156",
	}
	log.Printf(people.HelloWorld())
	log.Printf("%+v", jsonconv.ObjectToJson(people))

	student := Student{
		//Name:  "jack",
		//Sex:   "1",
		//Age:   0,
		//phone: "156",
		//People: people,
		P:     people,
		Class: "123",
		Email: "xxx@qq.com",
	}

	//log.Printf(student.HelloWorld())
	log.Printf("%+v", jsonconv.ObjectToJson(student))

}
