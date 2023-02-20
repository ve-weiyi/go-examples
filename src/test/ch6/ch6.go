package ch6

type Sayer interface {
	HelloWorld() string
}

type People struct {
	name string
}

func (m *People) HelloWorld() string {
	return "hello world"
}
