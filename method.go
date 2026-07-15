package main
import "fmt"

// method là hàm hoạt động bên trong struct
//  

type Counter int 

func (c Counter) String() string {
	return fmt.Sprintf("Counter(%d)", c)
}

func main() {
	var c Counter = 3
	fmt.Println(c.String())
	
}