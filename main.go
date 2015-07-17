package c

// int test() { return 42; }
import "C"
import (
	"fmt"
)

func Test() int {
	fmt.Println("ok")
	return int(C.test())
}
