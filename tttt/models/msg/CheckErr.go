package msg

import (
	"log"
)

//用于检查函数中对应的err

func CheckErr(err error) string {
	if err != nil {
		log.Printf("存在错误：%s", err)

	}
	return ""
}
