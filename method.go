package main
import "fmt"

// method là hàm hoạt động bên trong struct
//  

type Counter int 

// Counter là recever phần nhận vào
// String() là tên hàm, chỉ dùn 1 recever cụ thể
// method là function nhưng dữ liệu nhận vào của method là kiểu struct hoặc non-struct

// lưu ý:
// Nếu muốn dùng tên của func giống nhau thì phải có recever khác nhau
// Các method có thể trùng tên nếu xác định trên các dữ liệu khác nhau trong khi các hàm thì không được phép trùng tên
func (c Counter) String2() string {
	return fmt.Sprintf("Counter(%d) 111", c)
}

func main() {
	var c Counter = 2
	fmt.Println(c.String2())
	

}