package main

import "com.py/sourcego"

func main() {
	for a:=0;a<1000 ;a++  {
		a := []string{"a", "b", "c"}
		var bs []string
		sourcego.SliceCopyOfApi(bs,a)

	}
}