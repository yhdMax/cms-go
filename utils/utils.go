package utils

import (
	"fmt"
	"github.com/google/uuid"
	"log"
)

func Uuid() string {
	id, err := uuid.NewUUID()

	if err != nil {
		// 返回错误信息并中断程序
		log.Fatal(err)
	}
	return fmt.Sprintf("%v", id)
}
