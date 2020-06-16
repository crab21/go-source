package sourceslice

import "fmt"

/**
slice make过程中的问题
*/

func SliceMake1() {
	var num []int
	add1 := fmt.Sprintf("%p", num)
	fmt.Println(num, " addr: ", &num, " addr pointer:", add1) //[]  addr:  &[]  addr pointer: 0x0

	//append 1
	num = append(num, 1, 2, 3)
	add2 := fmt.Sprintf("%p", num)
	fmt.Println(num, " addr: ", &num, " addr pointer:", add2) //[1 2 3]  addr:  &[1 2 3]  addr pointer: 0xc000018160

	//append 2
	num = append(num, 1, 2, 3)
	add3 := fmt.Sprintf("%p", num)
	fmt.Println(num, " addr: ", &num, " addr pointer:", add3) //[1 2 3 1 2 3]  addr:  &[1 2 3 1 2 3]  addr pointer: 0xc00001e100
}

//初始化容量为10，长度为0
func SliceMake2() {
	var num = make([]int, 0, 10)
	add1 := fmt.Sprintf("%p", num)
	fmt.Println(num, " addr: ", &num, " addr pointer:", add1) //[]  addr:  &[]  addr pointer: 0xc00001c1e0

	//append 1
	num = append(num, 1, 2, 3)
	add2 := fmt.Sprintf("%p", num)
	fmt.Println(num, " addr: ", &num, " addr pointer:", add2) //[1 2 3]  addr:  &[1 2 3]  addr pointer: 0xc00001c1e0

	//append 2
	num = append(num, 1, 2, 3)
	add3 := fmt.Sprintf("%p", num)
	fmt.Println(num, " addr: ", &num, " addr pointer:", add3) //[1 2 3 1 2 3]  addr:  &[1 2 3 1 2 3]  addr pointer: 0xc00001c1e0

	//append 3
	num = append(num, 1, 2, 3, 4) //刚好到cap容量大小时，并未扩容
	add4 := fmt.Sprintf("%p", num)
	fmt.Println(num, " addr: ", &num, " addr pointer:", add4) //[1 2 3 1 2 3 1 2 3 4]  addr:  &[1 2 3 1 2 3 1 2 3 4]  addr pointer: 0xc00001c1e0

	//append 4
	num = append(num, 111) //超过cap的容量一个值
	add5 := fmt.Sprintf("%p", num)
	fmt.Println(num, " addr: ", &num, " addr pointer:", add5) //[1 2 3 1 2 3 1 2 3 4 111]  addr:  &[1 2 3 1 2 3 1 2 3 4 111]  addr pointer: 0xc0000f2000
}

//初始化容量为10，长度为10
func SliceMake3() {
	var num = make([]int, 10)
	add1 := fmt.Sprintf("%p", num)
	fmt.Println(num, " addr: ", &num, " addr pointer:", add1) //[0 0 0 0 0 0 0 0 0 0]  addr:  &[0 0 0 0 0 0 0 0 0 0]  addr pointer: 0xc00001c1e0

	//append 1
	num = append(num, 1, 2, 3) //扩容到原来两倍
	add2 := fmt.Sprintf("%p", num)
	fmt.Println(num, " addr: ", &num, " addr pointer:", add2) //[0 0 0 0 0 0 0 0 0 0 1 2 3]  addr:  &[0 0 0 0 0 0 0 0 0 0 1 2 3]  addr pointer: 0xc00005e000

	//append 2
	num = append(num, 1, 2, 3)
	add3 := fmt.Sprintf("%p", num)
	fmt.Println(num, " addr: ", &num, " addr pointer:", add3) //[0 0 0 0 0 0 0 0 0 0 1 2 3 1 2 3]  addr:  &[0 0 0 0 0 0 0 0 0 0 1 2 3 1 2 3]  addr pointer: 0xc00005e000

	//append 3
	num = append(num, 1, 2, 3, 4) //刚好到cap容量大小时，并未扩容
	add4 := fmt.Sprintf("%p", num)
	fmt.Println(num, " addr: ", &num, " addr pointer:", add4) //[0 0 0 0 0 0 0 0 0 0 1 2 3 1 2 3 1 2 3 4]  addr:  &[0 0 0 0 0 0 0 0 0 0 1 2 3 1 2 3 1 2 3 4]  addr pointer: 0xc00005e000

	//append 4
	num = append(num, 111) //超过cap的容量一个值
	add5 := fmt.Sprintf("%p", num)
	fmt.Println(num, " addr: ", &num, " addr pointer:", add5) //[0 0 0 0 0 0 0 0 0 0 1 2 3 1 2 3 1 2 3 4 111]  addr:  &[0 0 0 0 0 0 0 0 0 0 1 2 3 1 2 3 1 2 3 4 111]  addr pointer: 0xc000060000
}

//初始化cap:10  len:0
func SLiceMake4() {
	var num = make([]int, 0, 10)

	var addNum = make([]int, 0, 1025)
	for i := 0; i < 1024; i++ {
		addNum = append(addNum, i)
	}
	fmt.Println("addNum len:", len(addNum), "  cap:", cap(addNum)) //addNum len: 1024   cap: 1025
	fmt.Println("num len:", len(num), "  cap:", cap(num))          //num len: 0   cap: 10

	num = append(num, addNum...)
	fmt.Println("num len:", len(num), "  cap:", cap(num)) //num len: 1024   cap: 1024

	num = append(num, 1, 2, 3)
	// 若大于1024  则扩容1/4
	fmt.Println("num len:", len(num), "  cap:", cap(num)) //num len: 1027   cap: 1280
}
