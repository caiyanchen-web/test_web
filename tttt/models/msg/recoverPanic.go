package msg

import "fmt"

func RecoverPanic() {
	if r := recover(); r != nil {
		fmt.Println("发生panic", r)
	}
}
