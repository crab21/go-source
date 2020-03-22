package sourcego

import (
	"fmt"
	"unsafe"
)

func SourceMapValues() {
	type Entity struct {
		Desc string
	}

	entityMap := make(map[string]Entity, 0)
	entityMap["Cup"] = Entity{Desc: "This is a Cup"}

	// 这里会编译报错，因为entityMap["Cup"]是不可寻址的，所以不能直接访问内部变量
	//entityMap["Cup"].Desc = "This is a special Cup"
	count := **(**int)(unsafe.Pointer(&entityMap))

	fmt.Println(count)
}

func Fnv32(key string) uint64 {
	hash := uint64(14695981039346656037)
	prime32 := uint64(1099511628211)
	for i := 0; i < len(key); i++ {
		hash *= prime32
		hash ^= uint64(key[i])
	}
	return hash
}