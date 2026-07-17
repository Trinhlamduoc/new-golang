package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X,Y float64
}

// Đây là receiver value
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X * v.Y + v.X * v.Y)
}

// Đây là receiver pointer
// thay đổi luôn struct của (thao tác trực tiếp trên biến gốc thay vì tạo bản sao)

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3,4}
	v.Scale(10)
	fmt.Println(v.Abs())
}