package sourcestruct

import (
	"fmt"
	"unsafe"
)

func StructSizeof() {
	type T1 struct {
		a struct{}
		x int64
	}
	fmt.Println(unsafe.Sizeof(T1{})) // 8

	type T2 struct {
		x int64
		a struct{}
	}
	fmt.Println(unsafe.Sizeof(T2{})) // 16


	//new 
	t := new(T1)
	fmt.Println(unsafe.Sizeof(t))//8


	t2 := new(T2)
	fmt.Println(unsafe.Sizeof(t2))//8

}

