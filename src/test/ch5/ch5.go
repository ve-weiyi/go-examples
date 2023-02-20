package ch5

type People struct {
	Name  string
	Sex   string
	Age   int
	phone string
}

func (m *People) HelloWorld() string {
	return m.Name + ":hello world 1"
}

type Student struct {
	//Name  string
	//Sex   string
	//Age   int
	//phone string

	//People   //嵌套
	P     People //组合
	Class string
	Email string
}

//func (m *Student) HelloWorld() string {
//	return m.Name + ":hello world 2"
//}
