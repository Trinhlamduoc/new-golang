package main

import (
	"fmt"
)

// interface là 1 kiểu được định nghĩa bởi tập hợp các method,
// interface có thể chứa bất kỳ giá trị gì miễn nó có implement các method này

// khai báo interface
type Speaker interface {
	Speak() string
}

// cú pháp interface rỗng
// interface{}

// switch case type
func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println("Got an integer:", v)
	case string:
		fmt.Println("Got a string:", v)
	default:
		fmt.Println("Got something else:", v)
	}
}

type Foo struct{}

func (Foo) Speak() string {
	return "Hello, I am Foo"
}

func main() {
	var someSpeak Speaker = Foo{}
	fmt.Println(someSpeak.Speak())
}
