package main

import "fmt"

// đây là cấu trúc, có X,Y kiểu dữ liệu là int
// khai báo 

// type
// tên struct, 
// chữ hoa viết đầu: exported, package bên ngoài có thể nhìn thấy
// chữ hoa không viết đầu: unexported
type Vertex struct {
	X, Y int
}

func main() {
	v := Vertex{X: 1, Y: 2}
	p:= v

	p.X = 1223

	fmt.Println(p)
}