package main

import "fmt"

type Vertex struct {
	X, Y int
}

func main() {
	v := Vertex{X: 1, Y: 2}
	p:= &v
	p.X = 1e9
	fmt.Println(*p)
}