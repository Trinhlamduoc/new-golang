package main

import (
	"fmt"
	"math"
)


type MyFloat float64
// trong method Abs không thể khai báo receiver trực tiếp là float64 được, phải thông qua tạo kiểu mới là type
// vì trong Go không cho phép khai báo method trên kiểu được định nghĩa ở package khác, float64 là kiểu dựng sẵn của Go

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}