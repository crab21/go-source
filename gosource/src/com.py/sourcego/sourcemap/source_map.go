package sourcemap

import "fmt"

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



