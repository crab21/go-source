package sourcemap

import (
	"fmt"
	"testing"
)

//slice|map在函数入参之前，若只有声明，没有赋值，则函数结束后不影响本身值。
/**
map[1:100 2:200]
map[]

[300 20000]

[]

 len:  2  cap:  2
ChangeSliceAppend=== len:  4  cap:  4
ChangeSliceAppend data: [300 20000 999 888]
[3 333]

 len:  2  cap:  20
ChangeSliceAppend=== len:  4  cap:  20
ChangeSliceAppend data: [300 20000 999 888]
[300 20000]

 */
func TestChangeMapValue(t *testing.T) {
	var sp = map[int]int{1:20}
	ChangeMapValue(sp)
	//map[1:100 2:200]
	fmt.Println(sp)

	var spNil map[int]int
	ChangeMapValueNil(spNil)
	// []
	fmt.Println(spNil)
	fmt.Println()

	var sliceTestData = []int{3,333}
	ChangeSlice(sliceTestData)
	//[300 20000]
	fmt.Println(sliceTestData)
	fmt.Println()

	var sliceTestDataNil []int
	ChangeSliceNil(sliceTestDataNil)
	// []
	fmt.Println(sliceTestDataNil)
	fmt.Println()

	var sliceTestDataAppend = []int{3,333}
	fmt.Println(" len: ",len(sliceTestDataAppend)," cap: ",cap(sliceTestDataAppend))
	ChangeSliceAppend(sliceTestDataAppend)
	//[3 333]
	fmt.Println(sliceTestDataAppend)
	fmt.Println()

	var sliceTestDataAppendMake = make([]int,0,20)
	sliceTestDataAppendMake = append(sliceTestDataAppendMake, 111,222)

	//len:  2  cap:  20
	fmt.Println(" len: ",len(sliceTestDataAppendMake)," cap: ",cap(sliceTestDataAppendMake))
	ChangeSliceAppend(sliceTestDataAppendMake)
	//[300 20000]   cap足够大，则add新元素，属于值拷贝，第0位和第1位值的改变，依旧会影响原slice值。
	fmt.Println(sliceTestDataAppendMake)

}

