package sourcego

// []byte复制

//赋值方式
func SliceCopyOfNormal(des, source []string) {
	des = source
}
//copy方式   可以剔除原有的引用，产生一个新的[]byte
func SliceCopyOfApi(des,source []string) {
	copy(des,source)
}

//slice append方式

//make后再添加
func SliceAppendMake(source []string) {
	a:=make([]string,len(source),len(source))

	source = append(a,source...)

}

//直接添加 √
func SliceAppendNormal(source []string) {
	var b []string
	source = append(b,source...)
}