package sourceslice

import (
	"fmt"
	"unsafe"
)

/**
slice截取的各种情况
 */

func SliceInte_1() {
	number := make([]int, 20)
	fmt.Println(number) //[0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]

	number_inte:=number[1:9:11]
	fmt.Println(number_inte) //[0 0 0 0 0 0 0 0]

	number_inte[0]=20
	//before: [0 20 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]  cap: 20  len: 20
	fmt.Println("before:",number," cap:",cap(number)," len:",len(number))
	//now [20 0 0 0 0 0 0 0]  cap: 10  len: 8
	fmt.Println("now",number_inte," cap:",cap(number_inte)," len:",len(number_inte))

	//addr:  0xc00000c0c0   0xc00000c080
	fmt.Println("addr: ",unsafe.Pointer(&number_inte)," ",unsafe.Pointer(&number))
	number_inte1 := append(number_inte,1)
	//append later now: [20 0 0 0 0 0 0 0 1]  cap: 10  len: 9
	fmt.Println("append later now:",number_inte1," cap:",cap(number_inte1)," len:",len(number_inte1))
	//addr:  0xc00000c180
	fmt.Println("addr: ",unsafe.Pointer(&number_inte1))

	number_inte2 := append(number_inte,11,12,13)
	//2-append later now: [20 0 0 0 0 0 0 0 11 12 13]  cap: 20  len: 11
	fmt.Println("2-append later now:",number_inte2," cap:",cap(number_inte2)," len:",len(number_inte2))
	//2-addr:  0xc00000c1c0
	fmt.Println("2-addr: ",unsafe.Pointer(&number_inte2))
}

