package sourcebufio

import (
	"bufio"
	"fmt"
	"os"
)

func BufferRead() {
	open, _ := os.Open("/data.log")
	reader := bufio.NewReader(open)
	//声明缓存部分大小
	buf := make([]byte, 1024)
	read, _ := reader.Read(buf)
	//读取多少字节
	sp := buf[0:read]
	fmt.Println(sp)
	bufio.NewWriter(open)
}

const (
	mutexLocked = 1 << iota // mutex is locked
	mutexWoken
	mutexStarving
	mutexWaiterShift = iota
)

func PrintIotaInfo() {
	fmt.Println(mutexLocked)
	fmt.Println(mutexWoken)
	fmt.Println(mutexStarving)
	fmt.Println(mutexWaiterShift)
}
