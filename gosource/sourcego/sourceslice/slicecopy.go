package sourceslice
func SliceCP() {
	xForMakeCopy := make([]int,10)
	yForMakeCopy :=make([]int,10)
	copy(yForMakeCopy, xForMakeCopy)
}
