package msg

import "fmt"

//用于检查函数中对应的err

func CheckErr(err error) string {
	if err != nil {
		fmt.Printf("存在错误：%s", err)
		fmt.Println("")
	}
	return ""
}
