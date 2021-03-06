package sourceslice

import (
	"fmt"
)

func SliceCopyOfNormal(des, source []string) {
	// []byte复制
	//赋值方式
	des = source
}

//copy方式   可以剔除原有的引用，产生一个新的[]byte
func SliceCopyOfApi(des, source []string) {
	copy(des, source)
}

//高效复制   复用原地址，减少内存开辟时间 ✅✅✅✅✅✅✅✅✅
//后续通过汇编查看两个的重要区别
func SliceCopyOfSelf() {
	a := make([]int, 10)
	fmt.Println(a)
	b := append(a[:0:0], a...)
	fmt.Println(b)
	b[2] = 111
	fmt.Println(a, "   ", b)
}

//slice append方式

//make后再添加 √ [长度已知]
func SliceAppendMake(source []string) {
	a := make([]string, 0, len(source))

	a = append(a, source...)

}

//直接添加
func SliceAppendNormal(source []string) {
	var b []string
	b = append(b, source...)
}

//截取
func MapInit() {
	a := []string{"1", "2", "3", "4", "5"}
	b := make([]string, len(a))

	//深度拷贝  不仅是值得拷贝，附带底层地址的改变
	copy(b, a)
	//0xc0000b4080
	fmt.Printf(" %p \n", &a)
	//0xc0000b40a0
	fmt.Printf(" %p \n", &b)
	//[1 2 3]
	//3
	//4
	fmt.Println(b)
	fmt.Println(len(b))
	fmt.Println(cap(b))
	sourceSliceInterception(b)
	fmt.Println(b) //[wwwwwww 2 3 4 5]
	fmt.Println(a) //[1 2 3 4 5]
	bam := Bam{Age: 10}
	bam.BamChange()
	fmt.Println(bam)
	BamChange(&bam)
	fmt.Println(bam)
}

type Bam struct {
	Age int
}

func (bm *Bam) BamChange() {
	bm.Age = -1
}
func BamChange(bm *Bam) {
	bm.Age = -2
}
func sourceSliceInterception(sl []string) {
	sl[0] = "wwwwwww"

}

/**
边界检查分析
//go tool compile -d ssa/prove/debug=2 source_slice.go
*/

func NumSameBytes_1(x, y string) int {
	if len(x) > len(y) {
		x, y = y, x
	}
	for i := 0; i < len(x); i++ {
		if x[i] != y[i] {
			return i
		}
	}
	return len(x)
}

func NumSameBytes_2(x, y string) int {
	if len(x) > len(y) {
		x, y = y, x
	}
	if len(x) <= len(y) { // 虽然代码多了，但是效率提高了
		for i := 0; i < len(x); i++ {
			if x[i] != y[i] { // 边界检查被消除了
				return i
			}
		}
	}
	return len(x)
}

/**
slice1: [2 6 7 8]
arr1: [1 2 6 7 8]
slice2: [2 3 6 7 8]
arr2: [1 2 3 4 5]


*/

//当原数组容量够用，则在原来的基础上扩容  ✅✅✅✅✅✅✅✅✅✅✅✅✅
//详细原因参考go-plan9: https://plan9.io/sources/contrib/ericvh/go-plan9/src/pkg/
func AppendSliceError() {
	arr1 := [5]int{1, 2, 3, 4, 5}
	slice1 := arr1[1:2]
	slice1 = append(slice1, 6, 7, 8)
	fmt.Println("slice1:", slice1)
	fmt.Println("arr1:", arr1)
	arr2 := [5]int{1, 2, 3, 4, 5}
	slice2 := arr2[1:3]
	slice2 = append(slice2, 6, 7, 8)
	fmt.Println("slice2:", slice2)
	fmt.Println("arr2:", arr2)
}

//逃逸分析
//去掉最后一行打印  则输出32，32
//go run -gcflags="-m -N" **.go
func EscapeSliceSize() {
	s := "a"
	x := []byte(s)              // len(s) == 1
	fmt.Println(cap([]byte(s))) // 32
	fmt.Println(cap(x))         // 8
	fmt.Println(x)
}

func SliceAppendSlice() {
	num1 := make([]int, 0, 100)
	num1 = append(num1, 1)
	num2 := []int{11, 22, 33, 44, 55}
	num1s := append(num1, num2...)
	num2[2] = 6666
	num1[0] = 9999
	fmt.Println(num1s)
}

func MakeSliceIntercept(){
	var x = [20]int{3: 1, 9: 20}
	fmt.Println(x)

}
