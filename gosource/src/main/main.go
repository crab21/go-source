package main

import (
	"com.py/sourcego"
	"com.py/sourcego/sourcehttp"
)

func smain() {
	for a:=0;a<1000 ;a++  {
		a := []string{"a", "b", "c"}
		var bs []string
		sourcego.SliceCopyOfApi(bs,a)

	}
}
func main() {
	sourcehttp.WebsocketServerStart()
}