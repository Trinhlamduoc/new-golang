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

type Foo struct{}

func(Foo) Speak() string {
	return "Hello, I am Foo"
}

func main() {
	var someSpeak Speaker = Foo{}
	fmt.Println(someSpeak.Speak())
}