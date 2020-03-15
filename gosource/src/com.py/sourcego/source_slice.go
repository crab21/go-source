package sourcego


func SliceCopyOfNormal(des, source []string) {
	des = source
}
func SliceCopyOfApi(des,source []string) {
	copy(des,source)
}

func SliceAppendMake(source []string) {
	a:=make([]string,len(source),len(source))

	source = append(a,source...)

}

func SliceAppendNormal(source []string) {
	var b []string
	source = append(b,source...)
}