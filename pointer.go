package main

import (
	"fmt"
	// "runtime"
	// "math"
)

func main() {
		number := 10
		
		//  &number là tham chiếu ô nhớ của number ví dụ: 0x24be030960c0
		pointer := &number
		// đọc giá trị *pointer tương đương đang đọc giá trị của biến number
		fmt.Println(*pointer)
		// tham trị
		*pointer = 21
		fmt.Println(*pointer)
}
