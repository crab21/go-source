package sourcemap

import (
	"fmt"
	"unsafe"
)

func ChangeMapValue(a map[int]int) {
	a[1] = 100
	a[2] = 200
}

func ChangeMapValueNil(a map[int]int) {
	a = make(map[int]int)
	a[1] = 100
	a[2] = 200
}

func ChangeSlice(a []int) {
	a[0] = 300
	a[1] = 20000
}

/**
slice append:
/usr/local/go/src/reflect/value.go

func grow(s Value, extra int) (Value, int, int) {
	.....
	t := MakeSlice(s.Type(), i1, m)
	Copy(t, s)
	......
}
 */
func ChangeSliceAppend(a []int) {
	//当原本a容量不足时，会扩容  导致下面赋值语句作用于新的slice对象

	a = append(a, 999,888)
	fmt.Println("ChangeSliceAppend=== len: ",len(a)," cap: ",cap(a))
	a[0] = 300
	a[1] = 20000
	fmt.Println("ChangeSliceAppend data:",a)
}
func ChangeSliceNil(a []int) {
	a = make([]int, 2)
	a[0] = 300
	a[1] = 20000

}


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


