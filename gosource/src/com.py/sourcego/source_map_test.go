package sourcego

import (
	"fmt"
	"testing"
)

func TestSourceMapValues(t *testing.T) {
	SourceMapValues()
}

//32---1598013607
func TestFnv32(t *testing.T) {
	fnv32 := Fnv32("2020012300000610")
	fmt.Println(fnv32)
	sp :=fnv32%1024
	fmt.Println(sp)
	var ss map[int]int = nil
	if ss == nil {
		fmt.Println("ok")
	}
	fmt.Println(len(ss))
}