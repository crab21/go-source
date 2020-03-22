package sourcego

import "fmt"

// []byte复制

//赋值方式
func SliceCopyOfNormal(des, source []string) {
	des = source
}

//copy方式   可以剔除原有的引用，产生一个新的[]byte
func SliceCopyOfApi(des, source []string) {
	copy(des, source)
}

//slice append方式

//make后再添加
func SliceAppendMake(source []string) {
	a := make([]string, len(source), len(source))

	source = append(a, source...)

}

//直接添加 √
func SliceAppendNormal(source []string) {
	var b []string
	source = append(b, source...)
}

//截取
func MapInit() {
	a := []string{"1", "2", "3", "4", "5",}
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
	bam := Bam{Age: 10,}
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
